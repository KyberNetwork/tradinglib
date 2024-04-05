package eth

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type Simulator struct {
	c *rpc.Client
}

func NewSimulator(c *rpc.Client) *Simulator {
	return &Simulator{
		c: c,
	}
}

func (s *Simulator) EstimateGasWithOverrides(
	ctx context.Context, msg ethereum.CallMsg, block *big.Int, blockNumber *big.Int,
	overrides *map[common.Address]gethclient.OverrideAccount,
) (uint64, error) {
	var hex hexutil.Uint64
	err := s.c.CallContext(
		ctx, &hex, "eth_estimateGas", toCallArg(msg),
		toBlockNumArg(blockNumber), overrides,
	)

	return uint64(hex), err
}

func toCallArg(msg ethereum.CallMsg) interface{} {
	arg := map[string]interface{}{
		"from": msg.From,
		"to":   msg.To,
	}
	if len(msg.Data) > 0 {
		arg["input"] = hexutil.Bytes(msg.Data)
	}
	if msg.Value != nil {
		arg["value"] = (*hexutil.Big)(msg.Value)
	}
	if msg.Gas != 0 {
		arg["gas"] = hexutil.Uint64(msg.Gas)
	}
	if msg.GasPrice != nil {
		arg["gasPrice"] = (*hexutil.Big)(msg.GasPrice)
	}
	return arg
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
