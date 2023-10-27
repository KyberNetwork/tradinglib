package paraswap

import (
	"errors"
	"strings"

	rfqTypes "github.com/KyberNetwork/tradinglib/pkg/rfq/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	OrderFilledEvent = "OrderFilled"
)

var ErrInvalidRFQTopic = errors.New("invalid RfqFilled topic")

type Parser struct {
	abi       *abi.ABI
	ps        *OrderFilledFilterer
	eventHash string
}

func MustNewParser() *Parser {
	ps, err := NewOrderFilledFilterer(common.Address{}, nil)
	if err != nil {
		panic(err)
	}
	ABI, err := OrderFilledMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	event, ok := ABI.Events[OrderFilledEvent]
	if !ok {
		panic("no such event: OrderFilled")
	}
	return &Parser{
		ps:        ps,
		abi:       ABI,
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
	o, err := p.ps.ParseOrderFilled(log)
	if err != nil {
		return rfqTypes.RFQTrade{}, err
	}
	res := rfqTypes.RFQTrade{
		OrderHash:        common.Hash(o.OrderHash).String(),
		Maker:            strings.ToLower(o.Maker.Hex()),
		Taker:            strings.ToLower(o.Taker.Hex()),
		MakerToken:       strings.ToLower(o.MakerAsset.Hex()),
		TakerToken:       strings.ToLower(o.TakerAsset.Hex()),
		MakerTokenAmount: o.MakerAmount.String(),
		TakerTokenAmount: o.TakerAmount.String(),
		ContractAddress:  strings.ToLower(o.Raw.Address.String()),
		BlockNumber:      o.Raw.BlockNumber,
		TxHash:           o.Raw.TxHash.String(),
		LogIndex:         uint64(o.Raw.Index),
		Timestamp:        blockTime * 1000,
		EventHash:        p.eventHash,
	}
	return res, nil
}
