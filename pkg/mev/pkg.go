package mev

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/core/types"
)

type BundleSenderType int

const (
	BundleSenderTypeFlashbot BundleSenderType = iota + 1
	BundleSenderTypeBeaver
	BundleSenderTypeRsync
	BundleSenderTypeTitan
	BundleSenderTypeBloxroute
	BundleSenderTypeAll
)

const (
	JSONRPC2                        = "2.0"
	SendBundleID                    = 1
	BloxrouteSubmitBundleMethod     = "blxr_submit_bundle"
	BloxrouteSimulationBundleMethod = "blxr_simulate_bundle"
	ETHSendBundleMethod             = "eth_sendBundle"
	EthCallBundleMethod             = "eth_callBundle"
	ETHCancelBundleMethod           = "eth_cancelBundle"
)

type IBundleSender interface {
	SendBundle(
		ctx context.Context,
		uuid *string,
		blockNumber uint64,
		tx ...*types.Transaction,
	) (SendBundleResponse, error)
	CancelBundle(
		ctx context.Context, bundleUUID string,
	) error
	SimulateBundle(ctx context.Context, blockNumber uint64, txs ...*types.Transaction) (SendBundleResponse, error)
	GetSenderType() BundleSenderType
}

var (
	_ IBundleSender = &Client{}
	_ IBundleSender = &BloxrouteClient{}
)

var defaultHeaders = [][2]string{ // nolint: gochecknoglobals
	{"Content-Type", "application/json"},
	{"Accept", "application/json"},
}

func txToRlp(tx *types.Transaction) string {
	var buff bytes.Buffer
	_ = tx.EncodeRLP(&buff)

	rlp := hex.EncodeToString(buff.Bytes())

	if rlp[:2] == "b9" {
		rlp = rlp[6:]
	} else if rlp[:2] == "b8" {
		rlp = rlp[4:]
	}

	return rlp
}

func doRequest[T any](c *http.Client, req *http.Request, headers ...[2]string) (T, error) {
	var t T

	for _, h := range defaultHeaders {
		req.Header.Add(h[0], h[1])
	}
	for _, h := range headers {
		req.Header.Add(h[0], h[1])
	}
	httpResp, err := c.Do(req)
	if err != nil {
		return t, fmt.Errorf("do request error: %w", err)
	}
	defer httpResp.Body.Close()

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return t, fmt.Errorf("read response error: %w", err)
	}

	if httpResp.StatusCode != http.StatusOK {
		return t, fmt.Errorf("not OK status, status: [%d], data: [%s]",
			httpResp.StatusCode, string(respBody))
	}

	if err := json.Unmarshal(respBody, &t); err != nil {
		return t, fmt.Errorf("unmarshal response error: %w, data: [%s]", err, string(respBody))
	}

	return t, nil
}

type SendBundleResponse struct {
	Jsonrpc string           `json:"jsonrpc,omitempty"`
	ID      int              `json:"id,omitempty"`
	Result  SendBundleResult `json:"result,omitempty"`
	Error   SendBundleError  `json:"error,omitempty"`
}

type SendBundleError struct {
	Code     int    `json:"code,omitempty"`
	Messange string `json:"message,omitempty"`
}

type SendBundleResult struct {
	BundleGasPrice    string              `json:"bundleGasPrice,omitempty"`
	BundleHash        string              `json:"bundleHash,omitempty"`
	CoinbaseDiff      string              `json:"coinbaseDiff,omitempty"`
	EthSentToCoinbase string              `json:"ethSentToCoinbase,omitempty"`
	GasFees           string              `json:"gasFees,omitempty"`
	Results           []SendBundleResults `json:"results,omitempty"`
	StateBlockNumber  int                 `json:"stateBlockNumber,omitempty"`
	TotalGasUsed      int                 `json:"totalGasUsed,omitempty"`
}

type SendBundleResults struct {
	GasUsed int    `json:"gasUsed,omitempty"`
	TxHash  string `json:"txHash,omitempty"`
	Value   string `json:"value,omitempty"`
}

type FlashbotCancelBundleResponse struct {
	Jsonrpc string          `json:"jsonrpc,omitempty"`
	ID      int             `json:"id,omitempty"`
	Result  []string        `json:"result,omitempty"`
	Error   SendBundleError `json:"error,omitempty"`
}

func (resp FlashbotCancelBundleResponse) ToSendBundleResponse() SendBundleResponse {
	r := SendBundleResponse{
		Jsonrpc: resp.Jsonrpc,
		ID:      resp.ID,
		Error:   resp.Error,
	}
	if len(resp.Result) != 0 {
		r.Result.BundleHash = resp.Result[0]
	}

	return r
}

type TitanCancelBundleResponse struct {
	Jsonrpc string          `json:"jsonrpc,omitempty"`
	ID      int             `json:"id,omitempty"`
	Result  int             `json:"result,omitempty"`
	Error   SendBundleError `json:"error,omitempty"`
}
