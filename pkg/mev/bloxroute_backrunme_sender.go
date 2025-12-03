package mev

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/flashbots/mev-share-node/mevshare"
)

type BloxrouteBackrunmeSender struct {
	httpClient *http.Client
	endpoint   string
	authHeader string
}

type backrunmeRequest struct {
	Method string                 `json:"method"`
	ID     string                 `json:"id"`
	Params backrunmeRequestParams `json:"params"`
}

type backrunmeRequestParams struct {
	TransactionHash  string   `json:"transaction_hash"`
	Transaction      []string `json:"transaction"`
	BlockNumber      string   `json:"block_number"`
	StateBlockNumber string   `json:"state_block_number,omitempty"`
	Timestamp        *uint64  `json:"timestamp,omitempty"`
	MinTimestamp     *uint64  `json:"min_timestamp,omitempty"`
	MaxTimestamp     *uint64  `json:"max_timestamp,omitempty"`
}

type backrunmeResponse struct {
	JSONRPC string                  `json:"jsonrpc"`
	ID      string                  `json:"id"`
	Result  backrunmeResponseResult `json:"result,omitempty"`
	Error   *backrunmeResponseError `json:"error,omitempty"`
}

type backrunmeResponseResult struct {
	BundleHash string `json:"bundleHash"`
}

type simulateResult struct {
	BloxrouteDiff     string             `json:"bloxrouteDiff"`
	BundleGasPrice    string             `json:"bundleGasPrice"`
	BundleHash        string             `json:"bundleHash"`
	EthSentToCoinbase string             `json:"ethSentToCoinbase"`
	GasFees           string             `json:"gasFees"`
	MinerDiff         string             `json:"minerDiff"`
	Results           []simulateTxResult `json:"results"`
	SenderDiff        string             `json:"senderDiff"`
	StateBlockNumber  uint64             `json:"stateBlockNumber"`
	TotalGasUsed      uint64             `json:"totalGasUsed"`
	Status            string             `json:"status"`
}

type simulateTxResult struct {
	GasUsed uint64 `json:"gasUsed"`
	TxHash  string `json:"txHash"`
	Value   string `json:"value"`
}

type simulateResponse struct {
	JSONRPC string                  `json:"jsonrpc"`
	ID      string                  `json:"id"`
	Result  simulateResult          `json:"result,omitempty"`
	Error   *backrunmeResponseError `json:"error,omitempty"`
}

type backrunmeResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const (
	DefaultEndpoint             = "https://backrunme.blxrbdn.com"
	SubmitArbOnlyBundleMethod   = "submit_arb_only_bundle"
	SimulateArbOnlyBundleMethod = "simulate_arb_only_bundle"
)

