package mev

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/flashbots/mev-share-node/mevshare"
)

type BackrunPublicClient struct {
	c           *http.Client
	endpoint    string
	flashbotKey *ecdsa.PrivateKey
	senderType  BundleSenderType
}

func NewBackrunPublicClient(
	c *http.Client,
	endpoint string,
	flashbotKey *ecdsa.PrivateKey,
	senderType BundleSenderType,
) *BackrunPublicClient {
	return &BackrunPublicClient{
		c:           c,
		endpoint:    endpoint,
		flashbotKey: flashbotKey,
		senderType:  senderType,
	}
}

func (b BackrunPublicClient) SendBackrunBundle(
	ctx context.Context,
	uuid *string,
	blockNumber uint64,
	_ uint64,
	pendingTxHash common.Hash,
	_ []string,
	txs ...*types.Transaction,
) (SendBundleResponse, error) {
	req := SendBundleRequest{
		ID:      SendBundleID,
		JSONRPC: JSONRPC2,
		Method:  ETHSendBundleMethod,
	}
	p := new(SendBundleParams).SetBlockNumber(blockNumber).SetTransactions(txs...).SetPendingTxHash(pendingTxHash)
	if uuid != nil {
		p.SetUUID(*uuid, b.senderType)
	}
	req.Params = append(req.Params, p)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return SendBundleResponse{}, fmt.Errorf("marshal json error: %w", err)
	}

	var headers [][2]string
	if b.flashbotKey != nil {
		signature, err := requestSignature(b.flashbotKey, reqBody)
		if err != nil {
			return SendBundleResponse{}, fmt.Errorf("sign flashbot request error: %w", err)
		}
		headers = append(headers, [2]string{"X-Flashbots-Signature", signature})
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, b.endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return SendBundleResponse{}, fmt.Errorf("new http request error: %w", err)
	}

	resp, err := doRequest[SendBundleResponse](b.c, httpReq, headers...)
	if err != nil {
		return SendBundleResponse{}, err
	}

	if len(resp.Error.Messange) != 0 {
		return SendBundleResponse{}, fmt.Errorf("response error, code: [%d], message: [%s]",
			resp.Error.Code, resp.Error.Messange)
	}

	return resp, nil
}

func (b BackrunPublicClient) MevSimulateBundle(
	_ uint64, _ common.Hash, _ *types.Transaction,
) (*mevshare.SimMevBundleResponse, error) {
	return nil, ErrMethodNotSupport
}

func (b BackrunPublicClient) GetSenderType() BundleSenderType {
	return b.senderType
}
