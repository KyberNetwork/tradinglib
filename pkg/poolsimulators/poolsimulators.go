package poolsimulators

import (
	"errors"
	"fmt"
	"math/big"

	ksent "github.com/KyberNetwork/kyberswap-dex-lib/pkg/entity"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/pooltypes"
	pkgpool "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/swaplimit"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/valueobject"
)

var ErrPoolTypeNotSupported = errors.New("pool type is not supported")

// PoolSimulatorFromPool init pool sim.
func PoolSimulatorFromPool(pool ksent.Pool, chainID uint) (pkgpool.IPoolSimulator, error) {
	factoryFn := pkgpool.Factory(pool.Type)
	if factoryFn == nil {
		return nil, fmt.Errorf("%w: %s", ErrPoolTypeNotSupported, pool.Type)
	}

	return factoryFn(pkgpool.FactoryParams{
		EntityPool:  pool,
		ChainID:     valueobject.ChainID(chainID),
		BasePoolMap: nil,
		EthClient:   nil,
	})
}

func newSwapLimit(dex string, limit map[string]*big.Int) pkgpool.SwapLimit {
	switch dex {
	case pooltypes.PoolTypes.Synthetix,
		pooltypes.PoolTypes.NativeV1,
		pooltypes.PoolTypes.Dexalot,
		pooltypes.PoolTypes.RingSwap,
		pooltypes.PoolTypes.MxTrading,
		pooltypes.PoolTypes.LO1inch,
		pooltypes.PoolTypes.KyberPMM:
		return swaplimit.NewInventory(dex, limit)

	case pooltypes.PoolTypes.LimitOrder:
		return swaplimit.NewInventoryWithAllowedSenders(
			dex,
			limit,
			// here just for usecase of kyberswap only, there are some client have priority private limit orders.
			"",
		)

	case pooltypes.PoolTypes.Bebop:
		return swaplimit.NewSingleSwapLimit(dex)
	}

	return nil
}

func NewSwapLimit(limits map[string]map[string]*big.Int) map[string]pkgpool.SwapLimit {
	limitMap := make(map[string]pkgpool.SwapLimit, len(limits))

	for dex, limit := range limits {
		limitMap[dex] = newSwapLimit(dex, limit)
	}

	return limitMap
}
