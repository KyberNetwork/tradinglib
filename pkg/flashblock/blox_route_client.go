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

type messageHandler func(ctx context.Context, data []byte)

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

// WebSocketParsedBdnFlashBlockResponse represents the WebSocket response
type WebSocketParsedBdnFlashBlockResponse struct {
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Subscription string                                `json:"subscription"`
		Result       *GetParsedBdnFlashBlockStreamResponse `json:"result"`
	} `json:"params"`
}

type WebSocketFlashBlockResponse struct {
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Subscription string                          `json:"subscription"`
		Result       *GetBdnFlashBlockStreamResponse `json:"result"`
	} `json:"params"`
}

// NewBloxRouteClient creates a new bloxroute client
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

func (c *Client) ListenParsedBdnFlashBlock(ctx context.Context) error {
	subscribeMsg := WebSocketMessage{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "subscribe",
		Params:  []interface{}{"GetParsedBdnFlashBlockStream", map[string]interface{}{}},
	}

	return c.Listen(ctx, subscribeMsg, c.processParsedFlashblock)
}

func (c *Client) ListenFlashBlock(ctx context.Context) error {
	subscribeMsg := WebSocketMessage{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "subscribe",
		Params:  []interface{}{"GetBdnFlashBlockStream", map[string]interface{}{}},
	}

	return c.Listen(ctx, subscribeMsg, c.processFlashBlock)
}

// Listen connects to the WebSocket stream and listens for bloxroute with automatic retry
func (c *Client) Listen(ctx context.Context, subscribeMsg interface{}, messageHandler messageHandler) error {
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
			err := c.connectAndListen(ctx, resetRetryWait, subscribeMsg, messageHandler)
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
// nolint: lll
func (c *Client) connectAndListen(ctx context.Context, resetRetryDelay func(), subscribeMsg interface{}, messageHandler messageHandler) error {
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
			_, data, err := conn.ReadMessage()
			if err != nil {
				c.l.Errorw("Error reading bloxroute flashblock", "error", err, "data", data)
				continue
			}

			messageHandler(ctx, data)
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

// processParsedFlashblock processes a received flashblock and publishes to subscribers
// nolint: lll
func (c *Client) processParsedFlashblock(ctx context.Context, data []byte) {
	var response WebSocketParsedBdnFlashBlockResponse
	if err := json.Unmarshal(data, &response); err != nil {
		c.l.Errorw("Error unmarshaling bloxroute flashblock", "error", err, "data", string(data))
	}

	if response.Params.Result == nil {
		// First message is subscription confirmation
		c.l.Debugw("Received subscription confirmation")
		return
	}

	if response.Params.Result.Metadata == nil {
		return
	}

	blockNumber := response.Params.Result.Metadata.BlockNumber
	c.l.Debugw("Processing flashblock", "blockNumber", blockNumber, "index", response.Params.Result.Index)

	flashBlock, err := convertBloxRouteFlashBlock(*response.Params.Result)
	if err != nil {
		c.l.Errorw("Error converting bloxroute flashblock", "error", err, "blockNumber", blockNumber, "index", response.Params.Result.Index)
		return
	}
	// PublishFlashBlock the flashblock data to subscribers
	if c.publisher != nil {
		if err := c.publisher.PublishFlashBlock(ctx, BloxRouteDataSource, flashBlock); err != nil {
			c.l.Errorw("Failed to publish flashblock", "error", err, "blockNumber", blockNumber)
		}
	}
}

func convertBloxRouteFlashBlock(bloxrouteFlashblock GetParsedBdnFlashBlockStreamResponse) (Flashblock, error) {
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

func (c *Client) processFlashBlock(ctx context.Context, data []byte) {
	var response WebSocketFlashBlockResponse
	if err := json.Unmarshal(data, &response); err != nil {
		c.l.Errorw("Error unmarshaling bloxroute flashblock", "error", err, "data", string(data))
	}

	if response.Params.Result == nil {
		// First message is subscription confirmation
		c.l.Debugw("Received subscription confirmation")
		return
	}

	jsonBytes, err := DecompressBrotli(response.Params.Result.BdnFlashBlock)
	if err != nil {
		c.l.Errorw("Error decompressing flashblock", "error", err)
		return
	}

	var flashBlock Flashblock
	if err := json.Unmarshal(jsonBytes, &flashBlock); err != nil {
		c.l.Errorw("Error parsing flashblock", "error", err, "data", string(jsonBytes))
		return
	}

	if err := c.publisher.PublishFlashBlock(ctx, BloxRouteDataSource, flashBlock); err != nil {
		c.l.Errorw("Error publishing flashblock", "error", err)
	}
}
