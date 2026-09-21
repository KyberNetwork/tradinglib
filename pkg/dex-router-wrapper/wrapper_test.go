//nolint:testpackage
package dexrouterwrapper

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncodeWrapCalldata(t *testing.T) {
	testCases := []struct {
		name        string
		target      common.Address
		data        []byte
		outputToken common.Address
		recipient   common.Address
	}{
		{
			name:        "erc20 output",
			target:      common.HexToAddress("0x1111111111111111111111111111111111111111"),
			data:        []byte{0xde, 0xad, 0xbe, 0xef},
			outputToken: common.HexToAddress("0x2222222222222222222222222222222222222222"),
			recipient:   common.HexToAddress("0x3333333333333333333333333333333333333333"),
		},
		{
			name:        "native output",
			target:      common.HexToAddress("0x4444444444444444444444444444444444444444"),
			data:        []byte{},
			outputToken: NativeTokenAddress,
			recipient:   common.HexToAddress("0x5555555555555555555555555555555555555555"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			packed, err := EncodeWrapCalldata(tc.target, tc.data, tc.outputToken, tc.recipient)
			require.NoError(t, err)

			method := WrapperABI.Methods[methodWrap]
			require.Equal(t, method.ID, packed[:4])

			unpacked, err := method.Inputs.Unpack(packed[4:])
			require.NoError(t, err)

			var args struct {
				Target      common.Address
				Data        []byte
				OutputToken common.Address
				Recipient   common.Address
			}
			require.NoError(t, method.Inputs.Copy(&args, unpacked))

			assert.Equal(t, tc.target, args.Target)
			assert.Equal(t, tc.data, args.Data)
			assert.Equal(t, tc.outputToken, args.OutputToken)
			assert.Equal(t, tc.recipient, args.Recipient)
		})
	}
}

func TestDecodeWrapOutput(t *testing.T) {
	testCases := []struct {
		name             string
		wantReturnAmount *big.Int
		wantGasUsed      uint64
	}{
		{name: "zero output", wantReturnAmount: big.NewInt(0), wantGasUsed: 0},
		{name: "typical output", wantReturnAmount: big.NewInt(123456789), wantGasUsed: 87654},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			packedOutput, err := WrapperABI.Methods[methodWrap].Outputs.Pack(
				tc.wantReturnAmount, new(big.Int).SetUint64(tc.wantGasUsed),
			)
			require.NoError(t, err)

			gotReturnAmount, gotGasUsed, err := DecodeWrapOutput(packedOutput)
			require.NoError(t, err)

			assert.Equal(t, 0, tc.wantReturnAmount.Cmp(gotReturnAmount))
			assert.Equal(t, tc.wantGasUsed, gotGasUsed)
		})
	}
}

func TestDecodeWrapOutput_InvalidData(t *testing.T) {
	_, _, err := DecodeWrapOutput([]byte{0x01, 0x02})
	assert.Error(t, err)
}

func TestNativeTokenAddress(t *testing.T) {
	// NativeTokenAddress must be the EIP-55 checksum of the all-0xEE
	// 20-byte address, not an arbitrary hardcoded string.
	want := common.HexToAddress("0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee")
	assert.Equal(t, want, NativeTokenAddress)
}

func TestDeployedBytecode(t *testing.T) {
	code := DeployedBytecode()
	require.NotEmpty(t, code)

	// The returned slice must be a copy: mutating it must not corrupt the
	// package's cached bytecode used by subsequent calls.
	code[0] ^= 0xFF
	again := DeployedBytecode()
	assert.NotEqual(t, code, again)
}

func TestCreationBytecode(t *testing.T) {
	code := CreationBytecode()
	require.NotEmpty(t, code)
	assert.NotEqual(t, code, DeployedBytecode())
}

func TestStateOverride(t *testing.T) {
	overrideAddr := common.HexToAddress("0x9999999999999999999999999999999999999999")

	overrides := StateOverride(overrideAddr)

	require.Contains(t, overrides, overrideAddr)
	assert.Equal(t, DeployedBytecode(), overrides[overrideAddr].Code)
}
