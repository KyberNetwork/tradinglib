package flashblock

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Flashblock struct {
	PayloadID string          `json:"payload_id"`
	Index     int64           `json:"index"`
	Base      *FlashblockBase `json:"base"`
	Diff      *FlashblockDiff `json:"diff"`
	Metadata  *FlashblockMeta `json:"metadata"`
}

type FlashblockBase struct {
	ParentHash    common.Hash    `json:"parent_hash"`
	FeeRecipient  common.Address `json:"fee_recipient"`
	BlockNumber   uint64         `json:"block_number"`
	GasLimit      uint64         `json:"gas_limit"`
	Timestamp     uint64         `json:"timestamp"`
	BaseFeePerGas *hexutil.Big   `json:"base_fee_per_gas"`
}

type FlashblockDiff struct {
	StateRoot    common.Hash `json:"state_root"`
	BlockHash    common.Hash `json:"block_hash"`
	GasUsed      uint64      `json:"gas_used"`
	Transactions []string    `json:"transactions"`
	Withdrawals  []string    `json:"withdrawals"`
}

type FlashblockMeta struct {
	BlockNumber        uint64                          `json:"block_number"`
	NewAccountBalances map[common.Address]*hexutil.Big `json:"new_account_balances"`
	Receipts           map[common.Hash]*Receipt        `json:"receipts"`
}

// --- Unmarshal Implementations ---
func (b *FlashblockBase) UnmarshalJSON(data []byte) error {
	type Alias FlashblockBase
	aux := &struct {
		BlockNumber hexutil.Uint64 `json:"block_number"`
		GasLimit    hexutil.Uint64 `json:"gas_limit"`
		Timestamp   hexutil.Uint64 `json:"timestamp"`
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
	return nil
}

func (d *FlashblockDiff) UnmarshalJSON(data []byte) error {
	type Alias FlashblockDiff
	aux := &struct {
		GasUsed hexutil.Uint64 `json:"gas_used"`
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

// --- Marshal Implementations ---

func (b FlashblockBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ParentHash    common.Hash    `json:"parent_hash"`
		FeeRecipient  common.Address `json:"fee_recipient"`
		BlockNumber   hexutil.Uint64 `json:"block_number"`
		GasLimit      hexutil.Uint64 `json:"gas_limit"`
		Timestamp     hexutil.Uint64 `json:"timestamp"`
		BaseFeePerGas *hexutil.Big   `json:"base_fee_per_gas"`
	}{
		ParentHash:    b.ParentHash,
		FeeRecipient:  b.FeeRecipient,
		BlockNumber:   hexutil.Uint64(b.BlockNumber),
		GasLimit:      hexutil.Uint64(b.GasLimit),
		Timestamp:     hexutil.Uint64(b.Timestamp),
		BaseFeePerGas: b.BaseFeePerGas,
	})
}

func (d FlashblockDiff) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		StateRoot    common.Hash    `json:"state_root"`
		BlockHash    common.Hash    `json:"block_hash"`
		GasUsed      hexutil.Uint64 `json:"gas_used"`
		Transactions []string       `json:"transactions"`
		Withdrawals  []string       `json:"withdrawals"`
	}{
		StateRoot:    d.StateRoot,
		BlockHash:    d.BlockHash,
		GasUsed:      hexutil.Uint64(d.GasUsed),
		Transactions: d.Transactions,
		Withdrawals:  d.Withdrawals,
	})
}
