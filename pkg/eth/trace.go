package eth

import (
	"cmp"
	"context"
	"fmt"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
)

type TraceClient struct {
	rpcClient *ethrpc.Client
}

func NewTraceClient(ctx context.Context, httpClient *http.Client, rpcURL string) (*TraceClient, error) {
	rpcClient, err := ethrpc.DialOptions(ctx, rpcURL, ethrpc.WithHTTPClient(httpClient))
	if err != nil {
		return nil, fmt.Errorf("dial options: %w", err)
	}

	return &TraceClient{
		rpcClient: rpcClient,
	}, nil
}

func (c *TraceClient) DebugTraceTransaction(ctx context.Context, txHash string) (CallFrame, error) {
	const (
		method = "debug_traceTransaction"
		tracer = "callTracer"
	)
	var result CallFrame
	if err := c.rpcClient.CallContext(ctx, &result,
		method,
		txHash,
		map[string]interface{}{
			"tracer": tracer,
			"tracerConfig": map[string]interface{}{
				"withLog": true,
			},
		},
	); err != nil {
		return CallFrame{}, fmt.Errorf("call context: %w", err)
	}

	if len(result.Error) != 0 {
		return CallFrame{}, fmt.Errorf("error response: %s, reason: %s",
			result.Error, result.RevertReason)
	}

	return result, nil
}

func (c *TraceClient) DebugTraceCall(
	ctx context.Context,
	from, to string,
	gas uint64,
	gasPrice *big.Int,
	value *big.Int,
	encodedData string,
	block *big.Int,
) (CallFrame, error) {
	const (
		method = "debug_traceCall"
		tracer = "callTracer"
	)

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

	var result CallFrame
	if err := c.rpcClient.CallContext(ctx, &result,
		method,
		paramData,
		blockStr,
		map[string]any{
			"tracer": tracer,
			"tracerConfig": map[string]interface{}{
				"onlyTopCall": false,
				"withLog":     true,
			},
		},
	); err != nil {
		return CallFrame{}, fmt.Errorf("call context: %w", err)
	}

	if len(result.Error) != 0 {
		return CallFrame{}, fmt.Errorf("error response: %s, reason: %s",
			result.Error, result.RevertReason)
	}

	return result, nil
}

type CallLog struct {
	Address common.Address `json:"address"`
	Topics  []common.Hash  `json:"topics"`
	Data    string         `json:"data"`
}

func (l CallLog) ToEthereumLog() ethtypes.Log {
	return ethtypes.Log{
		Address: l.Address,
		Topics:  l.Topics,
		Data:    common.Hex2Bytes(l.Data),
	}
}

type CallFrame struct {
	From         string      `json:"from"`
	Gas          string      `json:"gas"`
	GasUsed      string      `json:"gasUsed"`
	To           string      `json:"to"`
	Input        string      `json:"input"`
	Output       string      `json:"output"`
	Calls        []CallFrame `json:"calls"`
	Value        string      `json:"value"`
	Type         string      `json:"type"`
	Logs         []CallLog   `json:"logs"`
	Error        string      `json:"error"`
	RevertReason string      `json:"revertReason"`
}
