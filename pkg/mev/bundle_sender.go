package mev

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// Client https://beaverbuild.org/docs.html; https://rsync-builder.xyz/docs;
// https://docs.flashbots.net/flashbots-auction/advanced/rpc-endpoint#eth_sendbundle
type Client struct {
	c                    *http.Client
	endpoint             string
	flashbotKey          *ecdsa.PrivateKey
	cancelBySendBundle   bool
	senderType           BundleSenderType
	enableSendPrivateRaw bool
}

// NewClient set the flashbotKey to nil will skip adding the signature header.
func NewClient(
	c *http.Client,
	endpoint string,
	flashbotKey *ecdsa.PrivateKey,
	senderType BundleSenderType,
	cancelBySendBundle bool,
	enableSendPrivateRaw bool,
) (*Client, error) {
	return &Client{
		c:                    c,
		endpoint:             endpoint,
		flashbotKey:          flashbotKey,
		cancelBySendBundle:   cancelBySendBundle,
		senderType:           senderType,
		enableSendPrivateRaw: enableSendPrivateRaw,
	}, nil
}

func (s *Client) GetSenderType() BundleSenderType {
	return s.senderType
}

func (s *Client) SendBundle(
	ctx context.Context,
	uuid *string,
	blockNumber uint64,
	txs ...*types.Transaction,
) (SendBundleResponse, error) {
	return s.sendBundle(ctx, ETHSendBundleMethod, uuid, blockNumber, txs, nil)
}

func (s *Client) SendBundleV2(
	ctx context.Context,
	req SendBundleV2Request,
	txs ...*types.Transaction,
) (SendBundleResponse, error) {
	p := new(SendBundleParams).
		SetTransactions(txs...)

	if req.UUID != nil {
		p.SetUUID(*req.UUID, s.senderType)
	}
	if req.BlockNumber != nil {
		p.SetBlockNumber(*req.BlockNumber)
	}
	if req.MinTimestamp != nil {
		p.MinTimestamp = req.MinTimestamp
	}
	if req.MaxTimestamp != nil {
		p.MaxTimestamp = req.MaxTimestamp
	}
	if req.RevertingTxs != nil {
		p.RevertingTxs = req.RevertingTxs
	}
	if s.senderType == BundleSenderTypeFlashbot {
		p = p.SetStateBlockNumber("latest")
	}

	if err := p.Err(); err != nil {
		return SendBundleResponse{}, err
	}

	return s.sendRawBundle(ctx, ETHSendBundleMethod, p)
}

func (s *Client) SendBundleHex(
	ctx context.Context,
	uuid *string,
	blockNumber uint64,
	hexEncodedTxs ...string,
) (SendBundleResponse, error) {
	return s.sendBundle(ctx, ETHSendBundleMethod, uuid, blockNumber, nil, hexEncodedTxs)
}

// getGetBundleStatsMethod
// nolint: unparam
func (s *Client) getGetBundleStatsMethod() string {
	switch s.senderType {
	case BundleSenderTypeFlashbot:
		return FlashbotGetBundleStatsMethod
	default:
		return FlashbotGetBundleStatsMethod
	}
}

func (s *Client) CancelBundle(
	ctx context.Context, bundleUUID string,
) error {
	if s.cancelBySendBundle {
		if _, err := s.sendBundle(ctx, ETHSendBundleMethod, &bundleUUID, 0, nil, nil); err != nil {
			return fmt.Errorf("cancel by send bundle error: %w", err)
		}
		return nil
	}

	// build request
	p := CancelBundleParams{
		ReplacementUUID: bundleUUID,
	}
	req := SendRequest{
		ID:      SendBundleID,
		JSONRPC: JSONRPC2,
		Method:  ETHCancelBundleMethod,
		Params:  []any{p},
	}
	reqBody, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("marshal json error: %w", err)
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, s.endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("new http request error: %w", err)
	}

	var headers [][2]string
	if s.flashbotKey != nil {
		signature, err := requestSignature(s.flashbotKey, reqBody)
		if err != nil {
			return fmt.Errorf("sign flashbot request error: %w", err)
		}

		headers = append(headers, [2]string{"X-Flashbots-Signature", signature})
	}

	// do
	var errResp ErrorResponse
	switch s.senderType {
	case BundleSenderTypeFlashbot:
		resp, err := doRequest[FlashbotCancelBundleResponse](s.c, httpReq, headers...)
		if err != nil {
			return err
		}
		errResp = resp.Error
	case BundleSenderTypeTitan:
		resp, err := doRequest[TitanCancelBundleResponse](s.c, httpReq, headers...)
		if err != nil {
			return err
		}
		errResp = resp.Error
	default:
		resp, err := doRequest[SendBundleResponse](s.c, httpReq, headers...)
		if err != nil {
			return err
		}
		errResp = resp.Error
	}

	// check
	if len(errResp.Messange) != 0 {
		return fmt.Errorf("response error, code: [%d], message: [%s]", errResp.Code, errResp.Messange)
	}

	return nil
}

