package flashblock_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/KyberNetwork/tradinglib/pkg/flashblock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRPCListener(t *testing.T) {
	const nodeRPC = ""
	if nodeRPC == "" {
		t.Skip()
	}

	c, err := flashblock.DialRPCListener(t.Context(), nodeRPC)
	require.NoError(t, err)

	ch := make(chan *flashblock.NewFlashblock)
	ctx, cancel := context.WithTimeout(t.Context(), time.Second*3)
	defer cancel()

	sub, err := c.SubNewFlashblocks(ctx, ch)
	require.NoError(t, err)

	for {
		select {
		case <-sub.Err():
			return

		case <-ctx.Done():
			return

		case v, ok := <-ch:
			assert.True(t, ok)

			str, err := json.MarshalIndent(v, "", "   ")
			require.NoError(t, err)

			t.Log(string(str))
		}
	}
}

func TestRPCListenerSubPendingLogs(t *testing.T) {
	const nodeRPC = ""
	if nodeRPC == "" {
		t.Skip()
	}

	c, err := flashblock.DialRPCListener(t.Context(), nodeRPC)
	require.NoError(t, err)

	ch := make(chan types.Log)
	ctx, cancel := context.WithTimeout(t.Context(), time.Second*3)
	defer cancel()

	sub, err := c.SubPendingLogs(ctx, ch)
	require.NoError(t, err)

	received := false

loop:
	for {
		select {
		case <-sub.Err():
			break loop

		case <-ctx.Done():
			break loop

		case v, ok := <-ch:
			assert.True(t, ok)
			received = true

			// BlockHash is intentionally zero: pending logs aren't part of a
			// mined block yet.
			assert.NotEqual(t, common.Address{}, v.Address)
			assert.NotEmpty(t, v.Topics)
			assert.NotEqual(t, common.Hash{}, v.TxHash)
			assert.NotZero(t, v.BlockNumber)

			str, err := json.MarshalIndent(v, "", "   ")
			require.NoError(t, err)

			t.Log(string(str))
		}
	}

	assert.True(t, received, "expected to receive at least one pending log")
}
