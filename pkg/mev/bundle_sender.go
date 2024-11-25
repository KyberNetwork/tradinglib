package mev

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
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
	return s.sendBundle(ctx, ETHSendBundleMethod, uuid, blockNumber, txs...)
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
		if _, err := s.sendBundle(ctx, ETHSendBundleMethod, &bundleUUID, 0); err != nil {
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
	return s.sendBundle(ctx, EthCallBundleMethod, nil, blockNumber, txs...)
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
	txs ...*types.Transaction,
) (SendBundleResponse, error) {
	if !s.enableSendPrivateRaw {
		return SendBundleResponse{}, nil
	}

	req := SendRequest{
		ID:      SendBundleID,
		JSONRPC: JSONRPC2,
		Method:  method,
	}
	p := new(SendBundleParams).SetBlockNumber(blockNumber).SetTransactions(txs...)
	if s.senderType == BundleSenderTypeFlashbot {
		p = p.SetStateBlockNumber("latest")
	}
	if uuid != nil {
		p.SetUUID(*uuid, s.senderType)
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

	return resp, nil
}

func (s *Client) SendPrivateRawTransaction(
	ctx context.Context,
	tx *types.Transaction,
) (SendPrivateRawTransactionResponse, error) {
	req := SendRequest{
		ID:      SendBundleID,
		JSONRPC: JSONRPC2,
		Method:  ETHSendPrivateRawTransaction,
		Params:  []any{"0x" + txToRlp(tx)},
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

func (p *SendBundleParams) SetTransactions(txs ...*types.Transaction) *SendBundleParams {
	if len(txs) == 0 {
		return p
	}

	transactions := make([]string, 0, len(txs))
	for _, tx := range txs {
		transactions = append(transactions, "0x"+txToRlp(tx))
	}

	p.Txs = transactions

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
	if senderType == BundleSenderTypeBeaver {
		p.UUID = uuid
		return p
	}
	p.ReplacementUUID = uuid

	return p
}

type CancelBundleParams struct {
	ReplacementUUID string `json:"replacementUuid"`
}
