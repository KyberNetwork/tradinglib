package utils

import (
	"math"
	"math/big"
)

func CalcAmountPrice(amount *big.Int, decimals uint8, price float64) float64 {
	amountFloat, _ := amount.Float64()
	return amountFloat * price / math.Pow10(int(decimals))
}

// CalcAmountUSD returns amount in from usd amount
// amount = (amountUSD / price) * 10^decimals
func CalcAmountFromPrice(amountUSD float64, decimals uint8, price float64) *big.Int {
	amountUSDBI := new(big.Float).SetFloat64(amountUSD)
	priceUSDBI := new(big.Float).SetFloat64(price)

	amount := amountUSDBI.Mul(amountUSDBI, new(big.Float).SetFloat64(math.Pow10(int(decimals))))
	result, _ := amount.Quo(amount, priceUSDBI).Int(nil)

	return result
}
