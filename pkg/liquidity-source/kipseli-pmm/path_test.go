//nolint:dupl,testpackage
package kipselipmm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	epsilon = 10e-5
)

// nolint: funlen
func TestPath_CalcAmountOutAndIn(t *testing.T) {
	t.Run("empty path", func(t *testing.T) {
		path := Path{}
		_, _, err := path.CalcAmountIn(100)
		require.ErrorIs(t, err, ErrNoPathFound)

		_, _, err = path.CalcAmountOut(100)
		require.ErrorIs(t, err, ErrNoPathFound)
	})

	t.Run("single pair", func(t *testing.T) {
		t.Run("sell: receive quote, return base", func(t *testing.T) {
			pair := examplePair(t)
			t.Run("normal", func(t *testing.T) {
				path := Path{}.AddPair(true, pair, PairState{}, 0)
				amountIn := 305.0 // 100 at rate 2 + 50 at rate 2.1 = 200 + 105 = 305
				amountOut := 150.0
				expectedLogs := []SwapLog{
					{
						Base:           pair.Base,
						Quote:          pair.Quote,
						IsSell:         true,
						AmountIn:       amountIn,
						AmountOut:      amountOut,
						SoldBaseAmount: amountOut,
					},
				}

				gotAmountOut, logs, err := path.CalcAmountOut(amountIn)
				require.NoError(t, err)

				assert.InDelta(t, amountOut, gotAmountOut, epsilon)
				assertSwapLogs(t, expectedLogs, logs)

				gotAmountIn, logs, err := path.CalcAmountIn(amountOut)
				require.NoError(t, err)

				assert.InDelta(t, amountIn, gotAmountIn, epsilon)
				assertSwapLogs(t, expectedLogs, logs)
			})

			t.Run("lower than min trade amount", func(t *testing.T) {
				path := Path{}.AddPair(true, pair, PairState{}, 200)

				_, _, err := path.CalcAmountOut(150)
				require.ErrorIs(t, err, ErrAmountOutTooSmall)

				_, _, err = path.CalcAmountIn(150) // baseAmount = 150 < 200
				require.ErrorIs(t, err, ErrAmountOutTooSmall)
			})

			t.Run("too large amountIn", func(t *testing.T) {
				path := Path{}.AddPair(true, pair, PairState{}, 0)

				_, _, err := path.CalcAmountOut(1281)
				require.ErrorIs(t, err, ErrInsufficientLiquidity)

				_, _, err = path.CalcAmountIn(601)
				require.ErrorIs(t, err, ErrInsufficientLiquidity)
			})

			t.Run("worse rate due to pair state", func(t *testing.T) {
				pairState := PairState{SoldAmount: 150}
				path := Path{}.AddPair(true, pair, pairState, 0)
				amountIn := 315.0 // 150 at rate 2.1 = 315
				amountOut := 150.0
				expectedLogs := []SwapLog{
					{
						Base:           pair.Base,
						Quote:          pair.Quote,
						IsSell:         true,
						AmountIn:       amountIn,
						AmountOut:      amountOut,
						SoldBaseAmount: amountOut,
					},
				}

				gotAmountOut, logs, err := path.CalcAmountOut(amountIn)
				require.NoError(t, err)

				assert.InDelta(t, amountOut, gotAmountOut, epsilon)
				assertSwapLogs(t, expectedLogs, logs)

				gotAmountIn, logs, err := path.CalcAmountIn(amountOut)
				require.NoError(t, err)

				assert.InDelta(t, amountIn, gotAmountIn, epsilon)
				assertSwapLogs(t, expectedLogs, logs)
			})

			t.Run("not enough liquidity due to pair state", func(t *testing.T) {
				pairState := PairState{SoldAmount: 500}
				path := Path{}.AddPair(true, pair, pairState, 0)

				_, _, err := path.CalcAmountOut(221)
				require.ErrorIs(t, err, ErrInsufficientLiquidity)

				_, _, err = path.CalcAmountIn(201)
				require.ErrorIs(t, err, ErrInsufficientLiquidity)
			})
		})

		t.Run("buy: receive base, return quote", func(t *testing.T) {
			pair := examplePair(t)
			t.Run("normal", func(t *testing.T) {
				path := Path{}.AddPair(false, pair, PairState{}, 0)
				amountIn := 150.0
				amountOut := 280.0 // 100 at rate 1.9 + 50 at rate 1.8 = 190 + 90 = 280
				swapLogs := []SwapLog{
					{
						Base:             pair.Base,
						Quote:            pair.Quote,
						IsSell:           false,
						AmountIn:         amountIn,
						AmountOut:        amountOut,
						BoughtBaseAmount: amountIn,
					},
				}

				gotAmountOut, logs, err := path.CalcAmountOut(amountIn)
				require.NoError(t, err)

				assert.InDelta(t, amountOut, gotAmountOut, epsilon)
				assertSwapLogs(t, swapLogs, logs)

				gotAmountIn, logs, err := path.CalcAmountIn(amountOut)
				require.NoError(t, err)

				assert.InDelta(t, amountIn, gotAmountIn, epsilon)
				assertSwapLogs(t, swapLogs, logs)
			})

			t.Run("lower than min trade amount", func(t *testing.T) {
				path := Path{}.AddPair(false, pair, PairState{}, 200)

				_, _, err := path.CalcAmountOut(150) // baseAmount = 150 < 200
				require.ErrorIs(t, err, ErrAmountInTooSmall)

				_, _, err = path.CalcAmountIn(280) // baseAmount = 150 < 200
				require.ErrorIs(t, err, ErrAmountInTooSmall)
			})
			t.Run("too large amountIn", func(t *testing.T) {
				path := Path{}.AddPair(false, pair, PairState{}, 0)

				_, _, err := path.CalcAmountOut(1061) // need 1060 to buy all
				require.ErrorIs(t, err, ErrInsufficientLiquidity)

				_, _, err = path.CalcAmountIn(1061) // need 1060 to buy all
				require.ErrorIs(t, err, ErrInsufficientLiquidity)
			})

			t.Run("worse rate due to pair state", func(t *testing.T) {
				pairState := PairState{BoughtAmount: 150}
				path := Path{}.AddPair(false, pair, pairState, 0)
				amountIn := 200.0
				amountOut := 355.0 // 150 at rate 1.8 + 50 at rate 1.7 = 270 + 85 = 355
				expectedLogs := []SwapLog{
					{
						Base:             pair.Base,
						Quote:            pair.Quote,
						IsSell:           false,
						AmountIn:         amountIn,
						AmountOut:        amountOut,
						BoughtBaseAmount: amountIn,
					},
				}

				gotAmountOut, logs, err := path.CalcAmountOut(amountIn)
				require.NoError(t, err)

				assert.InDelta(t, amountOut, gotAmountOut, epsilon)
				assertSwapLogs(t, expectedLogs, logs)

				gotAmountIn, logs, err := path.CalcAmountIn(amountOut)
				require.NoError(t, err)

				assert.InDelta(t, amountIn, gotAmountIn, epsilon)
				assertSwapLogs(t, expectedLogs, logs)
			})

			t.Run("not enough liquidity due to pair state", func(t *testing.T) {
				pairState := PairState{BoughtAmount: 400}
				path := Path{}.AddPair(false, pair, pairState, 0)

				_, _, err := path.CalcAmountOut(500)
				require.ErrorIs(t, err, ErrInsufficientLiquidity)

				_, _, err = path.CalcAmountIn(500)
				require.ErrorIs(t, err, ErrInsufficientLiquidity)
			})
		})
	})

	t.Run("two pairs", func(t *testing.T) {
		pair1 := examplePair(t)
		pair2 := examplePair(t)
		pair1.Base = "A"
		pair1.Quote = "mid"
		pair2.Base = "mid"
		pair2.Quote = "B"
		t.Run("buy-sell", func(t *testing.T) {
			path := Path{}.
				AddPair(false, pair1, PairState{}, 0).
				AddPair(true, pair2, PairState{}, 0)
			amountIn := 100.0
			// first pair: 100 at rate 1.9 = 190
			mid := 190.0
			// second pair: 190 quote at rate 2 = 95 base
			amountOut := 95.0
			// A (100) -> mid (190) -> B (95)
			// => bought A = 100, sold A = 0
			// => bought B = 0, sold B = 95
			expectedLogs := []SwapLog{
				{
					Base:             pair1.Base,
					Quote:            pair1.Quote,
					IsSell:           false,
					AmountIn:         amountIn,
					AmountOut:        mid,
					BoughtBaseAmount: amountIn,
				},
				{
					Base:           pair2.Base,
					Quote:          pair2.Quote,
					IsSell:         true,
					AmountIn:       mid,
					AmountOut:      amountOut,
					SoldBaseAmount: amountOut,
				},
			}

			gotAmountOut, logs, err := path.CalcAmountOut(amountIn)
			require.NoError(t, err)

			t.Logf("Logs: %+v", logs)
			assert.InDelta(t, amountOut, gotAmountOut, epsilon)
			assertSwapLogs(t, expectedLogs, logs)

			gotAmountIn, logs, err := path.CalcAmountIn(amountOut)
			require.NoError(t, err)

			t.Logf("Logs: %+v", logs)
			assert.InDelta(t, amountIn, gotAmountIn, epsilon)
			assertSwapLogs(t, expectedLogs, logs)
		})
	})
}

