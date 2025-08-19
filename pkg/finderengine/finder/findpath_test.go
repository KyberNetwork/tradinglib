package finder_test

import (
	"math/big"
	"testing"

	dexlibPool "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
	"github.com/KyberNetwork/tradinglib/pkg/finderengine/entity"
	"github.com/KyberNetwork/tradinglib/pkg/finderengine/finder"
	"github.com/stretchr/testify/assert"
)

func TestFindBestPath_BasicGraph(t *testing.T) {
	f := &finder.Finder{
		FindHops: finder.FindHops,
	}

	pools := PoolTest()

	edges := map[string]map[string][]dexlibPool.IPoolSimulator{
		"A": {
			"B": []dexlibPool.IPoolSimulator{pools["AB1"], pools["AB2"]},
			"C": []dexlibPool.IPoolSimulator{pools["AC1"], pools["AC2"]},
		},
		"B": {
			"A": []dexlibPool.IPoolSimulator{pools["AB1"], pools["AB2"]},
			"C": []dexlibPool.IPoolSimulator{pools["BC1"], pools["BC2"]},
		},
		"C": {
			"A": []dexlibPool.IPoolSimulator{pools["AC1"], pools["AC2"]},
			"B": []dexlibPool.IPoolSimulator{pools["BC1"], pools["BC2"]},
		},
	}

	minHops := map[string]uint64{
		"A": 1,
		"B": 1,
		"C": 0,
	}

	params := &entity.FinderParams{
		TokenIn:       "A",
		TargetToken:   "C",
		MaxHop:        5,
		NumHopSplits:  5,
		NumPathSplits: 5,
		AmountIn:      big.NewInt(12),
		GasPrice:      big.NewInt(0),
		Tokens: map[string]entity.SimplifiedToken{
			"A": {}, "B": {}, "C": {},
		},
		WhitelistHopTokens: map[string]struct{}{
			"B": {}, "C": {},
		},
	}

	results := f.FindBestPathsOptimized(params, minHops, edges)

	expectedResult := &entity.Path{
		AmountIn:    big.NewInt(12),
		AmountOut:   big.NewInt(43),
		TokenOrders: []string{"A", "B", "C"},
		HopOrders: []entity.Hop{
			{
				TokenIn: "A", TokenOut: "B", AmountIn: big.NewInt(12), AmountOut: big.NewInt(23), Fee: big.NewInt(0),
				Splits: []entity.HopSplit{
					{ID: "AB1", AmountIn: big.NewInt(10), AmountOut: big.NewInt(20), Fee: big.NewInt(0)},
					{ID: "AB2", AmountIn: big.NewInt(2), AmountOut: big.NewInt(3), Fee: big.NewInt(0)},
				},
			},
			{
				TokenIn: "B", TokenOut: "C", AmountIn: big.NewInt(23), AmountOut: big.NewInt(43), Fee: big.NewInt(0),
				Splits: []entity.HopSplit{
					{ID: "BC1", AmountIn: big.NewInt(20), AmountOut: big.NewInt(40), Fee: big.NewInt(0)},
					{ID: "BC2", AmountIn: big.NewInt(3), AmountOut: big.NewInt(3), Fee: big.NewInt(0)},
				},
			},
		},
	}
	assert.NotEmpty(t, results)
	assert.Equal(t, expectedResult, results)
}
