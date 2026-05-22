package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
)

type SimulatedPrivateMempoolLog struct {
	// address of the contract that generated the event
	Address common.Address `json:"address"`
	// list of topics provided by the contract.
	Topics []common.Hash `json:"topics"`
	// supplied by the contract, usually ABI-encoded
	Data hexutil.Bytes `json:"data"`
}

type FlashbotMevShareTxHint struct {
	Hash                 *common.Hash          `json:"hash,omitempty"`
	To                   *common.Address       `json:"to,omitempty"`
	FunctionSelector     *hexutil.Bytes        `json:"functionSelector,omitempty"`
	CallData             *hexutil.Bytes        `json:"callData,omitempty"`
	From                 *common.Address       `json:"from,omitempty"`
	Value                *hexutil.Big          `json:"value,omitempty"`
	MaxFeePerGas         *hexutil.Big          `json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas *hexutil.Big          `json:"maxPriorityFeePerGas,omitempty"`
	Nonce                *hexutil.Uint64       `json:"nonce,omitempty"`
	ChainID              *hexutil.Big          `json:"chainId,omitempty"`
	AccessList           *gethtypes.AccessList `json:"accessList,omitempty"`
	Gas                  *hexutil.Uint64       `json:"gas,omitempty"`
	Type                 *hexutil.Uint64       `json:"type,omitempty"`
}

type FlashbotMevshareEvent struct {
	Hash        common.Hash                  `json:"hash"`
	Logs        []SimulatedPrivateMempoolLog `json:"logs"`
	Txs         []FlashbotMevShareTxHint     `json:"txs"`
	MevGasPrice *hexutil.Big                 `json:"mevGasPrice,omitempty"`
	GasUsed     *hexutil.Uint64              `json:"gasUsed,omitempty"`
}
