package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type FlashbotMevshareLog struct {
	// address of the contract that generated the event
	Address common.Address `json:"address"`
	// list of topics provided by the contract.
	Topics []common.Hash `json:"topics"`
	// supplied by the contract, usually ABI-encoded
	Data hexutil.Bytes `json:"data"`
}

type FlashbotMevShareTxHint struct {
	Hash             *common.Hash    `json:"hash,omitempty"`
	To               *common.Address `json:"to,omitempty"`
	FunctionSelector *hexutil.Bytes  `json:"functionSelector,omitempty"`
	CallData         *hexutil.Bytes  `json:"callData,omitempty"`
}

type FlashbotMevshareEvent struct {
	Hash        common.Hash              `json:"hash"`
	Logs        []FlashbotMevshareLog    `json:"logs"`
	Txs         []FlashbotMevShareTxHint `json:"txs"`
	MevGasPrice *hexutil.Big             `json:"mevGasPrice,omitempty"`
	GasUsed     *hexutil.Uint64          `json:"gasUsed,omitempty"`
}
