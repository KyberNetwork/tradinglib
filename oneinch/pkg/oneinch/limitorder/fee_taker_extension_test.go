//nolint:testpackage
package limitorder

import (
	"math/big"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestCalculateTakingAmount(t *testing.T) {
	testCase := []struct {
		name         string
		feeTakerEx   FeeTakerExtension
		takerAddress common.Address
		takingAmount *big.Int
		expected     *big.Int
	}{
		{
			name: "normal case, taker is in white list",
			feeTakerEx: FeeTakerExtension{
				Address: common.HexToAddress("0x1"),
				Fees: Fees{
					Resolver: ResolverFee{
						Receiver:          common.HexToAddress("0x2"),
						Fee:               1,
						WhitelistDiscount: 0,
					},
				},
				feeCalculator: FeeCalculator{
					fees: Fees{
						Resolver: ResolverFee{
							Receiver:          common.HexToAddress("0x2"),
							Fee:               100,
							WhitelistDiscount: 0,
						},
					},
					whitelist: Whitelist{
						Addresses: []util.AddressHalf{util.HalfAddressFromAddress(common.HexToAddress("0x3"))},
					},
				},
			},
			takerAddress: common.HexToAddress("0x3"),
			takingAmount: big.NewInt(100_000_000),
			expected:     big.NewInt(101_000_000),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			takingAmount := tc.feeTakerEx.GetTakingAmount(tc.takerAddress, tc.takingAmount)
			assert.True(t, tc.expected.Cmp(takingAmount) == 0)
		})
	}
}

func TestCalculateResolverFee(t *testing.T) {
	testCase := []struct {
		name         string
		feeTakerEx   FeeTakerExtension
		takerAddress common.Address
		resolverFee  *big.Int
		expected     *big.Int
	}{
		{
			name: "normal case, taker is in white list",
			feeTakerEx: FeeTakerExtension{
				Address: common.HexToAddress("0x1"),
				Fees: Fees{
					Resolver: ResolverFee{
						Receiver:          common.HexToAddress("0x2"),
						Fee:               1,
						WhitelistDiscount: 0,
					},
				},
				feeCalculator: FeeCalculator{
					fees: Fees{
						Resolver: ResolverFee{
							Receiver: common.HexToAddress("0x2"),
							Fee:      100,
						},
					},
					whitelist: Whitelist{
						Addresses: []util.AddressHalf{util.HalfAddressFromAddress(common.HexToAddress("0x3"))},
					},
				},
			},
			takerAddress: common.HexToAddress("0x3"),
			resolverFee:  big.NewInt(100_000_000),
			expected:     big.NewInt(1_000_000),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			takingAmount := tc.feeTakerEx.GetResolverFee(tc.takerAddress, tc.resolverFee)
			assert.True(t, tc.expected.Cmp(takingAmount) == 0)
		})
	}
}

func TestCalculateIntegratorFee(t *testing.T) {
	testCase := []struct {
		name          string
		feeTakerEx    FeeTakerExtension
		takerAddress  common.Address
		integratorFee *big.Int
		expected      *big.Int
	}{
		{
			name: "normal case, taker is in white list",
			feeTakerEx: FeeTakerExtension{
				Address: common.HexToAddress("0x1"),
				Fees: Fees{
					Integrator: IntegratorFee{
						Integrator: common.HexToAddress("0x2"),
						Protocol:   common.HexToAddress("0x3"),
						Fee:        1,
						Share:      100,
					},
				},
				feeCalculator: FeeCalculator{
					fees: Fees{
						Integrator: IntegratorFee{
							Integrator: common.HexToAddress("0x2"),
							Protocol:   common.HexToAddress("0x3"),
							Fee:        500,
							Share:      1000,
						},
					},
					whitelist: Whitelist{
						Addresses: []util.AddressHalf{util.HalfAddressFromAddress(common.HexToAddress("0x4"))},
					},
				},
			},
			takerAddress:  common.HexToAddress("0x4"),
			integratorFee: big.NewInt(100_000_000),
			expected:      big.NewInt(500_000),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			integrator := tc.feeTakerEx.GetIntegratorFee(tc.takerAddress, tc.integratorFee)
			assert.True(t, tc.expected.Cmp(integrator) == 0)
		})
	}
}

func TestGetProtocolFee(t *testing.T) {
	testCase := []struct {
		name         string
		feeTakerEx   FeeTakerExtension
		takerAddress common.Address
		protocolFee  *big.Int
		expected     *big.Int
	}{
		{
			name: "normal case, taker is in white list",
			feeTakerEx: FeeTakerExtension{
				Address: common.HexToAddress("0x1"),
				Fees: Fees{
					Integrator: IntegratorFee{
						Integrator: common.HexToAddress("0x2"),
						Protocol:   common.HexToAddress("0x3"),
						Fee:        1,
						Share:      100,
					},
				},
				feeCalculator: FeeCalculator{
					fees: Fees{
						Resolver: ResolverFee{
							Receiver: common.HexToAddress("0x2"),
							Fee:      100,
						},
						Integrator: IntegratorFee{
							Integrator: common.HexToAddress("0x2"),
							Protocol:   common.HexToAddress("0x3"),
							Fee:        500,
							Share:      1000,
						},
					},
					whitelist: Whitelist{
						Addresses: []util.AddressHalf{util.HalfAddressFromAddress(common.HexToAddress("0x4"))},
					},
				},
			},
			takerAddress: common.HexToAddress("0x4"),
			protocolFee:  big.NewInt(100_000_000),
			expected:     big.NewInt(500_000 + 1_000_000),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			protocolFee := tc.feeTakerEx.GetProtocolFee(tc.takerAddress, tc.protocolFee)
			assert.True(t, tc.expected.Cmp(protocolFee) == 0)
		})
	}
}

