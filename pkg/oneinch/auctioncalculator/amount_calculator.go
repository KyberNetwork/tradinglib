package auctioncalculator

import (
	"math/big"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/fusionorder/fusionextention"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/limitorder"
	"github.com/ethereum/go-ethereum/common"
)

// nolint:godox
type AmountCalculator struct {
	auctionCalculator Calculator
	// TODO: need to use an other FeeCalculator for FusionOrder
	feeCalculator limitorder.FeeCalculator
}

func NewAmountCalculator(
	auctionCalculator Calculator,
	feeCalculator limitorder.FeeCalculator,
) AmountCalculator {
	return AmountCalculator{
		auctionCalculator: auctionCalculator,
		feeCalculator:     feeCalculator,
	}
}

func NewAmountCalculatorFromFusionExtension(extension fusionextention.FusionExtension) AmountCalculator {
	auctionCalculator := NewCalculatorFromAuctionData(extension.AuctionDetails)
	feeCalculator := limitorder.NewFeeCalculator(extension.Extra.Fees, extension.Whitelist)
	return NewAmountCalculator(auctionCalculator, feeCalculator)
}

func (c AmountCalculator) CalcAuctionTakingAmount(baseTakingAmount *big.Int, rate int64, fee int64) *big.Int {
	beforeFee := CalcAuctionTakingAmount(baseTakingAmount, rate)
	if fee == 0 {
		return beforeFee
	}

	afterFee := beforeFee.Mul(beforeFee, big.NewInt(limitorder.Base10000+fee))
	return afterFee.Div(afterFee, big.NewInt(limitorder.Base10000))
}

func (c AmountCalculator) ExtractFeeAmount(requiredTakingAmount *big.Int, fee int64) *big.Int {
	if fee == 0 {
		return big.NewInt(0)
	}

	afterFee := big.NewInt(limitorder.Base10000)
	afterFee.Mul(requiredTakingAmount, afterFee)
	afterFee.Add(afterFee, big.NewInt(limitorder.Base10000+fee-1))
	afterFee.Div(afterFee, big.NewInt(limitorder.Base10000+fee))
	return afterFee.Sub(requiredTakingAmount, afterFee)
}

func (c AmountCalculator) GetRequiredTakingAmount(
	taker common.Address,
	takingAmount *big.Int,
	ts *big.Int,
	blockBaseFee *big.Int,
) *big.Int {
	withFee := c.feeCalculator.GetTakingAmount(taker, takingAmount)
	return c.getAuctionBumpedAmount(withFee, ts, blockBaseFee)
}

func (c AmountCalculator) GetBestRequiredTakingAmount(
	takingAmount *big.Int,
	ts *big.Int,
	blockBaseFee *big.Int,
) *big.Int {
	return c.getAuctionBumpedAmount(big.NewInt(0), ts, blockBaseFee)
}

func (c AmountCalculator) GetRequiredMakingAmount(
	taker common.Address,
	makingAmount *big.Int,
	ts *big.Int,
	blockBaseFee *big.Int,
) *big.Int {
	withFee := c.feeCalculator.GetMakingAmount(taker, makingAmount)
	rateBump := c.auctionCalculator.CalcRateBump(ts, blockBaseFee)
	return CalcAuctionMakingAmount(withFee, rateBump)
}

func (c AmountCalculator) getAuctionBumpedAmount(takingAmount *big.Int, ts *big.Int, blockBaseFee *big.Int) *big.Int {
	rateBump := c.auctionCalculator.CalcRateBump(ts, blockBaseFee)
	return CalcAuctionTakingAmount(takingAmount, rateBump)
}

func (c AmountCalculator) GetFee(taker common.Address) (int64, int64) {
	return c.feeCalculator.GetFeesForTaker(taker)
}
