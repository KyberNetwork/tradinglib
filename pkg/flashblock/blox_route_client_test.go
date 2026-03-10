package flashblock

import (
	"context"
	"os"
	"testing"
	"time"

	"go.uber.org/zap"
)

var (
	bloxRouteWsUrl = "wss://base.blxrbdn.com:5005/ws"
	authHeader     = os.Getenv("AUTH_HEADER")
)

type logPublisher struct {
	l *zap.SugaredLogger
}

func (p *logPublisher) Publish(_ context.Context, source DataSource, data Flashblock) error {
	p.l.Infow("received flashblock", "source", source, "payloadID", data.PayloadID, "index", data.Index, "blockNumber", data.Metadata.BlockNumber)
	return nil
}

func TestListenFlashBlock(t *testing.T) {
	if authHeader == "" {
		t.Skip("AUTH_HEADER env not set")
	}

	logger, _ := zap.NewDevelopment()
	defer logger.Sync() //nolint:errcheck
	zap.ReplaceGlobals(logger)

	publisher := &logPublisher{l: zap.S()}

	client, err := NewBloxRouteClient(Config{
		WebSocketURL: bloxRouteWsUrl,
		AuthHeader:   authHeader,
	}, publisher)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.ListenFlashBlock(ctx)
	if err != nil && err != context.DeadlineExceeded {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestListenParsedBdnFlashBlock(t *testing.T) {
	if authHeader == "" {
		t.Skip("AUTH_HEADER env not set")
	}

	logger, _ := zap.NewDevelopment()
	defer logger.Sync() //nolint:errcheck
	zap.ReplaceGlobals(logger)

	publisher := &logPublisher{l: zap.S()}

	client, err := NewBloxRouteClient(Config{
		WebSocketURL: bloxRouteWsUrl,
		AuthHeader:   authHeader,
	}, publisher)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.ListenParsedBdnFlashBlock(ctx)
	if err != nil && err != context.DeadlineExceeded {
		t.Errorf("unexpected error: %v", err)
	}
}
