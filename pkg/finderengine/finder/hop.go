package finder

import (
	"container/heap"
	"fmt"
	"math/big"

	dexlibPool "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
	"github.com/KyberNetwork/tradinglib/pkg/finderengine/entity"
	"github.com/KyberNetwork/tradinglib/pkg/finderengine/isolated"
	"github.com/KyberNetwork/tradinglib/pkg/finderengine/utils"
	"github.com/oleiade/lane/v2"
)

const maxHopWorker = 8

type FindHopFunc func(
	tokenIn string,
	tokenInPrice float64,
	tokenInDecimals uint8,
	tokenOut string,
	amountIn *big.Int,
	pools []*isolated.Pool,
	numSplits uint64,
	minThresholdUSD float64,
) *entity.Hop

type poolItem struct {
	id     uint64
	addr   string
	amtOut *big.Int
	fee    *big.Int
	gas    int64
}

type poolMaxHeap []poolItem

func (h *poolMaxHeap) Len() int {
	return len(*h)
}
func (h *poolMaxHeap) Less(i, j int) bool { return (*h)[i].amtOut.Cmp((*h)[j].amtOut) > 0 }
func (h *poolMaxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *poolMaxHeap) Push(x any) {
	it, ok := x.(poolItem)
	if !ok {
		panic(fmt.Sprintf("poolMaxHeap: Push got %T, want poolItem", x))
	}
	*h = append(*h, it)
}

func (h *poolMaxHeap) Pop() any {
	old := *h
	n := len(old)
	it := old[n-1]
	*h = old[:n-1]
	return it
}