func (s *Client) SimulateBundle(
	ctx context.Context, blockNumber uint64, txs ...*types.Transaction,
) (SendBundleResponse, error) {
	return s.sendBundle(ctx, EthCallBundleMethod, nil, blockNumber, txs, nil)
}

func (s *Client) GetBundleStats(
	ctx context.Context, blockNumber uint64, bundleHash common.Hash,
) (GetBundleStatsResponse, error) {
	req := GetBundleStatsRequest{
		ID:      GetBundleStatsID,
		JSONRPC: JSONRPC2,
		Method:  s.getGetBundleStatsMethod(),
	}
	p := new(GetBundleStatsParams).SetBlockNumber(blockNumber).SetBundleHash(bundleHash)
	req.Params = append(req.Params, p)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return GetBundleStatsResponse{}, fmt.Errorf("marshal json error: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, s.endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return GetBundleStatsResponse{}, fmt.Errorf("new http request error: %w", err)
	}

	var headers [][2]string
	if s.flashbotKey != nil {
		signature, err := requestSignature(s.flashbotKey, reqBody)
		if err != nil {
			return GetBundleStatsResponse{}, fmt.Errorf("sign flashbot request error: %w", err)
		}
		headers = append(headers, [2]string{"X-Flashbots-Signature", signature})
	}

	resp, err := doRequest[GetBundleStatsResponse](s.c, httpReq, headers...)
	if err != nil {
		return GetBundleStatsResponse{}, err
	}

	return resp, nil
}

func (s *Client) sendBundle(
	ctx context.Context,
	method string,
	uuid *string,
	blockNumber uint64,
	txs []*types.Transaction,
	hexEncodedTxs []string,
) (SendBundleResponse, error) {
	p := new(SendBundleParams).
		SetBlockNumber(blockNumber).
		SetTransactions(txs...).
		SetTransactionsHex(hexEncodedTxs...)
	if s.senderType == BundleSenderTypeFlashbot {
		p = p.SetStateBlockNumber("latest")
	}
	if err := p.Err(); err != nil {
		return SendBundleResponse{}, err
	}

	if uuid != nil {
		p.SetUUID(*uuid, s.senderType)
	}

	return s.sendRawBundle(ctx, method, p)
}

func (s *Client) sendRawBundle(ctx context.Context, method string, p *SendBundleParams) (SendBundleResponse, error) {
	req := SendRequest{
		ID:      SendBundleID,
		JSONRPC: JSONRPC2,
		Method:  method,
	}

	req.Params = append(req.Params, p)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return SendBundleResponse{}, fmt.Errorf("marshal json error: %w", err)
	}

	var headers [][2]string
	if s.flashbotKey != nil {
		signature, err := requestSignature(s.flashbotKey, reqBody)
		if err != nil {
			return SendBundleResponse{}, fmt.Errorf("sign flashbot request error: %w", err)
		}
		headers = append(headers, [2]string{"X-Flashbots-Signature", signature})
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, s.endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return SendBundleResponse{}, fmt.Errorf("new http request error: %w", err)
	}

	resp, err := doRequest[SendBundleResponse](s.c, httpReq, headers...)
	if err != nil {
		return SendBundleResponse{}, err
	}

	if len(resp.Error.Messange) != 0 {
		return SendBundleResponse{}, fmt.Errorf("response error, code: [%d], message: [%s]",
			resp.Error.Code, resp.Error.Messange)
	}

	// for some case, blink builder resp contains "" wrap around the bundle hash like
	/*
		2024-12-24T03:58:38Z	info	operator/broadcaster.go:465	send bundle (multiple)
		success	{"RequestID": "ctl32rfdqqbc73cb4m80", "id": "builder-blink",
		"tx": ["0xd629cbb2b4b741f6e71f8daafdbe1d484c5b53020b9cca9a407e0a7cb65f394c"],
		"uuid": null, "block": 21469801, "response":
		{"jsonrpc":"2.0","id":1,
		"result":{"bundleHash":"\"0xf72e9e8afd22af2904857e03575eb6f125cabc0d18fe7fb89ee1f8c6861687ae\""},
		"error":{}},
		"time": "661.127521ms", "start": "2024-12-24 03:58:37.475914455 +0000 UTC"}

		so we need to strip before save to db
	*/
	resp.Result.BundleHash = CleanBundleHash(resp.Result.BundleHash)
	return resp, nil
}

