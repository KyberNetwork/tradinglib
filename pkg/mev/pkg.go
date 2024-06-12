package mev

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/flashbots/mev-share-node/mevshare"
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
	GetBundleStatsID                = 1
	SendBundleID                    = 1
	BloxrouteSubmitBundleMethod     = "blxr_submit_bundle"
	BloxrouteSimulationBundleMethod = "blxr_simulate_bundle"
	// FlashbotGetBundleStatsMethod
	// nolint: gosec
	FlashbotGetBundleStatsMethod = "flashbots_getBundleStats"
	ETHSendBundleMethod          = "eth_sendBundle"
	EthCallBundleMethod          = "eth_callBundle"
	ETHCancelBundleMethod        = "eth_cancelBundle"
	ETHEstimateGasBundleMethod   = "eth_estimateGasBundle"
	MevSendBundleMethod          = "mev_sendBundle"
	MaxBlockFromTarget           = 3
)

type IBundleSender interface {
	SendBundle(
		ctx context.Context,
		uuid *string,
		blockNumber uint64,
		tx ...*types.Transaction,
	) (SendBundleResponse, error)
	SendBackrunBundle(
		ctx context.Context,
		uuid *string,
		blockNumber uint64,
		pendingTxHash common.Hash,
		tx ...*types.Transaction,
	) (SendBundleResponse, error)
	CancelBundle(
		ctx context.Context, bundleUUID string,
	) error
	SimulateBundle(ctx context.Context, blockNumber uint64, txs ...*types.Transaction) (SendBundleResponse, error)
	EstimateBundleGas(
		ctx context.Context,
		messages []ethereum.CallMsg,
		overrides *map[common.Address]gethclient.OverrideAccount,
	) ([]uint64, error)
	MevSimulateBundle(
		blockNumber uint64,
		pendingTxHash common.Hash,
		tx *types.Transaction) (*mevshare.SimMevBundleResponse, error)
	GetSenderType() BundleSenderType
	GetBundleStats(
		ctx context.Context, blockNumber uint64, bundleHash common.Hash,
	) (GetBundleStatsResponse, error)
}

type IGasBundleEstimator interface {
	// EstimateBundleGas is used to estimate the gas for a bundle of transactions
	// Note that this method is expected only works with custom ethereum node which
	// supports estimate bundles gas via CallMsgs,
	// and using eth_estimateGasBundle method.
	EstimateBundleGas(
		ctx context.Context,
		messages []ethereum.CallMsg,
		overrides *map[common.Address]gethclient.OverrideAccount,
	) ([]uint64, error)
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

type GetBundleStatsResult struct {
	IsHighPriority bool      `json:"isHighPriority,omitempty"`
	IsSimulated    bool      `json:"isSimulated,omitempty"`
	SimulatedAt    time.Time `json:"simulatedAt,omitempty"`
	ReceivedAt     time.Time `json:"receivedAt,omitempty"`

	ConsideredByBuildersAt []*struct {
		Pubkey    string    `json:"pubkey,omitempty"`
		Timestamp time.Time `json:"timestamp"`
	} `json:"consideredByBuildersAt,omitempty"`
	SealedByBuildersAt []*struct {
		Pubkey    string    `json:"pubkey,omitempty"`
		Timestamp time.Time `json:"timestamp"`
	} `json:"sealedByBuildersAt,omitempty"`
}

type GetBundleStatsResponse struct {
	Jsonrpc string               `json:"jsonrpc,omitempty"`
	ID      int                  `json:"id,omitempty"`
	Result  GetBundleStatsResult `json:"result,omitempty"`
	Error   GetBundleStatsError  `json:"error,omitempty"`
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

type GetBundleStatsError struct {
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

func ToCallArg(msg ethereum.CallMsg) interface{} {
	arg := map[string]interface{}{
		"from": msg.From,
		"to":   msg.To,
	}
	if len(msg.Data) > 0 {
		arg["input"] = hexutil.Bytes(msg.Data)
	}
	if msg.Value != nil {
		arg["value"] = (*hexutil.Big)(msg.Value)
	}
	if msg.Gas != 0 {
		arg["gas"] = hexutil.Uint64(msg.Gas)
	}
	if msg.GasPrice != nil {
		arg["gasPrice"] = (*hexutil.Big)(msg.GasPrice)
	}
	return arg
}
