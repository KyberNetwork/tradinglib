//nolint:testpackage
package kipselipmm

import (
	"math/big"
	"testing"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
	"github.com/KyberNetwork/tradinglib/pkg/convert"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// nolint: funlen,maintidx
func TestPoolSimulator_for_single_pair(t *testing.T) {
	pricingData := PricingData{
		Tokens: map[string]Token{
			"ENS": {
				Symbol:         "ENS",
				Address:        "0xc18360217d8f7ab5e7c516566761ea12ce7f9d72",
				Decimals:       18,
				MinTradeAmount: 0.7025215869872952,
			},
			"USDT": {
				Symbol:         "USDT",
				Address:        "0xdac17f958d2ee523a2206206994597c13d831ec7",
				Decimals:       6,
				MinTradeAmount: 11.0,
			},
		},
		Pairs: []Pair{
			{
				Base:  "ENS",
				Quote: "USDT",
				Sell: []PriceLevel{
					{Rate: 15.652117523955797, Quantity: 125.21656894068956},
					{Rate: 15.662881770167372, Quantity: 375.623309584493},
					{Rate: 15.67439941823168, Quantity: 412.929739792515},
					{Rate: 15.6884233681878, Quantity: 244.732037911044361},
				},
				Buy: []PriceLevel{
					{Rate: 15.621030557013967, Quantity: 125.34444095357534},
					{Rate: 15.608694347724224, Quantity: 376.1931805840387},
					{Rate: 15.59745489881375, Quantity: 413.9014673459965},
					{Rate: 15.582109594358709, Quantity: 455.4711322720011},
					{Rate: 15.578273831762877, Quantity: 15.578273831762877},
					{Rate: 15.551558659824217, Quantity: 551.477130420391},
					{Rate: 15.541438433637838, Quantity: 607.0198635664137},
					{Rate: 15.537870376842847, Quantity: 667.8751830006263},
					{Rate: 15.535869442872722, Quantity: 376.56833189752433},
				},
			},
		},
	}

	priorityQuotes := []string{"USDT"}

	poolSimulator, err := NewPoolSimulator(pricingData, priorityQuotes)
	require.NoError(t, err)

	usdtToken := pricingData.Tokens["USDT"]
	ensToken := pricingData.Tokens["ENS"]
	pair := pricingData.Pairs[0]

	t.Run("buy ENS with USDT but not enough min trade amount", func(t *testing.T) {
		amountIn := convert.MustFloatToWei(0.5, ensToken.Decimals)
		_, err = poolSimulator.CalcAmountOut(pool.CalcAmountOutParams{
			TokenAmountIn: pool.TokenAmount{
				Token:  ensToken.Address,
				Amount: amountIn,
			},
			TokenOut: usdtToken.Address,
			Limit:    nil,
		})
		require.Error(t, err)

		assert.ErrorIs(t, err, ErrAmountInTooSmall)
	})

	t.Run("sell ENS to buy USDT but not enough min trade amount", func(t *testing.T) {
		amountOut := convert.MustFloatToWei(0.5, ensToken.Decimals) // less than min trade amount
		_, err = poolSimulator.CalcAmountIn(pool.CalcAmountInParams{
			TokenAmountOut: pool.TokenAmount{
				Token:  ensToken.Address,
				Amount: amountOut,
			},
			TokenIn: usdtToken.Address,
			Limit:   nil,
		})
		require.Error(t, err)

		assert.ErrorIs(t, err, ErrAmountOutTooSmall)
	})

	t.Run("buy ENS with USDT but exceed available liquidity", func(t *testing.T) {
		maxBuyEnsAmount := 0.0
		for _, pl := range pair.Buy {
			maxBuyEnsAmount += pl.Quantity
		}
		t.Logf("maxBuyEnsAmount: %f", maxBuyEnsAmount)

		amountIn := convert.MustFloatToWei(maxBuyEnsAmount+1, ensToken.Decimals)
		_, err = poolSimulator.CalcAmountOut(pool.CalcAmountOutParams{
			TokenAmountIn: pool.TokenAmount{
				Token:  ensToken.Address,
				Amount: amountIn,
			},
			TokenOut: usdtToken.Address,
			Limit:    nil,
		})
		require.Error(t, err)

		assert.ErrorIs(t, err, ErrInsufficientLiquidity)
	})

	t.Run("sell ENS to buy USDT but exceed available liquidity", func(t *testing.T) {
		maxSellEnsAmount := 0.0
		for _, pl := range pair.Sell {
			maxSellEnsAmount += pl.Quantity
		}
		t.Logf("maxSellEnsAmount: %f", maxSellEnsAmount)

		amountOut := convert.MustFloatToWei(maxSellEnsAmount+1, ensToken.Decimals)
		_, err = poolSimulator.CalcAmountIn(pool.CalcAmountInParams{
			TokenAmountOut: pool.TokenAmount{
				Token:  ensToken.Address,
				Amount: amountOut,
			},
			TokenIn: usdtToken.Address,
			Limit:   nil,
		})
		require.Error(t, err)

		assert.ErrorIs(t, err, ErrInsufficientLiquidity)
	})

	t.Run("buy ENS with USDT successfully", func(t *testing.T) {
		amountInF := 100.0
		amountIn := convert.MustFloatToWei(amountInF, ensToken.Decimals)
		amountOutF := 1562.103055 // first level of buy price levels

		result, err := poolSimulator.CalcAmountOut(pool.CalcAmountOutParams{
			TokenAmountIn: pool.TokenAmount{
				Token:  ensToken.Address,
				Amount: amountIn,
			},
			TokenOut: usdtToken.Address,
			Limit:    nil,
		})
		require.NoError(t, err)

		expectedSwapInfo := SwapExtra{
			TakerAsset:   ensToken.Address,
			TakingAmount: amountIn.String(),
			MakerAsset:   usdtToken.Address,
			MakingAmount: result.TokenAmountOut.Amount.String(),
			IsExactOut:   false,
			SwapLogs: []SwapLog{
				{
					Base:             ensToken.Symbol,
					Quote:            usdtToken.Symbol,
					AmountIn:         amountInF,
					AmountOut:        amountOutF,
					IsSell:           false,
					BoughtBaseAmount: amountInF,
				},
			},
		}

		assertSwapInfo(t, expectedSwapInfo, result.SwapInfo, ensToken, usdtToken)

		gotAmountOutF := convert.WeiToFloat(result.TokenAmountOut.Amount, usdtToken.Decimals)
		t.Logf("USDT amount: %f", gotAmountOutF)

		assert.InDelta(t, amountOutF, gotAmountOutF, epsilon)

		// now calculate back amount in should be close to original amount in
		resultIn, err := poolSimulator.CalcAmountIn(pool.CalcAmountInParams{
			TokenAmountOut: pool.TokenAmount{
				Token:  usdtToken.Address,
				Amount: result.TokenAmountOut.Amount,
			},
			TokenIn: ensToken.Address,
			Limit:   nil,
		})
		require.NoError(t, err)

		expectedSwapInfo.IsExactOut = true
		assertSwapInfo(t, expectedSwapInfo, resultIn.SwapInfo, ensToken, usdtToken)

		gotAmountInF := convert.WeiToFloat(resultIn.TokenAmountIn.Amount, ensToken.Decimals)
		t.Logf("ENS amount in (calculated back): %f", gotAmountInF)

		assert.InDelta(t, amountInF, gotAmountInF, epsilon)
	})

	t.Run("sell ENS to buy USDT successfully", func(t *testing.T) {
		amountOutF := 100.0
		amountOut := convert.MustFloatToWei(amountOutF, ensToken.Decimals)
		amountInF := 1565.211752 // first level of sell price levels

		result, err := poolSimulator.CalcAmountIn(pool.CalcAmountInParams{
			TokenAmountOut: pool.TokenAmount{
				Token:  ensToken.Address,
				Amount: amountOut,
			},
			TokenIn: usdtToken.Address,
			Limit:   nil,
		})
		require.NoError(t, err)

		expectedSwapInfo := SwapExtra{
			TakerAsset:   usdtToken.Address,
			TakingAmount: result.TokenAmountIn.Amount.String(),
			MakerAsset:   ensToken.Address,
			MakingAmount: amountOut.String(),
			IsExactOut:   true,
			SwapLogs: []SwapLog{
				{
					Base:           ensToken.Symbol,
					Quote:          usdtToken.Symbol,
					AmountIn:       amountInF,
					AmountOut:      amountOutF,
					IsSell:         true,
					SoldBaseAmount: amountOutF,
				},
			},
		}

		assertSwapInfo(t, expectedSwapInfo, result.SwapInfo, usdtToken, ensToken)

		gotAmountInF := convert.WeiToFloat(result.TokenAmountIn.Amount, usdtToken.Decimals)
		t.Logf("USDT amount in: %f", gotAmountInF)

		assert.InDelta(t, amountInF, gotAmountInF, epsilon)

		// now calculate back amount out should be close to original amount out
		resultOut, err := poolSimulator.CalcAmountOut(pool.CalcAmountOutParams{
			TokenAmountIn: pool.TokenAmount{
				Token:  usdtToken.Address,
				Amount: result.TokenAmountIn.Amount,
			},
			TokenOut: ensToken.Address,
			Limit:    nil,
		})
		require.NoError(t, err)

		expectedSwapInfo.IsExactOut = false
		assertSwapInfo(t, expectedSwapInfo, resultOut.SwapInfo, usdtToken, ensToken)

		gotAmountOutF := convert.WeiToFloat(resultOut.TokenAmountOut.Amount, ensToken.Decimals)
		t.Logf("ENS amount out (calculated back): %f", gotAmountOutF)

		assert.InDelta(t, amountOutF, gotAmountOutF, epsilon)
	})

	t.Run("update balance and clone state", func(t *testing.T) {
		// update sold amounts
		poolSimulator.UpdateBalance(pool.UpdateBalanceParams{
			SwapInfo: SwapExtra{
				SwapLogs: []SwapLog{
					{
						Base:             ensToken.Symbol,
						Quote:            usdtToken.Symbol,
						BoughtBaseAmount: 100.0,
					},
					{
						Base:           ensToken.Symbol,
						Quote:          usdtToken.Symbol,
						SoldBaseAmount: 50.0,
					},
				},
			},
		})

		state := poolSimulator.pairStates[ensToken.Symbol]
		assert.InDelta(t, 50.0, state.SoldAmount, epsilon)
		assert.InDelta(t, 100.0, state.BoughtAmount, epsilon)

		// clone state
		cloneSimulator, ok := poolSimulator.CloneState().(*PoolSimulator)
		require.True(t, ok)

		assert.Len(t, cloneSimulator.tokens, len(poolSimulator.tokens))
		assert.Len(t, cloneSimulator.tokenAddressesToTokenSymbols, len(cloneSimulator.tokenAddressesToTokenSymbols))
		assert.Len(t, cloneSimulator.pairs, len(poolSimulator.pairs))
		assert.Len(t, cloneSimulator.pairStates, len(poolSimulator.pairStates))
		assert.ElementsMatch(t, priorityQuotes, cloneSimulator.priorityQuotes)

		cloneState := cloneSimulator.pairStates[ensToken.Symbol]
		assert.InDelta(t, state.SoldAmount, cloneState.SoldAmount, epsilon)
		assert.InDelta(t, state.BoughtAmount, cloneState.BoughtAmount, epsilon)
	})
}

func assertSwapInfo(t *testing.T, expectedSwapExtra SwapExtra, swapInfo any, inToken, outToken Token) {
	t.Helper()

	actualSwapExtra, ok := swapInfo.(SwapExtra)
	require.True(t, ok)

	assert.Equal(t, expectedSwapExtra.TakerAsset, actualSwapExtra.TakerAsset)
	assert.Equal(t, expectedSwapExtra.MakerAsset, actualSwapExtra.MakerAsset)
	assert.InDelta(t,
		convert.WeiToFloat(mustBigIntFromString(t, expectedSwapExtra.TakingAmount), inToken.Decimals),
		convert.WeiToFloat(mustBigIntFromString(t, actualSwapExtra.TakingAmount), inToken.Decimals),
		epsilon)
	assert.InDelta(t,
		convert.WeiToFloat(mustBigIntFromString(t, expectedSwapExtra.MakingAmount), outToken.Decimals),
		convert.WeiToFloat(mustBigIntFromString(t, actualSwapExtra.MakingAmount), outToken.Decimals),
		epsilon)

	assert.Equal(t, expectedSwapExtra.IsExactOut, actualSwapExtra.IsExactOut)

	assertSwapLogs(t, expectedSwapExtra.SwapLogs, actualSwapExtra.SwapLogs)
}

func mustBigIntFromString(t *testing.T, s string) *big.Int {
	t.Helper()

	v, ok := new(big.Int).SetString(s, 10)
	require.True(t, ok, "invalid big integer number: %s", s)

	return v
}
