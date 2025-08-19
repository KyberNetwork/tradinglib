package finder_test

import (
	"math/big"
	"testing"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
	"github.com/KyberNetwork/tradinglib/pkg/finderengine/entity"
	"github.com/KyberNetwork/tradinglib/pkg/finderengine/finder"
	"github.com/KyberNetwork/tradinglib/pkg/finderengine/isolated"
	"github.com/stretchr/testify/assert"
)

func Test_FindHops(t *testing.T) {
	pools := []pool.IPoolSimulator{
		&mockPool{
			address: "AB1", tokenIn: "A", tokenOut: "B",
			asks: []Order{{A: big.NewInt(20), R: big.NewInt(200)}, {A: big.NewInt(20), R: big.NewInt(90)}}, // 1 A = 2 B
		},
		&mockPool{
			address: "AB2", tokenIn: "A", tokenOut: "B",
			asks: []Order{{A: big.NewInt(20), R: big.NewInt(150)}, {A: big.NewInt(30), R: big.NewInt(120)}},
		},
	}

	isolatedPools := isolated.NewIsolatedPools(pools)
	amountIn := big.NewInt(80)
	numSplits := uint64(5)
	hop := finder.FindHops("A", 1, 18, "B", amountIn, isolatedPools, numSplits, 1)
	assert.Len(t, hop.Splits, 2)
	expectedHop := &entity.Hop{
		TokenIn:       "A",
		TokenOut:      "B",
		AmountIn:      amountIn,
		AmountOut:     big.NewInt(113),
		GasUsed:       0,
		GasFeePrice:   0,
		L1GasFeePrice: 0,
		Fee:           big.NewInt(0),
		Splits: []entity.HopSplit{
			{
				ID:            "AB1",
				AmountIn:      big.NewInt(32),
				AmountOut:     big.NewInt(50),
				Fee:           big.NewInt(0),
				GasUsed:       0,
				GasFeePrice:   0,
				L1GasFeePrice: 0,
			},
			{
				ID:            "AB2",
				AmountIn:      big.NewInt(48),
				AmountOut:     big.NewInt(63),
				Fee:           big.NewInt(0),
				GasUsed:       0,
				GasFeePrice:   0,
				L1GasFeePrice: 0,
			},
		},
	}

	expectedPools := []pool.IPoolSimulator{
		&mockPool{
			address: "AB1", tokenIn: "A", tokenOut: "B",
			asks: []Order{{A: big.NewInt(20), R: big.NewInt(200)}, {A: big.NewInt(20), R: big.NewInt(90)}}, // 1 A = 2 B
		},
		&mockPool{
			address: "AB2", tokenIn: "A", tokenOut: "B",
			asks: []Order{{A: big.NewInt(20), R: big.NewInt(150)}, {A: big.NewInt(30), R: big.NewInt(120)}},
		},
	}

	assert.Equal(t, expectedHop, hop)
	assert.Equal(t, expectedPools, pools)
}
