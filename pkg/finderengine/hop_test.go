package finderengine

import (
	"math/big"
	"testing"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
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
	// simulate decreasing rate after each call
	mp.count++

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

func (mp *mockPool) CloneState() pool.IPoolSimulator { return mp }

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
		&mockPool{address: "A", tokenIn: "A", tokenOut: "B", rate: big.NewInt(110), subRate: big.NewInt(10)},
		&mockPool{address: "B", tokenIn: "A", tokenOut: "B", rate: big.NewInt(100), subRate: big.NewInt(3)},
	}

	amountIn := big.NewInt(1000000)
	numSplits := uint64(6)
	hop := FindHops("A", 1, 18, "B", amountIn, pools, numSplits)
	// Assert each pool got used
	assert.Len(t, hop.Splits, 2)
}
