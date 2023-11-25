package convert

import (
	"errors"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

var (
	ErrMaxExponent   = errors.New("reach max exponent")
	ErrInvalidNumber = errors.New("number is not valid")
)

const (
	maxBPS = 10000
)

// Exp10 ...
func Exp10(n int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(10), big.NewInt(n), nil) // nolint: gomnd
}

func BPS(amount *big.Int, bps int64) *big.Int {
	if bps == 0 {
		return common.Big0
	}
	bpsInt := big.NewInt(bps)
	return new(big.Int).Quo(new(big.Int).Mul(amount, bpsInt), big.NewInt(maxBPS))
}

func AddBPS(amount *big.Int, bps int64) *big.Int {
	if bps == 0 {
		return amount
	}
	diff := BPS(amount, bps)
	newAmount := new(big.Int).Add(amount, diff)
	return newAmount
}

// WeiToFloat ..
func WeiToFloat(amount *big.Int, decimals int64) float64 {
	amountFloat := big.NewFloat(0).SetInt(amount)
	amountFloat.Quo(amountFloat, big.NewFloat(0).SetInt(Exp10(decimals)))
	output, _ := amountFloat.Float64()
	return output
}

// FloatToWei ...
func FloatToWei(amount float64, decimals int64) (*big.Int, error) {
	if math.IsNaN(amount) || math.IsInf(amount, 0) {
		return nil, fmt.Errorf("%w: %f", ErrInvalidNumber, amount)
	}

	if decimals > math.MaxInt32 {
		return nil, ErrMaxExponent
	}
	d := decimal.NewFromFloat(amount)
	expo := decimal.New(1, int32(decimals))
	return d.Mul(expo).BigInt(), nil
}

// MustFloatToWei same as FloatToWei but will panic if decimals > maxInt32.
func MustFloatToWei(amount float64, decimals int64) *big.Int {
	d, err := FloatToWei(amount, decimals)
	if err != nil {
		panic(err)
	}
	return d
}

// IntToWei ...
func IntToWei(amount int64, decimals int32) *big.Int {
	return decimal.New(amount, decimals).BigInt()
}

// Round rounds `value` up or down 1 `tickSize`.
func Round(value float64, tickSize float64, roundUp bool) float64 {
	tickSizeD := decimal.NewFromFloat(tickSize)
	valueD := decimal.NewFromFloat(value)

	valueD = valueD.Div(tickSizeD)

	if roundUp {
		valueD = valueD.Ceil()
	} else {
		valueD = valueD.Floor()
	}

	return valueD.Mul(tickSizeD).InexactFloat64()
}

// RoundUp rounds `value` up 1 `tickSize`.
func RoundUp(value float64, tickSize float64) float64 {
	return Round(value, tickSize, true)
}

// RoundDown rounds `value` down 1 `tickSize`.
func RoundDown(value float64, tickSize float64) float64 {
	return Round(value, tickSize, false)
}