func (s *Client) SendPrivateRawTransaction(
	ctx context.Context,
	tx *types.Transaction,
) (SendPrivateRawTransactionResponse, error) {
	if !s.enableSendPrivateRaw {
		return SendPrivateRawTransactionResponse{}, nil
	}

	txBin, err := tx.MarshalBinary()
	if err != nil {
		return SendPrivateRawTransactionResponse{}, fmt.Errorf("marshal tx binary: %w", err)
	}

	req := SendRequest{
		ID:      SendBundleID,
		JSONRPC: JSONRPC2,
		Method:  ETHSendPrivateRawTransaction,
		Params:  []any{hexutil.Encode(txBin)},
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return SendPrivateRawTransactionResponse{}, fmt.Errorf("marshal json error: %w", err)
	}

	var headers [][2]string
	if s.flashbotKey != nil {
		signature, err := requestSignature(s.flashbotKey, reqBody)
		if err != nil {
			return SendPrivateRawTransactionResponse{}, fmt.Errorf("sign flashbot request error: %w", err)
		}
		headers = append(headers, [2]string{"X-Flashbots-Signature", signature})
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, s.endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return SendPrivateRawTransactionResponse{}, fmt.Errorf("new http request error: %w", err)
	}

	resp, err := doRequest[SendPrivateRawTransactionResponse](s.c, httpReq, headers...)
	if err != nil {
		return SendPrivateRawTransactionResponse{}, err
	}

	if len(resp.Error.Messange) != 0 {
		return SendPrivateRawTransactionResponse{}, fmt.Errorf("response error, code: [%d], message: [%s]",
			resp.Error.Code, resp.Error.Messange)
	}

	return resp, nil
}

