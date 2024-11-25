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

type MerkleClient struct {
	c                *http.Client
	endpoint         string
	flashbotKey      *ecdsa.PrivateKey
	senderType       BundleSenderType
	xSignatureHeader string
}

func NewMerkleClient(
	c *http.Client,
	endpoint string,
	flashbotKey *ecdsa.PrivateKey,
	senderType BundleSenderType,
	xSignatureHeader string,
) *MerkleClient {
	return &MerkleClient{
		c:                c,
		endpoint:         endpoint,
		flashbotKey:      flashbotKey,
		senderType:       senderType,
		xSignatureHeader: xSignatureHeader,
	}
}

func (b MerkleClient) SendBackrunBundle(
	ctx context.Context,
	uuid *string,
	blockNumber uint64,
	_ uint64,
	pendingTxHashes []common.Hash,
	_ []string,
	txs ...*types.Transaction,
) (SendBundleResponse, error) {
	req := SendRequest{
		ID:      SendBundleID,
		JSONRPC: JSONRPC2,
		Method:  ETHSendBundleMethod,
	}

	// we set transactions first, then prepend pendingTxHashes
	p := new(SendBundleParams).SetBlockNumber(blockNumber).SetTransactions(txs...).SetPendingTxHashes(pendingTxHashes...)
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
		headers = append(headers, [2]string{b.xSignatureHeader, signature})
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, b.endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return SendBundleResponse{}, fmt.Errorf("new http request error: %w", err)
	}

	resp, err := doRequest[MerkleSendBundleResponse](b.c, httpReq, headers...)
	if err != nil {
		return SendBundleResponse{}, err
	}

	if len(resp.Error.Messange) != 0 {
		return SendBundleResponse{}, fmt.Errorf("response error, code: [%d], message: [%s]",
			resp.Error.Code, resp.Error.Messange)
	}

	return SendBundleResponse{
		Result: SendBundleResult{
			BundleGasPrice:    "",
			BundleHash:        resp.Result,
			CoinbaseDiff:      "",
			EthSentToCoinbase: "",
			GasFees:           "",
			Results:           nil,
			StateBlockNumber:  0,
			TotalGasUsed:      0,
			Message:           "",
		},
	}, nil
}

func (b MerkleClient) MevSimulateBundle(
	_ uint64, _ common.Hash, _ *types.Transaction,
) (*mevshare.SimMevBundleResponse, error) {
	return nil, ErrMethodNotSupport
}

func (b MerkleClient) GetSenderType() BundleSenderType {
	return b.senderType
}
