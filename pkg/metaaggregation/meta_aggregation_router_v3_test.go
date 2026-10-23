//nolint:testpackage
package metaaggregation

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const swapV3Signature = "swap((address[],uint256[],(bytes,address[],uint256[],address[],uint256[])[]," +
	"address[],(uint256,address[],uint256[])[],bytes,address,bytes,address,uint256,bytes))"

func newSwapV3Params() SwapV3Params {
	return SwapV3Params{
		InputTokens:  []common.Address{common.HexToAddress("0x1"), common.HexToAddress("0x2")},
		InputAmounts: []*big.Int{big.NewInt(100), big.NewInt(200)},
		InputData: []SwapV3InputTokenData{
			{
				PermitData:    []byte{0x00},
				FeeRecipients: []common.Address{common.HexToAddress("0x3")},
				Fees:          []*big.Int{big.NewInt(1)},
				Targets:       []common.Address{common.HexToAddress("0x4")},
				Amounts:       []*big.Int{big.NewInt(99)},
			},
			{
				PermitData:    []byte{},
				FeeRecipients: []common.Address{},
				Fees:          []*big.Int{},
				Targets:       []common.Address{common.HexToAddress("0x5"), common.HexToAddress("0x6")},
				Amounts:       []*big.Int{big.NewInt(150), big.NewInt(50)},
			},
		},
		OutputTokens: []common.Address{common.HexToAddress("0x7")},
		OutputData: []SwapV3OutputTokenData{
			{
				MinAmount:     big.NewInt(42),
				FeeRecipients: []common.Address{common.HexToAddress("0x8")},
				Fees:          []*big.Int{big.NewInt(2)},
			},
		},
		Permit2Data:  []byte("permit2"),
		Executor:     common.HexToAddress("0x9"),
		ExecutorData: []byte("executor"),
		Recipient:    common.HexToAddress("0xa"),
		Deadline:     big.NewInt(1_800_000_000),
		ClientData:   []byte("client"),
	}
}

func TestRouterV3SwapSelector(t *testing.T) {
	method := RouterV3ABI.Methods[methodSwap]
	assert.Equal(t, swapV3Signature, method.Sig)
	assert.Equal(t, crypto.Keccak256([]byte(swapV3Signature))[:4], method.ID)
	assert.NotEqual(t, RouterV2ABI.Methods[methodSwap].ID, method.ID)
}

func TestEncodeDecodeMetaAggregationRouterSwapV3(t *testing.T) {
	params := newSwapV3Params()

	encoded, err := EncodeMetaAggregationRouterSwapData(&metaAggregationSwapV3Inputs{SwapV3Params: params})
	require.NoError(t, err)
	assert.Equal(t, RouterV3ABI.Methods[methodSwap].ID, encoded[:4])

	decoded, err := DecodeMetaAggregationRouterSwapData(encoded)
	require.NoError(t, err)
	require.Equal(t, SwapV3, decoded.Type())
	// nolint: forcetypeassert
	assert.Equal(t, params, decoded.(*metaAggregationSwapV3Inputs).SwapV3Params)
}

func TestMetaAggregationRouterSwapV3MinReturnAmount(t *testing.T) {
	encoded, err := packSwapV3Inputs(newSwapV3Params())
	require.NoError(t, err)
	input, err := DecodeMetaAggregationRouterSwapData(encoded)
	require.NoError(t, err)

	got := input.GetMinReturnAmount()
	assert.Equal(t, big.NewInt(42), got)
	got.SetInt64(1)
	assert.Equal(t, big.NewInt(42), input.GetMinReturnAmount(), "getter must return a copy")

	input.SetMinReturnAmount(big.NewInt(7))
	reencoded, err := EncodeMetaAggregationRouterSwapData(input)
	require.NoError(t, err)
	redecoded, err := DecodeMetaAggregationRouterSwapData(reencoded)
	require.NoError(t, err)
	assert.Equal(t, big.NewInt(7), redecoded.GetMinReturnAmount())
}

