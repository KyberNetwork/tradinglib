package auction_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/KyberNetwork/tradinglib/pkg/convert"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/auction"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/fusionorder"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuctionCalculator(t *testing.T) {
	t.Run("should be created successfully from suffix and salt", func(t *testing.T) {
		auctionStartTime := int64(1708448252)
		postInteraction, err := fusionorder.NewSettlementPostInteractionDataFromSettlementSuffixData(
			fusionorder.SettlementSuffixData{
				IntegratorFee: fusionorder.IntegratorFee{
					Ratio:    fusionorder.BpsToRatioFormat(1).Int64(),
					Receiver: common.BigToAddress(big.NewInt(1)),
				},
				BankFee:            0,
				ResolvingStartTime: auctionStartTime,
				Whitelist: []fusionorder.AuctionWhitelistItem{
					{
						Address:   common.BigToAddress(big.NewInt(0)),
						AllowFrom: auctionStartTime,
					},
				},
			})
		require.NoError(t, err)

		actionDetails, err := fusionorder.NewAuctionDetails(
			auctionStartTime,
			50_000,
			120,
			nil,
			fusionorder.AuctionGasCostInfo{},
		)
		require.NoError(t, err)

		calculator := auction.NewCalculatorFromAuctionData(postInteraction.IntegratorFee.Ratio, actionDetails)

		takingAmount, ok := new(big.Int).SetString("1420000000", 10)
		require.True(t, ok)

		rate := calculator.CalcRateBump(big.NewInt(auctionStartTime+60), big.NewInt(0))
		auctionTakingAmount := calculator.CalcAuctionTakingAmount(takingAmount, rate)

		assert.Equal(t, int64(25000), rate)
		assert.Zero(t, auctionTakingAmount.Cmp(big.NewInt(1423692355))) // 1423550000 from rate + 142355 (1bps) integrator fee
	})
}

func TestCalculator_GasBump(t *testing.T) {
	now := time.Now().Unix()
	duration := int64(1800) // 30 minutes
	takingAmount := parseEther(t, 1)
	calculator := auction.NewCalculator(
		now-60,
		duration,
		1000000,
		[]fusionorder.AuctionPoint{
			{
				Delay:       60,
				Coefficient: 500000,
			},
		},
		0,
		fusionorder.AuctionGasCostInfo{
			GasBumpEstimate:  10_000,
			GasPriceEstimate: 1000,
		},
	)

	t.Run("0 gwei = no gas fee", func(t *testing.T) {
		bump := calculator.CalcRateBump(big.NewInt(now), big.NewInt(0))
		auctionTakingAmount := calculator.CalcAuctionTakingAmount(takingAmount, bump)
		assert.Zero(t, auctionTakingAmount.Cmp(parseEther(t, 1.05)))
	})

	t.Run("0.1 gwei == 0.01% gas fee", func(t *testing.T) {
		bump := calculator.CalcRateBump(big.NewInt(now), parseUnits(t, 1, 8))
		auctionTakingAmount := calculator.CalcAuctionTakingAmount(takingAmount, bump)
		assert.Zero(t, auctionTakingAmount.Cmp(parseEther(t, 1.0499)))
	})

	t.Run("15 gwei == 1.5% gas fee", func(t *testing.T) {
		bump := calculator.CalcRateBump(big.NewInt(now), parseUnits(t, 15, 9))
		auctionTakingAmount := calculator.CalcAuctionTakingAmount(takingAmount, bump)
		assert.Zero(t, auctionTakingAmount.Cmp(parseEther(t, 1.035)))
	})

	t.Run("100 gwei == 10% gas fee", func(t *testing.T) {
		bump := calculator.CalcRateBump(big.NewInt(now), parseUnits(t, 100, 9))
		auctionTakingAmount := calculator.CalcAuctionTakingAmount(takingAmount, bump)
		assert.Zero(t, auctionTakingAmount.Cmp(parseEther(t, 1)))
	})
}

func parseEther(t *testing.T, ether float64) *big.Int {
	t.Helper()
	return parseUnits(t, ether, 18)
}

func parseUnits(t *testing.T, amount float64, decimals int64) *big.Int {
	t.Helper()
	v, err := convert.FloatToWei(amount, decimals)
	require.NoError(t, err)

	return v
}
