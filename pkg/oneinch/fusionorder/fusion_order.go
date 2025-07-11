package fusionorder

import (
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/auctioncalculator"
	"math/big"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/fusionorder/fusionextention"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/limitorder"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/util"
	"github.com/ethereum/go-ethereum/common"
)

type FusionOrder struct {
	Order           limitorder.LimitOrderV4         `json:"order"`
	Extension       limitorder.Extension            `json:"extension"`
	FusionExtension fusionextention.FusionExtension `json:"fusionExtension"`
}

func (o FusionOrder) GetCalculator() auctioncalculator.AmountCalculator {
	return auctioncalculator.NewAmountCalculatorFromExtension(o.FusionExtension)
}

func (o FusionOrder) CalcTakingAmount(
	taker common.Address,
	makingAmount *big.Int,
	blockTime,
	baseFee *big.Int,
) *big.Int {
	takingAmount := util.CalcTakingAmount(makingAmount, o.Order.MakingAmount, o.Order.TakingAmount)
	return o.GetCalculator().GetRequiredTakingAmount(taker, takingAmount, blockTime, baseFee)
}

func (o FusionOrder) CalcMakingAmount(
	taker common.Address,
	takingAmount *big.Int,
	blockTime,
	baseFee *big.Int,
) *big.Int {
	makingAmount := util.CalcMakingAmount(takingAmount, o.Order.MakingAmount, o.Order.TakingAmount)
	return o.GetCalculator().GetRequiredMakingAmount(taker, makingAmount, blockTime, baseFee)
}
