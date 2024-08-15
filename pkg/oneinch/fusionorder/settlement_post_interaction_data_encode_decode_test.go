package fusionorder_test

import (
	"math/big"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/fusionorder"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/fusionutils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Those tests are copied from
// https://github.com/1inch/fusion-sdk/blob/8721c62612b08cc7c0e01423a1bdd62594e7b8d0/src/fusion-order/settlement-post-interaction-data/settlement-post-interaction-data.spec.ts#L6
func TestSettlementPostInteractionData(t *testing.T) {
	t.Run("Should encode/decode with bank fee and whitelist", func(t *testing.T) {
		data, err := fusionorder.NewSettlementPostInteractionDataFromSettlementSuffixData(
			fusionorder.SettlementSuffixData{
				BankFee:            big.NewInt(1),
				ResolvingStartTime: big.NewInt(1708117482),
				Whitelist: []fusionorder.AuctionWhitelistItem{
					{
						Address:   common.BigToAddress(big.NewInt(0)),
						AllowFrom: big.NewInt(0),
					},
				},
			},
		)
		require.NoError(t, err)

		encoded := data.Encode()

		assert.Equal(t, 44, len(encoded))

		decoded, err := fusionorder.DecodeSettlementPostInteractionData(encoded)
		require.NoError(t, err)

		assertSettlementPostInteractionDataEqual(t, data, decoded)
	})

	t.Run("Should encode/decode with no fees and whitelist", func(t *testing.T) {
		data, err := fusionorder.NewSettlementPostInteractionDataFromSettlementSuffixData(
			fusionorder.SettlementSuffixData{
				ResolvingStartTime: big.NewInt(1708117482),
				Whitelist: []fusionorder.AuctionWhitelistItem{
					{
						Address:   common.BigToAddress(big.NewInt(0)),
						AllowFrom: big.NewInt(0),
					},
				},
			},
		)
		require.NoError(t, err)

		encoded := data.Encode()

		assert.Equal(t, 36, len(encoded))

		decoded, err := fusionorder.DecodeSettlementPostInteractionData(encoded)
		require.NoError(t, err)

		assertSettlementPostInteractionDataEqual(t, data, decoded)
	})

	t.Run("Should encode/decode with fees and whitelist", func(t *testing.T) {
		data, err := fusionorder.NewSettlementPostInteractionDataFromSettlementSuffixData(
			fusionorder.SettlementSuffixData{
				BankFee:            big.NewInt(0),
				ResolvingStartTime: big.NewInt(1708117482),
				Whitelist: []fusionorder.AuctionWhitelistItem{
					{
						Address:   common.BigToAddress(big.NewInt(0)),
						AllowFrom: big.NewInt(0),
					},
				},
				IntegratorFee: fusionorder.IntegratorFee{
					Ratio:    fusionutils.BpsToRatioFormat(10),
					Receiver: common.BigToAddress(big.NewInt(1)),
				},
			},
		)
		require.NoError(t, err)

		decoded, err := fusionorder.DecodeSettlementPostInteractionData(data.Encode())
		require.NoError(t, err)

		assertSettlementPostInteractionDataEqual(t, data, decoded)
	})
	t.Run("Should encode/decode with fees, custom receiver and whitelist", func(t *testing.T) {
		data, err := fusionorder.NewSettlementPostInteractionDataFromSettlementSuffixData(
			fusionorder.SettlementSuffixData{
				BankFee:            big.NewInt(0),
				ResolvingStartTime: big.NewInt(1708117482),
				Whitelist: []fusionorder.AuctionWhitelistItem{
					{
						Address:   common.BigToAddress(big.NewInt(0)),
						AllowFrom: big.NewInt(0),
					},
				},
				IntegratorFee: fusionorder.IntegratorFee{
					Ratio:    fusionutils.BpsToRatioFormat(10),
					Receiver: common.BigToAddress(big.NewInt(1)),
				},
				CustomReceiver: common.BigToAddress(big.NewInt(1337)),
			},
		)
		require.NoError(t, err)

		decoded, err := fusionorder.DecodeSettlementPostInteractionData(data.Encode())
		require.NoError(t, err)

		assertSettlementPostInteractionDataEqual(t, data, decoded)
	})

	t.Run("Should generate correct whitelist", func(t *testing.T) {
		start := big.NewInt(1708117482)

		data, err := fusionorder.NewSettlementPostInteractionDataFromSettlementSuffixData(
			fusionorder.SettlementSuffixData{
				ResolvingStartTime: start,
				Whitelist: []fusionorder.AuctionWhitelistItem{
					{
						Address:   common.BigToAddress(big.NewInt(2)),
						AllowFrom: new(big.Int).Add(start, big.NewInt(1_000)),
					},
					{
						Address:   common.BigToAddress(big.NewInt(0)),
						AllowFrom: new(big.Int).Sub(start, big.NewInt(10)), // should be set to start
					},
					{
						Address:   common.BigToAddress(big.NewInt(1)),
						AllowFrom: new(big.Int).Add(start, big.NewInt(10)),
					},
					{
						Address:   common.BigToAddress(big.NewInt(3)),
						AllowFrom: new(big.Int).Add(start, big.NewInt(10)),
					},
				},
			},
		)
		require.NoError(t, err)

		expectedWhitelist := []fusionorder.WhitelistItem{
			{
				AddressHalf: fusionorder.AddressHalf{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				Delay:       0,
			},
			{
				AddressHalf: fusionorder.AddressHalf{0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
				Delay:       10,
			},
			{
				AddressHalf: fusionorder.AddressHalf{0, 0, 0, 0, 0, 0, 0, 0, 0, 3},
				Delay:       0,
			},
			{
				AddressHalf: fusionorder.AddressHalf{0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
				Delay:       990,
			},
		}

		assert.ElementsMatch(t, expectedWhitelist, data.Whitelist)

		assert.True(t,
			data.CanExecuteAt(common.BigToAddress(big.NewInt(1)),
				new(big.Int).Add(start, big.NewInt(10))),
		)
		assert.False(t,
			data.CanExecuteAt(common.BigToAddress(big.NewInt(1)),
				new(big.Int).Add(start, big.NewInt(9))),
		)
		assert.True(t,
			data.CanExecuteAt(common.BigToAddress(big.NewInt(3)),
				new(big.Int).Add(start, big.NewInt(10))),
		)
		assert.False(t,
			data.CanExecuteAt(common.BigToAddress(big.NewInt(3)),
				new(big.Int).Add(start, big.NewInt(9))),
		)
		assert.False(t,
			data.CanExecuteAt(common.BigToAddress(big.NewInt(2)),
				new(big.Int).Add(start, big.NewInt(50))),
		)
	})
}

func assertSettlementPostInteractionDataEqual(t *testing.T, expected, actual fusionorder.SettlementPostInteractionData) {
	assert.ElementsMatch(t, expected.Whitelist, actual.Whitelist)
	assert.Equal(t, expected.IntegratorFee.Ratio, actual.IntegratorFee.Ratio)
	assert.Equal(t, expected.IntegratorFee.Receiver, actual.IntegratorFee.Receiver)
	assert.Equal(t, expected.BankFee, actual.BankFee)
	assert.Equal(t, expected.ResolvingStartTime, actual.ResolvingStartTime)
	assert.Equal(t, expected.CustomReceiver, actual.CustomReceiver)
}
