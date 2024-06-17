package mev

import (
	"context"
	"crypto/ecdsa"

	"github.com/duoxehyon/mev-share-go/rpc"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/flashbots/mev-share-node/mevshare"
)

type FlashbotMevShareSender struct {
	client      rpc.MevAPIClient
	endpoint    string
	flashbotKey *ecdsa.PrivateKey
}

func NewMevShareSender(
	endpoint string,
	flashbotKey *ecdsa.PrivateKey,
) (*FlashbotMevShareSender, error) {
	if flashbotKey == nil {
		return nil, ErrMissingPrivKey
	}

	return &FlashbotMevShareSender{
		client:      rpc.NewClient(endpoint, flashbotKey),
		endpoint:    endpoint,
		flashbotKey: flashbotKey,
	}, nil
}

func (m FlashbotMevShareSender) SendBundle(
	_ context.Context,
	_ *string,
	_ uint64,
	_ ...*types.Transaction,
) (SendBundleResponse, error) {
	return SendBundleResponse{}, ErrMethodNotSupport
}

func (m FlashbotMevShareSender) SendBackrunBundle(
	_ context.Context,
	_ *string,
	blockNumber uint64,
	pendingTxHash common.Hash,
	targetBuilders []string,
	txs ...*types.Transaction,
) (SendBundleResponse, error) {
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
		Privacy: &mevshare.MevBundlePrivacy{
			Builders: targetBuilders,
		},
	}
	if len(targetBuilders) == 0 {
		req.Privacy = &mevshare.MevBundlePrivacy{
			Builders: []string{
				FlashbotBuilderRegistrationFlashbot,
				FlashbotBuilderRegistrationBeaverBuild,
				FlashbotBuilderRegistrationTitan,
				FlashbotBuilderRegistrationRsync,
				FlashbotBuilderRegistrationBobaBuilder,
				FlashbotBuilderRegistrationBuilder0x69,
				FlashbotBuilderRegistrationBTCS,
				FlashbotBuilderRegistrationPenguinBuild,
			},
		}
	}

	// Send bundle
	res, err := m.client.SendBundle(req)
	if err != nil {
		return SendBundleResponse{}, err
	}

	return SendBundleResponse{
		Result: SendBundleResult{
			BundleHash: res.BundleHash.String(),
		},
	}, err
}

func (m FlashbotMevShareSender) CancelBundle(context.Context, string) error {
	return ErrMethodNotSupport
}

func (m FlashbotMevShareSender) SimulateBundle(
	context.Context,
	uint64, ...*types.Transaction,
) (SendBundleResponse, error) {
	return SendBundleResponse{}, ErrMethodNotSupport
}

func (m FlashbotMevShareSender) EstimateBundleGas(
	context.Context,
	[]ethereum.CallMsg,
	*map[common.Address]gethclient.OverrideAccount,
) ([]uint64, error) {
	return nil, ErrMethodNotSupport
}

func (m FlashbotMevShareSender) MevSimulateBundle(
	blockNumber uint64,
	pendingTxHash common.Hash,
	tx *types.Transaction,
) (*mevshare.SimMevBundleResponse, error) {
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

func (m FlashbotMevShareSender) GetSenderType() BundleSenderType {
	return BundleSenderTypeMevShare
}

func (m FlashbotMevShareSender) GetBundleStats(
	context.Context,
	uint64,
	common.Hash,
) (GetBundleStatsResponse, error) {
	return GetBundleStatsResponse{}, ErrMethodNotSupport
}
