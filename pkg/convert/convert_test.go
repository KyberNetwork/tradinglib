package convert_test

import (
	"math"
	"math/big"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/convert"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExp10(t *testing.T) {
	tests := []struct {
		n      int64
		expect string
	}{
		{
			n:      3,
			expect: "1000",
		},
		{
			n:      18,
			expect: "1000000000000000000",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, convert.Exp10(test.n).String())
	}
}

func TestBPS(t *testing.T) {
	tests := []struct {
		amount *big.Int
		bps    int64
		expect *big.Int
	}{
		{
			amount: big.NewInt(100000),
			bps:    50,
			expect: big.NewInt(500),
		},
		{
			amount: big.NewInt(123456),
			bps:    10,
			expect: big.NewInt(123),
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, convert.BPS(test.amount, test.bps))
	}
}

func TestAddBPS(t *testing.T) {
	amount := big.NewInt(10000)
	add50bps := convert.AddBPS(amount, 50)
	require.Equal(t, big.NewInt(10050), add50bps)

	addn50bps := convert.AddBPS(amount, -50)
	require.Equal(t, big.NewInt(9950), addn50bps)
}

func TestWeiToFloat(t *testing.T) {
	amountHighPrec, _ := new(big.Int).SetString("123456789123456789123", 10)
	tests := []struct {
		amount   *big.Int
		decimals int64
		expect   float64
	}{
		{
			amount:   big.NewInt(123456789),
			decimals: 6,
			expect:   123.456789,
		},
		{
			amount:   amountHighPrec,
			decimals: 18,
			expect:   123.45678912345,
		},
	}

	for _, test := range tests {
		assert.True(t, floatEqual(test.expect, convert.WeiToFloat(test.amount, test.decimals)))
	}
}

func TestFloatToWei(t *testing.T) {
	tests := []struct {
		amount    float64
		decimals  int64
		expect    string
		expectErr error
	}{
		{
			amount:   12.3456789,
			decimals: 6,
			expect:   "12345678",
		},
		{
			amount:   1533.572643,
			decimals: 18,
			expect:   "1533572643000000000000",
		},
		{
			amount:    123.456,
			decimals:  0xfffffffff,
			expectErr: convert.ErrMaxExponent,
		},
	}

	for _, test := range tests {
		wei, err := convert.FloatToWei(test.amount, test.decimals)
		if assert.ErrorIs(t, err, test.expectErr) && err == nil {
			assert.Equal(t, test.expect, wei.String())
		}
	}
}

func TestMustFloatToWei(t *testing.T) {
	tests := []struct {
		amount   float64
		decimals int64
		expect   string
		isPanic  bool
	}{
		{
			amount:   12.3456789,
			decimals: 6,
			expect:   "12345678",
		},
		{
			amount:   1533.572643,
			decimals: 18,
			expect:   "1533572643000000000000",
		},
		{
			amount:   123.456,
			decimals: 0xfffffffff,
			isPanic:  true,
		},
	}

	for _, test := range tests {
		var wei *big.Int
		funcCanPanic := func() {
			wei = convert.MustFloatToWei(test.amount, test.decimals)
		}
		if test.isPanic {
			assert.Panics(t, funcCanPanic)
		} else if assert.NotPanics(t, funcCanPanic) {
			assert.Equal(t, test.expect, wei.String())
		}
	}
}

func TestIntToWei(t *testing.T) {
	tests := []struct {
		amount   int64
		decimals int32
		expect   string
	}{
		{
			amount:   12,
			decimals: 6,
			expect:   "12000000",
		},
		{
			amount:   12452,
			decimals: 18,
			expect:   "12452000000000000000000",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, convert.IntToWei(test.amount, test.decimals).String())
	}
}

func TestRound(t *testing.T) {
	tests := []struct {
		value    float64
		tickSize float64
		roundUp  bool
		expect   float64
	}{
		{
			value:    1234.5678,
			tickSize: 0.1,
			roundUp:  false,
			expect:   1234.5,
		},
		{
			value:    1234.5678,
			tickSize: 1,
			roundUp:  true,
			expect:   1235.0,
		},
		{
			value:    1234.5678,
			tickSize: 10,
			roundUp:  true,
			expect:   1240.0,
		},
		{
			value:    1234.5678,
			tickSize: 0.5,
			roundUp:  false,
			expect:   1234.5,
		},
		{
			value:    149.95,
			tickSize: 0.3,
			roundUp:  false,
			expect:   149.7,
		},
		{
			value:    44.985,
			tickSize: 0.3,
			roundUp:  false,
			expect:   44.7,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, convert.Round(test.value, test.tickSize, test.roundUp))
	}
}

func TestRoundUp(t *testing.T) {
	tests := []struct {
		value    float64
		tickSize float64
		expect   float64
	}{
		{
			value:    1234.5678,
			tickSize: 0.1,
			expect:   1234.6,
		},
		{
			value:    1234.5678,
			tickSize: 1,
			expect:   1235.0,
		},
		{
			value:    1234.5678,
			tickSize: 10,
			expect:   1240.0,
		},
		{
			value:    1234.5678,
			tickSize: 0.5,
			expect:   1235.0,
		},
		{
			value:    44.985,
			tickSize: 0.3,
			expect:   45,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, convert.RoundUp(test.value, test.tickSize))
	}
}

func TestRoundDown(t *testing.T) {
	tests := []struct {
		value    float64
		tickSize float64
		expect   float64
	}{
		{
			value:    1234.5678,
			tickSize: 0.1,
			expect:   1234.5,
		},
		{
			value:    1234.5678,
			tickSize: 1,
			expect:   1234.0,
		},
		{
			value:    1234.5678,
			tickSize: 10,
			expect:   1230.0,
		},
		{
			value:    1234.5678,
			tickSize: 0.5,
			expect:   1234.5,
		},
		{
			value:    44.985,
			tickSize: 0.3,
			expect:   44.7,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, convert.RoundDown(test.value, test.tickSize))
	}
}

func floatEqual(f1, f2 float64) bool {
	return math.Abs(f1-f2) < 1e-9
}
