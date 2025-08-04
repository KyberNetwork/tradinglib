package limitorder

import (
	"math/big"

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
