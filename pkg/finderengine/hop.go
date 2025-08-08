package finderengine

import (
	"math/big"

	dexlibPool "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
	"github.com/KyberNetwork/tradinglib/pkg/finderengine/entity"
	"github.com/oleiade/lane/v2"
)

type FindHopFunc func(
	tokenIn string,
	tokenInPrice float64,
	tokenInDecimals uint8,
	tokenOut string,
	amountIn *big.Int,
	pools []dexlibPool.IPoolSimulator,
	numSplits uint64,
) *entity.Hop

type PoolHeap struct {
	ID              uint64
	Pool            string
	AmountIn        *big.Int
	AmountOutResult *dexlibPool.CalcAmountOutResult
}

func FindHops(
	tokenIn string,
	tokenInPrice float64,
	tokenInDecimals uint8,
	tokenOut string,
	amountIn *big.Int,
	pools []dexlibPool.IPoolSimulator,
	numSplits uint64,
) *entity.Hop {
	splits := splitAmount(amountIn, numSplits)
	baseSplit := splits[0]
	baseCalcParams := dexlibPool.CalcAmountOutParams{
		TokenAmountIn: dexlibPool.TokenAmount{Token: tokenIn, Amount: baseSplit},
		TokenOut:      tokenOut,
	}

	maxHeap := New(func(x, y *PoolHeap) bool {
		return x.AmountOutResult.TokenAmountOut.Amount.Cmp(y.AmountOutResult.TokenAmountOut.Amount) > 0
	})

	for id, pool := range pools {
		// Implement parallel
		if result, err := pool.CalcAmountOut(baseCalcParams); err == nil {
			maxHeap.Push(&PoolHeap{
				ID:              uint64(id),
				Pool:            pool.GetAddress(),
				AmountIn:        baseSplit,
				AmountOutResult: result,
			})
		}
	}

	hopSplitMap := make(map[string]*entity.HopSplit, len(pools))

	for i := uint64(0); i < numSplits && maxHeap.Len() > 0; i++ {
		chunk := splits[i]
		isLast := i == numSplits-1

		best, _ := maxHeap.Pop()
		pool := pools[best.ID]

		if isLast {
			lastChunk := splits[len(splits)-1]
			if result, err := pool.CalcAmountOut(dexlibPool.CalcAmountOutParams{
				TokenAmountIn: dexlibPool.TokenAmount{Token: tokenIn, Amount: lastChunk},
				TokenOut:      tokenOut,
			}); err == nil {
				best.AmountIn = lastChunk
				best.AmountOutResult = result
			}
		}

		split := hopSplitMap[best.Pool]
		if split == nil {
			split = &entity.HopSplit{
				ID:        best.Pool,
				AmountIn:  big.NewInt(0),
				AmountOut: big.NewInt(0),
			}
			hopSplitMap[best.Pool] = split
		}

		split.AmountIn.Add(split.AmountIn, chunk)
		split.AmountOut.Add(split.AmountOut, best.AmountOutResult.TokenAmountOut.Amount)

		pool.UpdateBalance(dexlibPool.UpdateBalanceParams{
			TokenAmountIn:  dexlibPool.TokenAmount{Token: tokenIn, Amount: chunk},
			TokenAmountOut: *best.AmountOutResult.TokenAmountOut,
			Fee:            *best.AmountOutResult.Fee,
		})

		if !isLast {
			if result, err := pool.CalcAmountOut(baseCalcParams); err == nil {
				maxHeap.Push(&PoolHeap{
					ID:              best.ID,
					Pool:            best.Pool,
					AmountIn:        baseCalcParams.TokenAmountIn.Amount,
					AmountOutResult: result,
				})
			}
		}
	}

	splitsOut := make([]*entity.HopSplit, 0, len(hopSplitMap))
	totalAmountIn := big.NewInt(0)
	totalAmountOut := big.NewInt(0)
	for _, s := range hopSplitMap {
		splitsOut = append(splitsOut, s)
		totalAmountIn.Add(totalAmountIn, s.AmountIn)
		totalAmountOut.Add(totalAmountOut, s.AmountOut)
	}

	return &entity.Hop{
		TokenIn:   tokenIn,
		TokenOut:  tokenOut,
		AmountIn:  totalAmountIn,
		AmountOut: totalAmountOut,
		Splits:    splitsOut,
	}
}

func (f *Finder) minHopsToTokenOut(
	tokenIn string,
	tokenOut string,
	edges map[string]map[string][]dexlibPool.IPoolSimulator,
	whitelistedHopTokens map[string]struct{},
) map[string]uint64 {
	minHops := make(map[string]uint64)
	queue := lane.NewQueue[string]()

	minHops[tokenOut] = 0
	queue.Enqueue(tokenOut)

	for queue.Size() > 0 {
		token, _ := queue.Dequeue()
		if minHops[token] == f.MaxHop {
			continue
		}

		if edges[token] == nil {
			continue
		}

		for tokenFrom := range edges[token] {
			if _, visited := minHops[tokenFrom]; visited {
				continue
			}

			_, isWhitelisted := whitelistedHopTokens[tokenFrom]
			isHopToken := tokenFrom != tokenIn
			if isHopToken && !isWhitelisted {
				continue
			}

			minHops[tokenFrom] = minHops[token] + 1
			queue.Enqueue(tokenFrom)
		}
	}

	return minHops
}
