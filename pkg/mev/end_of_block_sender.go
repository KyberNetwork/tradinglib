package mev

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

// IEndOfBlockBundleSender submits a bundle that the builder simulates against the
// END-OF-BLOCK state instead of the state the block was being built on.
//
// After a block is built, an end-of-block bundle whose target pool was touched by that
// block is simulated on the final state and appended when the simulation succeeds.
// Atomicity and revert guarantees are the same as for a normal bundle.
//
// https://docs.titanbuilder.xyz/api/eth_sendendofblockbundle
// https://docs.quasar.win/api/eth_sendendofblockbundle
type IEndOfBlockBundleSender interface {
	SendEndOfBlockBundle(
		ctx context.Context,
		req SendEndOfBlockBundleRequest,
		txs ...*types.Transaction,
	) (SendBundleResponse, error)
	GetSenderType() BundleSenderType
}

// nolint: gochecknoglobals
var (
	_ IEndOfBlockBundleSender = &Client{}
	_ IEndOfBlockBundleSender = &BloxrouteClient{}
)

type SendEndOfBlockBundleRequest struct {
	// TargetPools is required: the pool addresses whose end-of-block state this bundle
	// targets. The bundle is only simulated for a block that modified one of them.
	TargetPools []common.Address
	// (Optional) the block this bundle is valid on. 0 leaves it out, and the builder
	// defaults to the current block.
	BlockNumber uint64
	// (Optional) tx hashes that are allowed to revert or be discarded.
	RevertingTxHashes []common.Hash
	// (Optional) an arbitrary string used to replace or cancel this bundle. Passing an
	// empty tx list with a replacement UUID cancels the bundle that holds it. nil leaves
	// the field out.
	ReplacementUUID *string
	// (Optional) monotonically increasing sequence for bundles sharing one
	// ReplacementUUID. A bundle whose sequence is not higher than the last one is
	// dropped. 0 falls back to ordering by builder receive time.
	ReplacementSeqNumber *uint64
}

// endOfBlockBundleParams is the param object Titan and Quasar document. It is
// deliberately not SendBundleParams: txs and targetPools are required and must be
// emitted even when empty, and none of the builderNet/refund fields apply.
type endOfBlockBundleParams struct {
	Txs                  []string `json:"txs"`
	BlockNumber          string   `json:"blockNumber,omitempty"`
	RevertingTxHashes    []string `json:"revertingTxHashes,omitempty"`
	TargetPools          []string `json:"targetPools"`
	ReplacementUUID      *string  `json:"replacementUuid,omitempty"`
	ReplacementSeqNumber *uint64  `json:"replacementSeqNumber,omitempty"`
}

// SendEndOfBlockBundle sends eth_sendEndOfBlockBundle. Titan and Quasar are the
// builders that document the method, but the request is not gated on the sender type:
// the caller decides which endpoint gets it. postBundle attaches the
// X-Flashbots-Signature header only when the client was built with a flashbot key.
func (s *Client) SendEndOfBlockBundle(
	ctx context.Context,
	req SendEndOfBlockBundleRequest,
	txs ...*types.Transaction,
) (SendBundleResponse, error) {
	if len(req.TargetPools) == 0 {
		return SendBundleResponse{}, ErrMissingTargetPools
	}

	rawTxs, err := marshalTxs(txs)
	if err != nil {
		return SendBundleResponse{}, err
	}

	p := endOfBlockBundleParams{
		Txs:                  rawTxs,
		TargetPools:          hexAddresses(req.TargetPools),
		RevertingTxHashes:    hexHashes(req.RevertingTxHashes),
		ReplacementUUID:      req.ReplacementUUID,
		ReplacementSeqNumber: req.ReplacementSeqNumber,
	}
	if req.BlockNumber != 0 {
		p.BlockNumber = hexutil.EncodeUint64(req.BlockNumber)
	}

	return s.postBundle(ctx, ETHSendEndOfBlockBundleMethod, p)
}

// SendEndOfBlockBundle sends the bundle through bloXroute.
//
// bloXroute exposes no eth_sendEndOfBlockBundle method, and no end-of-block RPC of any
// name. Its equivalent is the Target Pool Backrunning feature of blxr_submit_bundle:
// target_addresses (the pools) plus bottom=true (bottom-of-block placement, as opposed
// to bottom=false for intrablock). Setting them routes the bundle only to the builders
// that support target-pool backrunning.
// https://docs.bloxroute.com/eth/submit-bundles/bundle-submission
//
// bloXroute also accepts target_slots, hex storage slots per address index-aligned with
// target_addresses, to narrow the trigger below pool granularity. Titan and Quasar have
// no such parameter, so SendEndOfBlockBundleRequest does not carry one; set
// BLXRSubmitBundleParams.TargetSlots directly if you need it.
//
// ReplacementSeqNumber has no bloXroute equivalent and is dropped. A bundle sharing a
// UUID is ordered by bloXroute receive time, so a delayed submission can replace a
// newer one.
func (s *BloxrouteClient) SendEndOfBlockBundle(
	ctx context.Context,
	req SendEndOfBlockBundleRequest,
	txs ...*types.Transaction,
) (SendBundleResponse, error) {
	if len(req.TargetPools) == 0 {
		return SendBundleResponse{}, ErrMissingTargetPools
	}
	p := new(BLXRSubmitBundleParams).
		SetBlockNumber(req.BlockNumber).
		SetTransactions(txs...).
		SetTargetAddresses(req.TargetPools...).
		SetBottom(true)
	if req.ReplacementUUID != nil {
		p.SetUUID(*req.ReplacementUUID)
	}
	if len(req.RevertingTxHashes) != 0 {
		reverting := hexHashes(req.RevertingTxHashes)
		p.RevertingHashes = &reverting
	}
	p.BlockchainNetwork = s.blxrBlockchainNetwork

	if err := p.Err(); err != nil {
		return SendBundleResponse{}, err
	}

	return s.sendBundle(ctx, p)
}

// marshalTxs returns the RLP of every tx, hex encoded. The result is never nil: an
// empty tx list is how a bundle is cancelled, so it has to marshal to [] and not null.
func marshalTxs(txs []*types.Transaction) ([]string, error) {
	out := make([]string, 0, len(txs))
	for _, tx := range txs {
		raw, err := tx.MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("marshal tx binary: %w", err)
		}
		out = append(out, hexutil.Encode(raw))
	}

	return out, nil
}

func hexAddresses(addresses []common.Address) []string {
	out := make([]string, 0, len(addresses))
	for _, a := range addresses {
		out = append(out, a.Hex())
	}

	return out
}

func hexHashes(hashes []common.Hash) []string {
	if len(hashes) == 0 {
		return nil
	}
	out := make([]string, 0, len(hashes))
	for _, h := range hashes {
		out = append(out, h.Hex())
	}

	return out
}
