package encode

import (
	"math/big"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/valueobject"
)

/*
	struct Order {
		uint256 salt;
		uint256 maker;
		uint256 receiver;
		uint256 makerAsset;
		uint256 takerAsset;
		uint256 makingAmount;
		uint256 takingAmount;
		uint256 makerTraits;
	}
*/
type OneInchV6Order struct {
	Salt         *big.Int
	Maker        *big.Int
	Receiver     *big.Int
	MakerAsset   *big.Int
	TakerAsset   *big.Int
	MakingAmount *big.Int
	TakingAmount *big.Int
	MakerTraits  *big.Int
}

type EncodingSwap struct {
	Pool              string
	TokenIn           string
	TokenOut          string
	SwapAmount        *big.Int
	AmountOut         *big.Int
	LimitReturnAmount *big.Int
	Exchange          valueobject.Exchange
	PoolLength        int
	PoolType          string
	PoolExtra         interface{}
	Extra             interface{}

	Flags []struct {
		Value uint32
	}

	CollectAmount *big.Int

	Recipient string
}
