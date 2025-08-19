package finder_test

import (
	"fmt"
	"math/big"
	"testing"

	dexlibPool "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
	"github.com/KyberNetwork/tradinglib/pkg/finderengine/entity"
	"github.com/KyberNetwork/tradinglib/pkg/finderengine/finder"
	"github.com/KyberNetwork/tradinglib/pkg/finderengine/isolated"
)

func GenTest() (map[string]entity.SimplifiedToken, map[string]struct{}, map[string]map[string][]dexlibPool.IPoolSimulator) {
	tokens := make(map[string]entity.SimplifiedToken)
	for i := 0; i < 1000; i++ {
		addr := fmt.Sprintf("token%d", i)
		tokens[addr] = entity.SimplifiedToken{Address: addr, Decimals: 18}
	}
	whitelist := make(map[string]struct{})
	for i := 0; i < 50; i++ {
		addr := fmt.Sprintf("token%d", i)
		whitelist[addr] = struct{}{}
	}

	edges := make(map[string]map[string][]dexlibPool.IPoolSimulator)
	for i := 0; i < 1000; i++ {
		from := fmt.Sprintf("token%d", i)
		edges[from] = make(map[string][]dexlibPool.IPoolSimulator)
		for j := 0; j < 5; j++ {
			to := fmt.Sprintf("token%d", (i+j+1)%1000)
			edges[from][to] = []dexlibPool.IPoolSimulator{&mockPool{}}
		}
	}
	return tokens, whitelist, edges
}

func BenchmarkFindBestPathsOptimized(b *testing.B) {
	tokens, whitelist, edges := GenTest()
	params := &entity.FinderParams{
		MaxHop:             5,
		NumHopSplits:       2,
		NumPathSplits:      2,
		TokenIn:            "token0",
		TargetToken:        "token999",
		AmountIn:           big.NewInt(1_000_000_000_000_000_000),
		GasPrice:           big.NewInt(0),
		WhitelistHopTokens: whitelist,
		Tokens:             tokens,
		GasIncluded:        false,
	}
	finder := &finder.Finder{
		FindHops: func(tokenIn string, tokenInPrice float64, tokenInDecimals uint8, tokenOut string, amountIn *big.Int, pools []*isolated.Pool, numSplits uint64, minThresholdUSD float64) *entity.Hop {
			return &entity.Hop{
				TokenIn:   tokenIn,
				TokenOut:  tokenOut,
				AmountIn:  amountIn,
				AmountOut: new(big.Int).Add(amountIn, big.NewInt(1)),
				Fee:       big.NewInt(0),
				Splits: []entity.HopSplit{
					{ID: fmt.Sprintf("MockPool-%s-%s", tokenIn, tokenOut)},
				},
			}
		},
	}

	// Dummy minHops
	minHops := make(map[string]uint64)
	for k := range tokens {
		minHops[k] = 1
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = finder.FindBestPathsOptimized(params, minHops, edges)
	}
}
