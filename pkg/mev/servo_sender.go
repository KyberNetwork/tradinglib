package mev

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/flashbots/mev-share-node/mevshare"
)

// XAPIKeyHeader authenticates every SERVO request; SERVO signs nothing.
// nolint: gosec // header name, not a credential
const XAPIKeyHeader = "X-Api-Key"

// nolint: gochecknoglobals
var (
	ErrMissingAPIKey      = fmt.Errorf("missing api key")
	ErrMissingEndpoint    = fmt.Errorf("missing endpoint")
	ErrMissingPendingTxs  = fmt.Errorf("missing pending tx hashes")
	ErrMissingBlockNumber = fmt.Errorf("missing block number")
)

// ServoSender bids into the SERVO 2.0 order-flow auction. A bid is an eth_sendBundle
// whose txs are the victim hashes in arrival order followed by our raw backrun, and
// whose value to SERVO is the backrun's PRIORITY FEE (not a coinbase transfer).
type ServoSender struct {
	c        *http.Client
	endpoint string
	apiKey   string
}

func NewServoSender(c *http.Client, endpoint, apiKey string) (*ServoSender, error) {
	if endpoint == "" {
		return nil, ErrMissingEndpoint
	}
	if apiKey == "" {
		return nil, ErrMissingAPIKey
	}
	return &ServoSender{c: c, endpoint: endpoint, apiKey: apiKey}, nil
}

// servoBundleParams is the exact param object SERVO documents. It is deliberately
// not SendBundleParams: that struct emits `ReplacementUuid` (capital R) and always
// emits allowBuilderNetRefunds, neither of which SERVO specifies.
type servoBundleParams struct {
	Txs             []string `json:"txs"`
	BlockNumber     string   `json:"blockNumber"`
	ReplacementUUID string   `json:"replacementUuid,omitempty"`
}

// SendBackrunBundle submits one bid. blockNumber/maxBlockNumber are collapsed into
// SERVO's single blockNumber, which is a DEADLINE, not a target: SERVO retries the
// bid every block up to it. Omitting it would let SERVO retry for 5 blocks, so a
// statically-priced backrun is deliberately never sent without one.
// targetBuilders and coinbaseProfit are ignored — SERVO picks the builders.
func (s *ServoSender) SendBackrunBundle(
	ctx context.Context,
	uuid *string,
	blockNumber uint64,
	maxBlockNumber uint64,
	pendingTxHashes []common.Hash,
	_ []string,
	_ *big.Int,
	txs ...*types.Transaction,
) (SendBundleResponse, error) {
	if len(pendingTxHashes) == 0 {
		return SendBundleResponse{}, ErrMissingPendingTxs
	}
	if len(txs) == 0 {
		return SendBundleResponse{}, ErrInvalidLenTx
	}
	deadline := max(maxBlockNumber, blockNumber)
	if deadline == 0 {
		return SendBundleResponse{}, ErrMissingBlockNumber
	}

	bundleTxs := make([]string, 0, len(pendingTxHashes)+len(txs))
	for _, h := range pendingTxHashes {
		bundleTxs = append(bundleTxs, h.Hex())
	}
	for _, tx := range txs {
		raw, err := tx.MarshalBinary()
		if err != nil {
			return SendBundleResponse{}, fmt.Errorf("marshal tx: %w", err)
		}
		bundleTxs = append(bundleTxs, hexutil.Encode(raw))
	}

	params := servoBundleParams{Txs: bundleTxs, BlockNumber: hexutil.EncodeUint64(deadline)}
	if uuid != nil {
		params.ReplacementUUID = *uuid
	}

	reqBody, err := json.Marshal(SendRequest{
		ID:      SendBundleID,
		JSONRPC: JSONRPC2,
		Method:  ETHSendBundleMethod,
		Params:  []any{params},
	})
	if err != nil {
		return SendBundleResponse{}, fmt.Errorf("marshal json error: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, s.endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return SendBundleResponse{}, fmt.Errorf("new http request error: %w", err)
	}

	resp, err := doRequest[SendBundleResponse](s.c, httpReq, [2]string{XAPIKeyHeader, s.apiKey})
	if err != nil {
		return SendBundleResponse{}, err
	}
	if len(resp.Error.Messange) != 0 {
		return SendBundleResponse{}, fmt.Errorf("response error, code: [%d], message: [%s]",
			resp.Error.Code, resp.Error.Messange)
	}

	return resp, nil
}

func (s *ServoSender) MevSimulateBundle(
	_ context.Context, _ uint64, _ common.Hash, _ *types.Transaction,
) (*mevshare.SimMevBundleResponse, error) {
	return nil, ErrMethodNotSupport
}

func (s *ServoSender) GetSenderType() BundleSenderType {
	return BundleSenderTypeServo
}
