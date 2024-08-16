package fusionorder_test

import (
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/decode"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/fusionorder"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// nolint: lll,funlen
func TestAuctionDetail(t *testing.T) {
	t.Run("should encode/decode", func(t *testing.T) {
		auctionDetail, err := fusionorder.NewAuctionDetails(
			1_673_548_149,
			50_000,
			180,
			[]fusionorder.AuctionPoint{
				{
					Delay:       10,
					Coefficient: 10_000,
				},
				{
					Delay:       20,
					Coefficient: 5_000,
				},
			},
			fusionorder.AuctionGasCostInfo{},
		)
		require.NoError(t, err)

		encodedAuctionDetail := auctionDetail.Encode()

		decodedAuctionDetail, err := fusionorder.DecodeAuctionDetails(encodedAuctionDetail)
		require.NoError(t, err)

		assertAuctionDetailsEqual(t, auctionDetail, decodedAuctionDetail)
	})

	t.Run("decode", func(t *testing.T) {
		makingAmountData, err := hexutil.Decode(
			"0xfb2809a5314473e1165f6b58018e20ed8f07b84000f1b8000005e566bb30120000b401def800f1b800b4",
		)
		require.NoError(t, err)
		decodeAuctionDetails, err := fusionorder.DecodeAuctionDetails(
			makingAmountData[common.AddressLength:],
		)
		require.NoError(t, err)

		t.Logf("AuctionDetails: %+v", decodeAuctionDetails)
	})

	// nolint: lll
	t.Run("decode real data", func(t *testing.T) {
		// This data is get from
		// https://app.blocksec.com/explorer/tx/eth/0x73e317981af9c352f26bac125b1a6d3e1d31076b87c679a4f771b4a5c5a7f76f?line=4&debugLine=4
		extraData, err := hexutil.Decode("0x01e9f1000005d866bda8b40000b404fa0103d477003c01e9f10078")
		require.NoError(t, err)
		decodeAuctionDetails, err := fusionorder.DecodeAuctionDetails(extraData)
		require.NoError(t, err)

		// those value is collected from
		// https://app.blocksec.com/explorer/tx/eth/0x73e317981af9c352f26bac125b1a6d3e1d31076b87c679a4f771b4a5c5a7f76f?line=79&debugLine=79
		expectedStartTime := uint64(1_723_705_524)
		expectedDuration := 1_723_705_704 - expectedStartTime
		initialRateBump := 326_145
		// those value is collected from running decode function in fusion-sdk with this extraData
		// https://github.com/1inch/fusion-sdk/blob/8721c62612b08cc7c0e01423a1bdd62594e7b8d0/src/fusion-order/auction-details/auction-details.ts#L76
		points := []fusionorder.AuctionPoint{
			{
				Delay:       60,
				Coefficient: 250_999,
			},
			{
				Delay:       120,
				Coefficient: 125_425,
			},
		}
		gasBumpEstimate := 125_425
		gasPriceEstimate := 1496

		assert.EqualValues(t, expectedStartTime, decodeAuctionDetails.StartTime)
		assert.EqualValues(t, expectedDuration, decodeAuctionDetails.Duration)
		assert.EqualValues(t, initialRateBump, decodeAuctionDetails.InitialRateBump)
		assert.ElementsMatch(t, points, decodeAuctionDetails.Points)
		assert.EqualValues(t, gasBumpEstimate, decodeAuctionDetails.GasCost.GasBumpEstimate)
		assert.EqualValues(t, gasPriceEstimate, decodeAuctionDetails.GasCost.GasPriceEstimate)
	})

	t.Run("should return error when data invalid", func(t *testing.T) {
		extraData, err := hexutil.Decode("0x01e9f1000005d866bda8b40000b404fa0103d4")
		require.NoError(t, err)

		_, err = fusionorder.DecodeAuctionDetails(extraData)

		require.ErrorIs(t, err, decode.ErrInvalidDataLength)
	})
}

func assertAuctionDetailsEqual(t *testing.T, expected, actual fusionorder.AuctionDetails) {
	t.Helper()
	assert.Equal(t, expected.StartTime, actual.StartTime)
	assert.Equal(t, expected.Duration, actual.Duration)
	assert.Equal(t, expected.InitialRateBump, actual.InitialRateBump)
	assert.ElementsMatch(t, expected.Points, actual.Points)
	assert.Equal(t, expected.GasCost.GasBumpEstimate, actual.GasCost.GasBumpEstimate)
	assert.Equal(t, expected.GasCost.GasPriceEstimate, actual.GasCost.GasPriceEstimate)
}
