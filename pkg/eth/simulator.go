package eth

import (
	"context"
	"fmt"
	"math/big"

	"github.com/KyberNetwork/tradinglib/pkg/mev"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type Simulator struct {
	c          *rpc.Client
	gethClient *gethclient.Client
}

func NewSimulator(c *rpc.Client) *Simulator {
	return &Simulator{
		c:          c,
		gethClient: gethclient.New(c),
	}
}

func (s *Simulator) EstimateGasWithOverrides(
	ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int,
	overrides *map[common.Address]gethclient.OverrideAccount,
) (uint64, error) {
	var hex hexutil.Uint64
	err := s.c.CallContext(
		ctx, &hex, "eth_estimateGas", mev.ToCallArg(msg),
		toBlockNumArg(blockNumber), overrides,
	)

	return uint64(hex), err
}

func (s *Simulator) CallContract(
	ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int,
	overrides *map[common.Address]gethclient.OverrideAccount,
) ([]byte, error) {
	return s.gethClient.CallContract(ctx, msg, blockNumber, overrides)
}

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	if number.Sign() >= 0 {
		return hexutil.EncodeBig(number)
	}
	// It's negative.
	if number.IsInt64() {
		return rpc.BlockNumber(number.Int64()).String()
	}
	// It's negative and large, which is invalid.
	return fmt.Sprintf("<invalid %d>", number)
}
