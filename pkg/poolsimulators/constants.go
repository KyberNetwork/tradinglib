package poolsimulators

import (
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/valueobject"
)

// DefaultGasAlgebra
// nolint: gochecknoglobals, gomnd
var (
	DefaultGasAlgebra = map[valueobject.Exchange]int64{
		valueobject.ExchangeQuickSwapV3: 280000,
		valueobject.ExchangeSynthSwapV3: 280000,
		valueobject.ExchangeSwapBasedV3: 280000,
		valueobject.ExchangeLynex:       280000,
		valueobject.ExchangeCamelotV3:   280000,
		valueobject.ExchangeZyberSwapV3: 280000,
		valueobject.ExchangeThenaFusion: 280000,
	}

	DefaultGasAlgebraIntegral = map[valueobject.Exchange]int64{
		valueobject.ExchangeHorizonIntegral: 280000,
		valueobject.ExchangeSwapsicle:       280000,
		valueobject.ExchangeScribe:          280000,
		valueobject.ExchangeSilverSwap:      280000,
		valueobject.ExchangeFenix:           280000,
		valueobject.ExchangeBlade:           280000,
	}
)
