package util

import "math/big"

// CalcTakingAmount https://github.com/1inch/limit-order-protocol/blob/23d655844191dea7960a186652307604a1ed480a/contracts/libraries/AmountCalculatorLib.sol#L6
func CalcTakingAmount(swapMakerAmount, orderMakerAmount, orderTakerAmount *big.Int) *big.Int {
	amount := new(big.Int).Mul(swapMakerAmount, orderTakerAmount)
	amount.Add(amount, orderMakerAmount)
	amount.Sub(amount, big.NewInt(1))
	return amount.Div(amount, orderMakerAmount)
}

// CalcMakingAmount https://github.com/1inch/limit-order-protocol/blob/23d655844191dea7960a186652307604a1ed480a/contracts/libraries/AmountCalculatorLib.sol#L6
func CalcMakingAmount(swapTakerAmount, orderMakerAmount, orderTakerAmount *big.Int) *big.Int {
	amount := new(big.Int).Mul(swapTakerAmount, orderMakerAmount)
	return amount.Div(amount, orderTakerAmount)
}
