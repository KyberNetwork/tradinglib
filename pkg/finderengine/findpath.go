package finderengine

import (
	dexlibPool "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
	"github.com/KyberNetwork/tradinglib/pkg/finderengine/entity"
)

func (f *Finder) findBestPathsOptimized(
	params *entity.FinderParams,
	minHops map[string]uint64,
	edges map[string]map[string][]dexlibPool.IPoolSimulator,
) *entity.Path {
	startNode := entity.NewPath(params.AmountIn)
	layer := map[string]*entity.Path{
		params.TokenIn: startNode,
	}

	for hop := uint64(0); hop < params.MaxHop; hop++ {
		newLayer := f.generateNextLayer(params, layer, minHops, hop, edges)
		if layer[params.TargetToken] != nil {
			if newLayer[params.TargetToken] == nil {
				newLayer[params.TargetToken] = layer[params.TargetToken]
			} else if newLayer[params.TargetToken].Cmp(layer[params.TargetToken], true) <= 0 {
				newLayer[params.TargetToken] = layer[params.TargetToken]
			}
		}

		layer = newLayer
	}

	return layer[params.TargetToken]
}

func (f *Finder) generateNextLayer(
	params *entity.FinderParams,
	currentLayer map[string]*entity.Path,
	minHops map[string]uint64,
	currentHop uint64,
	edges map[string]map[string][]dexlibPool.IPoolSimulator,
) map[string]*entity.Path {
	var newPaths []*entity.Path

	for tokenIn, path := range currentLayer {
		tokenInEdges := edges[tokenIn]
		tokenInInfo := params.Tokens[tokenIn]
		tokenInPrice := params.Prices[tokenIn]
		for tokenOut, pools := range tokenInEdges {
			if _, exists := params.WhitelistHopTokens[tokenOut]; tokenOut != params.TargetToken && !exists {
				continue
			}

			if _, exists := minHops[tokenOut]; !exists {
				continue
			}

			if currentHop+1+minHops[tokenOut] >= params.MaxHop {
				continue
			}

			hop := f.FindHops(tokenIn, tokenInPrice, tokenInInfo.Decimals, tokenOut, path.AmountOut, pools, params.NumHopSplits)
			newPath := f.generateNextPath(params, path, hop)
			newPaths = append(newPaths, newPath)
		}
	}

	nextLayer := make(map[string]*entity.Path)
	for _, path := range newPaths {
		lastToken := path.TokenOrders[len(path.TokenOrders)-1]
		if nextLayer[lastToken] == nil {
			nextLayer[lastToken] = path
			continue
		}
		if nextLayer[lastToken].Cmp(path, true) <= 0 {
			nextLayer[lastToken] = path
		}
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
	return nextPath
}

func updatePoolState(path *entity.Path, pools map[string]dexlibPool.IPoolSimulator) {
	for _, hop := range path.HopOrders {
		for _, hopSplit := range hop.Splits {
			pool := pools[hopSplit.ID]
			pool.UpdateBalance(dexlibPool.UpdateBalanceParams{
				TokenAmountIn: dexlibPool.TokenAmount{
					Token:  hop.TokenIn,
					Amount: hopSplit.AmountIn,
				},
				TokenAmountOut: dexlibPool.TokenAmount{
					Token:  hop.TokenOut,
					Amount: hopSplit.AmountOut,
				},
			})
		}
	}
}
