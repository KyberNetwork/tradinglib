package flashblock

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

type RPCListener struct {
	c *rpc.Client
}

func NewRPCListener(c *rpc.Client) *RPCListener {
	return &RPCListener{
		c: c,
	}
}

func DialRPCListener(ctx context.Context, url string) (*RPCListener, error) {
	c, err := rpc.DialContext(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("dial rpc: %w", err)
	}

	return &RPCListener{c: c}, nil
}

// https://docs.base.org/base-chain/api-reference/flashblocks-api/newFlashblocks
func (l *RPCListener) SubNewFlashblocks(
	ctx context.Context, ch chan<- NewFlashblock,
) (ethereum.Subscription, error) {
	const subscriptionType = "newFlashblocks"
	sub, err := l.c.EthSubscribe(ctx, ch, subscriptionType)
	if err != nil {
		return nil, fmt.Errorf("sub %s: %w", subscriptionType, err)
	}

	return sub, nil
}

func (l *RPCListener) SubPendingLogs(
	ctx context.Context, ch chan<- types.Log,
) (ethereum.Subscription, error) {
	const subscriptionType = "pendingLogs"
	sub, err := l.c.EthSubscribe(ctx, ch, subscriptionType)
	if err != nil {
		return nil, fmt.Errorf("sub %s: %w", subscriptionType, err)
	}

	return sub, nil
}
