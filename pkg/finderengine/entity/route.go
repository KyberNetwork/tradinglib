package entity

import "math/big"

type Route struct {
	TokenIn  string
	TokenOut string

	AmountIn       *big.Int
	AmountOut      *big.Int
	AmountOutPrice float64

	GasUsed     int64
	GasFeePrice float64

	L1GasFeePrice float64

	Paths []*Path
}

func NewConstructRoute(tokenIn, tokenOut string) *Route {

	return &Route{
		TokenIn:  tokenIn,
		TokenOut: tokenOut,

		AmountIn:  big.NewInt(0),
		AmountOut: big.NewInt(0),
		Paths:     []*Path{},
	}
}

type BestRouteResult struct {
	BestRoutes   []*Route
	AMMBestRoute *Route
}

func (res *BestRouteResult) IsRouteNotFound() bool {
	return res == nil || len(res.BestRoutes) == 0
}
