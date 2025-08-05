package limitorder

import (
	"math/big"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/util"
	"github.com/ethereum/go-ethereum/common"
)

type LimitOrderWithFee struct {
	LimitOrder        LimitOrderV4 `json:"limit_order"`
	Extension         Extension    `json:"extension"`
	feeTakerExtension FeeTakerExtension
}

func NewLimitOrderWithFee(
	limitOrder LimitOrderV4,
	extension Extension,
) (LimitOrderWithFee, error) {
	loWithFee := LimitOrderWithFee{
		LimitOrder: limitOrder,
		Extension:  extension,
	}

	if extension.IsEmpty() {
		return loWithFee, nil
	}
	feeTakerExtension, err := NewFeeTakerFromExtension(extension)
	if err != nil {
		return LimitOrderWithFee{}, err
	}
	loWithFee.feeTakerExtension = feeTakerExtension
	return loWithFee, nil
}

func (l LimitOrderWithFee) CalcTakingAmount(
	taker common.Address,
	makingAmount *big.Int,
) *big.Int {
	takingAmount := l.LimitOrder.CalcTakingAmount(makingAmount)
	if l.Extension.IsEmpty() {
		return takingAmount
	}
	return l.feeTakerExtension.GetTakingAmount(taker, takingAmount)
}

func (l LimitOrderWithFee) CalcMakingAmount(
	taker common.Address,
	takingAmount *big.Int,
) *big.Int {
	makingAmount := util.CalcMakingAmount(takingAmount, l.LimitOrder.MakingAmount, l.LimitOrder.TakingAmount)
	if l.Extension.IsEmpty() {
		return makingAmount
	}
	return l.feeTakerExtension.GetMakingAmount(taker, makingAmount)
}