func examplePair(t *testing.T) Pair {
	t.Helper()
	return Pair{
		Base:  "base",
		Quote: "quote",
		Sell: []PriceLevel{
			{Rate: 2, Quantity: 100},   // 200 quotes
			{Rate: 2.1, Quantity: 200}, // 420 quotes
			{Rate: 2.2, Quantity: 300}, // 660 quotes
		},
		Buy: []PriceLevel{
			{Rate: 1.9, Quantity: 100}, // 190 quotes
			{Rate: 1.8, Quantity: 200}, // 360 quotes
			{Rate: 1.7, Quantity: 300}, // 510 quotes
		},
	}
}

func assertSwapLogs(t *testing.T, expected, actual []SwapLog) {
	t.Helper()
	require.Equal(t, len(expected), len(actual), "number of swap logs mismatch") // nolint: testifylint
	for i := range expected {
		assert.Equal(t, expected[i].Base, actual[i].Base, "log %d: base mismatch", i)
		assert.Equal(t, expected[i].Quote, actual[i].Quote, "log %d: quote mismatch", i)
		assert.Equal(t, expected[i].IsSell, actual[i].IsSell, "log %d: isSell mismatch", i)
		assert.InDeltaf(t, expected[i].AmountIn, actual[i].AmountIn, epsilon, "log %d: amountIn mismatch", i)
		assert.InDeltaf(t, expected[i].AmountOut, actual[i].AmountOut, epsilon, "log %d: amountOut mismatch", i)
		assert.InDeltaf(t, expected[i].SoldBaseAmount, actual[i].SoldBaseAmount, epsilon,
			"log %d: soldBaseAmount mismatch", i)
		assert.InDeltaf(t, expected[i].BoughtBaseAmount, actual[i].BoughtBaseAmount, epsilon,
			"log %d: boughtBaseAmount mismatch", i)
	}
}

