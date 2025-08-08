package finderengine

import (
	dexlibPool "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
	"github.com/KyberNetwork/tradinglib/pkg/finderengine/entity"
)

type Finder struct {
	MaxHop              uint64
	DistributionPercent uint64
	NumPathSplits       uint64
	NumHopSplits        uint64
	findHops            FindHopFunc
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

	minHops := f.minHopsToTokenOut(params.TokenIn, params.TokenOut, edges, params.WhitelistHopTokens)
	_ = f.findBestPathsOptimized(&params, minHops, edges)
	// Optimize Route: TODO

	return nil, nil
}

func (f *Finder) validateParameters(params entity.FinderParams) error {
	if _, exist := params.Tokens[params.TokenIn]; !exist {
		return ErrTokenInNotFound
	}
	if _, exist := params.Tokens[params.TokenOut]; !exist {
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
