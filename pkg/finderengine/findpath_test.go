package finderengine

import (
	"fmt"
	"math/big"
	"testing"

	dexlibPool "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
	"github.com/KyberNetwork/tradinglib/pkg/finderengine/entity"
	"github.com/stretchr/testify/require"
)

func TestFindBestPaths_ComplexGraph(t *testing.T) {
	f := &Finder{
		FindHops: func(tokenIn string, tokenInPrice float64, tokenInDecimals uint8, tokenOut string, amountIn *big.Int, pools []dexlibPool.IPoolSimulator, numSplits uint64) *entity.Hop {
			return &entity.Hop{
				TokenIn:   tokenIn,
				TokenOut:  tokenOut,
				AmountIn:  amountIn,
				AmountOut: new(big.Int).Add(amountIn, big.NewInt(1)),
				Splits: []*entity.HopSplit{
					{ID: fmt.Sprintf("MockPool-%s-%s", tokenIn, tokenOut)},
				},
			}
		},
	}

	edges := map[string]map[string][]dexlibPool.IPoolSimulator{
		"A": {"B": {}, "C": {}, "G": {}},
		"B": {"A": {}, "D": {}},
		"C": {"A": {}, "D": {}, "E": {}},
		"D": {"B": {}, "C": {}, "F": {}},
		"E": {"C": {}, "F": {}},
		"F": {"D": {}, "E": {}, "G": {}},
		"G": {"A": {}, "F": {}},
	}

	minHops := map[string]uint64{
		"A": 3,
		"B": 2,
		"C": 2,
		"G": 1,
		"D": 1,
		"E": 1,
		"F": 0,
	}
	params := &entity.FinderParams{
		TokenIn:       "A",
		TargetToken:   "F",
		MaxHop:        5,
		NumHopSplits:  1,
		NumPathSplits: 1,
		AmountIn:      big.NewInt(100),
		Tokens: map[string]entity.SimplifiedToken{
			"A": {}, "B": {}, "C": {}, "D": {}, "E": {}, "F": {}, "G": {},
		},
		WhitelistHopTokens: map[string]struct{}{
			"B": {}, "C": {}, "D": {}, "E": {}, "F": {}, "G": {},
		},
	}

	results := f.findBestPathsOptimized(params, minHops, edges)

	require.NotEmpty(t, results)
}
