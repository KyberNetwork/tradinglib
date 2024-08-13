package fusionutils

import "math/big"

const (
	FeeBase             = 100_000
	BpsBase             = 10_000
	BpsToRatioNumerator = FeeBase / BpsBase
)

func BpsToRatioFormat(bps int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(bps), big.NewInt(BpsToRatioNumerator))
}

func AddRatioToAmount(amount *big.Int, ratio *big.Int) *big.Int {
	return new(big.Int).Add(amount, new(big.Int).Div(new(big.Int).Mul(amount, ratio), big.NewInt(FeeBase)))
}
