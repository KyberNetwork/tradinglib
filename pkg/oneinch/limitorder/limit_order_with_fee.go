package limitorder

import (
	"math/big"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/util"
	"github.com/ethereum/go-ethereum/common"
)

type LimitOrderWithFee struct {
	LimitOrder        LimitOrderV4 `json:"limit_order"`
	extension         Extension
	feeTakerExtension *FeeTakerExtension
}

func NewLimitOrderWithFee(
	limitOrder LimitOrderV4,
	extension Extension,
) (LimitOrderWithFee, error) {
	loWithFee := LimitOrderWithFee{
		LimitOrder: limitOrder,
		extension:  extension,
	}

	if extension.IsEmpty() {
		return loWithFee, nil
	}
	feeTakerExtension, err := NewFeeTakerFromExtension(extension)
	if err != nil {
		return LimitOrderWithFee{}, err
	}
	loWithFee.feeTakerExtension = &feeTakerExtension
	return loWithFee, nil
}

func (l LimitOrderWithFee) CalcTakingAmount(
	taker common.Address,
	makingAmount *big.Int,
) *big.Int {
	takingAmount := l.LimitOrder.CalcTakingAmount(makingAmount)
	if l.feeTakerExtension == nil {
		return takingAmount
	}
	return l.feeTakerExtension.GetTakingAmount(taker, takingAmount)
}

func (l LimitOrderWithFee) CalcMakingAmount(
	taker common.Address,
	takingAmount *big.Int,
) *big.Int {
	makingAmount := util.CalcMakingAmount(takingAmount, l.LimitOrder.MakingAmount, l.LimitOrder.TakingAmount)
	if l.feeTakerExtension == nil {
		return makingAmount
	}
	return l.feeTakerExtension.GetMakingAmount(taker, makingAmount)
}

func (l LimitOrderWithFee) GetFee(taker common.Address) (int64, int64) {
	if l.feeTakerExtension == nil {
		return 0, 0
	}
	return l.feeTakerExtension.feeCalculator.GetFeesForTaker(taker)
}
