package mev

import (
	"github.com/duoxehyon/mev-share-go/rpc"
	"crypto/ecdsa"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/flashbots/mev-share-node/mevshare"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type MevShareSender struct {
	client      rpc.MevAPIClient
	endpoint    string
	flashbotKey *ecdsa.PrivateKey
}

func (m MevShareSender) SendBundle(_ context.Context, _ *string, _ uint64, _ ...*types.Transaction) (SendBundleResponse, error) {
	return SendBundleResponse{}, ErrMethodNotSupport
}

func (m MevShareSender) SendBackrunBundle(_ context.Context, _ *string, blockNumber uint64, pendingTxHash common.Hash, txs ...*types.Transaction) (SendBundleResponse, error) {
	if m.client == nil {
		return SendBundleResponse{}, ErrMevShareClientNil
	}

	if len(txs) != 1 {
		return SendBundleResponse{}, ErrInvalidLenTx
	}

	rlpEncodedTx, err := txs[0].MarshalBinary()
	if err != nil {
		return SendBundleResponse{}, err
	}

	txBytes := hexutil.Bytes(rlpEncodedTx)
	// Define the bundle transactions
	mevBundleBody := []mevshare.MevBundleBody{
		{
			Hash: &pendingTxHash,
		},
		{
			Tx: &txBytes,
		},
	}
	inclusion := mevshare.MevBundleInclusion{
		BlockNumber: hexutil.Uint64(blockNumber),
	}

	// Make the bundle
	req := mevshare.SendMevBundleArgs{
		Body:      mevBundleBody,
		Inclusion: inclusion,
	}
	// Send bundle
	res, err := m.client.SendBundle(req)
	return SendBundleResponse{
		Result: SendBundleResult{
			BundleHash: res.BundleHash.String(),
		},
	}, err
}

func (m MevShareSender) CancelBundle(context.Context, string) error {
	return ErrMethodNotSupport
}

func (m MevShareSender) SimulateBundle(context.Context, uint64, ...*types.Transaction) (SendBundleResponse, error) {
	return SendBundleResponse{}, ErrMethodNotSupport
}

func (m MevShareSender) EstimateBundleGas(context.Context, []ethereum.CallMsg, *map[common.Address]gethclient.OverrideAccount) ([]uint64, error) {
	return nil, ErrMethodNotSupport
}

func (m MevShareSender) MevSimulateBundle(blockNumber uint64, pendingTxHash common.Hash, tx *types.Transaction) (*mevshare.SimMevBundleResponse, error) {
	if m.client == nil {
		return nil, ErrMevShareClientNil
	}

	rlpEncodedTx, err := tx.MarshalBinary()
	if err != nil {
		return nil, err
	}
	txBytes := hexutil.Bytes(rlpEncodedTx)

	// Define the bundle transactions
	txs := []mevshare.MevBundleBody{
		{
			Hash: &pendingTxHash,
		},
		{
			Tx: &txBytes,
		},
	}
	inclusion := mevshare.MevBundleInclusion{
		BlockNumber: hexutil.Uint64(blockNumber),
	}

	// Make the bundle
	req := mevshare.SendMevBundleArgs{
		Body:      txs,
		Inclusion: inclusion,
	}
	// Send bundle
	res, err := m.client.SimBundle(req, mevshare.SimMevBundleAuxArgs{})
	
	return res, err
}

func (m MevShareSender) GetSenderType() BundleSenderType {
	return BundleSenderTypeMevShare
}

func (m MevShareSender) GetBundleStats(context.Context, uint64, common.Hash) (GetBundleStatsResponse, error) {
	return GetBundleStatsResponse{}, ErrMethodNotSupport
}