func FindHops(
	tokenIn string,
	tokenInPrice float64,
	tokenInDecimals uint8,
	tokenOut string,
	amountIn *big.Int,
	pools []*isolated.Pool,
	numSplits uint64,
	minThresholdUSD float64,
) *entity.Hop {
	if len(pools) == 0 || amountIn.Sign() <= 0 || numSplits == 0 {
		return &entity.Hop{
			TokenIn:   tokenIn,
			TokenOut:  tokenOut,
			AmountIn:  amountIn,
			AmountOut: big.NewInt(0),
			Fee:       big.NewInt(0),
			Splits:    nil,
		}
	}

	splits := utils.SplitAmountThreshold(amountIn, tokenInDecimals, numSplits, minThresholdUSD, tokenInPrice)
	baseSplit := splits[0]

	baseParams := dexlibPool.CalcAmountOutParams{
		TokenAmountIn: dexlibPool.TokenAmount{Token: tokenIn, Amount: baseSplit},
		TokenOut:      tokenOut,
	}

	type input struct{ idx int }
	type out struct {
		idx int
		res *dexlibPool.CalcAmountOutResult
		ok  bool
	}

	n := len(pools)
	data := make(chan input, n)
	outs := make(chan out, n)

	for w := 0; w < maxHopWorker; w++ {
		go func(data <-chan input, results chan<- out) {
			for d := range data {
				res, err := pools[d.idx].CalcAmountOut(baseParams)
				if err != nil || res == nil || res.TokenAmountOut == nil || res.TokenAmountOut.Amount == nil {
					results <- out{idx: d.idx, ok: false}
					continue
				}
				results <- out{idx: d.idx, res: res, ok: true}
			}
		}(data, outs)
	}

	for i := 0; i < n; i++ {
		data <- input{idx: i}
	}
	close(data)

	h := make(poolMaxHeap, 0, n)
	for i := 0; i < n; i++ {
		o := <-outs
		if !o.ok {
			continue
		}
		addr := pools[o.idx].GetAddress()
		h = append(h, poolItem{
			id:     uint64(o.idx),
			addr:   addr,
			amtOut: new(big.Int).Set(o.res.TokenAmountOut.Amount),
			fee:    new(big.Int).Set(o.res.Fee.Amount),
			gas:    o.res.Gas,
		})
	}

	if len(h) == 0 {
		return &entity.Hop{
			TokenIn:   tokenIn,
			TokenOut:  tokenOut,
			AmountIn:  amountIn,
			AmountOut: big.NewInt(0),
			Fee:       big.NewInt(0),
			GasUsed:   0,
			Splits:    nil,
		}
	}
	heap.Init(&h)

	hopSplitMap := make(map[string]*entity.HopSplit, len(h))
	totalIn := big.NewInt(0)
	totalOut := big.NewInt(0)
	totalFee := big.NewInt(0)
	totalGas := int64(0)

	for i := 0; i < len(splits) && h.Len() > 0; i++ {
		chunk := splits[i]
		isLast := i == len(splits)-1
		best, _ := heap.Pop(&h).(poolItem)
		p := pools[best.id]

		var item poolItem
		if isLast && chunk.Cmp(baseSplit) != 0 {
			r, err := p.CalcAmountOut(dexlibPool.CalcAmountOutParams{
				TokenAmountIn: dexlibPool.TokenAmount{Token: tokenIn, Amount: chunk},
				TokenOut:      tokenOut,
			})
			if err == nil && r != nil {
				item = poolItem{addr: p.GetAddress(), amtOut: r.TokenAmountOut.Amount, gas: r.Gas, fee: r.Fee.Amount}
			} else {
				item = poolItem{amtOut: big.NewInt(0), gas: 0, fee: big.NewInt(0)}
			}
		} else {
			item = best
		}

		s := hopSplitMap[best.addr]
		if s == nil {
			s = &entity.HopSplit{
				ID:        best.addr,
				AmountIn:  big.NewInt(0),
				AmountOut: big.NewInt(0),
				Fee:       big.NewInt(0),
			}
			hopSplitMap[item.addr] = s
		}
		s.AmountIn.Add(s.AmountIn, chunk)
		s.AmountOut.Add(s.AmountOut, item.amtOut)
		s.Fee.Add(s.Fee, item.fee)

		totalIn.Add(totalIn, chunk)
		totalOut.Add(totalOut, item.amtOut)
		totalFee.Add(totalFee, item.fee)
		totalGas += item.gas

		p.UpdateBalance(dexlibPool.UpdateBalanceParams{
			TokenAmountIn:  dexlibPool.TokenAmount{Token: tokenIn, Amount: chunk},
			TokenAmountOut: dexlibPool.TokenAmount{Token: tokenOut, Amount: item.amtOut},
			Fee:            dexlibPool.TokenAmount{Token: tokenIn, Amount: item.fee},
		})

		if !isLast {
			newRes, err := p.CalcAmountOut(baseParams)
			if err == nil && newRes != nil && newRes.TokenAmountOut != nil && newRes.TokenAmountOut.Amount != nil {
				best.amtOut = new(big.Int).Set(newRes.TokenAmountOut.Amount)
				best.fee = new(big.Int).Set(newRes.Fee.Amount)
				best.gas = newRes.Gas
				heap.Push(&h, best)
			}
		}
	}

	splitsOut := make([]entity.HopSplit, 0, len(hopSplitMap))
	for _, s := range hopSplitMap {
		splitsOut = append(splitsOut, *s)
	}
	return &entity.Hop{
		TokenIn:   tokenIn,
		TokenOut:  tokenOut,
		Fee:       totalFee,
		AmountIn:  totalIn,
		AmountOut: totalOut,
		Splits:    splitsOut,
	}
}

func (f *Finder) minHopsToTokenOut(
	tokenIn string,
	tokenOut string,
	edges map[string]map[string][]dexlibPool.IPoolSimulator,
	whitelistedHopTokens map[string]struct{},
	maxHop uint64,
) map[string]uint64 {
	minHops := make(map[string]uint64)
	queue := lane.NewQueue[string]()

	minHops[tokenOut] = 0
	queue.Enqueue(tokenOut)

	for queue.Size() > 0 {
		token, _ := queue.Dequeue()
		if minHops[token] == maxHop {
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
