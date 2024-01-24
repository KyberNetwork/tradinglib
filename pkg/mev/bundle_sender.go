package mev

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// Client https://beaverbuild.org/docs.html; https://rsync-builder.xyz/docs;
// https://docs.flashbots.net/flashbots-auction/advanced/rpc-endpoint#eth_sendbundle
type Client struct {
	c                  *http.Client
	endpoint           string
	flashbotKey        *ecdsa.PrivateKey
	cancelBySendBundle bool
	senderType         BundleSenderType
}

// NewClient set the flashbotKey to nil will skip adding the signature header.
func NewClient(
	c *http.Client,
	endpoint string,
	flashbotKey *ecdsa.PrivateKey,
	cancelBySendBundle bool,
	senderType BundleSenderType,
) *Client {
	return &Client{
		c:                  c,
		endpoint:           endpoint,
		flashbotKey:        flashbotKey,
		cancelBySendBundle: cancelBySendBundle,
		senderType:         senderType,
	}
}

func (s *Client) SendBundle(
	ctx context.Context,
	uuid *string,
	blockNumber uint64,
	txs ...*types.Transaction,
) (SendBundleResponse, error) {
	return s.sendBundle(ctx, uuid, blockNumber, txs...)
}

func (s *Client) CancelBundle(
	ctx context.Context, bundleUUID string,
) error {
	if s.cancelBySendBundle {
		if _, err := s.sendBundle(ctx, &bundleUUID, 0); err != nil {
			return fmt.Errorf("cancel by send bundle error: %w", err)
		}
		return nil
	}

	// build request
	p := CancelBundleParams{
		ReplacementUUID: bundleUUID,
	}
	req := SendBundleRequest{
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
	var errResp SendBundleError
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

func (s *Client) sendBundle(
	ctx context.Context,
	uuid *string,
	blockNumber uint64,
	txs ...*types.Transaction,
) (SendBundleResponse, error) {
	req := SendBundleRequest{
		ID:      SendBundleID,
		JSONRPC: JSONRPC2,
		Method:  ETHSendBundleMethod,
	}
	p := new(SendBundleParams).SetBlockNumber(blockNumber).SetTransactions(txs...)
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

func requestSignature(key *ecdsa.PrivateKey, body []byte) (string, error) {
	hashed := crypto.Keccak256Hash(body).Hex()
	signature, err := crypto.Sign(accounts.TextHash([]byte(hashed)), key)
	if err != nil {
		return "", fmt.Errorf("sign crypto error: %w", err)
	}

	return fmt.Sprintf("%s:%s", crypto.PubkeyToAddress(key.PublicKey), hexutil.Encode(signature)), nil
}

type SendBundleRequest struct {
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
	UUID string `json:"uuid,omitempty"`
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
