package poolsimulators

// DEPRECATED: pls refer to https://github.com/KyberNetwork/reserve-taker/blob/v4.0.4/internal/pool/parse_pool.go#L10-L23

// var ErrPoolTypeNotSupported = errors.New("pool type is not supported")

// // PoolSimulatorFromPool init pool sim.
// func PoolSimulatorFromPool(pool ksent.Pool, chainID uint) (pkgpool.IPoolSimulator, error) {
// 	factoryFn := pkgpool.Factory(pool.Type)
// 	if factoryFn == nil {
// 		return nil, fmt.Errorf("%w: %s", ErrPoolTypeNotSupported, pool.Type)
// 	}

// 	return factoryFn(pkgpool.FactoryParams{
// 		EntityPool:  pool,
// 		ChainID:     valueobject.ChainID(chainID),
// 		BasePoolMap: nil,
// 		EthClient:   nil,
// 	})
// }

// func newSwapLimit(dex string, limit map[string]*big.Int) pkgpool.SwapLimit {
// 	switch dex {
// 	case pooltypes.PoolTypes.Synthetix,
// 		pooltypes.PoolTypes.NativeV1,
// 		pooltypes.PoolTypes.Dexalot,
// 		pooltypes.PoolTypes.RingSwap,
// 		pooltypes.PoolTypes.LO1inch,
// 		pooltypes.PoolTypes.KyberPMM,
// 		pooltypes.PoolTypes.Pmm1,
// 		pooltypes.PoolTypes.Pmm2:
// 		return swaplimit.NewInventory(dex, limit)

// 	case pooltypes.PoolTypes.LimitOrder:
// 		return swaplimit.NewInventoryWithAllowedSenders(
// 			dex,
// 			limit,
// 			// here just for usecase of kyberswap only, there are some client have priority private limit orders.
// 			"",
// 		)

// 	case pooltypes.PoolTypes.Bebop:
// 		return swaplimit.NewSingleSwapLimit(dex)
// 	}

// 	return nil
// }

// func NewSwapLimit(limits map[string]map[string]*big.Int) map[string]pkgpool.SwapLimit {
// 	limitMap := make(map[string]pkgpool.SwapLimit, len(limits))

// 	for dex, limit := range limits {
// 		limitMap[dex] = newSwapLimit(dex, limit)
// 	}

// 	return limitMap
// }
