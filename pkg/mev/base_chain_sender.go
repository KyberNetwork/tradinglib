package mev

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

type BaseChainSender struct {
	c          *http.Client
	endpoint   string
	senderType BundleSenderType
}

func NewBaseChainSender(
	c *http.Client,
	endpoint string,
	senderType BundleSenderType,
) *BaseChainSender {
	return &BaseChainSender{
		c:          c,
		endpoint:   endpoint,
		senderType: senderType,
	}
}

func (s *BaseChainSender) GetSenderType() BundleSenderType {
	return s.senderType
}

func (s *BaseChainSender) SendRawTransaction(
	ctx context.Context,
	tx *types.Transaction,
) (SendRawTransactionResponse, error) {
	txBin, err := tx.MarshalBinary()
	if err != nil {
		return SendRawTransactionResponse{}, fmt.Errorf("marshal tx binary: %w", err)
	}

	req := SendRequest{
		ID:      SendBundleID,
		JSONRPC: JSONRPC2,
		Method:  ETHSendRawTransaction,
		Params:  []any{hexutil.Encode(txBin)},
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return SendRawTransactionResponse{}, fmt.Errorf("marshal json error: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, s.endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return SendRawTransactionResponse{}, fmt.Errorf("new http request error: %w", err)
	}

	resp, err := doRequest[SendRawTransactionResponse](s.c, httpReq)
	if err != nil {
		return SendRawTransactionResponse{}, err
	}

	if len(resp.Error.Messange) != 0 {
		return SendRawTransactionResponse{}, fmt.Errorf("response error, code: [%d], message: [%s]",
			resp.Error.Code, resp.Error.Messange)
	}

	return resp, nil
}
