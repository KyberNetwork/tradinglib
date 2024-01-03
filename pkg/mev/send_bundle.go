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

// SendBundle https://docs.flashbots.net/flashbots-auction/advanced/rpc-endpoint#eth_sendbundle,
// https://beaverbuild.org/docs.html, https://rsync-builder.xyz/docs
func SendBundle( // nolint: cyclop
	ctx context.Context, c *http.Client, endpoint string, param *SendBundleParams, options ...SendBundleOption,
) (SendBundleResponse, error) {
	var opts sendBundleOpts
	for _, fn := range options {
		if fn == nil {
			continue
		}
		fn(&opts)
	}

	req := SendBundleRequest{
		ID:      SendBundleID,
		JSONRPC: JSONRPC2,
		Method:  ETHSendBundleMethod,
	}
	req.Params = append(req.Params, param)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return SendBundleResponse{}, fmt.Errorf("marshal json error: %w", err)
	}

	headers := make(map[string]string)
	if opts.flashbotSignKey != nil {
		signature, err := signRequest(opts.flashbotSignKey, reqBody)
		if err != nil {
			return SendBundleResponse{}, fmt.Errorf("sign flashbot request error: %w", err)
		}

		// for flashbot only
		headers["X-Flashbots-Signature"] = fmt.Sprintf("%s:%s",
			crypto.PubkeyToAddress(opts.flashbotSignKey.PublicKey), hexutil.Encode(signature))
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return SendBundleResponse{}, fmt.Errorf("new http request error: %w", err)
	}

	for k, v := range headers {
		httpReq.Header.Add(k, v)
	}
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Accept", "application/json")

	resp, err := doRequest[SendBundleResponse](c, httpReq)
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

type sendBundleOpts struct {
	flashbotSignKey *ecdsa.PrivateKey
}

type SendBundleOption func(*sendBundleOpts)

func WithFlashbotSignature(key *ecdsa.PrivateKey) SendBundleOption {
	return func(sbo *sendBundleOpts) {
		if key == nil {
			return
		}

		sbo.flashbotSignKey = key
	}
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
