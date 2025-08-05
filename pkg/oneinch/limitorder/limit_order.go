package limitorder

import (
	"math/big"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/util"
	"github.com/ethereum/go-ethereum/common"
)

type LimitOrderV4 struct {
	OrderHash    string         `json:"orderHash"`
	Salt         *big.Int       `json:"salt"`
	Maker        common.Address `json:"maker"`
	Receiver     common.Address `json:"receiver"`
	MakerAsset   common.Address `json:"makerAsset"`
	TakerAsset   common.Address `json:"takerAsset"`
	MakingAmount *big.Int       `json:"makingAmount"`
	TakingAmount *big.Int       `json:"protocolFee"`
	MakerTraits  *MakerTraits   `json:"makerTraits"`
}

func (l LimitOrderV4) CalcTakingAmount(
	makingAmount *big.Int,
) *big.Int {
	return util.CalcTakingAmount(makingAmount, l.MakingAmount, l.TakingAmount)
}

func (l LimitOrderV4) CalcMakingAmount(
	takingAmount *big.Int,
) *big.Int {
	return util.CalcMakingAmount(takingAmount, l.MakingAmount, l.TakingAmount)
}
