package mev

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	bloxroute "github.com/bloXroute-Labs/bloxroute-sdk-go"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

type BloxrouteBackrunmeSender struct {
	client *bloxroute.Client
}

type BackrunmeSendBundleResponse struct {
	BundleHash string `json:"bundleHash"`
}

func NewBloxrouteBackrunmeSender(authHeader, wsGatewayUrl string) (*BloxrouteBackrunmeSender, error) {
	c, err := bloxroute.NewClient(context.Background(), &bloxroute.Config{
		AuthHeader:   authHeader,
		WSGatewayURL: wsGatewayUrl,
	})
	if err != nil {
		return nil, fmt.Errorf("new bloxroute client: %w", err)
	}

	return &BloxrouteBackrunmeSender{client: c}, nil
}

func (s *BloxrouteBackrunmeSender) SendBackrunBundle(
	ctx context.Context,
	uuid *string,
	blockNumber uint64,
	maxBlockNumber uint64,
	pendingTxHashes []common.Hash,
	targetBuilders []string,
	txs ...*types.Transaction,
) (SendBundleResponse, error) {
	transactions := make([]string, 0, len(txs))
	for _, tx := range txs {
		txBin, err := tx.MarshalBinary()
		if err != nil {
			return SendBundleResponse{}, fmt.Errorf("marshal tx: %w", err)
		} else {
			hexTx := hexutil.Encode(txBin)
			// remove 0x prefix
			hexTx = strings.TrimPrefix(hexTx, "0x")
			transactions = append(transactions, hexTx)
		}
	}

	res, err := s.client.SendEthBundle(ctx, &bloxroute.SendEthBundleParams{
		BlockNumber:  fmt.Sprintf("0x%x", blockNumber),
		Transactions: transactions,
	})
	if err != nil {
		return SendBundleResponse{}, fmt.Errorf("send eth bundle: %w", err)
	}

	marshaled, err := res.MarshalJSON()
	if err != nil {
		return SendBundleResponse{}, fmt.Errorf("marshal res: %w", err)
	}

	var backrunmeSendBundleResponse BackrunmeSendBundleResponse
	if err := json.Unmarshal(marshaled, &backrunmeSendBundleResponse); err != nil {
		return SendBundleResponse{}, fmt.Errorf("unmarshal res: %w", err)
	}

	return SendBundleResponse{
		Result: SendBundleResult{
			BundleHash: backrunmeSendBundleResponse.BundleHash,
		},
	}, nil
}

func (s *BloxrouteBackrunmeSender) MevSimulateBundle(
	ctx context.Context,
	blockNumber uint64,
	pendingTxHash common.Hash,
	tx *types.Transaction,
) (SendBundleResponse, error) {
	return SendBundleResponse{}, nil
}

func (s *BloxrouteBackrunmeSender) GetSenderType() BundleSenderType {
	return BundleSenderTypeBloxrouteBackrunme
}
