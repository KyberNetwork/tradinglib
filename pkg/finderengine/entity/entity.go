package entity

import (
	"math/big"

	dexlibPool "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
)

type Token struct {
	Address  string `json:"address"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Decimals uint8  `json:"decimals"`
}

func (t Token) GetAddress() string {
	return t.Address
}

type SimplifiedToken struct {
	Address  string `json:"address"`
	Decimals uint8  `json:"decimals"`
}

func (t SimplifiedToken) GetAddress() string {
	return t.Address
}

type FinderParams struct {
	// TokenIn is the token to be swapped
	TokenIn string
	// TokenOut is the token to be received
	TargetToken string
	// AmountIn is the amount of TokenIn to be swapped
	AmountIn *big.Int

	// WhitelistHopTokens is the list of tokens that can be used as intermediate tokens
	// when finding the best route.
	WhitelistHopTokens map[string]struct{}

	// Pools is a mapping between pool address and its simulator.
	// The pathfinder will use these pools to find the best route.
	Pools map[string]dexlibPool.IPoolSimulator

	// SwapLimits is a mapping between pool type and its swap limit (inventory).
	// The pathfinder will use these limits to find the best route.
	SwapLimits map[string]dexlibPool.SwapLimit

	// Tokens is a mapping between token address and its information.
	// TokenIn, TokenOut, WhitelistTokens (& GasToken if GasInclude = true)
	// should be included in this map.
	Tokens map[string]SimplifiedToken

	// Prices is a mapping between token address and its price.
	// The price can be USD price or Native price (from the On-chain price feed).
	// If GasIncluded is true, the pathfinder will use the price information to find the best route.
	Prices map[string]float64

	// GasIncluded is the flag to indicate whether the gas fee is included in finding the best route or not.
	// If true, the gas fee will be accounted for in the final result (the best route is the one with the
	// highest price of TokenOut after deducting the gas fee).
	// If false, the gas fee will be ignored (the best route is the one with the highest amount of TokenOut).
	GasIncluded bool

	// GasToken is the token used to pay for the gas fee. Required if GasIncluded is true.
	GasToken string

	// GasPrice is the gas price in WEI. Required if GasIncluded is true.
	// This field should be differentiated from the price of the gas token:
	// GasFee = GasPrice * GasUsed;
	// GasFeePrice = GasFee * Price[GasToken] / 10^Tokens[GasToken].Decimals;
	GasPrice *big.Int

	// L1GasFeePriceOverhead estimated L1 gas fee for an empty route summary data (without a pool)
	// in Price value (USD/Native).
	L1GasFeePriceOverhead float64

	// L1GasFeePricePerPool estimated L1 gas fee for each pool in Price value (USD/Native).
	L1GasFeePricePerPool float64

	MaxHop        uint64
	NumPathSplits uint64
	NumHopSplits  uint64
}
