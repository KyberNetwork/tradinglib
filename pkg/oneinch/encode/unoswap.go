package encode

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"math/big"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/pooltypes"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var (
	//go:embed OneInchAggregationRouterV6.json
	oneInchAggregationRouterV6JSON []byte

	OneInchAggregationRouterV6ABI abi.ABI
)

const (
	MethodUnoswap  = "unoswap"
	MethodUnoswap2 = "unoswap2"
	MethodUnoswap3 = "unoswap3"
)

const (
	unoswapProtocolOffset   = 253
	uniswapZeroForOneOffset = 247
)

func init() {
	parsed, err := abi.JSON(bytes.NewReader(oneInchAggregationRouterV6JSON))
	if err != nil {
		log.Println("[ERROR] parse abi json: %w", err)
		return
	}

	OneInchAggregationRouterV6ABI = parsed
}

func EncodeUnoswap(
	poolTypes []string, pools []common.Address, tokens []common.Address, amountIn, amountOut *big.Int,
) ([]byte, error) {
	if len(pools) == 0 || len(pools) > 3 {
		return nil, fmt.Errorf("%w: %d", ErrInvalidSwapLength, len(pools))
	}
	if len(poolTypes) != len(pools) {
		return nil, fmt.Errorf("%w: %d/%d", ErrInvalidPoolsAndPoolTypesLength, len(poolTypes), len(pools))
	}
	if len(pools) != len(tokens)-1 {
		return nil, fmt.Errorf("%w: %d/%d", ErrInvalidPoolsAndTokensLength, len(pools), len(tokens))
	}

	dexes := make([]*big.Int, 0, len(pools))
	for i := range pools {
		dex, err := GetUnoswapDex(poolTypes[i], tokens[i].String(), tokens[i+1].String(), pools[i])
		if err != nil {
			return nil, err
		}

		dexes = append(dexes, dex)
	}

	switch len(dexes) {
	case 1: //nolint:gomnd
		return EncodeUnoswap1(
			tokens[0].Big(), amountIn, amountOut, dexes[0],
		)

	case 2: //nolint:gomnd
		return EncodeUnoswap2(
			tokens[0].Big(), amountIn, amountOut, dexes[0], dexes[1],
		)

	case 3: //nolint:gomnd
		return EncodeUnoswap3(
			tokens[0].Big(), amountIn, amountOut, dexes[0], dexes[1], dexes[2],
		)

	default:
		return nil, fmt.Errorf("%w: %d", ErrInvalidSwapLength, len(dexes))
	}
}

func EncodeUnoswap1(
	token *big.Int,
	amount *big.Int,
	minReturn *big.Int,
	dex *big.Int,
) ([]byte, error) {
	return OneInchAggregationRouterV6ABI.Pack(MethodUnoswap, token, amount, minReturn, dex)
}

func EncodeUnoswap2(
	token *big.Int,
	amount *big.Int,
	minReturn *big.Int,
	dex1 *big.Int,
	dex2 *big.Int,
) ([]byte, error) {
	return OneInchAggregationRouterV6ABI.Pack(MethodUnoswap2, token, amount, minReturn, dex1, dex2)
}

func EncodeUnoswap3(
	token *big.Int,
	amount *big.Int,
	minReturn *big.Int,
	dex1 *big.Int,
	dex2 *big.Int,
	dex3 *big.Int,
) ([]byte, error) {
	return OneInchAggregationRouterV6ABI.Pack(MethodUnoswap3, token, amount, minReturn, dex1, dex2, dex3)
}

func GetUnoswapDex(poolType string, tokenIn, tokenOut string, poolAddress common.Address) (*big.Int, error) {
	var dex *big.Int
	switch poolType {
	case pooltypes.PoolTypes.UniswapV2:
		dex = big.NewInt(0)

	case pooltypes.PoolTypes.UniswapV3:
		dex = big.NewInt(1)

	default:
		return nil, fmt.Errorf("%w: %s", ErrNotSupportedDex, poolType)
	}
	dex.Lsh(dex, unoswapProtocolOffset)

	zeroForOne := tokenIn < tokenOut
	if zeroForOne {
		dex.SetBit(dex, uniswapZeroForOneOffset, 1)
	}
	dex.Or(dex, poolAddress.Big())

	return dex, nil
}

func CanEncodeUnoswap(poolType string) bool {
	switch poolType {
	case pooltypes.PoolTypes.UniswapV2, pooltypes.PoolTypes.UniswapV3:
		return true

	default:
		return false
	}
}