func TestPath_CalcAmountOut(t *testing.T) {
	t.Run("AAVE/USDT", func(t *testing.T) {
		pair := Pair{
			Base:  "AAVE",
			Quote: "USDT",
			Buy: []PriceLevel{
				{Rate: 240.78546237552965, Quantity: 8.311422745967896},
				{Rate: 240.72491476246017, Quantity: 33.24806746344619},
				{Rate: 240.69320086087652, Quantity: 36.57511963392522},
				{Rate: 240.6323653056927, Quantity: 40.23825016838933},
				{Rate: 240.5986385938008, Quantity: 44.26642687577312},
				{Rate: 240.5025346337481, Quantity: 48.69758804591842},
				{Rate: 240.467427845288, Quantity: 53.575167358805686},
				{Rate: 240.4173154641419, Quantity: 58.944967973331245},
				{Rate: 240.20110580717315, Quantity: 64.84699077964109},
				{Rate: 240.17113586546103, Quantity: 71.34059103792214},
				{Rate: 240.14628627756846, Quantity: 78.48277045348874},
				{Rate: 240.09961832134167, Quantity: 86.34782759025904},
				{Rate: 240.05938206961264, Quantity: 94.99853034452508},
				{Rate: 239.74162011208472, Quantity: 104.51243120167317},
				{Rate: 239.68065293508, Quantity: 114.99291744428558},
				{Rate: 239.58550359732072, Quantity: 126.54244448991267},
				{Rate: 239.53621022261726, Quantity: 139.22533377109426},
				{Rate: 239.53135271990118, Quantity: 153.15097286347714},
				{Rate: 239.40717812332448, Quantity: 168.52654917856353},
				{Rate: 239.00505636358574, Quantity: 185.47382194902434},
				{Rate: 238.8795923685774, Quantity: 204.12835986654204},
				{Rate: 238.8746544945989, Quantity: 224.54583743449098},
				{Rate: 238.81491932989667, Quantity: 247.06220379555043},
				{Rate: 237.91356316966397, Quantity: 271.91703460374447},
				{Rate: 237.83356475610685, Quantity: 293.73435070354844},
			},
		}

		// current AAVE price around 242.5 USDT, so this is a buy AAVE scenario
		rate := 240.7

		path := Path{}.AddPair(false, pair, PairState{}, 0.04560189742912006)
		amountIn := 2.5
		amountOut := amountIn * rate

		gotAmountOut, logs, err := path.CalcAmountOut(amountIn)
		require.NoError(t, err)
		require.Len(t, logs, 1)

		t.Logf("amountOut: %f", amountOut)
		t.Logf("logs: %+v", logs)

		assert.InDelta(t, amountOut, gotAmountOut, 5)
		assert.InDelta(t, logs[0].BoughtBaseAmount, amountIn, epsilon)

		gotAmountIn, logs, err := path.CalcAmountIn(gotAmountOut)
		require.NoError(t, err)
		require.Len(t, logs, 1)

		assert.InDelta(t, amountIn, gotAmountIn, epsilon)
		assert.InDelta(t, logs[0].BoughtBaseAmount, amountIn, epsilon)

		t.Logf("amountIn: %f", gotAmountIn)
		t.Logf("logs: %+v", logs)
	})
}
