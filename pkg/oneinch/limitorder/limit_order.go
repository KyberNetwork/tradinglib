package limitorder

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const noPartialFillsFlag = 255

type LimitOrderV4 struct {
	OrderHash    string         `json:"orderHash"`
	Salt         *big.Int       `json:"salt"`
	Maker        common.Address `json:"maker"`
	Receiver     common.Address `json:"receiver"`
	MakerAsset   common.Address `json:"makerAsset"`
	TakerAsset   common.Address `json:"takerAsset"`
	MakingAmount *big.Int       `json:"makingAmount"`
	TakingAmount *big.Int       `json:"takingAmount"`
	MakerTraits  *big.Int       `json:"makerTraits"`
}

func (o LimitOrderV4) IsPartialFillAllowed() bool {
	return o.MakerTraits.Bit(noPartialFillsFlag) == 0
}
