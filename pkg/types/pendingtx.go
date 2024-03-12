package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
)

type PendingTxSource string

const (
	PendingTxBloxroute    PendingTxSource = "bloxroute"
	PendingTxFlashbots    PendingTxSource = "flashbots"
	PendingTxBeaverBuild  PendingTxSource = "beaverbuild"
	PendingTxRsyncBuilder PendingTxSource = "rsync-builder"
	PendingTxTitan        PendingTxSource = "titan"
	PendingTxPublicRPC    PendingTxSource = "public-rpc"
)

type Message struct {
	PendingBlockNumber uint64       `json:"pending_block_number"`
	TxHash             string       `json:"tx_hash"`
	SimDebugInfo       SimDebugInfo `json:"sim_debug_info"`
	InternalTx         *CallFrame   `json:"internal_txs"`
	BaseFee            *big.Int     `json:"base_fee"`
	CurrentBlockTime   uint64       `json:"current_block_time"`
	GasFeeCap          *big.Int     `json:"gas_fee_cap"`
	GasPrice           *big.Int     `json:"gas_price"`
	GasTip             *big.Int     `json:"gas_tip"`
	Gas                uint64       `json:"gas"`
	GasUsed            uint64       `json:"gas_used"`
	From               string       `json:"from"`
	Nonce              uint64       `json:"nonce"`
}

type SimDebugInfo struct {
	E2EMs                 int64 `json:"e2e_ms"`
	DetectTimeMs          int64 `json:"detect_time_ms"`
	StartTraceTimeMs      int64 `json:"start_trace_time_ms"`
	EndTraceTimeMs        int64 `json:"end_trace_time_ms"`
	StartSimulationTimeMs int64 `json:"start_simulation_time_ms"`
	EndSimulationTimeMs   int64 `json:"end_simulation_time_ms"`
}

type CallFrame struct {
	Type         vm.OpCode       `json:"-"`
	From         common.Address  `json:"from"`
	Gas          uint64          `json:"gas"`
	GasUsed      uint64          `json:"gasUsed"`
	To           *common.Address `json:"to,omitempty"`
	Input        []byte          `json:"input"`
	Output       []byte          `json:"output,omitempty"`
	Error        string          `json:"error,omitempty"`
	RevertReason string          `json:"revertReason,omitempty"`
	Calls        []*CallFrame    `json:"calls,omitempty"`
	Logs         []*types.Log    `json:"logs,omitempty"`
	// Placed at end on purpose. The RLP will be decoded to 0 instead of
	// nil if there are non-empty elements after in the struct.
	Value *big.Int `json:"value,omitempty"`

	// contract call fields
	ContractCall *ContractCall `json:"contract_call,omitempty"`
}

type ContractCallParam struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
}

type ContractCall struct {
	ContractType string              `json:"contract_type,omitempty"`
	Name         string              `json:"name"`
	Params       []ContractCallParam `json:"params"`
}

type Transaction struct {
	From  common.Address `json:"from"`
	Nonce uint64         `json:"nonce"`
}

type MinedBlock struct {
	BlockNumber  int64         `json:"block_number"`
	BlockHash    string        `json:"block_hash"`
	BlockTime    int64         `json:"block_time"`
	Transactions []Transaction `json:"transactions"`
}
