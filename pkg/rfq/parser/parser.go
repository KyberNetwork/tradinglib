package parser

import (
	rfqTypes "github.com/KyberNetwork/tradinglib/pkg/rfq/types"
	"github.com/ethereum/go-ethereum/core/types"
)

type Parser interface {
	Parse(log types.Log, blockTime uint64) (rfqTypes.RFQTrade, error)
	Topics() []string
}
