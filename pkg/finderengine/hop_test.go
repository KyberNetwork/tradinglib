package finderengine

import (
	"math/big"
	"testing"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
	"github.com/KyberNetwork/tradinglib/pkg/finderengine/entity"
	"github.com/stretchr/testify/assert"
)

type mockPool struct {
	address  string
	tokenIn  string
	tokenOut string
	rate     *big.Int
	subRate  *big.Int
	count    int
}

func (mp *mockPool) CalcAmountOut(params pool.CalcAmountOutParams) (*pool.CalcAmountOutResult, error) {
	out := new(big.Int).Div(new(big.Int).Mul(params.TokenAmountIn.Amount, mp.rate), big.NewInt(1))

	return &pool.CalcAmountOutResult{
		TokenAmountOut: &pool.TokenAmount{
			Token:  mp.tokenOut,
			Amount: out,
		},
		Fee: &pool.TokenAmount{
			Token:  mp.tokenIn,
			Amount: big.NewInt(1),
		},
	}, nil
}

func (mp *mockPool) UpdateBalance(params pool.UpdateBalanceParams) {
	mp.rate = new(big.Int).Sub(mp.rate, mp.subRate)
}

func (mp *mockPool) CloneState() pool.IPoolSimulator {
	return &mockPool{
		address:  mp.address,
		tokenIn:  mp.tokenIn,
		tokenOut: mp.tokenOut,
		rate:     new(big.Int).Set(mp.rate),
		subRate:  new(big.Int).Set(mp.subRate),
		count:    mp.count,
	}
}

func (mp *mockPool) CanSwapFrom(address string) []string {
	if address == mp.tokenIn {
		return []string{mp.tokenOut}
	}
	return nil
}
func (mp *mockPool) GetTokens() []string     { return []string{mp.tokenIn, mp.tokenOut} }
func (mp *mockPool) GetReserves() []*big.Int { return nil }
func (mp *mockPool) GetAddress() string      { return mp.address }
func (mp *mockPool) GetExchange() string {
	return ""
}
func (mp *mockPool) GetType() string                                  { return "" }
func (mp *mockPool) GetMetaInfo(tokenIn, tokenOut string) interface{} { return nil }
func (mp *mockPool) GetTokenIndex(address string) int                 { return 0 }
func (mp *mockPool) CalculateLimit() map[string]*big.Int              { return nil }
func (mp *mockPool) CanSwapTo(address string) []string                { return nil }

func Test_FindHops(t *testing.T) {
	pools := []pool.IPoolSimulator{
		&mockPool{address: "1", tokenIn: "A", tokenOut: "B", rate: big.NewInt(110), subRate: big.NewInt(10)},
		&mockPool{address: "2", tokenIn: "A", tokenOut: "B", rate: big.NewInt(100), subRate: big.NewInt(3)},
	}

	amountIn := big.NewInt(1000000)
	numSplits := uint64(6)
	hop := FindHops("A", 1, 18, "B", amountIn, pools, numSplits)
	assert.Len(t, hop.Splits, 2)
	expectedHop := &entity.Hop{
		TokenIn:       "A",
		TokenOut:      "B",
		AmountIn:      amountIn,
		AmountOut:     big.NewInt(98666636),
		GasUsed:       0,
		GasFeePrice:   0,
		L1GasFeePrice: 0,
		Fee:           big.NewInt(6),
		Splits: []*entity.HopSplit{
			{
				ID:            "1",
				AmountIn:      big.NewInt(333332),
				AmountOut:     big.NewInt(34999860),
				Fee:           big.NewInt(2),
				GasUsed:       0,
				GasFeePrice:   0,
				L1GasFeePrice: 0,
			},
			{
				ID:            "2",
				AmountIn:      big.NewInt(666668),
				AmountOut:     big.NewInt(63666776),
				Fee:           big.NewInt(4),
				GasUsed:       0,
				GasFeePrice:   0,
				L1GasFeePrice: 0,
			},
		},
	}

	expectedPools := []pool.IPoolSimulator{
		&mockPool{address: "1", tokenIn: "A", tokenOut: "B", rate: big.NewInt(110), subRate: big.NewInt(10)},
		&mockPool{address: "2", tokenIn: "A", tokenOut: "B", rate: big.NewInt(100), subRate: big.NewInt(3)},
	}

	assert.Equal(t, expectedHop, hop)
	assert.Equal(t, expectedPools, pools)
}
