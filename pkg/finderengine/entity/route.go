package entity

import "math/big"

type Route struct {
	TokenIn  string
	TokenOut string

	AmountIn       *big.Int
	AmountInPrice  float64
	AmountOut      *big.Int
	AmountOutPrice float64

	GasUsed     int64
	GasFeePrice float64
	GasPrice    *big.Int
	GasFee      *big.Int

	L1GasFeePrice float64

	Paths []*Path
}

type Swap struct {
	Pool      string
	TokenIn   string
	TokenOut  string
	AmountIn  *big.Int
	AmountOut *big.Int
}

type FinalizedRoute struct {
	TokenIn  string
	TokenOut string

	AmountIn       *big.Int
	AmountInPrice  float64
	AmountOut      *big.Int
	AmountOutPrice float64

	GasUsed     int64
	GasFeePrice float64
	GasPrice    *big.Int
	GasFee      *big.Int

	L1GasFeePrice float64

	Route [][]Swap
}

func NewConstructRoute(tokenIn, tokenOut string) *Route {
	return &Route{
		TokenIn:   tokenIn,
		TokenOut:  tokenOut,
		AmountIn:  big.NewInt(0),
		AmountOut: big.NewInt(0),
		Paths:     []*Path{},
	}
}

type BestRouteResult struct {
	AMMBestRoute *Route
}

func (res *BestRouteResult) IsRouteNotFound() bool {
	return res == nil || res.AMMBestRoute == nil
}
