package flashblock

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log,
) (ethereum.Subscription, error) {
	const subscriptionType = "pendingLogs"
	arg, err := toFilterArg(q)
	if err != nil {
		return nil, fmt.Errorf("parse arg: %w", err)
	}

	sub, err := l.c.EthSubscribe(ctx, ch, subscriptionType, arg)
	if err != nil {
		return nil, fmt.Errorf("sub %s: %w", subscriptionType, err)
	}

	return sub, nil
}

func (l *RPCListener) ResubNewFlashblocks(
	ctx context.Context,
	delay time.Duration,
	ch chan<- NewFlashblock,
) error {
	// first try
	sub, err := l.SubNewFlashblocks(ctx, ch)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil

		case <-sub.Err():
			for ctx.Err() == nil {
				<-time.After(delay)
				newSub, err := l.SubNewFlashblocks(ctx, ch)
				if err == nil {
					sub = newSub
					break
				}
			}
		}
	}
}

func (l *RPCListener) ResubPendingLogs(
	ctx context.Context,
	delay time.Duration,
	q ethereum.FilterQuery,
	ch chan<- types.Log,
) error {
	// first try
	sub, err := l.SubPendingLogs(ctx, q, ch)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil

		case <-sub.Err():
			for ctx.Err() == nil {
				<-time.After(delay)
				newSub, err := l.SubPendingLogs(ctx, q, ch)
				if err == nil {
					sub = newSub
					break
				}
			}
		}
	}
}

func toFilterArg(q ethereum.FilterQuery) (interface{}, error) {
	arg := map[string]interface{}{}
	// Only include "address" when there are actual address filters.
	// An empty slice is treated the same as nil (no filter), and omitting
	// the field avoids sending "address":[] to nodes that reject empty arrays
	// (e.g. Hedera, some non-Geth implementations).
	if len(q.Addresses) > 0 {
		arg["address"] = q.Addresses
	}
	if q.Topics != nil {
		arg["topics"] = q.Topics
	}
	if q.BlockHash != nil {
		arg["blockHash"] = *q.BlockHash
		if q.FromBlock != nil || q.ToBlock != nil {
			return nil, errors.New("cannot specify both BlockHash and FromBlock/ToBlock")
		}
	} else {
		if q.FromBlock == nil {
			arg["fromBlock"] = "0x0"
		} else {
			arg["fromBlock"] = toBlockNumArg(q.FromBlock)
		}
		arg["toBlock"] = toBlockNumArg(q.ToBlock)
	}
	return arg, nil
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