func NewBloxrouteBackrunmeSender(authHeader, endpoint string) (*BloxrouteBackrunmeSender, error) {
	if endpoint == "" {
		endpoint = DefaultEndpoint
	}

	// Create HTTP client with insecure TLS (as shown in curl --insecure)
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	return &BloxrouteBackrunmeSender{
		httpClient: httpClient,
		endpoint:   endpoint,
		authHeader: authHeader,
	}, nil
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
	// Validate inputs
	if len(pendingTxHashes) == 0 {
		return SendBundleResponse{}, fmt.Errorf("at least one pending transaction hash is required")
	}
	if len(txs) == 0 {
		return SendBundleResponse{}, fmt.Errorf("at least one backrun transaction is required")
	}

	// Encode backrun transactions (without 0x prefix as per API spec)
	transactions := make([]string, 0, len(txs))
	for _, tx := range txs {
		txBin, err := tx.MarshalBinary()
		if err != nil {
			return SendBundleResponse{}, fmt.Errorf("marshal tx: %w", err)
		}
		hexTx := hexutil.Encode(txBin)
		// Remove 0x prefix as required by the API
		hexTx = strings.TrimPrefix(hexTx, "0x")
		transactions = append(transactions, hexTx)
	}

	// First, simulate the bundle using MevSimulateBundle
	_, err := s.MevSimulateBundle(ctx, blockNumber, pendingTxHashes[0], txs[0])
	if err != nil {
		return SendBundleResponse{}, fmt.Errorf("simulate bundle failed: %w", err)
	}

	// If simulation passed, proceed with submission
	// Build request params
	params := backrunmeRequestParams{
		TransactionHash: pendingTxHashes[0].Hex(),
		Transaction:     transactions,
		BlockNumber:     fmt.Sprintf("0x%x", blockNumber),
	}

	// Build request
	req := backrunmeRequest{
		Method: SubmitArbOnlyBundleMethod,
		ID:     "1",
		Params: params,
	}

	// Marshal request body
	reqBody, err := json.Marshal(req)
	if err != nil {
		return SendBundleResponse{}, fmt.Errorf("marshal request: %w", err)
	}

	// Create HTTP request
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, s.endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return SendBundleResponse{}, fmt.Errorf("create http request: %w", err)
	}

	// Set headers
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", s.authHeader)

	// Send request
	resp, err := s.httpClient.Do(httpReq)
	if err != nil {
		return SendBundleResponse{}, fmt.Errorf("send http request: %w", err)
	}
	defer resp.Body.Close()

	// Parse response
	var backrunmeResp backrunmeResponse
	if err := json.NewDecoder(resp.Body).Decode(&backrunmeResp); err != nil {
		return SendBundleResponse{}, fmt.Errorf("decode response: %w", err)
	}

	// Check for errors
	if backrunmeResp.Error != nil && backrunmeResp.Error.Message != "" {
		return SendBundleResponse{}, fmt.Errorf("backrunme error [%d]: %s",
			backrunmeResp.Error.Code, backrunmeResp.Error.Message)
	}

	return SendBundleResponse{
		Result: SendBundleResult{
			BundleHash: backrunmeResp.Result.BundleHash,
		},
	}, nil
}

func (s *BloxrouteBackrunmeSender) MevSimulateBundle(
	ctx context.Context,
	blockNumber uint64,
	pendingTxHash common.Hash,
	tx *types.Transaction,
) (*mevshare.SimMevBundleResponse, error) {
	// Encode backrun transaction (without 0x prefix as per API spec)
	txBin, err := tx.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("marshal tx: %w", err)
	}
	hexTx := hexutil.Encode(txBin)
	// Remove 0x prefix as required by the API
	hexTx = strings.TrimPrefix(hexTx, "0x")
	transactions := []string{hexTx}

	// Build request params
	params := backrunmeRequestParams{
		TransactionHash:  pendingTxHash.Hex(),
		Transaction:      transactions,
		BlockNumber:      fmt.Sprintf("0x%x", blockNumber),
		StateBlockNumber: "latest",
	}

	// Build request
	req := backrunmeRequest{
		Method: SimulateArbOnlyBundleMethod,
		ID:     tx.ChainId().String(),
		Params: params,
	}

	// Marshal request body
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal simulate request: %w", err)
	}

	// Create HTTP request
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, s.endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("create simulate http request: %w", err)
	}

	// Set headers
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", s.authHeader)

	// Send request
	resp, err := s.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("send simulate http request: %w", err)
	}
	defer resp.Body.Close()

	// Parse response
	var simResp simulateResponse
	if err := json.NewDecoder(resp.Body).Decode(&simResp); err != nil {
		return nil, fmt.Errorf("decode simulate response: %w", err)
	}

	// Check for errors
	if simResp.Error != nil && simResp.Error.Message != "" {
		return nil, fmt.Errorf("simulate error [%d]: %s",
			simResp.Error.Code, simResp.Error.Message)
	}

	// Check simulation status
	if simResp.Result.Status != "good" {
		return nil, fmt.Errorf("simulation status not good: %s", simResp.Result.Status)
	}

	// Return simulation result as bundle hash for consistency
	return &mevshare.SimMevBundleResponse{
		Success:    true,
		StateBlock: hexutil.Uint64(simResp.Result.StateBlockNumber),
		GasUsed:    hexutil.Uint64(simResp.Result.TotalGasUsed),
	}, nil
}

func (s *BloxrouteBackrunmeSender) GetSenderType() BundleSenderType {
	return BundleSenderTypeBloxrouteBackrunme
}