func TestMetaAggregationRouterSwapV3RequiresSingleOutput(t *testing.T) {
	testCases := []struct {
		name   string
		modify func(p *SwapV3Params)
		// The ABI cannot pack a nil uint256, so such params have no calldata to decode.
		skipDecode bool
	}{
		{
			name: "no output tokens",
			modify: func(p *SwapV3Params) {
				p.OutputTokens = []common.Address{}
				p.OutputData = []SwapV3OutputTokenData{}
			},
			skipDecode: false,
		},
		{
			name: "two output tokens",
			modify: func(p *SwapV3Params) {
				p.OutputTokens = append(p.OutputTokens, common.HexToAddress("0xb"))
				p.OutputData = append(p.OutputData, p.OutputData[0])
			},
			skipDecode: false,
		},
		{
			name: "mismatched output data",
			modify: func(p *SwapV3Params) {
				p.OutputData = append(p.OutputData, p.OutputData[0])
			},
			skipDecode: false,
		},
		{
			name: "nil min amount",
			modify: func(p *SwapV3Params) {
				p.OutputData[0].MinAmount = nil
			},
			skipDecode: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			params := newSwapV3Params()
			tc.modify(&params)

			_, err := EncodeMetaAggregationRouterSwapData(&metaAggregationSwapV3Inputs{SwapV3Params: params})
			require.ErrorIs(t, err, ErrNotSingleOutput)

			if tc.skipDecode {
				return
			}
			calldata, err := RouterV3ABI.Pack(methodSwap, params)
			require.NoError(t, err)
			_, err = DecodeMetaAggregationRouterSwapData(calldata)
			require.ErrorIs(t, err, ErrNotSingleOutput)
		})
	}
}

func TestDecodeMetaAggregationRouterSwapDataInvalidInput(t *testing.T) {
	testCases := []struct {
		name string
		data []byte
	}{
		{name: "empty", data: nil},
		{name: "short", data: []byte{0x01, 0x02}},
		{name: "unknown selector", data: []byte{0xde, 0xad, 0xbe, 0xef, 0x00}},
		{name: "V3 selector only", data: RouterV3ABI.Methods[methodSwap].ID},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := DecodeMetaAggregationRouterSwapData(tc.data)
			require.Error(t, err)
		})
	}
}

func TestDecodeMetaAggregationSwapV3Output(t *testing.T) {
	outputs := RouterV3ABI.Methods[methodSwap].Outputs
	pack := func(t *testing.T, amounts []*big.Int, gas int64) []byte {
		t.Helper()
		data, err := outputs.Pack(amounts, big.NewInt(gas))
		require.NoError(t, err)
		return data
	}

	t.Run("single output", func(t *testing.T) {
		amount, gas, err := DecodeMetaAggregationSwapV3Output(pack(t, []*big.Int{big.NewInt(42)}, 100000))
		require.NoError(t, err)
		assert.Equal(t, big.NewInt(42), amount)
		assert.Equal(t, uint64(100000), gas)
	})

	t.Run("no output amounts", func(t *testing.T) {
		_, _, err := DecodeMetaAggregationSwapV3Output(pack(t, []*big.Int{}, 100000))
		require.ErrorIs(t, err, ErrNotSingleOutput)
	})

	t.Run("two output amounts", func(t *testing.T) {
		_, _, err := DecodeMetaAggregationSwapV3Output(pack(t, []*big.Int{big.NewInt(1), big.NewInt(2)}, 100000))
		require.ErrorIs(t, err, ErrNotSingleOutput)
	})

	t.Run("zero gas", func(t *testing.T) {
		_, _, err := DecodeMetaAggregationSwapV3Output(pack(t, []*big.Int{big.NewInt(42)}, 0))
		require.EqualError(t, err, "gas is zero")
	})

	t.Run("truncated output", func(t *testing.T) {
		data := pack(t, []*big.Int{big.NewInt(42)}, 100000)
		_, _, err := DecodeMetaAggregationSwapV3Output(data[:64])
		require.Error(t, err)
	})
}
