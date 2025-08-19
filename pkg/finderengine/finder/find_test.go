package finder_test

import (
	"math/big"
	"testing"

	dexlibPool "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
	"github.com/KyberNetwork/tradinglib/pkg/finderengine/entity"
	"github.com/KyberNetwork/tradinglib/pkg/finderengine/finder"
	"github.com/stretchr/testify/assert"
)

func PoolTest() map[string]dexlibPool.IPoolSimulator {
	return map[string]dexlibPool.IPoolSimulator{
		"AB1": &mockPool{
			address: "AB1", tokenIn: "A", tokenOut: "B",
			bids: []Order{{A: big.NewInt(100), R: big.NewInt(50)}, {A: big.NewInt(100), R: big.NewInt(20)}},  // 1 B = 1/2 A
			asks: []Order{{A: big.NewInt(100), R: big.NewInt(200)}, {A: big.NewInt(100), R: big.NewInt(90)}}, // 1 A = 2 B
		},
		"AB2": &mockPool{
			address: "AB2", tokenIn: "A", tokenOut: "B",
			bids: []Order{{A: big.NewInt(100), R: big.NewInt(50)}},
			asks: []Order{{A: big.NewInt(100), R: big.NewInt(150)}},
		},
		"AC1": &mockPool{
			address: "AC1", tokenIn: "A", tokenOut: "C",
			bids: []Order{{A: big.NewInt(100), R: big.NewInt(30)}, {A: big.NewInt(100), R: big.NewInt(20)}},
			asks: []Order{{A: big.NewInt(200), R: big.NewInt(300)}, {A: big.NewInt(100), R: big.NewInt(100)}},
		},
		"AC2": &mockPool{
			address: "AC2", tokenIn: "A", tokenOut: "C",
			bids: []Order{{A: big.NewInt(100), R: big.NewInt(50)}},
			asks: []Order{{A: big.NewInt(100), R: big.NewInt(250)}},
		},

		"BC1": &mockPool{
			address: "BC1", tokenIn: "B", tokenOut: "C",
			bids: []Order{{A: big.NewInt(100), R: big.NewInt(50)}, {A: big.NewInt(100), R: big.NewInt(20)}},
			asks: []Order{{A: big.NewInt(200), R: big.NewInt(200)}, {A: big.NewInt(100), R: big.NewInt(100)}},
		},
		"BC2": &mockPool{
			address: "BC2", tokenIn: "B", tokenOut: "C",
			bids: []Order{{A: big.NewInt(10), R: big.NewInt(50)}, {A: big.NewInt(100), R: big.NewInt(20)}},
			asks: []Order{{A: big.NewInt(10), R: big.NewInt(100)}, {A: big.NewInt(100), R: big.NewInt(50)}},
		},
	}
}

func Test_Find(t *testing.T) {
	f := &finder.Finder{
		FindHops: finder.FindHops,
	}

	pools := PoolTest()

	params := entity.FinderParams{
		TokenIn:       "A",
		TargetToken:   "C",
		MaxHop:        5,
		NumHopSplits:  5,
		NumPathSplits: 5,
		AmountIn:      big.NewInt(555),
		GasPrice:      big.NewInt(0),
		Tokens: map[string]entity.SimplifiedToken{
			"A": {}, "B": {}, "C": {},
		},
		WhitelistHopTokens: map[string]struct{}{
			"B": {}, "C": {},
		},
		Pools: pools,
	}

	bestRoute, err := f.Find(params)
	assert.NoError(t, err)
	assert.NotEmpty(t, bestRoute)

	expectedPools := map[string]dexlibPool.IPoolSimulator{
		"AB1": &mockPool{
			address: "AB1", tokenIn: "A", tokenOut: "B",
			bids: []Order{{A: big.NewInt(100), R: big.NewInt(50)}, {A: big.NewInt(100), R: big.NewInt(20)}},  // 1 B = 1/2 A
			asks: []Order{{A: big.NewInt(100), R: big.NewInt(200)}, {A: big.NewInt(100), R: big.NewInt(90)}}, // 1 A = 2 B
		},
		"AB2": &mockPool{
			address: "AB2", tokenIn: "A", tokenOut: "B",
			bids: []Order{{A: big.NewInt(100), R: big.NewInt(50)}},
			asks: []Order{{A: big.NewInt(100), R: big.NewInt(150)}},
		},
		"AC1": &mockPool{
			address: "AC1", tokenIn: "A", tokenOut: "C",
			bids: []Order{{A: big.NewInt(100), R: big.NewInt(30)}, {A: big.NewInt(100), R: big.NewInt(20)}},
			asks: []Order{{A: big.NewInt(200), R: big.NewInt(300)}, {A: big.NewInt(100), R: big.NewInt(100)}},
		},
		"AC2": &mockPool{
			address: "AC2", tokenIn: "A", tokenOut: "C",
			bids: []Order{{A: big.NewInt(100), R: big.NewInt(50)}},
			asks: []Order{{A: big.NewInt(100), R: big.NewInt(250)}},
		},

		"BC1": &mockPool{
			address: "BC1", tokenIn: "B", tokenOut: "C",
			bids: []Order{{A: big.NewInt(100), R: big.NewInt(50)}, {A: big.NewInt(100), R: big.NewInt(20)}},
			asks: []Order{{A: big.NewInt(200), R: big.NewInt(200)}, {A: big.NewInt(100), R: big.NewInt(100)}},
		},
		"BC2": &mockPool{
			address: "BC2", tokenIn: "B", tokenOut: "C",
			bids: []Order{{A: big.NewInt(10), R: big.NewInt(50)}, {A: big.NewInt(100), R: big.NewInt(20)}},
			asks: []Order{{A: big.NewInt(10), R: big.NewInt(100)}, {A: big.NewInt(100), R: big.NewInt(50)}},
		},
	}

	assert.Equal(t, expectedPools, pools)
}
