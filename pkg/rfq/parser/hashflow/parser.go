package hashflow

import (
	"errors"
	"strings"

	rfqTypes "github.com/KyberNetwork/tradinglib/pkg/rfq/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	TradeEvent = "Trade"
)

var ErrTradeTopic = errors.New("invalid Trade topic")

type Parser struct {
	abi       *abi.ABI
	ps        *HashflowFilterer
	eventHash string
}

func MustNewParser() *Parser {
	ps, err := NewHashflowFilterer(common.Address{}, nil)
	if err != nil {
		panic(err)
	}
	ab, err := HashflowMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	event, ok := ab.Events[TradeEvent]
	if !ok {
		panic("no such event: Trade")
	}
	return &Parser{
		ps:        ps,
		abi:       ab,
		eventHash: event.ID.String(),
	}
}

func (p *Parser) Topics() []string {
	return []string{
		p.eventHash,
	}
}

func (p *Parser) Parse(log types.Log, blockTime uint64) (rfqTypes.RFQTrade, error) {
	if len(log.Topics) > 0 && log.Topics[0].Hex() != p.eventHash {
		return rfqTypes.RFQTrade{}, ErrTradeTopic
	}
	o, err := p.ps.ParseTrade(log)
	if err != nil {
		return rfqTypes.RFQTrade{}, err
	}
	res := rfqTypes.RFQTrade{
		OrderHash:        common.Hash(o.Txid).String(),
		Maker:            strings.ToLower(log.Address.Hex()),
		Taker:            strings.ToLower(o.Trader.Hex()),
		MakerToken:       strings.ToLower(o.QuoteToken.Hex()),
		TakerToken:       strings.ToLower(o.BaseToken.Hex()),
		MakerTokenAmount: o.QuoteTokenAmount.String(),
		TakerTokenAmount: o.BaseTokenAmount.String(),
		ContractAddress:  strings.ToLower(o.Raw.Address.String()),
		BlockNumber:      o.Raw.BlockNumber,
		TxHash:           o.Raw.TxHash.String(),
		LogIndex:         uint64(o.Raw.Index),
		Timestamp:        blockTime * 1000,
		EventHash:        p.eventHash,
	}
	return res, nil
}
