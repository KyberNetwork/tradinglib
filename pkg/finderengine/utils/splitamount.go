package utils

import (
	"math"
	"math/big"
)

const float64EqualityThreshold = 1e-9

func AlmostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func SplitAmount(amount *big.Int, splitNums uint64) []*big.Int {
	splitNumsBI := new(big.Int).SetUint64(splitNums)
	base := new(big.Int).Div(amount, splitNumsBI)
	remainder := new(big.Int).Sub(amount, new(big.Int).Mul(splitNumsBI, base))

	splits := make([]*big.Int, 0, splitNums)
	for i := uint64(0); i < splitNums; i++ {
		splits = append(splits, new(big.Int).Set(base))
	}
	if remainder.Cmp(big.NewInt(0)) != 0 {
		splits = append(splits, remainder)
	}

	return splits
}

func SplitAmountThreshold(
	amount *big.Int, decimals uint8, splitNums uint64, minThresholdUsd, price float64,
) []*big.Int {
	if amount == nil || amount.Sign() <= 0 || splitNums == 0 {
		return []*big.Int{new(big.Int).Set(amount)}
	}

	if minThresholdUsd <= 0 || price <= 0 {
		return SplitAmount(amount, splitNums)
	}

	scale := math.Pow10(int(decimals))
	minUnits := int64(math.Ceil((minThresholdUsd / price) * scale))
	if minUnits <= 0 {
		return SplitAmount(amount, splitNums)
	}

	maxSplits := new(big.Int).Quo(new(big.Int).Set(amount), big.NewInt(minUnits)).Uint64()
	if maxSplits == 0 {
		return []*big.Int{new(big.Int).Set(amount)}
	}
	if splitNums > maxSplits {
		splitNums = maxSplits
	}
	return SplitAmount(amount, splitNums)
}
