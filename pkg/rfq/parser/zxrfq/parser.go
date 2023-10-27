package zxrfq

import (
	"errors"
	"strings"

	rfqTypes "github.com/KyberNetwork/tradinglib/pkg/rfq/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	RfqOrderFilledEvent = "RfqOrderFilled"
)

var ErrInvalidRFQTopic = errors.New("invalid RfqFilled topic")

type Parser struct {
	abi       *abi.ABI
	ps        *RfqOrderFilledFilterer
	eventHash string
}

func MustNewParser() *Parser {
	ps, err := NewRfqOrderFilledFilterer(common.Address{}, nil)
	if err != nil {
		panic(err)
	}
	ab, err := RfqOrderFilledMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	event, ok := ab.Events[RfqOrderFilledEvent]
	if !ok {
		panic("no such event: RfqOrderFilled")
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
		return rfqTypes.RFQTrade{}, ErrInvalidRFQTopic
	}
	o, err := p.ps.ParseRfqOrderFilled(log)
	if err != nil {
		return rfqTypes.RFQTrade{}, err
	}
	res := rfqTypes.RFQTrade{
		OrderHash:        common.Hash(o.OrderHash).String(),
		Maker:            strings.ToLower(o.Maker.Hex()),
		Taker:            strings.ToLower(o.Taker.Hex()),
		MakerToken:       strings.ToLower(o.MakerToken.Hex()),
		TakerToken:       strings.ToLower(o.TakerToken.Hex()),
		MakerTokenAmount: o.MakerTokenFilledAmount.String(),
		TakerTokenAmount: o.TakerTokenFilledAmount.String(),
		ContractAddress:  strings.ToLower(o.Raw.Address.String()),
		BlockNumber:      o.Raw.BlockNumber,
		TxHash:           o.Raw.TxHash.String(),
		LogIndex:         uint64(o.Raw.Index),
		Timestamp:        blockTime * 1000,
		EventHash:        p.eventHash,
	}
	return res, nil
}
