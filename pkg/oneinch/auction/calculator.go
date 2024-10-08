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
	takerFeeRatio   *big.Int
	gasCost         fusionorder.AuctionGasCostInfo
}

func NewCalculator(
	startTime int64,
	duration int64,
	initialRateBump int64,
	points []fusionorder.AuctionPoint,
	takerFeeRatio int64,
	gasCost fusionorder.AuctionGasCostInfo,
) Calculator {
	return Calculator{
		startTime:       big.NewInt(startTime),
		duration:        big.NewInt(duration),
		initialRateBump: big.NewInt(initialRateBump),
		points:          points,
		takerFeeRatio:   big.NewInt(takerFeeRatio),
		gasCost:         gasCost,
	}
}

func NewCalculatorFromAuctionData(
	takerFeeRatio int64,
	auctionDetails fusionorder.AuctionDetails,
) Calculator {
	return NewCalculator(
		auctionDetails.StartTime,
		auctionDetails.Duration,
		auctionDetails.InitialRateBump,
		auctionDetails.Points,
		takerFeeRatio,
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
	zeroBigInt := big.NewInt(0)
	if zeroBigInt.Cmp(blockBaseFee) == 0 {
		return zeroBigInt
	}
	if c.gasCost.GasPriceEstimate == 0 {
		return zeroBigInt
	}
	if c.gasCost.GasBumpEstimate == 0 {
		return zeroBigInt
	}

	return new(big.Int).Div(
		new(big.Int).Div(
			new(big.Int).Mul(
				big.NewInt(c.gasCost.GasBumpEstimate), blockBaseFee,
			),
			big.NewInt(c.gasCost.GasPriceEstimate),
		),
		big.NewInt(GasPriceBase),
	)
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
			return new(big.Int).Div(
				new(big.Int).Add(
					new(big.Int).Mul(
						new(big.Int).Sub(blockTime, currentPointTime),
						nextRateBump,
					),
					new(big.Int).Mul(
						new(big.Int).Sub(nextPointTime, blockTime),
						currentRateBump,
					),
				),
				new(big.Int).Sub(nextPointTime, currentPointTime),
			)
		}

		currentPointTime = nextPointTime
		currentRateBump = nextRateBump
	}

	return new(big.Int).Div(
		new(big.Int).Mul(
			new(big.Int).Sub(auctionFinishTime, blockTime),
			currentRateBump,
		),
		new(big.Int).Sub(auctionFinishTime, currentPointTime),
	)
}

func (c Calculator) CalcAuctionTakingAmount(takingAmount *big.Int, rate int64) *big.Int {
	return CalcAuctionTakingAmount(takingAmount, rate, c.takerFeeRatio)
}

func CalcAuctionTakingAmount(takingAmount *big.Int, rate int64, takerFeeRatio *big.Int) *big.Int {
	rateBumpDenominator := big.NewInt(RateBumpDenominator)
	auctionTakingAmount := new(big.Int).Div(
		new(big.Int).Mul(
			takingAmount,
			new(big.Int).Add(big.NewInt(rate), rateBumpDenominator),
		),
		rateBumpDenominator,
	)

	if takerFeeRatio.Cmp(big.NewInt(0)) == 0 {
		return auctionTakingAmount
	}

	return fusionorder.AddRatioToAmount(auctionTakingAmount, takerFeeRatio)
}

func CalcInitialRateBump(startAmount *big.Int, endAmount *big.Int) int64 {
	rateBumpDenominator := big.NewInt(RateBumpDenominator)
	bump := new(big.Int).Mul(
		new(big.Int).Div(
			new(big.Int).Mul(rateBumpDenominator, startAmount),
			endAmount,
		),
		rateBumpDenominator,
	)

	return bump.Int64()
}

func BaseFeeToGasPriceEstimate(baseFee *big.Int) *big.Int {
	return new(big.Int).Div(baseFee, big.NewInt(GasPriceBase))
}

func CalcGasBumpEstimate(endTakingAmount, gasCostInToToken *big.Int) *big.Int {
	return new(big.Int).Div(
		new(big.Int).Mul(
			gasCostInToToken,
			big.NewInt(RateBumpDenominator),
		),
		endTakingAmount,
	)
}
