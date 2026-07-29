package flashblock

import (
	"encoding/json"
	"fmt"
	"maps"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

// NewFlashblock is the payload delivered by the newFlashblocks subscription.
//
// Unlike the diff/base/metadata shape described in
// https://docs.base.org/base-chain/api-reference/flashblocks-api/newFlashblocks,
// this feed sends the full accumulated block (header + transactions
// confirmed so far) on every update, so its shape matches a regular
// eth_getBlockByNumber(full=true) response and reuses go-ethereum's block
// types directly.
type NewFlashblock struct {
	Header types.Header

	Uncles       []common.Hash     `json:"uncles"`
	Transactions []Transaction     `json:"transactions"`
	Withdrawals  types.Withdrawals `json:"withdrawals"`
}

// blockBody covers the fields of Block that aren't part of types.Header.
type blockBody struct {
	Uncles       []common.Hash     `json:"uncles"`
	Transactions []Transaction     `json:"transactions"`
	Withdrawals  types.Withdrawals `json:"withdrawals"`
}

// UnmarshalJSON is defined explicitly because Header is embedded logically,
// not anonymously: an anonymous types.Header field would promote its
// UnmarshalJSON method onto Block, which would hijack decoding of the
// whole struct and silently drop Uncles/Transactions/Withdrawals.
func (b *NewFlashblock) UnmarshalJSON(data []byte) error {
	if err := b.Header.UnmarshalJSON(data); err != nil {
		return fmt.Errorf("decode header: %w", err)
	}

	var body blockBody
	if err := json.Unmarshal(data, &body); err != nil {
		return fmt.Errorf("decode body: %w", err)
	}

	b.Uncles = body.Uncles
	b.Transactions = body.Transactions
	b.Withdrawals = body.Withdrawals

	return nil
}

func (b NewFlashblock) MarshalJSON() ([]byte, error) {
	headerJSON, err := b.Header.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("encode header: %w", err)
	}

	var merged map[string]json.RawMessage
	if err := json.Unmarshal(headerJSON, &merged); err != nil {
		return nil, fmt.Errorf("decode header fields: %w", err)
	}

	bodyJSON, err := json.Marshal(blockBody{
		Uncles:       b.Uncles,
		Transactions: b.Transactions,
		Withdrawals:  b.Withdrawals,
	})
	if err != nil {
		return nil, fmt.Errorf("encode body: %w", err)
	}

	var body map[string]json.RawMessage
	if err := json.Unmarshal(bodyJSON, &body); err != nil {
		return nil, fmt.Errorf("decode body fields: %w", err)
	}
	maps.Copy(merged, body)

	out, err := json.Marshal(merged)
	if err != nil {
		return nil, fmt.Errorf("encode merged fields: %w", err)
	}

	return out, nil
}

// Transaction wraps a single block transaction.
//
// go-ethereum's types.Transaction only decodes the transaction types it
// knows about (legacy, access list, dynamic fee, blob, set code) and errors
// on anything else. Base's deposit transactions (type 0x7e) aren't among
// them, so Tx is left nil for those and Raw always keeps the untouched
// payload as a fallback.
type Transaction struct {
	Tx  *types.Transaction
	Raw json.RawMessage
}

func (t *Transaction) UnmarshalJSON(data []byte) error {
	t.Raw = append(json.RawMessage(nil), data...)

	tx := new(types.Transaction)
	if err := tx.UnmarshalJSON(data); err == nil {
		t.Tx = tx
	}

	return nil
}

func (t Transaction) MarshalJSON() ([]byte, error) {
	return t.Raw, nil
}

// DepositTx captures the fields of an Optimism deposit transaction
// (type 0x7e), which is not supported by types.Transaction. Use it to
// decode Transaction.Raw when Transaction.Tx is nil.
type DepositTx struct {
	Type                  hexutil.Uint64 `json:"type"`
	SourceHash            common.Hash    `json:"sourceHash"`
	From                  common.Address `json:"from"`
	To                    common.Address `json:"to"`
	Mint                  *hexutil.Big   `json:"mint"`
	Value                 *hexutil.Big   `json:"value"`
	Gas                   hexutil.Uint64 `json:"gas"`
	Input                 hexutil.Bytes  `json:"input"`
	Hash                  common.Hash    `json:"hash"`
	BlockHash             *common.Hash   `json:"blockHash"`
	BlockNumber           *hexutil.Big   `json:"blockNumber"`
	TransactionIndex      hexutil.Uint64 `json:"transactionIndex"`
	BlockTimestamp        hexutil.Uint64 `json:"blockTimestamp"`
	DepositReceiptVersion hexutil.Uint64 `json:"depositReceiptVersion"`
	Nonce                 hexutil.Uint64 `json:"nonce"`
}
