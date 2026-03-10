// nolint: testpackage
package flashblock

import (
	"context"
	"testing"
	"time"

	"go.uber.org/zap"
)

const baseNode = "wss://mainnet.flashblocks.base.org/ws"

func TestNodeListenFlashBlock(t *testing.T) {
	t.Skip("skip for CI")

	logger, _ := zap.NewDevelopment()
	defer logger.Sync() //nolint:errcheck
	zap.ReplaceGlobals(logger)

	publisher := &logPublisher{l: zap.S()}

	client, err := NewNodeListenerClient(NodeBlockListenerConfig{
		WebSocketURL: baseNode,
		AuthHeader:   authHeader,
	}, publisher, NodeDataSource)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.ListenFlashBlocks(ctx)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