// order: 0x1bf7bba4b4140c5a02431c5854c5170f2cc15599c7dade20aa099ca0037d1cf7
// tx_hash: 0x1c988bf44efa704617494fbd0dfca5e1049801af1df6aa8809896a526724431f
func TestNewFeeTakerExtension(t *testing.T) {
	extension, err := DecodeExtension("0x000000d400000072000000720000007200000072000000390000000000000000c0dfdb9e7a392c3dbbe7c6fbe8fbc1789c9fe05e00000001f43203b09498030ae3416b66dc74db31d09524fa87b1f7d18bd45f0b94f54a968fc0dfdb9e7a392c3dbbe7c6fbe8fbc1789c9fe05e00000001f43203b09498030ae3416b66dc74db31d09524fa87b1f7d18bd45f0b94f54a968fc0dfdb9e7a392c3dbbe7c6fbe8fbc1789c9fe05e00000000000000000000000000000000000000000090cbe4bdd538d6e9b379bff5fe72c3d67a521de500000001f43203b09498030ae3416b66dc74db31d09524fa87b1f7d18bd45f0b94f54a968f")
	assert.NoError(t, err)
	feeTakerExtension, err := NewFeeTakerFromExtension(extension)
	assert.NoError(t, err)
	resolver := common.HexToAddress("0xbee3211ab312a8d065c4fef0247448e17a8da000")
	t.Log(feeTakerExtension.GetMakingAmount(resolver, big.NewInt(4371259792644370)))
}

func TestDecodeExtension(t *testing.T) {
	testCase := []struct {
		name      string
		extension string
		hasError  bool
	}{
		{
			name:      "normal case",
			extension: "0x000000e800000072000000720000007200000072000000390000000000000000c0dfdb9e7a392c3dbbe7c6fbe8fbc1789c9fe05e00000001f43203b09498030ae3416b66dc74db31d09524fa87b1f7d18bd45f0b94f54a968fc0dfdb9e7a392c3dbbe7c6fbe8fbc1789c9fe05e00000001f43203b09498030ae3416b66dc74db31d09524fa87b1f7d18bd45f0b94f54a968fc0dfdb9e7a392c3dbbe7c6fbe8fbc1789c9fe05e01000000000000000000000000000000000000000090cbe4bdd538d6e9b379bff5fe72c3d67a521de509cc0a79dfef324587c6c9bc814d3a5a072e71de00000001f43203b09498030ae3416b66dc74db31d09524fa87b1f7d18bd45f0b94f54a968f",
			hasError:  false,
		},
		{
			name:      "extension has only permit is non empty",
			extension: "0x000000f4000000f4000000f40000000000000000000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb480000000000000000000000006a4c4d5429324fa5dbd0321558385af577713050000000000000000000000000111111125421ca6dc452d289314280a0f8842a65ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000000000000000000000000000000000000000000076f69b73000000000000000000000000000000000000000000000000000000000000001c4a0096cbe47829b3fd81c23858f747f07ebdea9e59b5b61ac99f2b1aa3960a0d2b0315f388762b3d61f1058305635d01028599a1fbd7210c16e5b82ccef5324f",
			hasError:  false,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			extension, err := DecodeExtension(tc.extension)
			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, extension)
				_, err = NewFeeTakerFromExtension(extension)
				assert.NoError(t, err)
			}
		})
	}
}

func TestAnotherDecodeExtension(t *testing.T) {
	extension, err := DecodeExtension("0x000000d400000072000000720000007200000072000000390000000000000000c0dfdb9e7a392c3dbbe7c6fbe8fbc1789c9fe05e00000001f43203b09498030ae3416b66dc74db31d09524fa87b1f7d18bd45f0b94f54a968fc0dfdb9e7a392c3dbbe7c6fbe8fbc1789c9fe05e00000001f43203b09498030ae3416b66dc74db31d09524fa87b1f7d18bd45f0b94f54a968fc0dfdb9e7a392c3dbbe7c6fbe8fbc1789c9fe05e00000000000000000000000000000000000000000090cbe4bdd538d6e9b379bff5fe72c3d67a521de500000001f43203b09498030ae3416b66dc74db31d09524fa87b1f7d18bd45f0b94f54a968f")
	assert.NoError(t, err)
	feeTakerExtension, err := NewFeeTakerFromExtension(extension)
	assert.NoError(t, err)
	t.Log(feeTakerExtension.Fees)
}
