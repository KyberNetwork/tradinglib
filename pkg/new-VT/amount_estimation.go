package newvt

import (
	"errors"
	"math"
	"strconv"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/entity"
)

const esp = 1e-9

type AmmImpactResult struct {
	TokenIn   string
	TokenOut  string
	AmountIn  float64
	AmountOut float64
}

// spread must be > 0
func EstimateAmountFromSpreadAMM(pool entity.Pool, tokenTarget string, spread float64) (*AmmImpactResult, error) {
	if spread <= 0 {
		return nil, errors.New("spread must be > 0")
	}

	if len(pool.Tokens) != 2 || len(pool.Reserves) != 2 {
		return nil, errors.New("pool must contain exactly 2 tokens & 2 reserves")
	}

	// Determine which token is target (tokenOut)
	var outIdx, inIdx int
	if pool.Tokens[0].Address == tokenTarget {
		outIdx = 0
		inIdx = 1
	} else if pool.Tokens[1].Address == tokenTarget {
		outIdx = 1
		inIdx = 0
	} else {
		return nil, errors.New("target token not found in pool")
	}

	tokenOut := pool.Tokens[outIdx].Address
	tokenIn := pool.Tokens[inIdx].Address

	// Parse reserves
	reserveTokenOut, err := strconv.ParseFloat(pool.Reserves[outIdx], 64)
	if err != nil {
		return nil, err
	}
	reserveTokenIn, err := strconv.ParseFloat(pool.Reserves[inIdx], 64)
	if err != nil {
		return nil, err
	}

	x := reserveTokenOut // tokenOut
	y := reserveTokenIn  // tokenIn

	// Swap fee
	fee := pool.SwapFee
	if fee < 0 || fee >= 1 {
		return nil, errors.New("invalid pool fee")
	}
	feeFactor := 1 - fee

	// ---- Derivation based on AMM math ----
	// Target new price:
	// P1 = (1+spread) * (y/x)
	//
	// New X reserve:
	x1 := x / math.Sqrt(1+spread)
	if x1 <= 0 {
		return nil, errors.New("invalid computed x1")
	}

	// k = x*y
	k := x * y

	// new Y reserve:
	y1 := k / x1

	// Δy = amountInAfterFee
	deltaY := y1 - y
	if deltaY <= 0 {
		return nil, errors.New("deltaY <= 0, no swap needed")
	}

	// actual amountIn before fee applied
	amountIn := deltaY / feeFactor

	// compute amountOut = X taken from pool
	amountOut := x - x1

	return &AmmImpactResult{
		TokenIn:   tokenIn,
		TokenOut:  tokenOut,
		AmountIn:  amountIn,
		AmountOut: amountOut,
	}, nil
}

func EstimateAmountFromSpreadViaBS(targetSpread, maxAmount float64, token string) (float64, error) {
	var (
		l              = esp
		r              = maxAmount
		amountExpected = 0.0
	)
	for (r - l) > esp {
		mid := (l + r) / 2
		spread, err := calculateSpread(mid, token)
		if err != nil {
			return 0, err
		}
		if spread >= targetSpread {
			amountExpected = mid
			r = mid
		} else {
			l = mid
		}
	}

	return amountExpected, nil
}

func calculateSpread(amountIn float64, token string) (float64, error) {
	return 0, nil
}
