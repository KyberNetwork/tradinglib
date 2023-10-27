package parser

import (
	"github.com/KyberNetwork/tradinglib/pkg/rfq/parser/hashflow"
	hashflowv3 "github.com/KyberNetwork/tradinglib/pkg/rfq/parser/hashflow_v3"
	"github.com/KyberNetwork/tradinglib/pkg/rfq/parser/kyberswap"
	kyberswaprfq "github.com/KyberNetwork/tradinglib/pkg/rfq/parser/kyberswap_rfq"
	"github.com/KyberNetwork/tradinglib/pkg/rfq/parser/native"
	"github.com/KyberNetwork/tradinglib/pkg/rfq/parser/paraswap"
	"github.com/KyberNetwork/tradinglib/pkg/rfq/parser/tokenlon"
	"github.com/KyberNetwork/tradinglib/pkg/rfq/parser/zxotc"
	"github.com/KyberNetwork/tradinglib/pkg/rfq/parser/zxrfq"
	rfqTypes "github.com/KyberNetwork/tradinglib/pkg/rfq/types"
	"github.com/ethereum/go-ethereum/core/types"
)

type Parser interface {
	Parse(log types.Log, blockTime uint64) (rfqTypes.RFQTrade, error)
	Topics() []string
}

func NewRFQParsers() map[string]Parser {
	rfqParsers := make(map[string]Parser)

	addRFQParser := func(parser Parser) {
		for _, topic := range parser.Topics() {
			rfqParsers[topic] = parser
		}
	}

	addRFQParser(hashflow.MustNewParser())
	addRFQParser(hashflowv3.MustNewParser())
	addRFQParser(kyberswap.MustNewParser())
	addRFQParser(zxotc.MustNewParser())
	addRFQParser(zxrfq.MustNewParser())
	addRFQParser(tokenlon.MustNewParser())
	addRFQParser(paraswap.MustNewParser())
	addRFQParser(native.MustNewParser())
	addRFQParser(kyberswaprfq.MustNewParser())

	return rfqParsers
}
