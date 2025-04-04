package auction

import (
	"math/big"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/fusionorder"
)

const (
	GasPriceBase        = 1_000_000  // 1000 means 1 Gwei
	RateBumpDenominator = 10_000_000 // 100%
)

// Calculator is copied from
// nolint: lll
// https://github.com/1inch/fusion-sdk/blob/8721c62612b08cc7c0e01423a1bdd62594e7b8d0/src/auction-calculator/auction-calculator.ts#L10
type Calculator struct {
	startTime       *big.Int
	duration        *big.Int
	initialRateBump *big.Int
	points          []fusionorder.AuctionPoint
	gasCost         fusionorder.AuctionGasCostInfo
}

func NewCalculator(
	startTime int64,
	duration int64,
	initialRateBump int64,
	points []fusionorder.AuctionPoint,
	gasCost fusionorder.AuctionGasCostInfo,
) Calculator {
	return Calculator{
		startTime:       big.NewInt(startTime),
		duration:        big.NewInt(duration),
		initialRateBump: big.NewInt(initialRateBump),
		points:          points,
		gasCost:         gasCost,
	}
}

func NewCalculatorFromAuctionData(
	auctionDetails fusionorder.AuctionDetails,
) Calculator {
	return NewCalculator(
		auctionDetails.StartTime,
		auctionDetails.Duration,
		auctionDetails.InitialRateBump,
		auctionDetails.Points,
		auctionDetails.GasCost,
	)
}

func (c Calculator) FinishTime() *big.Int {
	return new(big.Int).Add(c.startTime, c.duration)
}

func (c Calculator) CalcRateBump(time, blockBaseFee *big.Int) int64 {
	gasBump := c.getGasPriceBump(blockBaseFee)
	auctionBump := c.getAuctionBump(time)

	final := big.NewInt(0)
	if auctionBump.Cmp(gasBump) > 0 { // auctionBump > gasBump
		final = new(big.Int).Sub(auctionBump, gasBump)
	}

	return final.Int64()
}

func (c Calculator) getGasPriceBump(blockBaseFee *big.Int) *big.Int {
	if blockBaseFee.Sign() == 0 || c.gasCost.GasPriceEstimate == 0 || c.gasCost.GasBumpEstimate == 0 {
		return big.NewInt(0)
	}

	gasPriceBump := big.NewInt(c.gasCost.GasBumpEstimate)
	gasPriceBump.Mul(gasPriceBump, blockBaseFee)
	gasPriceBump.Div(gasPriceBump, big.NewInt(c.gasCost.GasPriceEstimate))
	return gasPriceBump.Div(gasPriceBump, big.NewInt(GasPriceBase))
}

func (c Calculator) getAuctionBump(blockTime *big.Int) *big.Int {
	auctionFinishTime := c.FinishTime()

	if blockTime.Cmp(c.startTime) <= 0 { // blockTime <= startTime
		return c.initialRateBump
	}
	if blockTime.Cmp(auctionFinishTime) >= 0 { // blockTime >= finishTime
		return big.NewInt(0)
	}

	currentPointTime := c.startTime
	currentRateBump := c.initialRateBump

	for _, p := range c.points {
		nextRateBump := big.NewInt(p.Coefficient)
		nextPointTime := new(big.Int).Add(currentPointTime, big.NewInt(p.Delay))

		if blockTime.Cmp(nextPointTime) <= 0 { // blockTime <= nextPointTime
			// nolint: lll
			// This complicated formula below is copied from
			// smart_contract: https://github.com/1inch/limit-order-settlement/blob/2eef6f86bf0142024f9a8bf054a0256b41d8362a/contracts/extensions/BaseExtension.sol#L180
			// fusion_sdk: https://github.com/1inch/fusion-sdk/blob/8721c62612b08cc7c0e01423a1bdd62594e7b8d0/src/auction-calculator/auction-calculator.ts#L148
			diffToCurrent := new(big.Int).Sub(blockTime, currentPointTime)
			diffToNext := new(big.Int).Sub(nextPointTime, blockTime)
			totalDiff := new(big.Int).Sub(nextPointTime, currentPointTime)
			auctionBump := diffToCurrent.Mul(diffToCurrent, nextRateBump)
			auctionBump.Add(auctionBump, diffToNext.Mul(diffToNext, currentRateBump))
			return auctionBump.Div(auctionBump, totalDiff)
		}

		currentPointTime = nextPointTime
		currentRateBump = nextRateBump
	}

	auctionBump := new(big.Int).Sub(auctionFinishTime, blockTime)
	auctionBump.Mul(auctionBump, currentRateBump)
	return auctionBump.Div(auctionBump, new(big.Int).Sub(auctionFinishTime, currentPointTime))
}

func (c Calculator) CalcAuctionTakingAmount(takingAmount *big.Int, rate int64) *big.Int {
	return CalcAuctionTakingAmount(takingAmount, rate)
}

func (c Calculator) CalcAuctionMakingAmount(makingAmount *big.Int, rate int64) *big.Int {
	return CalcAuctionMakingAmount(makingAmount, rate)
}

func CalcAuctionTakingAmount(takingAmount *big.Int, rate int64) *big.Int {
	auctionTakingAmount := new(big.Int).Mul(takingAmount, big.NewInt(rate+RateBumpDenominator))
	return auctionTakingAmount.Div(auctionTakingAmount, big.NewInt(RateBumpDenominator))
}

func CalcAuctionMakingAmount(makingAmount *big.Int, rate int64) *big.Int {
	auctionMakingAmount := new(big.Int).Mul(makingAmount, big.NewInt(RateBumpDenominator))
	return auctionMakingAmount.Mul(auctionMakingAmount, big.NewInt(rate+RateBumpDenominator))
}

func CalcInitialRateBump(startAmount *big.Int, endAmount *big.Int) int64 {
	rateBumpDenominator := big.NewInt(RateBumpDenominator)
	bump := new(big.Int).Mul(rateBumpDenominator, startAmount)
	bump.Div(bump, endAmount)
	bump.Sub(bump, rateBumpDenominator)
	return bump.Int64()
}

func BaseFeeToGasPriceEstimate(baseFee *big.Int) *big.Int {
	return new(big.Int).Div(baseFee, big.NewInt(GasPriceBase))
}

func CalcGasBumpEstimate(endTakingAmount, gasCostInToToken *big.Int) *big.Int {
	gasBump := big.NewInt(RateBumpDenominator)
	gasBump.Mul(gasBump, gasCostInToToken)
	return gasBump.Div(gasBump, endTakingAmount)
}
