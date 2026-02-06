package flashblock

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// Config holds the configuration for the flashblock client
type Config struct {
	WebSocketURL string
	AuthHeader   string
}

// validate validates the configuration
func (c *Config) validate() error {
	if c.WebSocketURL == "" {
		return fmt.Errorf("WebSocketURL is required")
	}
	if c.AuthHeader == "" {
		return fmt.Errorf("AuthHeader is required")
	}

	return nil
}

// Client represents a WebSocket client for bloXroute Base streamer
type Client struct {
	config    Config
	l         *zap.SugaredLogger
	conn      *websocket.Conn
	publisher Publisher
}

// WebSocketMessage represents the WebSocket subscription message
type WebSocketMessage struct {
	JSONRPC string        `json:"jsonrpc"`
	ID      int           `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

// WebSocketResponse represents the WebSocket response
type WebSocketResponse struct {
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Subscription string                                `json:"subscription"`
		Result       *GetParsedBdnFlashBlockStreamResponse `json:"result"`
	} `json:"params"`
}

// NewNodeListenerClient creates a new bloxroute client
func NewBloxRouteClient(config Config, publisher Publisher) (*Client, error) {
	if err := config.validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	return &Client{
		config:    config,
		l:         zap.S().Named("blox-route-client"),
		publisher: publisher,
	}, nil
}

// Start starts the bloxroute stream listener
func (c *Client) Start(ctx context.Context) error {
	if err := c.Listen(ctx); err != nil {
		c.l.Errorw("Failed to listen for bloxroute", "error", err)
	}

	return nil
}

// Listen connects to the WebSocket stream and listens for bloxroute with automatic retry
func (c *Client) Listen(ctx context.Context) error {
	c.l.Infow("Starting bloxroute block listener with retry", "url", c.config.WebSocketURL)

	retryWait := 3 * time.Second
	retryCount := 0

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
			if err == nil {
				return nil // Normal exit
			}

			// Check if error is retryable
			if !c.isRetryableError(err) {
				c.l.Errorw("Non-retryable error, stopping flashblock listener", "error", err)
				return err
			}

			// Log retry attempt
			retryCount++
			c.l.Warnw("Flashblock connection error, retrying",
				"error", err,
				"retryCount", retryCount,
				"retryWait", retryWait)

			// Wait before retry
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(retryWait):
				// Exponential backoff with max delay
				retryWait *= 2
				if retryWait > maxRetryWait {
					retryWait = maxRetryWait
				}
			}
		}
	}
}

// connectAndListen establishes connection and listens for bloxroute
func (c *Client) connectAndListen(ctx context.Context, resetRetryDelay func()) error {
	c.l.Infow("Connecting to flashblock stream", "url", c.config.WebSocketURL)

	// Create WebSocket connection with authorization header
	header := make(map[string][]string)
	header["Authorization"] = []string{c.config.AuthHeader}

	conn, _, err := websocket.DefaultDialer.Dial(c.config.WebSocketURL, header)
	if err != nil {
		return fmt.Errorf("failed to connect to WebSocket: %w", err)
	}
	c.conn = conn
	defer conn.Close()

	// Subscribe to GetParsedBdnFlashBlockStream
	subscribeMsg := WebSocketMessage{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "subscribe",
		Params:  []interface{}{"GetParsedBdnFlashBlockStream", map[string]interface{}{}},
	}

	if err := conn.WriteJSON(subscribeMsg); err != nil {
		return fmt.Errorf("failed to send subscription message: %w", err)
	}

	c.l.Info("Connected to flashblock stream, listening for events...")

	// Reset retry delay after successful connection
	resetRetryDelay()

	// PING PONG
	conn.SetPingHandler(func(appData string) error {
		c.l.Debugw("Ping received", "data", appData)
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

	// Listen for bloxroute
	for {
		select {
		case <-ctx.Done():
			c.l.Info("Context cancelled, stopping flashblock listener")
			return ctx.Err()
		default:
			var response WebSocketResponse
			_, data, err := conn.ReadMessage()
			if err != nil {
				c.l.Errorw("Error reading bloxroute flashblock", "error", err, "data", data)
				continue
			}

			if err := json.Unmarshal(data, &response); err != nil {
				c.l.Errorw("Error unmarshaling bloxroute flashblock", "error", err, "data", string(data))
			}

			if response.Params.Result == nil {
				// First message is subscription confirmation
				c.l.Debugw("Received subscription confirmation")
				continue
			}

			c.processFlashblock(ctx, *response.Params.Result)
		}
	}
}

// isRetryableError checks if an error should trigger a retry
func (c *Client) isRetryableError(err error) bool {
	if err == nil {
		return false
	}

	// Don't retry on context cancellation
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return false
	}

	// Retry on WebSocket close errors
	if websocket.IsCloseError(err,
		websocket.CloseAbnormalClosure,
		websocket.CloseNormalClosure,
		websocket.CloseServiceRestart,
		websocket.CloseGoingAway) {
		return true
	}

	// Retry on connection reset
	if errors.Is(err, syscall.ECONNRESET) {
		return true
	}

	// Retry on connection refused
	if errors.Is(err, syscall.ECONNREFUSED) {
		return true
	}

	// Default to retry for other errors
	return true
}

// processFlashblock processes a received flashblock and publishes to subscribers
func (c *Client) processFlashblock(ctx context.Context, response GetParsedBdnFlashBlockStreamResponse) {
	if response.Metadata == nil {
		return
	}

	blockNumber := response.Metadata.BlockNumber
	c.l.Debugw("Processing flashblock", "blockNumber", blockNumber, "index", response.Index)

	flashBlock, err := convertBlockRouteFlashBlock(response)
	if err != nil {
		c.l.Errorw("Error converting bloxroute flashblock", "error", err, "blockNumber", blockNumber, "index", response.Index)
		return
	}
	// Publish the flashblock data to subscribers
	if c.publisher != nil {
		if err := c.publisher.Publish(ctx, BloxRouteDataSource, flashBlock); err != nil {
			c.l.Errorw("Failed to publish flashblock", "error", err, "blockNumber", blockNumber)
		}
	}
}

func convertBlockRouteFlashBlock(bloxrouteFlashblock GetParsedBdnFlashBlockStreamResponse) (Flashblock, error) {
	index, err := strconv.ParseInt(bloxrouteFlashblock.Index, 10, 64)
	if err != nil {
		return Flashblock{}, fmt.Errorf("invalid index: %w value %v", err, index)
	}
	return Flashblock{
		PayloadID: bloxrouteFlashblock.PayloadId,
		Index:     index,
		Base:      convertBloxRouteBaseToFlashblockBase(bloxrouteFlashblock.Base),
		Diff:      convertBloxRouteDiffToFlashblockDiff(bloxrouteFlashblock.Diff),
		Metadata:  convertBloxRouteMetadataToFlashblockMeta(bloxrouteFlashblock.Metadata),
	}, nil
}
