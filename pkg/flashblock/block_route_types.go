package flashblock

import (
	"encoding/json"
	"maps"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// This file contains the protobuf message types for bloXroute Base streamer API
// These types are based on the GetParsedBdnFlashBlockStream response structure

// GetParsedBdnFlashBlockStreamRequest is the request message
type GetParsedBdnFlashBlockStreamRequest struct{}

type GetBdnFlashBlockStreamResponse struct {
	BdnFlashBlock []byte `json:"bdnFlashBlock"`
}

// GetParsedBdnFlashBlockStreamResponse represents a parsed flashblock response
type GetParsedBdnFlashBlockStreamResponse struct {
	PayloadId string    `json:"payloadId"`
	Index     string    `json:"index"`
	Base      *Base     `json:"base,omitempty"`
	Diff      *Diff     `json:"diff,omitempty"`
	Metadata  *Metadata `json:"metadata,omitempty"`
}

// Base contains the base block information (only in index 0)
type Base struct {
	ParentBeaconBlockRoot string       `json:"parentBeaconBlockRoot"`
	ParentHash            string       `json:"parentHash"`
	FeeRecipient          string       `json:"feeRecipient"`
	PrevRandao            string       `json:"prevRandao"`
	BlockNumber           uint64       `json:"blockNumber"`
	GasLimit              uint64       `json:"gasLimit"`
	Timestamp             uint64       `json:"timestamp"`
	ExtraData             string       `json:"extraData"`
	BaseFeePerGas         *hexutil.Big `json:"baseFeePerGas"`
}

func (b *Base) UnmarshalJSON(data []byte) error {
	type Alias Base
	aux := &struct {
		BlockNumber   hexutil.Uint64 `json:"blockNumber"`
		GasLimit      hexutil.Uint64 `json:"gasLimit"`
		Timestamp     hexutil.Uint64 `json:"timestamp"`
		BaseFeePerGas *hexutil.Big   `json:"base_fee_per_gas"`
		*Alias
	}{
		Alias: (*Alias)(b),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	b.BlockNumber = uint64(aux.BlockNumber)
	b.GasLimit = uint64(aux.GasLimit)
	b.Timestamp = uint64(aux.Timestamp)
	if aux.BaseFeePerGas != nil {
		b.BaseFeePerGas = aux.BaseFeePerGas
	}
	return nil
}

// Diff contains the differential block information
type Diff struct {
	StateRoot       string   `json:"stateRoot"`
	ReceiptsRoot    string   `json:"receiptsRoot"`
	LogsBloom       string   `json:"logsBloom"`
	GasUsed         uint64   `json:"gasUsed"`
	BlockHash       string   `json:"blockHash"`
	Transactions    []string `json:"transactions"`
	Withdrawals     []string `json:"withdrawals"`
	WithdrawalsRoot string   `json:"withdrawalsRoot"`
}

func (d *Diff) UnmarshalJSON(data []byte) error {
	type Alias Diff
	aux := &struct {
		GasUsed hexutil.Uint64 `json:"gasUsed"`
		*Alias
	}{
		Alias: (*Alias)(d),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	d.GasUsed = uint64(aux.GasUsed)
	return nil
}

// Metadata contains block metadata including receipts
type Metadata struct {
	BlockNumber        uint64                          `json:"blockNumber"` // "0x2483f21"
	NewAccountBalances map[common.Address]*hexutil.Big `json:"newAccountBalances"`
	Receipts           map[common.Hash]*Receipt        `json:"receipts"`
}

func (b *Metadata) UnmarshalJSON(data []byte) error {
	type Alias Metadata
	aux := &struct {
		BlockNumber string `json:"blockNumber"`
		*Alias
	}{
		Alias: (*Alias)(b),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	blockNo, err := strconv.ParseInt(aux.BlockNumber, 10, 64)
	if err != nil {
		return err
	}
	b.BlockNumber = uint64(blockNo)
	return nil
}

// nolint: godox
// TODO: update logs later
// Receipt represents a transaction receipt
type Receipt struct {
	Eip1559 *Eip1559Receipt `json:"Eip1559,omitempty"`
	Legacy  *LegacyReceipt  `json:"Legacy,omitempty"`
	Deposit *DepositReceipt `json:"Deposit,omitempty"`
	Status  string          `json:"status"`
	Type    string          `json:"type"`
}

// Eip1559Receipt represents an EIP-1559 transaction receipt
type Eip1559Receipt struct {
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	Logs              []*Log `json:"logs"`
	Status            string `json:"status"`
}

// LegacyReceipt represents a legacy transaction receipt
type LegacyReceipt struct {
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	Logs              []*Log `json:"logs"`
	Status            string `json:"status"`
}

// DepositReceipt represents a deposit transaction receipt
type DepositReceipt struct {
	CumulativeGasUsed     string `json:"cumulativeGasUsed"`
	Logs                  []*Log `json:"logs"`
	Status                string `json:"status"`
	DepositNonce          string `json:"depositNonce"`
	DepositReceiptVersion string `json:"depositReceiptVersion"`
}

// Log represents an event log
type Log struct {
	Address string   `json:"address"`
	Topics  []string `json:"topics"`
	Data    string   `json:"data"`
}

// StreamerApiClient is the gRPC client interface
// nolint:lll
type StreamerApiClient interface {
	GetParsedBdnFlashBlockStream(ctx interface{}, in *GetParsedBdnFlashBlockStreamRequest, opts ...interface{}) (StreamerApi_GetParsedBdnFlashBlockStreamClient, error)
}

// StreamerApi_GetParsedBdnFlashBlockStreamClient is the stream client interface
type StreamerApi_GetParsedBdnFlashBlockStreamClient interface {
	Recv() (*GetParsedBdnFlashBlockStreamResponse, error)
}

func convertBloxRouteBaseToFlashblockBase(b *Base) *FlashblockBase {
	if b == nil {
		return nil
	}
	return &FlashblockBase{
		ParentHash:    common.HexToHash(b.ParentHash),
		FeeRecipient:  common.HexToAddress(b.FeeRecipient),
		BlockNumber:   b.BlockNumber,
		GasLimit:      b.GasLimit,
		Timestamp:     b.Timestamp,
		BaseFeePerGas: b.BaseFeePerGas,
	}
}

func convertBloxRouteDiffToFlashblockDiff(d *Diff) *FlashblockDiff {
	if d == nil {
		return nil
	}
	return &FlashblockDiff{
		StateRoot:    common.HexToHash(d.StateRoot),
		BlockHash:    common.HexToHash(d.BlockHash),
		GasUsed:      d.GasUsed,
		Transactions: d.Transactions,
		Withdrawals:  d.Withdrawals,
	}
}

func convertBloxRouteMetadataToFlashblockMeta(m *Metadata) *FlashblockMeta {
	if m == nil {
		return nil
	}
	return &FlashblockMeta{
		BlockNumber:        m.BlockNumber,
		NewAccountBalances: maps.Clone(m.NewAccountBalances),
		Receipts:           maps.Clone(m.Receipts),
	}
}
