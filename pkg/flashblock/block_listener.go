package flashblock

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

const (
	maxRetryWait = 30 * time.Second
)

type NodeBlockListenerConfig struct {
	WebSocketURL string
	AuthHeader   string
}

func (c NodeBlockListenerConfig) validate() error {
	if c.WebSocketURL == "" || (!strings.HasPrefix(c.WebSocketURL, "ws://") &&
		!strings.HasPrefix(c.WebSocketURL, "wss://")) {
		return fmt.Errorf("invalid WebSocketURL: %s", c.WebSocketURL)
	}
	return nil
}

type NodeClient struct {
	config     NodeBlockListenerConfig
	l          *zap.SugaredLogger
	conn       *websocket.Conn
	publisher  Publisher
	dataSource DataSource
}

// nolint:lll
func NewNodeListenerClient(config NodeBlockListenerConfig, publisher Publisher, dataSource DataSource) (*NodeClient, error) {
	if err := config.validate(); err != nil {
		return nil, err
	}

	return &NodeClient{
		config:     config,
		l:          zap.S().Named("block-listener"),
		publisher:  publisher,
		dataSource: dataSource,
	}, nil
}

func (c *NodeClient) ListenFlashBlocks(ctx context.Context) error {
	retryWait := time.Second

	resetRetryWait := func() {
		retryWait = time.Second
	}

	for {
		select {
		case <-ctx.Done():
			c.l.Info("Context cancelled, stopping flashblock listener")
			return ctx.Err()
		default:
			err := c.connectAndListen(ctx, resetRetryWait)
			if err != nil {
				c.l.Errorw("Flashblock listener error", "error", err)
			}

			if err == nil {
				return nil
			}

			c.l.Warnw("Retrying in", "wait", retryWait)
			time.Sleep(retryWait)
			retryWait *= 2
			if retryWait > maxRetryWait {
				retryWait = maxRetryWait
			}
		}
	}
}

// nolint:cyclop
func (c *NodeClient) connectAndListen(ctx context.Context, resetRetryDelay func()) error {
	c.l.Infow("Connecting to flashblock stream", "url", c.config.WebSocketURL)

	header := make(map[string][]string)
	if c.config.AuthHeader != "" {
		header["Authorization"] = []string{c.config.AuthHeader}
	}

	conn, _, err := websocket.DefaultDialer.Dial(c.config.WebSocketURL, header)
	if err != nil {
		return fmt.Errorf("failed to connect to WebSocket: %w", err)
	}
	c.conn = conn
	defer conn.Close()

	c.l.Info("Connected to flashblock stream, listening for events...")

	// Reset retry delay after successful connection
	resetRetryDelay()

	// PING PONG
	conn.SetPingHandler(func(appData string) error {
		c.l.Debugw("Ping received", "data", appData)
		// nolint: errcheck
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		return conn.WriteMessage(websocket.PongMessage, []byte{})
	})

	// Create a context for the ping goroutine that will be cancelled when this function exits
	pingCtx, cancelPing := context.WithCancel(ctx)
	defer cancelPing()

	go func() {
		pingTicker := time.NewTicker(5 * time.Second)
		defer pingTicker.Stop()
		for {
			select {
			case <-pingCtx.Done():
				return
			case <-pingTicker.C:
				if err := conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
					c.l.Errorw("Ping error", "error", err)
					_ = conn.Close()
					return
				}
			}
		}
	}()

	// Listen for flashblocks
	for {
		select {
		case <-ctx.Done():
			c.l.Info("Context cancelled, stopping flashblock listener")
			return nil
		default:
			// nolint:errcheck
			conn.SetReadDeadline(time.Now().Add(5 * time.Second))
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				c.l.Errorw("Error reading node flashblock", "error", err)
				return fmt.Errorf("stream error: %w", err)
			}

			if messageType != websocket.BinaryMessage && messageType != websocket.TextMessage {
				c.l.Info("Ignoring non-binary/text message", "type", messageType, "message", string(message))
				continue
			}

			jsonBytes, err := DecompressBrotli(message)
			if err != nil {
				c.l.Errorw("Error decompressing flashblock", "error", err)
			}

			var flashBlock Flashblock
			if err := json.Unmarshal(jsonBytes, &flashBlock); err != nil {
				c.l.Errorw("Error parsing flashblock", "error", err, "data", string(jsonBytes))
				continue
			}

			if err := c.publisher.Publish(ctx, c.dataSource, flashBlock); err != nil {
				c.l.Errorw("Error publishing flashblock", "error", err)
			}
		}
	}
}