func (s *Client) GetUserStats(
	ctx context.Context,
	useV2 bool,
	blockNumber uint64,
) (map[string]any, error) {
	if s.flashbotKey == nil {
		return nil, fmt.Errorf("not supported, nil key")
	}

	var method string
	switch s.senderType {
	case BundleSenderTypeTitan:
		method = TitanGetUserStats

	default:
		method = FlashbotGetUserStats
		if useV2 {
			method = FlashbotGetUserStatsV2
		}
	}

	params := GetUserStatsParams{}
	params.SetBlockNumber(blockNumber)

	req := SendRequest{
		ID:      1,
		JSONRPC: JSONRPC2,
		Method:  method,
		Params:  []any{},
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal json error: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, s.endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("new http request error: %w", err)
	}

	var headers [][2]string
	if s.flashbotKey != nil {
		signature, err := requestSignature(s.flashbotKey, reqBody)
		if err != nil {
			return nil, fmt.Errorf("sign flashbot request error: %w", err)
		}
		headers = append(headers, [2]string{"X-Flashbots-Signature", signature})
	}

	resp, err := doRequest[map[string]any](s.c, httpReq, headers...)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func requestSignature(key *ecdsa.PrivateKey, body []byte) (string, error) {
	hashed := crypto.Keccak256Hash(body).Hex()
	signature, err := crypto.Sign(accounts.TextHash([]byte(hashed)), key)
	if err != nil {
		return "", fmt.Errorf("sign crypto error: %w", err)
	}

	return fmt.Sprintf("%s:%s", crypto.PubkeyToAddress(key.PublicKey), hexutil.Encode(signature)), nil
}

type GetBundleStatsRequest struct {
	ID      int    `json:"id"`
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
}

type GetBundleStatsParams struct {
	BlockNumber hexutil.Uint64 `json:"blockNumber"`
	BundleHash  string         `json:"bundleHash"`
}

func (b *GetBundleStatsParams) SetBlockNumber(blockNumber uint64) *GetBundleStatsParams {
	b.BlockNumber = hexutil.Uint64(blockNumber)
	return b
}

func (b *GetBundleStatsParams) SetBundleHash(bundleHash common.Hash) *GetBundleStatsParams {
	b.BundleHash = bundleHash.Hex()
	return b
}

type SendRequest struct {
	ID      int    `json:"id"`
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
}

type SendBundleParams struct {
	// Array[String], A list of signed transactions to execute in an atomic bundle
	Txs []string `json:"txs,omitempty"`
	// String, a hex encoded block number for which this bundle is valid on
	BlockNumber string `json:"blockNumber,omitempty"`
	// (Optional) Number, the minimum timestamp for which this bundle is valid, in seconds since the unix epoch
	MinTimestamp *uint64 `json:"minTimestamp,omitempty"`
	// (Optional) Number, the maximum timestamp for which this bundle is valid, in seconds since the unix epoch
	MaxTimestamp *uint64 `json:"maxTimestamp,omitempty"`
	// (Optional) Array[String], A list of tx hashes that are allowed to revert
	RevertingTxs *[]string `json:"revertingTxHashes,omitempty"`
	// (Optional) String, UUID that can be used to cancel/replace this bundle
	ReplacementUUID string `json:"ReplacementUuid,omitempty"`
	// (Optional) String, UUID that can be used to cancel/replace this bundle (For beaverbuild)
	UUID             string `json:"uuid,omitempty"`
	StateBlockNumber string `json:"stateBlockNumber,omitempty"`

	Errors []error `json:"-"` // check when building bundle
}

func (p *SendBundleParams) SetStateBlockNumber(stateBlockNumber string) *SendBundleParams {
	p.StateBlockNumber = stateBlockNumber
	return p
}

func (p *SendBundleParams) SetPendingTxHash(txHash common.Hash) *SendBundleParams {
	if txHash == (common.Hash{}) {
		return p
	}

	p.Txs = append([]string{txHash.Hex()}, p.Txs...)
	return p
}

// SetPendingTxHashes will prepend the txHashes to the current list of transactions.
func (p *SendBundleParams) SetPendingTxHashes(txHashes ...common.Hash) *SendBundleParams {
	if len(txHashes) == 0 {
		return p
	}
	pendingTxs := make([]string, 0, len(txHashes))
	for _, txHash := range txHashes {
		pendingTxs = append(pendingTxs, txHash.Hex())
	}

	p.Txs = append(pendingTxs, p.Txs...)

	return p
}

// tested at: https://team-kyber.slack.com/archives/C03P04E6UAW/p1754280128232029?thread_ts=1754063210.955249&cid=C03P04E6UAW
func (p *SendBundleParams) SetTransactions(txs ...*types.Transaction) *SendBundleParams {
	if len(txs) == 0 {
		return p
	}

	transactions := make([]string, 0, len(txs))
	for _, tx := range txs {
		txBin, err := tx.MarshalBinary()
		if err != nil {
			p.Errors = append(p.Errors, fmt.Errorf("marshal tx: %w", err))
		} else {
			transactions = append(transactions, hexutil.Encode(txBin))
		}
	}

	p.Txs = transactions

	return p
}

func (p *SendBundleParams) SetTransactionsHex(txs ...string) *SendBundleParams {
	if len(txs) == 0 {
		return p
	}

	p.Txs = append(p.Txs, txs...)

	return p
}

func (p *SendBundleParams) SetBlockNumber(block uint64) *SendBundleParams {
	if block == 0 {
		return p
	}

	p.BlockNumber = fmt.Sprintf("0x%x", block)

	return p
}

func (p *SendBundleParams) SetUUID(uuid string, senderType BundleSenderType) *SendBundleParams {
	if senderType == BundleSenderTypeBeaver ||
		senderType == BundleSenderTypeLoki ||
		senderType == BundleSenderTypeJetbldr {
		p.UUID = uuid
		return p
	}
	p.ReplacementUUID = uuid

	return p
}

func (p *SendBundleParams) Err() error {
	if len(p.Errors) == 0 {
		return nil
	}

	return errors.Join(p.Errors...)
}

type CancelBundleParams struct {
	ReplacementUUID string `json:"replacementUuid"`
}

type GetUserStatsParams struct {
	BlockNumber string `json:"blockNumber"`
}

func (p *GetUserStatsParams) SetBlockNumber(block uint64) *GetUserStatsParams {
	if block == 0 {
		return p
	}

	p.BlockNumber = hexutil.EncodeUint64(block)

	return p
}
