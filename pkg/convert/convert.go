package convert

import (
	"errors"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

var ErrMaxExponent = errors.New("reach max exponent")

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
func WeiToFloat(amount *big.Int, decimals int64) (float64, big.Accuracy) {
	amountFloat := big.NewFloat(0).SetInt(amount)
	amountFloat.Quo(amountFloat, big.NewFloat(0).SetInt(Exp10(decimals)))
	output, acc := amountFloat.Float64()
	return output, acc
}

// FloatToWei ...
func FloatToWei(amount float64, decimals int64) (*big.Int, error) {
	if decimals > math.MaxInt32 {
		return nil, ErrMaxExponent
	}
	d := decimal.NewFromFloatWithExponent(amount, int32(decimals))
	return d.BigInt(), nil
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

func RoundUp(value float64, tickSize float64) float64 {
	places := int32(math.Abs(math.Round(math.Log10(tickSize))))
	v := decimal.NewFromFloat(value)
	rec := v.Round(places)
	if rec.LessThan(v) {
		rec = rec.Add(decimal.NewFromFloat(tickSize))
	}
	r, _ := rec.Float64()
	return r
}

func RoundDown(value float64, tickSize float64) float64 {
	places := int32(math.Abs(math.Round(math.Log10(tickSize))))
	v := decimal.NewFromFloat(value)
	rec := v.Round(places)
	if rec.GreaterThan(v) {
		rec = rec.Sub(decimal.NewFromFloat(tickSize))
		if rec.IsNegative() {
			rec = decimal.NewFromInt(0)
		}
	}
	r, _ := rec.Float64()
	return r
}
