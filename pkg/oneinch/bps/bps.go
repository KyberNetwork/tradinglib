package bps

import "math/big"

func FromFraction(val int, base *big.Int) uint16 {
	bps := new(big.Int).SetInt64(int64(val) * 10_000)
	return uint16(bps.Div(bps, base).Int64())
}

func FromPercent(val int, base *big.Int) uint16 {
	bps := new(big.Int).SetInt64(int64(val) * 100)
	return uint16(bps.Div(bps, base).Int64())
}
