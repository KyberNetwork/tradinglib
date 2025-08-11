package fusionorder

import (
	"math/big"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/auctioncalculator"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/fusionorder/fusionextention"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/limitorder"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/util"
	"github.com/ethereum/go-ethereum/common"
)

type FusionOrder struct {
	LimitOrder       limitorder.LimitOrderV4         `json:"limit_order"`
	FusionExtension  fusionextention.FusionExtension `json:"fusion_extension"`
	amountCalculator auctioncalculator.AmountCalculator
}

func NewFusionOrder(
	limitOrder limitorder.LimitOrderV4,
	fusionExtension fusionextention.FusionExtension,
) FusionOrder {
	return FusionOrder{
		LimitOrder:       limitOrder,
		FusionExtension:  fusionExtension,
		amountCalculator: auctioncalculator.NewAmountCalculatorFromFusionExtension(fusionExtension),
	}
}

func (o FusionOrder) CalcTakingAmount(
	taker common.Address,
	makingAmount *big.Int,
	blockTime,
	baseFee *big.Int,
) *big.Int {
	takingAmount := util.CalcTakingAmount(makingAmount, o.LimitOrder.MakingAmount, o.LimitOrder.TakingAmount)
	return o.amountCalculator.GetRequiredTakingAmount(taker, takingAmount, blockTime, baseFee)
}

func (o FusionOrder) CalcMakingAmount(
	taker common.Address,
	takingAmount *big.Int,
	blockTime,
	baseFee *big.Int,
) *big.Int {
	makingAmount := util.CalcMakingAmount(takingAmount, o.LimitOrder.MakingAmount, o.LimitOrder.TakingAmount)
	return o.amountCalculator.GetRequiredMakingAmount(taker, makingAmount, blockTime, baseFee)
}

func (o FusionOrder) GetFee(taker common.Address) (int64, int64) {
	return o.amountCalculator.GetFee(taker)
}
