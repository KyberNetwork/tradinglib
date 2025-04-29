package eth

import (
	"math/big"
)

const (
	defaultTargetDenominator        = 2
	defaultBaseFeeChangeDenominator = 8
)

// [STEAL] https://github.com/KyberNetwork/1inch-resolver/blob/4bac2bf061c311ad297967dafc7b881a9191fba9/internal/gas/tracker.go#L160C21-L179C2
func CalcNextBaseFee(parentBaseFee *big.Int, gasUsed uint64, gasLimit uint64) *big.Int {
	defaultTargetGas := gasLimit / defaultTargetDenominator
	if gasUsed == defaultTargetGas {
		return parentBaseFee
	}

	if gasUsed > defaultTargetGas {
		// delta = baseFee * (gasUsed - target) / (target * baseFeeChangeDenominator)
		delta := new(big.Int).SetUint64(gasUsed - defaultTargetGas)
		delta.Mul(parentBaseFee, delta)
		delta.Div(delta, new(big.Int).SetUint64(defaultTargetGas*defaultBaseFeeChangeDenominator))
		return delta.Add(parentBaseFee, delta)
	}

	// delta = baseFee * (target - gasUsed) / (target * baseFeeChangeDenominator)
	delta := new(big.Int).SetUint64(defaultTargetGas - gasUsed)
	delta.Mul(parentBaseFee, delta)
	delta.Div(delta, new(big.Int).SetUint64(defaultTargetGas*defaultBaseFeeChangeDenominator))
	return delta.Sub(parentBaseFee, delta)
}
