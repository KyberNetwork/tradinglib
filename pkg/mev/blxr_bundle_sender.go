package mev

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
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

func (s *BloxrouteClient) SimulateBundle(
	_ context.Context,
	_ uint64,
	_ ...*types.Transaction,
) (SendBundleResponse, error) {
	return SendBundleResponse{}, ErrMethodNotSupport
}

func (s *BloxrouteClient) EstimateBundleGas(
	_ context.Context,
	_ []ethereum.CallMsg,
	_ *map[common.Address]gethclient.OverrideAccount,
) ([]uint64, error) {
	return nil, nil
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

func (s *BloxrouteClient) GetSenderType() BundleSenderType {
	return BundleSenderTypeBloxroute
}

func (s *BloxrouteClient) SendBundle(
	ctx context.Context,
	uuid *string,
	blockNumber uint64,
	txs ...*types.Transaction,
) (SendBundleResponse, error) {
	p := new(BLXRSubmitBundleParams).SetBlockNumber(blockNumber).SetTransactions(txs...)
	if uuid != nil {
		p.SetUUID(*uuid)
	}
	if err := p.Err(); err != nil {
		return SendBundleResponse{}, err
	}

	return s.sendBundle(ctx, p)
}

func (s *BloxrouteClient) SendBundleHex(
	ctx context.Context,
	uuid *string,
	blockNumber uint64,
	hexEncodedTxs ...string,
) (SendBundleResponse, error) {
	p := new(BLXRSubmitBundleParams).SetBlockNumber(blockNumber).SetTransactionsHex(hexEncodedTxs...)
	if uuid != nil {
		p.SetUUID(*uuid)
	}
	if err := p.Err(); err != nil {
		return SendBundleResponse{}, err
	}

	return s.sendBundle(ctx, p)
}

func (s *BloxrouteClient) sendBundle(
	ctx context.Context,
	p *BLXRSubmitBundleParams,
) (SendBundleResponse, error) {
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

	resp, err := doRequest[BLXRSubmitBundleResponse](s.c, httpReq, [2]string{"Authorization", s.auth})
	if err != nil {
		return SendBundleResponse{}, err
	}

	if len(resp.Error.Messange) != 0 {
		return SendBundleResponse{}, fmt.Errorf("response error, code: [%d], message: [%s]",
			resp.Error.Code, resp.Error.Messange)
	}

	return SendBundleResponse(resp), nil
}

func (s *BloxrouteClient) SendPrivateRawTransaction(
	ctx context.Context,
	tx *types.Transaction,
) (SendPrivateRawTransactionResponse, error) {
	return SendPrivateRawTransactionResponse{}, nil
}

func (s *BloxrouteClient) CancelBundle(
	ctx context.Context, bundleUUID string,
) error {
	_, err := s.SendBundle(ctx, &bundleUUID, 0)
	if err != nil {
		return fmt.Errorf("cancel by send bundle error: %w", err)
	}

	return nil
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

	Errors []error `json:"-"`
}

func (p *BLXRSubmitBundleParams) SetTransactions(txs ...*types.Transaction) *BLXRSubmitBundleParams {
	if len(txs) == 0 {
		return p
	}

	transactions := make([]string, 0, len(txs))
	for _, tx := range txs {
		txBin, err := tx.MarshalBinary()
		if err != nil {
			p.Errors = append(p.Errors, err)
		} else {
			transactions = append(transactions, hexutil.Encode(txBin))
		}
	}

	p.Transaction = transactions

	return p
}

func (p *BLXRSubmitBundleParams) SetTransactionsHex(txs ...string) *BLXRSubmitBundleParams {
	if len(txs) == 0 {
		return p
	}

	p.Transaction = append(p.Transaction, txs...)

	return p
}

func (p *BLXRSubmitBundleParams) SetBlockNumber(block uint64) *BLXRSubmitBundleParams {
	if block == 0 {
		return p
	}

	p.BlockNumber = fmt.Sprintf("0x%x", block)

	return p
}

func (p *BLXRSubmitBundleParams) SetUUID(uuid string) *BLXRSubmitBundleParams {
	p.UUID = uuid

	return p
}

func (p *BLXRSubmitBundleParams) Err() error {
	if len(p.Errors) == 0 {
		return nil
	}

	return errors.Join(p.Errors...)
}

type BLXRSubmitBundleResponse struct {
	Jsonrpc string           `json:"jsonrpc,omitempty"`
	ID      int              `json:"id,string,omitempty"`
	Result  SendBundleResult `json:"result,omitempty"`
	Error   ErrorResponse    `json:"error,omitempty"`
}

func bloxrouteSignFlashbot(key *ecdsa.PrivateKey, p *BLXRSubmitBundleParams) (string, error) {
	param := new(SendBundleParams)
	param.Txs = p.Transaction
	param.BlockNumber = p.BlockNumber
	param.MinTimestamp = p.MinTimestamp
	param.MaxTimestamp = p.MaxTimestamp
	param.RevertingTxs = p.RevertingHashes

	req := SendRequest{
		ID:      SendBundleID,
		JSONRPC: JSONRPC2,
		Method:  ETHSendBundleMethod,
	}
	req.Params = append(req.Params, param)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("marshal json error: %w", err)
	}

	signature, err := requestSignature(key, reqBody)
	if err != nil {
		return "", fmt.Errorf("sign request error: %w", err)
	}

	return signature, nil
}

func (s *BloxrouteClient) GetBundleStats(_ context.Context, _ uint64, _ common.Hash) (GetBundleStatsResponse, error) {
	return GetBundleStatsResponse{}, fmt.Errorf("method not support")
}

func (s *BloxrouteClient) GetUserStats(
	ctx context.Context,
	useV2 bool,
	blockNumber uint64,
) (map[string]any, error) {
	return nil, fmt.Errorf("method not support")
}
