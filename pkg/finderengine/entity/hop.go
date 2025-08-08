package entity

import "math/big"

type HopSplit struct {
	ID            string
	AmountIn      *big.Int
	AmountOut     *big.Int
	GasUsed       *big.Int
	GasFeePrice   float64
	L1GasFeePrice float64
}

type Hop struct {
	TokenIn       string
	TokenOut      string
	AmountIn      *big.Int
	AmountOut     *big.Int
	GasUsed       int64
	GasFeePrice   float64
	L1GasFeePrice float64
	Splits        []*HopSplit
}
