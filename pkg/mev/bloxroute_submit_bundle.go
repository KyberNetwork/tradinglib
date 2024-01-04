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

type BlxrBuilder string

const (
	BuilderBloxroute    BlxrBuilder = "bloxroute"
	BuilderFlashbot     BlxrBuilder = "flashbots"
	BuilderBeaverBuild  BlxrBuilder = "beaverbuild"
	BuilderRsyncBuilder BlxrBuilder = "rsync-builder"
	BuilderAll          BlxrBuilder = "all"
)

type BloxrouteClient struct {
	c               *http.Client
	endpoint        string
	auth            string
	flashbotKey     *ecdsa.PrivateKey
	enabledBuilders []BlxrBuilder
}

// NewBloxrouteClient set flashbotKey to nil if you don't want to send to flashbot builders
// With BuilderAll still need to add the flashbot key & the flashbot builder separately
// https://docs.bloxroute.com/apis/mev-solution/bundle-submission
func NewBloxrouteClient(
	c *http.Client,
	endpoint, auth string,
	flashbotKey *ecdsa.PrivateKey,
	enabledBuilders ...BlxrBuilder,
) *BloxrouteClient {
	return &BloxrouteClient{
		c:               c,
		endpoint:        endpoint,
		auth:            auth,
		flashbotKey:     flashbotKey,
		enabledBuilders: enabledBuilders,
	}
}

func (s *BloxrouteClient) SendBundle(
	ctx context.Context, blockNumber uint64, txs ...*types.Transaction,
) (SendBundleResponse, error) {
	p := new(BLXRSubmitBundleParams).SetBlockNumber(blockNumber).SetTransactions(txs...)

	mevBuilders := make(map[BlxrBuilder]string)
	for _, b := range s.enabledBuilders {
		if b == BuilderFlashbot && s.flashbotKey != nil {
			sig, err := bloxrouteSignFlashbot(s.flashbotKey, p)
			if err != nil {
				return SendBundleResponse{}, fmt.Errorf("sign flashbot error: %w", err)
			}
			mevBuilders[BuilderFlashbot] = sig
			continue
		}
		mevBuilders[b] = ""
	}

	p.MEVBuilders = mevBuilders
	req := BLXRSubmitBundleRequest{
		ID:     strconv.Itoa(SendBundleID),
		Method: BloxrouteSubmitBundleMethod,
		Params: p,
	}
	reqBody, err := json.Marshal(req)
	if err != nil {
		return SendBundleResponse{}, fmt.Errorf("marshal json error: %w", err)
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, s.endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return SendBundleResponse{}, fmt.Errorf("new http request error: %w", err)
	}

	resp, err := doRequest[SendBundleResponse](s.c, httpReq, [2]string{"Authorization", s.auth})
	if err != nil {
		return SendBundleResponse{}, err
	}

	if len(resp.Error.Messange) != 0 {
		return SendBundleResponse{}, fmt.Errorf("response error, code: [%d], message: [%s]",
			resp.Error.Code, resp.Error.Messange)
	}

	return resp, nil
}

type BLXRSubmitBundleRequest struct {
	ID     string                  `json:"id,omitempty"`
	Method string                  `json:"method,omitempty"`
	Params *BLXRSubmitBundleParams `json:"params,omitempty"`
}

type BLXRSubmitBundleParams struct {
	Transaction     []string               `json:"transaction,omitempty"`
	BlockNumber     string                 `json:"block_number,omitempty"`
	MinTimestamp    *uint64                `json:"min_timestamp,omitempty"`
	MaxTimestamp    *uint64                `json:"max_timestamp,omitempty"`
	RevertingHashes *[]string              `json:"reverting_hashes,omitempty"`
	UUID            string                 `json:"uuid,omitempty"`
	MEVBuilders     map[BlxrBuilder]string `json:"mev_builders,omitempty"`
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
