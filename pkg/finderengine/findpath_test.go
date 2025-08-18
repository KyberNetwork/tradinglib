package finderengine_test

import (
	"math/big"
	"testing"

	dexlibPool "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
	"github.com/KyberNetwork/tradinglib/pkg/finderengine"
	"github.com/KyberNetwork/tradinglib/pkg/finderengine/entity"
	"github.com/stretchr/testify/assert"
)

func TestFindBestPath_BasicGraph(t *testing.T) {
	f := &finderengine.Finder{
		FindHops: finderengine.FindHops,
	}

	pools := map[string]dexlibPool.IPoolSimulator{
		"AB1": &mockPool{
			address: "AB1", tokenIn: "A", tokenOut: "B",
			bids: []Order{{A: big.NewInt(10), R: big.NewInt(50)}, {A: big.NewInt(10), R: big.NewInt(20)}},
			asks: []Order{{A: big.NewInt(10), R: big.NewInt(200)}, {A: big.NewInt(10), R: big.NewInt(90)}},
		},
		"AB2": &mockPool{
			address: "AB2", tokenIn: "A", tokenOut: "B",
			bids: []Order{{A: big.NewInt(10), R: big.NewInt(50)}},
			asks: []Order{{A: big.NewInt(10), R: big.NewInt(150)}},
		},
		"AC1": &mockPool{
			address: "AC1", tokenIn: "A", tokenOut: "C",
			bids: []Order{{A: big.NewInt(10), R: big.NewInt(30)}, {A: big.NewInt(10), R: big.NewInt(20)}},
			asks: []Order{{A: big.NewInt(20), R: big.NewInt(300)}, {A: big.NewInt(10), R: big.NewInt(100)}},
		},
		"AC2": &mockPool{
			address: "AC2", tokenIn: "A", tokenOut: "C",
			bids: []Order{{A: big.NewInt(10), R: big.NewInt(50)}},
			asks: []Order{{A: big.NewInt(10), R: big.NewInt(250)}},
		},

		"BC1": &mockPool{
			address: "BC1", tokenIn: "B", tokenOut: "C",
			bids: []Order{{A: big.NewInt(10), R: big.NewInt(50)}, {A: big.NewInt(10), R: big.NewInt(20)}},
			asks: []Order{{A: big.NewInt(20), R: big.NewInt(200)}, {A: big.NewInt(10), R: big.NewInt(100)}},
		},
		"BC2": &mockPool{
			address: "BC2", tokenIn: "B", tokenOut: "C",
			bids: []Order{{A: big.NewInt(10), R: big.NewInt(50)}, {A: big.NewInt(10), R: big.NewInt(20)}},
			asks: []Order{{A: big.NewInt(10), R: big.NewInt(100)}, {A: big.NewInt(10), R: big.NewInt(50)}},
		},
	}

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
