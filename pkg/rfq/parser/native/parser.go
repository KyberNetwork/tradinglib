package native

import (
	"errors"
	"strings"

	rfqTypes "github.com/KyberNetwork/tradinglib/pkg/rfq/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	SwapEvent = "Swap"
)

var ErrTradeTopic = errors.New("invalid Trade topic")

type Parser struct {
	abi       *abi.ABI
	ps        *NativeFilterer
	eventHash string
}

func MustNewParser() *Parser {
	ps, err := NewNativeFilterer(common.Address{}, nil)
	if err != nil {
		panic(err)
	}
	ab, err := NativeMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	event, ok := ab.Events[SwapEvent]
	if !ok {
		panic("no such event: Swap")
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
	o, err := p.ps.ParseSwap(log)
	if err != nil {
		return rfqTypes.RFQTrade{}, err
	}
	res := rfqTypes.RFQTrade{
		OrderHash:        common.BytesToHash(o.QuoteId[:]).String(),
		Maker:            strings.ToLower(o.Raw.Address.String()),
		Taker:            strings.ToLower(o.Recipient.String()),
		MakerToken:       strings.ToLower(o.TokenOut.String()),
		TakerToken:       strings.ToLower(o.TokenIn.String()),
		MakerTokenAmount: o.AmountOut.Abs(o.AmountOut).String(),
		TakerTokenAmount: o.AmountIn.String(),
		ContractAddress:  strings.ToLower(o.Raw.Address.String()),
		BlockNumber:      o.Raw.BlockNumber,
		TxHash:           o.Raw.TxHash.String(),
		LogIndex:         uint64(o.Raw.Index),
		Timestamp:        blockTime * 1000,
		EventHash:        p.eventHash,
	}
	return res, nil
}
