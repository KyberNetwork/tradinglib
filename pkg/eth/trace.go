package eth

import (
	"bytes"
	"cmp"
	"encoding/json"
	"io"
	"math/big"
	"net/http"

	"github.com/duoxehyon/mev-share-go/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

	b, err := json.Marshal(payload)
	if err != nil {
		return CallFrame{}, err
	}

	req, err := http.NewRequest(http.MethodPost, c.rpcURL, bytes.NewBuffer(b))
	if err != nil {
		return CallFrame{}, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return CallFrame{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return CallFrame{}, err
	}

	var rpcResponse CommomTraceResponse[CallFrame]
	if err := json.Unmarshal(body, &rpcResponse); err != nil {
		return CallFrame{}, err
	}

	return rpcResponse.Result, err
}

type CommomTraceResponse[T any] struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  T      `json:"result"`
}

type CallLog struct {
	Address common.Address `json:"address"`
	Topics  []common.Hash  `json:"topics"`
	Data    string         `json:"data"`
}

func (l CallLog) ToEthereumLog() types.Log {
	return types.Log{
		Address: l.Address,
		Topics:  l.Topics,
		Data:    common.Hex2Bytes(l.Data),
	}
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

	b, err := json.Marshal(payload)
	if err != nil {
		return CallFrame{}, err
	}

	req, err := http.NewRequest(http.MethodPost, c.rpcURL, bytes.NewBuffer(b))
	if err != nil {
		return CallFrame{}, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return CallFrame{}, err
	}
	defer res.Body.Close()

	text, err := io.ReadAll(res.Body)
	if err != nil {
		return CallFrame{}, err
	}

	var rpcResponse CommomTraceResponse[CallFrame]
	if err := json.Unmarshal(text, &rpcResponse); err != nil {
		return CallFrame{}, err
	}

	return rpcResponse.Result, err
}
