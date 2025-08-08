package finderengine

import (
	"sort"
	"sync"

	dexlibPool "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
	"github.com/KyberNetwork/tradinglib/pkg/finderengine/entity"
	mapset "github.com/deckarep/golang-set/v2"
)

func (f *Finder) findBestPathsOptimized(
	params *entity.FinderParams,
	minHops map[string]uint64,
	edges map[string]map[string][]dexlibPool.IPoolSimulator,
) []*entity.Path {
	startNode := entity.NewPath(params.AmountIn)
	layer := map[string][]*entity.Path{
		params.TokenIn: {startNode},
	}

	for hop := uint64(0); hop < f.MaxHop; hop++ {
		layer = f.generateNextLayer(params, layer, minHops, hop, edges)
	}

	bestPaths := layer[params.TokenOut]
	sort.Slice(bestPaths, func(i, j int) bool {
		return bestPaths[i].AmountOut.Cmp(bestPaths[j].AmountOut) >= 0
	})
	return layer[params.TokenOut]
}

func (f *Finder) generateNextLayer(
	params *entity.FinderParams,
	currentLayer map[string][]*entity.Path,
	minHops map[string]uint64,
	currentHop uint64,
	edges map[string]map[string][]dexlibPool.IPoolSimulator,
) map[string][]*entity.Path {
	var (
		wg         sync.WaitGroup
		newPaths   sync.Map
		interation int
	)
	for tokenIn, paths := range currentLayer {
		tokenInEdges := edges[tokenIn]
		tokenInInfo := params.Tokens[tokenIn]
		tokenInPrice := params.Prices[tokenIn]
		for _, path := range paths {
			usedTokens := mapset.NewThreadUnsafeSet(path.TokenOrders...)
			for tokenOut := range tokenInEdges {
				if usedTokens.Contains(tokenOut) {
					continue
				}

				if _, isWhitelisted := params.WhitelistHopTokens[tokenOut]; !isWhitelisted && tokenOut != params.TokenOut {
					continue
				}

				remainingHopToTokenOut, exist := minHops[tokenOut]
				if !exist {
					continue
				}
				if currentHop+1+remainingHopToTokenOut > f.MaxHop {
					continue
				}

				go func(
					iteration int,
					path *entity.Path,
					pool []dexlibPool.IPoolSimulator,
					fromToken string,
					toToken string,
				) {
					defer wg.Done()
					hop := f.findHops(
						tokenInInfo.Address,
						tokenInPrice,
						tokenInInfo.Decimals,
						tokenOut,
						path.AmountOut,
						tokenInEdges[tokenOut],
						f.NumHopSplits,
					)

					nextPath := f.generateNextPath(params, path, hop)
					newPaths.Store(interation, nextPath)
				}(interation, path, tokenInEdges[tokenOut], tokenIn, tokenOut)

				interation++
			}
		}
	}

	nextLayer := make(map[string][]*entity.Path)
	for i := 0; i < interation; i++ {
		_nextPath, ok := newPaths.Load(i)
		if !ok || _nextPath == nil {
			continue
		}

		nextPath := _nextPath.(*entity.Path)
		lastToken := nextPath.TokenOrders[len(nextPath.TokenOrders)-1]
		nextLayer[lastToken] = append(nextLayer[lastToken], nextPath)
	}

	return nextLayer
}

func (f *Finder) generateNextPath(params *entity.FinderParams, currentPath *entity.Path, hop *entity.Hop) *entity.Path {
	nextPath := entity.NewPath(currentPath.AmountIn)
	nextPath.TokenOrders = make([]string, 0, len(currentPath.TokenOrders)+1)
	nextPath.HopOrders = make([]*entity.Hop, 0, len(currentPath.HopOrders)+1)
	for _, token := range currentPath.TokenOrders {
		nextPath.AddToken(token)
	}
	nextPath.AddToken(hop.TokenOut)

	for _, hop := range currentPath.HopOrders {
		nextPath.AddHop(hop)
	}
	nextPath.AddHop(hop)
	nextPath.SetAmountOutAndPrice(
		hop.AmountOut,
		params.Tokens[hop.TokenOut].Decimals,
		params.Prices[hop.TokenOut],
	)
	nextPath.SetGasUsedAndPrice(
		currentPath.GasUsed+hop.GasUsed,
		params.GasPrice,
		params.Tokens[params.GasToken].Decimals,
		params.Prices[params.GasToken],
		params.L1GasFeePricePerPool,
	)
	return nil
}
