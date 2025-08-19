package finder

import (
	"math/big"

	dexlibPool "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
	"github.com/KyberNetwork/tradinglib/pkg/finderengine/entity"
	"github.com/KyberNetwork/tradinglib/pkg/finderengine/isolated"
	"github.com/KyberNetwork/tradinglib/pkg/finderengine/utils"
)

type Finder struct {
	FindHops        FindHopFunc
	MinThresholdUSD float64
}

func (f *Finder) Find(params entity.FinderParams) (*entity.BestRouteResult, error) {
	if err := f.validateParameters(params); err != nil {
		return nil, err
	}

	edges := make(map[string]map[string][]dexlibPool.IPoolSimulator)
	for i := range params.Pools {
		pool := params.Pools[i]
		tokens := pool.GetTokens()
		for i := range tokens {
			if edges[tokens[i]] == nil {
				edges[tokens[i]] = make(map[string][]dexlibPool.IPoolSimulator)
			}
			for j := range tokens {
				if i == j {
					continue
				}
				if edges[tokens[i]][tokens[j]] == nil {
					edges[tokens[i]][tokens[j]] = make([]dexlibPool.IPoolSimulator, 0)
				}
				edges[tokens[i]][tokens[j]] = append(edges[tokens[i]][tokens[j]], pool)
			}
		}
	}

	bestRoute := entity.Route{
		TokenIn:       params.TokenIn,
		TokenOut:      params.TargetToken,
		AmountIn:      new(big.Int).Set(params.AmountIn),
		AmountOut:     big.NewInt(0),
		GasUsed:       0,
		GasFeePrice:   0,
		L1GasFeePrice: 0,
		Paths:         nil,
	}

	minHops := f.minHopsToTokenOut(params.TokenIn, params.TargetToken, edges, params.WhitelistHopTokens, params.MaxHop)
	tokenInInfo := params.Tokens[params.TokenIn]
	tokenInPrice := params.Prices[params.TokenIn]
	splits := utils.SplitAmountThreshold(
		params.AmountIn, tokenInInfo.Decimals, params.NumPathSplits, tokenInPrice, f.MinThresholdUSD,
	)

	isolatedPools := make(map[string]*isolated.Pool, len(params.Pools))
	for address, pool := range params.Pools {
		isolatedPools[address] = isolated.NewIsolatedPool(pool)
	}
	for _, split := range splits {
		params.AmountIn = split
		bestPath := f.FindBestPathsOptimized(&params, minHops, edges)
		bestRoute.AmountOut.Add(bestPath.AmountOut, bestPath.AmountOut)
		bestRoute.Paths = append(bestRoute.Paths, bestPath)
		updatePoolState(bestPath, isolatedPools)
	}

	return &entity.BestRouteResult{
		AMMBestRoute: &bestRoute,
	}, nil
}

func (f *Finder) validateParameters(params entity.FinderParams) error {
	if _, exist := params.Tokens[params.TokenIn]; !exist {
		return ErrTokenInNotFound
	}
	if _, exist := params.Tokens[params.TargetToken]; !exist {
		return ErrTokenOutNotFound
	}

	if params.GasIncluded {
		if params.GasToken == "" {
			return ErrGasTokenRequired
		}
		if params.GasPrice == nil {
			return ErrGasPriceRequired
		}
		if _, exist := params.Tokens[params.GasToken]; !exist {
			return ErrGasTokenNotFound
		}
	}

	return nil
}
