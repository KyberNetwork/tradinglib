package util

import "math/big"

func CalcTakingAmount(swapMakerAmount, orderMakerAmount, orderTakerAmount *big.Int) *big.Int {
	amount := new(big.Int).Mul(swapMakerAmount, orderTakerAmount)
	amount.Add(amount, orderMakerAmount)
	amount.Sub(amount, big.NewInt(1))
	return amount.Div(amount, orderMakerAmount)
}

func CalcMakingAmount(swapTakerAmount, orderMakerAmount, orderTakerAmount *big.Int) *big.Int {
	amount := new(big.Int).Mul(swapTakerAmount, orderMakerAmount)
	return amount.Div(amount, orderTakerAmount)
}
