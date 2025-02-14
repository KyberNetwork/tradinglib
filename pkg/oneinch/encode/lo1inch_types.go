package encode

import (
	"math/big"
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
