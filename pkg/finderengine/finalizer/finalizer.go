package finalizer

import (
	"fmt"
	"math/big"

	dexlibPool "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
	"github.com/KyberNetwork/tradinglib/pkg/finderengine/entity"
	"github.com/KyberNetwork/tradinglib/pkg/finderengine/isolated"
	"github.com/KyberNetwork/tradinglib/pkg/finderengine/utils"
	"go.uber.org/zap"
)

const BASE_GAS int64 = 125000

type finalizer struct{}

func NewFinalizer() *finalizer {
	return &finalizer{}
}

func (f *finalizer) Finalize(params entity.FinderParams, route *entity.Route) (finalRoute *entity.FinalizedRoute, err error) {
	defer func() {
		if r := recover(); r != nil {
			finalRoute = nil
			err = fmt.Errorf("panic finalize route: %v", r)
		}
	}()

	// Build isolated pools safely
	isolatedPools := make(map[string]*isolated.Pool, len(params.Pools))
	for address, pool := range params.Pools {
		isolatedPools[address] = isolated.NewIsolatedPool(pool)
	}

	var (
		amountOut     = big.NewInt(0)
		gasUsed       = BASE_GAS
		l1GasFeePrice = params.L1GasFeePriceOverhead
	)

	finalizedRoute := make([][]entity.Swap, 0, len(route.Paths))

	for _, path := range route.Paths {
		if len(path.HopOrders) == 0 {
			return nil, fmt.Errorf("route contains an empty path")
		}
		if len(path.TokenOrders) == 0 {
			return nil, fmt.Errorf("path has no token orders")
		}

		finalizedPath := make([]entity.Swap, 0, len(path.HopOrders))
		currentAmountIn := new(big.Int).Set(path.AmountIn)

		for _, hop := range path.HopOrders {
			fromToken := hop.TokenIn
			toToken := hop.TokenOut

			hopAmountOut := big.NewInt(0)

			for _, split := range hop.Splits {
				hopAmountIn := new(big.Int).Set(split.AmountIn)
				if hopAmountIn.Cmp(currentAmountIn) > 0 {
					hopAmountIn = new(big.Int).Set(currentAmountIn)
				}

				// Decrease current available
				currentAmountIn.Sub(currentAmountIn, hopAmountIn)

				tokenAmountIn := dexlibPool.TokenAmount{Token: fromToken, Amount: hopAmountIn}

				pool, ok := isolatedPools[split.ID]
				if !ok || pool == nil {
					return nil, fmt.Errorf("unknown or nil pool id: %s", split.ID)
				}

				res, calcErr := dexlibPool.CalcAmountOut(pool, tokenAmountIn, toToken, nil)
				if calcErr != nil {
					zap.S().Warnf(
						"failed to swap %s %v to %v in pool %s: %v",
						hopAmountIn.String(), fromToken, toToken, pool.GetAddress(), calcErr,
					)
					return nil, fmt.Errorf("invalid swap: %w", calcErr)
				}
				if res == nil || !res.IsValid() || res.TokenAmountOut == nil {
					return nil, fmt.Errorf("invalid swap result: empty amountOut for pool %s", pool.GetAddress())
				}

				updateBalanceParams := dexlibPool.UpdateBalanceParams{
					TokenAmountIn:  tokenAmountIn,
					TokenAmountOut: *res.TokenAmountOut,
					Fee:            *res.Fee,
					SwapInfo:       res.SwapInfo,
				}
				pool.UpdateBalance(updateBalanceParams)

				finalizedPath = append(finalizedPath, entity.Swap{
					Pool:      pool.GetAddress(),
					TokenIn:   fromToken,
					TokenOut:  toToken,
					AmountIn:  hopAmountIn,
					AmountOut: res.TokenAmountOut.Amount,
				})

				hopAmountOut.Add(hopAmountOut, res.TokenAmountOut.Amount)
				gasUsed += res.Gas
			}

			l1GasFeePrice += params.L1GasFeePricePerPool * float64(len(hop.Splits))
			currentAmountIn = hopAmountOut
		}

		lastToken := path.TokenOrders[len(path.TokenOrders)-1]
		if lastToken == params.TargetToken {
			amountOut.Add(amountOut, currentAmountIn)
		}

		finalizedRoute = append(finalizedRoute, finalizedPath)
	}

	gasFee := new(big.Int).Mul(big.NewInt(gasUsed), params.GasPrice)
	if _, ok := params.Tokens[params.TokenIn]; !ok {
		return nil, fmt.Errorf("missing token metadata for input token %v", params.TokenIn)
	}
	if _, ok := params.Tokens[params.TargetToken]; !ok {
		return nil, fmt.Errorf("missing token metadata for output token %v", params.TargetToken)
	}
	if _, ok := params.Tokens[params.GasToken]; !ok {
		return nil, fmt.Errorf("missing token metadata for gas token %v", params.GasToken)
	}

	finalRoute = &entity.FinalizedRoute{
		TokenIn:  params.TokenIn,
		AmountIn: params.AmountIn,
		AmountInPrice: utils.CalcAmountPrice(
			params.AmountIn,
			params.Tokens[params.TokenIn].Decimals,
			params.Prices[params.TokenIn],
		),
		TokenOut:  params.TargetToken,
		AmountOut: amountOut,
		AmountOutPrice: utils.CalcAmountPrice(
			amountOut,
			params.Tokens[params.TargetToken].Decimals,
			params.Prices[params.TargetToken],
		),
		GasUsed:  gasUsed,
		GasPrice: params.GasPrice,
		GasFee:   gasFee,
		GasFeePrice: utils.CalcAmountPrice(
			gasFee,
			params.Tokens[params.GasToken].Decimals,
			params.Prices[params.GasToken],
		),
		L1GasFeePrice: l1GasFeePrice,
		Route:         finalizedRoute,
	}

	return finalRoute, nil
}
