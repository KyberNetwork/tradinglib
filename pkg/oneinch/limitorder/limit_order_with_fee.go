package limitorder

import (
	"math/big"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/util"
	"github.com/ethereum/go-ethereum/common"
)

type LimitOrderWithFee struct {
	LimitOrder        LimitOrderV4      `json:"limit_order"`
	FeeTakerExtension FeeTakerExtension `json:"fee_taker_extension"`
}

func (l LimitOrderWithFee) CalcTakingAmount(
	taker common.Address,
	makingAmount *big.Int,
) *big.Int {
	takingAmount := util.CalcTakingAmount(makingAmount, l.LimitOrder.MakingAmount, l.LimitOrder.TakingAmount)
	return l.FeeTakerExtension.GetTakingAmount(taker, takingAmount)
}

func (l LimitOrderWithFee) CalcMakingAmount(
	taker common.Address,
	takingAmount *big.Int,
) *big.Int {
	makingAmount := util.CalcMakingAmount(takingAmount, l.LimitOrder.MakingAmount, l.LimitOrder.TakingAmount)
	return l.FeeTakerExtension.GetMakingAmount(taker, makingAmount)
}
