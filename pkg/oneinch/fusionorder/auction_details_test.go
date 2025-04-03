package fusionorder_test

import (
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/decode"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/fusionorder"
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
		decodedAuctionDetail, err := fusionorder.DecodeAuctionDetails(decode.NewBytesIterator(encodedAuctionDetail))
		require.NoError(t, err)

		assert.Equal(t, auctionDetail, decodedAuctionDetail)
	})

	// nolint: lll
	t.Run("decode real data", func(t *testing.T) {
		// Decode from a fusion order should not return error.
		extraData, err := hexutil.Decode("0x00ca070000023367ed11940000b401934d0100ca0700b40000000000640db09498030ae3416b66dc5dcd8578ca14eec99e63972ad4499f120902631ad18bd45f0b94f54a968fd61b892b2ad62490118595770895ad27ad6b0d95339fb574bdc56763f995617556ed277ab32233786de5e0e428ac771d77b55b7d1434eae4a48b2c8626813bd1b091ea6bedbd00000000000000000000b8394f2220fac7e6ade6")
		require.NoError(t, err)
		decodeAuctionDetails, err := fusionorder.DecodeAuctionDetails(decode.NewBytesIterator(extraData))
		if assert.NoError(t, err) {
			expectedStartTime := uint64(1743589780)
			expectedDuration := int64(180)
			initialRateBump := int64(103245)
			points := []fusionorder.AuctionPoint{
				{
					Delay:       180,
					Coefficient: 51719,
				},
			}
			gasBumpEstimate := int64(51719)
			gasPriceEstimate := int64(563)

			assert.EqualValues(t, expectedStartTime, decodeAuctionDetails.StartTime)
			assert.EqualValues(t, expectedDuration, decodeAuctionDetails.Duration)
			assert.EqualValues(t, initialRateBump, decodeAuctionDetails.InitialRateBump)
			assert.ElementsMatch(t, points, decodeAuctionDetails.Points)
			assert.EqualValues(t, gasBumpEstimate, decodeAuctionDetails.GasCost.GasBumpEstimate)
			assert.EqualValues(t, gasPriceEstimate, decodeAuctionDetails.GasCost.GasPriceEstimate)
		}
	})

	t.Run("should return error when data invalid", func(t *testing.T) {
		extraData, err := hexutil.Decode("0x01e9f1000005d866bda8b40000b404fa0103d4")
		require.NoError(t, err)

		_, err = fusionorder.DecodeAuctionDetails(decode.NewBytesIterator(extraData))

		require.ErrorIs(t, err, decode.ErrOutOfData)
	})
}
