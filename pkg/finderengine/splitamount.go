package finderengine

import (
	"math"
	"math/big"
)

const float64EqualityThreshold = 1e-9

func AlmostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func splitAmount(amount *big.Int, splitNums uint64) []*big.Int {
	splitNumsBI := new(big.Int).SetUint64(splitNums)
	base := new(big.Int).Div(amount, splitNumsBI)
	remainder := new(big.Int).Sub(amount, new(big.Int).Mul(splitNumsBI, base))

	splits := make([]*big.Int, splitNums)
	for i := uint64(0); i < splitNums; i++ {
		splits[i] = new(big.Int).Set(base)
	}
	splits[splitNums-1].Add(splits[splitNums-1], remainder)
	return splits
}
