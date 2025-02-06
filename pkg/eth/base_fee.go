package eth

import (
	"math/big"
)

// [STEAL] https://github.com/KyberNetwork/1inch-resolver/blob/v0.2.13/internal/gas/tracker.go#L157
func CalcNextBaseFee(targetGas, baseFeeChangeDenominator, gasUsed uint64, baseFee *big.Int) *big.Int {
	if gasUsed == targetGas {
		return new(big.Int).Set(baseFee)
	}

	prodTargetGasAndDenominator := new(big.Int).SetUint64(targetGas * baseFeeChangeDenominator)

	if gasUsed > targetGas {
		// delta = baseFee * (gasUsed - target) / (target * baseFeeChangeDenominator)
		delta := new(big.Int).SetUint64(gasUsed - targetGas)
		delta.Mul(baseFee, delta)
		delta.Div(delta, prodTargetGasAndDenominator)
		return delta.Add(baseFee, delta)
	}

	// delta = baseFee * (target - gasUsed) / (target * baseFeeChangeDenominator)
	delta := new(big.Int).SetUint64(targetGas - gasUsed)
	delta.Mul(baseFee, delta)
	delta.Div(delta, prodTargetGasAndDenominator)

	return delta.Sub(baseFee, delta)
}
