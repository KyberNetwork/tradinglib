package encode

import (
	"bytes"
	_ "embed"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
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

func init() {
	parsed, err := abi.JSON(bytes.NewReader(oneInchAggregationRouterV6JSON))
	if err != nil {
		log.Println("[ERROR] parse abi json: %w", err)
		return
	}

	OneInchAggregationRouterV6ABI = parsed
}

func EncodeUnoswap(
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
