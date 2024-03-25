package types

import (
	"encoding/json"
	"math/big"

	"github.com/KyberNetwork/tradinglib/pkg/mev"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
)

type Message struct {
	PendingBlockNumber uint64               `json:"pending_block_number"`
	TxHash             string               `json:"tx_hash"`
	SimDebugInfo       SimDebugInfo         `json:"sim_debug_info"`
	InternalTx         *CallFrame           `json:"internal_txs"`
	BaseFee            *big.Int             `json:"base_fee"`
	CurrentBlockTime   uint64               `json:"current_block_time"`
	GasFeeCap          *big.Int             `json:"gas_fee_cap"`
	GasPrice           *big.Int             `json:"gas_price"`
	GasTip             *big.Int             `json:"gas_tip"`
	Gas                uint64               `json:"gas"`
	GasUsed            uint64               `json:"gas_used"`
	From               string               `json:"from"`
	Nonce              uint64               `json:"nonce"`
	Source             mev.BundleSenderType `json:"source"`
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

type CallLog struct {
	Address common.Address `json:"address"`
	Topics  []common.Hash  `json:"topics"`
	Data    hexutil.Bytes  `json:"data"`
}

func (c CallFrame) TypeString() string {
	return c.Type.String()
}

// MarshalJSON marshals as JSON.
func (c CallFrame) MarshalJSON() ([]byte, error) {
	type CallFrame0 struct {
		Type         vm.OpCode       `json:"-"`
		From         common.Address  `json:"from"`
		Gas          hexutil.Uint64  `json:"gas"`
		GasUsed      hexutil.Uint64  `json:"gasUsed"`
		To           *common.Address `json:"to,omitempty"`
		Input        hexutil.Bytes   `json:"input"`
		Output       hexutil.Bytes   `json:"output,omitempty"`
		Error        string          `json:"error,omitempty"`
		RevertReason string          `json:"revertReason,omitempty"`
		Calls        []*CallFrame    `json:"calls,omitempty"`
		Logs         []*CallLog      `json:"logs,omitempty"`
		Value        *hexutil.Big    `json:"value,omitempty"`
		TypeString   string          `json:"type"`
		ContractCall *ContractCall   `json:"contract_call,omitempty"`
	}
	var enc CallFrame0
	enc.Type = c.Type
	enc.From = c.From
	enc.Gas = hexutil.Uint64(c.Gas)
	enc.GasUsed = hexutil.Uint64(c.GasUsed)
	enc.To = c.To
	enc.Input = c.Input
	enc.Output = c.Output
	enc.Error = c.Error
	enc.RevertReason = c.RevertReason
	enc.Calls = c.Calls
	if c.Logs != nil {
		logs := make([]*CallLog, 0, len(c.Logs))
		for _, log := range c.Logs {
			logs = append(logs, &CallLog{
				Address: log.Address,
				Topics:  log.Topics,
				Data:    log.Data,
			})
		}
		enc.Logs = logs
	}
	enc.Value = (*hexutil.Big)(c.Value)
	enc.TypeString = c.TypeString()
	enc.ContractCall = c.ContractCall
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (c *CallFrame) UnmarshalJSON(input []byte) error {
	type CallFrame0 struct {
		Type         *vm.OpCode      `json:"-"`
		From         *common.Address `json:"from"`
		Gas          *hexutil.Uint64 `json:"gas"`
		GasUsed      *hexutil.Uint64 `json:"gasUsed"`
		To           *common.Address `json:"to,omitempty"`
		Input        *hexutil.Bytes  `json:"input"`
		Output       *hexutil.Bytes  `json:"output,omitempty"`
		Error        *string         `json:"error,omitempty"`
		RevertReason *string         `json:"revertReason,omitempty"`
		Calls        []*CallFrame    `json:"calls,omitempty"`
		Logs         []*CallLog      `json:"logs,omitempty"`
		Value        *hexutil.Big    `json:"value,omitempty"`
		ContractCall *ContractCall   `json:"contract_call,omitempty"`
	}
	var dec CallFrame0
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Type != nil {
		c.Type = *dec.Type
	}
	if dec.From != nil {
		c.From = *dec.From
	}
	if dec.Gas != nil {
		c.Gas = uint64(*dec.Gas)
	}
	if dec.GasUsed != nil {
		c.GasUsed = uint64(*dec.GasUsed)
	}
	if dec.To != nil {
		c.To = dec.To
	}
	if dec.Input != nil {
		c.Input = *dec.Input
	}
	if dec.Output != nil {
		c.Output = *dec.Output
	}
	if dec.Error != nil {
		c.Error = *dec.Error
	}
	if dec.RevertReason != nil {
		c.RevertReason = *dec.RevertReason
	}
	if dec.Calls != nil {
		c.Calls = dec.Calls
	}

	if dec.Value != nil {
		c.Value = (*big.Int)(dec.Value)
	}
	if dec.ContractCall != nil {
		c.ContractCall = dec.ContractCall
	}

	logs := make([]*types.Log, 0, len(dec.Logs))
	for _, log := range dec.Logs {
		logs = append(logs, &types.Log{
			Address: log.Address,
			Topics:  log.Topics,
			Data:    log.Data,
		})
	}
	c.Logs = logs
	return nil
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
