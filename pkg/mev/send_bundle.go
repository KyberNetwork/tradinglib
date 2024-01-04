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
	c           *http.Client
	endpoint    string
	flashbotKey *ecdsa.PrivateKey
}

// NewClient set the flashbotKey to nil will skip adding the signature header.
func NewClient(c *http.Client, endpoint string, flashbotKey *ecdsa.PrivateKey) *Client {
	return &Client{
		c:           c,
		endpoint:    endpoint,
		flashbotKey: flashbotKey,
	}
}

func (s *Client) SendBundle(
	ctx context.Context, blockNumber uint64, txs ...*types.Transaction,
) (SendBundleResponse, error) {
	req := SendBundleRequest{
		ID:      SendBundleID,
		JSONRPC: JSONRPC2,
		Method:  ETHSendBundleMethod,
	}
	p := new(SendBundleParams).SetBlockNumber(blockNumber).SetTransactions(txs...)
	req.Params = append(req.Params, p)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return SendBundleResponse{}, fmt.Errorf("marshal json error: %w", err)
	}

	var headers [][2]string
	if s.flashbotKey != nil {
		signature, err := signRequest(s.flashbotKey, reqBody)
		if err != nil {
			return SendBundleResponse{}, fmt.Errorf("sign flashbot request error: %w", err)
		}
		flashbotSig := fmt.Sprintf("%s:%s",
			crypto.PubkeyToAddress(s.flashbotKey.PublicKey), hexutil.Encode(signature))

		headers = append(headers, [2]string{"X-Flashbots-Signature", flashbotSig})
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

func signRequest(key *ecdsa.PrivateKey, body []byte) ([]byte, error) {
	hashed := crypto.Keccak256Hash(body).Hex()
	signature, err := crypto.Sign(accounts.TextHash([]byte(hashed)), key)
	if err != nil {
		return nil, fmt.Errorf("sign crypto error: %w", err)
	}

	return signature, nil
}

type SendBundleRequest struct {
	ID      int    `json:"id"`
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
}

type SendBundleParams struct {
	// Array[String], A list of signed transactions to execute in an atomic bundle
	Txs []string `json:"txs"`
	// String, a hex encoded block number for which this bundle is valid on
	BlockNumber string `json:"blockNumber"`
	// (Optional) Number, the minimum timestamp for which this bundle is valid, in seconds since the unix epoch
	MinTimestamp *uint64 `json:"minTimestamp,omitempty"`
	// (Optional) Number, the maximum timestamp for which this bundle is valid, in seconds since the unix epoch
	MaxTimestamp *uint64 `json:"maxTimestamp,omitempty"`
	// (Optional) Array[String], A list of tx hashes that are allowed to revert
	RevertingTxs *[]string `json:"revertingTxHashes,omitempty"`
}

func (p *SendBundleParams) SetTransactions(txs ...*types.Transaction) *SendBundleParams {
	transactions := make([]string, 0, len(txs))
	for _, tx := range txs {
		transactions = append(transactions, "0x"+txToRlp(tx))
	}

	p.Txs = transactions

	return p
}

func (p *SendBundleParams) SetBlockNumber(block uint64) *SendBundleParams {
	p.BlockNumber = fmt.Sprintf("0x%x", block)

	return p
}
