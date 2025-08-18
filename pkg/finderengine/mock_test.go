package finderengine_test

import (
	"errors"
	"math/big"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
)

type mockPool struct {
	address  string
	tokenIn  string
	tokenOut string
	bids     []Order
	asks     []Order
}

type Order struct {
	A *big.Int
	R *big.Int
}

func (mp *mockPool) CalcAmountOut(params pool.CalcAmountOutParams) (*pool.CalcAmountOutResult, error) {
	if params.TokenAmountIn.Amount == nil {
		return nil, errors.New("invalid amount")
	}

	remaining := new(big.Int).Set(params.TokenAmountIn.Amount)
	amountOut := big.NewInt(0)

	var book []Order
	if params.TokenAmountIn.Token == mp.tokenIn {
		book = mp.asks
	} else {
		book = mp.bids
	}

	for _, o := range book {
		if remaining.Sign() == 0 {
			break
		}

		if remaining.Cmp(o.A) >= 0 {
			cost := new(big.Int).Mul(o.A, o.R)
			cost.Div(cost, big.NewInt(100))
			amountOut.Add(amountOut, cost)
			remaining.Sub(remaining, o.A)
		} else {
			cost := new(big.Int).Mul(remaining, o.R)
			cost.Div(cost, big.NewInt(100))
			amountOut.Add(amountOut, cost)
			remaining.SetInt64(0)
		}
	}

	return &pool.CalcAmountOutResult{
		TokenAmountOut: &pool.TokenAmount{
			Token:  params.TokenOut,
			Amount: amountOut,
		},
		Fee: &pool.TokenAmount{
			Amount: big.NewInt(0),
		},
	}, nil
}

func (mp *mockPool) UpdateBalance(params pool.UpdateBalanceParams) {
	var book *[]Order
	switch inToken := params.TokenAmountIn.Token; inToken {
	case mp.tokenIn:
		book = &mp.asks
	case mp.tokenOut:
		book = &mp.bids
	default:
		book = &mp.bids
	}

	remaining := new(big.Int).Set(params.TokenAmountIn.Amount)
	newBook := make([]Order, 0, len(*book))
	for _, lv := range *book {
		if lv.A == nil {
			continue
		}
		switch lv.A.Cmp(remaining) {
		case 1:
			lv.A = new(big.Int).Sub(lv.A, remaining)
			remaining.SetInt64(0)
			newBook = append(newBook, lv)
		case 0:
			remaining.SetInt64(0)
		case -1:
			remaining.Sub(remaining, lv.A)
		}
	}
	*book = newBook
}

func (mp *mockPool) CloneState() pool.IPoolSimulator {
	newBids := make([]Order, 0, len(mp.bids))
	newAsks := make([]Order, 0, len(mp.asks))

	for i := range mp.bids {
		newBids = append(newBids, Order{
			A: new(big.Int).Set(mp.bids[i].A),
			R: new(big.Int).Set(mp.bids[i].R),
		})
	}

	for i := range mp.asks {
		newAsks = append(newAsks, Order{
			A: new(big.Int).Set(mp.asks[i].A),
			R: new(big.Int).Set(mp.asks[i].R),
		})
	}
	return &mockPool{
		address:  mp.address,
		tokenIn:  mp.tokenIn,
		tokenOut: mp.tokenOut,
		bids:     newBids,
		asks:     newAsks,
	}
}

func (mp *mockPool) CanSwapFrom(address string) []string {
	if address == mp.tokenIn {
		return []string{mp.tokenOut}
	}
	return nil
}
func (mp *mockPool) GetTokens() []string { return []string{mp.tokenIn, mp.tokenOut} }
func (mp *mockPool) GetReserves() []*big.Int {
	// for i := range mp.asks {
	// 	fmt.Println(mp.asks[i])
	// }

	return nil
}
func (mp *mockPool) GetAddress() string { return mp.address }
func (mp *mockPool) GetExchange() string {
	return ""
}
func (mp *mockPool) GetType() string                                  { return "" }
func (mp *mockPool) GetMetaInfo(tokenIn, tokenOut string) interface{} { return nil }
func (mp *mockPool) GetTokenIndex(address string) int                 { return 0 }
func (mp *mockPool) CalculateLimit() map[string]*big.Int              { return nil }
func (mp *mockPool) CanSwapTo(address string) []string                { return nil }
