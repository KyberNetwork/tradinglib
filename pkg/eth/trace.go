package eth

import (
	"bytes"
	"cmp"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

type TraceClient struct {
	httpClient *http.Client
	rpcURL     string
}

func NewTraceClient(httpClient *http.Client, rpcURL string) *TraceClient {
	return &TraceClient{
		httpClient: httpClient,
		rpcURL:     rpcURL,
	}
}

func (c *TraceClient) DebugTraceTransaction(txHash string) (CallFrame, error) {
	payload := map[string]interface{}{
		"method":  "debug_traceTransaction",
		"id":      1,
		"jsonrpc": "2.0",
		"params": []interface{}{
			txHash,
			map[string]interface{}{
				"tracer": "callTracer",
				"tracerConfig": map[string]interface{}{
					"withLog": true,
				},
			},
		},
	}

	var resp CommomTraceResponse[CallFrame]
	if err := c.post(payload, &resp); err != nil {
		return CallFrame{}, fmt.Errorf("post error: %w", err)
	}
	if resp.Error.Code != 0 {
		return CallFrame{}, fmt.Errorf("error response: code: %d, message: %s",
			resp.Error.Code, resp.Error.Message)
	}

	return resp.Result, nil
}

func (c *TraceClient) DebugTraceCall(
	from, to string,
	gas uint64,
	gasPrice *big.Int,
	value *big.Int,
	encodedData string,
	block *big.Int,
) (CallFrame, error) {
	paramData := map[string]any{
		"from": cmp.Or(from, "null"),
		"to":   to,
		"data": cmp.Or(encodedData, "null"),
	}
	if gas != 0 {
		paramData["gas"] = hexutil.EncodeUint64(gas)
	}
	if gasPrice != nil {
		paramData["gasPrice"] = hexutil.EncodeBig(gasPrice)
	}
	if value != nil {
		paramData["value"] = hexutil.EncodeBig(value)
	}
	blockStr := "latest"
	if block != nil {
		blockStr = hexutil.EncodeBig(block)
	}

	payload := map[string]any{
		"method":  "debug_traceCall",
		"id":      1,
		"jsonrpc": "2.0",
		"params": []any{
			paramData,
			blockStr,
			map[string]any{
				"tracer": "callTracer",
				"tracerConfig": map[string]interface{}{
					"onlyTopCall": false,
					"withLog":     true,
				},
			},
		},
	}

	var resp CommomTraceResponse[CallFrame]
	if err := c.post(payload, &resp); err != nil {
		return CallFrame{}, fmt.Errorf("post error: %w", err)
	}

	if resp.Error.Code != 0 {
		return CallFrame{}, fmt.Errorf("error response: code: %d, message: %s",
			resp.Error.Code, resp.Error.Message)
	}

	return resp.Result, nil
}

func (c *TraceClient) post(payload any, expect any) error {
	var body *bytes.Buffer
	if payload != nil {
		b, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("marshal error: %w", err)
		}
		body = bytes.NewBuffer(b)
	}

	req, err := http.NewRequest(http.MethodPost, c.rpcURL, body)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	text, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("not status OK, status: %d", resp.StatusCode)
	}

	if expect != nil {
		if err := json.Unmarshal(text, expect); err != nil {
			return fmt.Errorf("unmarshal error: %w, data: [%s]", err, text)
		}
	}

	return nil
}

type CommomTraceResponse[T any] struct {
	JSONRPC string        `json:"jsonrpc"`
	ID      int           `json:"id"`
	Result  T             `json:"result"`
	Error   ErrorResponse `json:"error"`
}

type CallLog struct {
	Address common.Address `json:"address"`
	Topics  []common.Hash  `json:"topics"`
	Data    string         `json:"data"`
}

func (l CallLog) ToEthereumLog() (types.Log, error) {
	dataBytes, err := hexutil.Decode(l.Data)
	if err != nil {
		return types.Log{}, fmt.Errorf("decode error: %w", err)
	}

	return types.Log{
		Address: l.Address,
		Topics:  l.Topics,
		Data:    dataBytes,
	}, nil
}

type CallFrame struct {
	From    string      `json:"from"`
	Gas     string      `json:"gas"`
	GasUsed string      `json:"gasUsed"`
	To      string      `json:"to"`
	Input   string      `json:"input"`
	Output  string      `json:"output"`
	Calls   []CallFrame `json:"calls"`
	Value   string      `json:"value"`
	Type    string      `json:"type"`
	Logs    []CallLog   `json:"logs"`
}

type ErrorResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}
