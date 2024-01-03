package mev

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type Builder string

const (
	BuilderBloxroute    Builder = "bloxroute"
	BuilderFlashbot     Builder = "flashbots"
	BuilderBeaverBuild  Builder = "beaverbuild"
	BuilderRsyncBuilder Builder = "rsync-builder"
	BuilderAll          Builder = "all"
)

// BloxrouteSubmitBundle https://docs.bloxroute.com/apis/mev-solution/bundle-submission
func BloxrouteSubmitBundle( // nolint: cyclop
	ctx context.Context, c *http.Client, auth, endpoint string,
	param *BLXRSubmitBundleParams, options ...BloxrouteSubmitBundleOption,
) (SendBundleResponse, error) {
	var opts blxrSubmitBundleOptions
	for _, fn := range options {
		if fn == nil {
			continue
		}
		fn(&opts)
	}

	mevBuilders := make(map[Builder]string)
	if opts.builderBloxroute {
		mevBuilders[BuilderBloxroute] = ""
	}
	if opts.builderFlashbot != nil {
		sig, err := bloxrouteSignFlashbot(opts.builderFlashbot, param)
		if err != nil {
			return SendBundleResponse{}, fmt.Errorf("sign flashbot error: %w", err)
		}
		mevBuilders[BuilderFlashbot] = sig
	}
	if opts.builderBeaverBuild {
		mevBuilders[BuilderBeaverBuild] = ""
	}
	if opts.builderRsyncBuilder {
		mevBuilders[BuilderRsyncBuilder] = ""
	}
	if opts.builderAll {
		mevBuilders[BuilderAll] = ""
	}
	param.MevBuilders = mevBuilders

	req := BLXRSubmitBundleRequest{
		ID:     strconv.Itoa(SendBundleID),
		Method: BloxrouteSubmitBundleMethod,
		Params: param,
	}
	reqBody, err := json.Marshal(req)
	if err != nil {
		return SendBundleResponse{}, fmt.Errorf("marshal json error: %w", err)
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return SendBundleResponse{}, fmt.Errorf("new http request error: %w", err)
	}
	httpReq.Header.Add("Authorization", auth)
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

type blxrSubmitBundleOptions struct {
	builderBloxroute    bool
	builderFlashbot     *ecdsa.PrivateKey
	builderBeaverBuild  bool
	builderRsyncBuilder bool
	builderAll          bool
}

type BloxrouteSubmitBundleOption func(*blxrSubmitBundleOptions)

func WithBuilderBloxroute() BloxrouteSubmitBundleOption {
	return func(bsbo *blxrSubmitBundleOptions) {
		bsbo.builderBloxroute = true
	}
}

func WithBuilderFlashbot(key *ecdsa.PrivateKey) BloxrouteSubmitBundleOption {
	return func(bsbo *blxrSubmitBundleOptions) {
		bsbo.builderFlashbot = key
	}
}

func WithBuilderBeaverBuild() BloxrouteSubmitBundleOption {
	return func(bsbo *blxrSubmitBundleOptions) {
		bsbo.builderBeaverBuild = true
	}
}

func WithBuilderRsyncBuilder() BloxrouteSubmitBundleOption {
	return func(bsbo *blxrSubmitBundleOptions) {
		bsbo.builderRsyncBuilder = true
	}
}

// WithBuilderAll still need to use the WithBuilderFlashbot to submit to flashbots.
func WithBuilderAll() BloxrouteSubmitBundleOption {
	return func(bsbo *blxrSubmitBundleOptions) {
		bsbo.builderAll = true
	}
}

type BLXRSubmitBundleRequest struct {
	ID     string                  `json:"id,omitempty"`
	Method string                  `json:"method,omitempty"`
	Params *BLXRSubmitBundleParams `json:"params,omitempty"`
}

type BLXRSubmitBundleParams struct {
	Transaction     []string           `json:"transaction,omitempty"`
	BlockNumber     string             `json:"block_number,omitempty"`
	MinTimestamp    *uint64            `json:"min_timestamp,omitempty"`
	MaxTimestamp    *uint64            `json:"max_timestamp,omitempty"`
	RevertingHashes *[]string          `json:"reverting_hashes,omitempty"`
	UUID            string             `json:"uuid,omitempty"`
	MevBuilders     map[Builder]string `json:"mev_builders,omitempty"`
}

func (p *BLXRSubmitBundleParams) SetTransactions(txs ...*types.Transaction) *BLXRSubmitBundleParams {
	transactions := make([]string, 0, len(txs))
	for _, tx := range txs {
		transactions = append(transactions, "0x"+txToRlp(tx))
	}

	p.Transaction = transactions

	return p
}

func (p *BLXRSubmitBundleParams) SetBlockNumber(block uint64) *BLXRSubmitBundleParams {
	p.BlockNumber = fmt.Sprintf("0x%x", block)

	return p
}

func bloxrouteSignFlashbot(key *ecdsa.PrivateKey, p *BLXRSubmitBundleParams) (string, error) {
	param := new(SendBundleParams)
	param.Txs = p.Transaction
	param.BlockNumber = p.BlockNumber
	param.MinTimestamp = p.MinTimestamp
	param.MaxTimestamp = p.MaxTimestamp
	param.RevertingTxs = p.RevertingHashes

	req := SendBundleRequest{
		ID:      SendBundleID,
		JSONRPC: JSONRPC2,
		Method:  ETHSendBundleMethod,
	}
	req.Params = append(req.Params, param)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("marshal json error: %w", err)
	}

	signature, err := signRequest(key, reqBody)
	if err != nil {
		return "", fmt.Errorf("sign request error: %w", err)
	}

	sig := fmt.Sprintf("%s:%s", crypto.PubkeyToAddress(key.PublicKey), hexutil.Encode(signature))

	return sig, nil
}
