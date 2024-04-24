package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type FlashbotMevshareTxHint struct {
	Hash             *common.Hash    `json:"hash,omitempty"`
	To               *common.Address `json:"to,omitempty"`
	FunctionSelector *hexutil.Bytes  `json:"functionSelector,omitempty"`
	CallData         *hexutil.Bytes  `json:"callData,omitempty"`
}
