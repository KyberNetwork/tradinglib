package tokenlon

import (
	"errors"
	"strings"

	rfqTypes "github.com/KyberNetwork/tradinglib/pkg/rfq/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	FillOrderEvent = "FillOrder"
)

var ErrInvalidRFQTopic = errors.New("invalid RfqFilled topic")

type Parser struct {
	abi       *abi.ABI
	ps        *FillOrderFilterer
	eventHash string
}

func MustNewParser() *Parser {
	ps, err := NewFillOrderFilterer(common.Address{}, nil)
	if err != nil {
		panic(err)
	}
	ab, err := FillOrderMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	event, ok := ab.Events[FillOrderEvent]
	if !ok {
		panic("no such event: FillOrder")
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
	o, err := p.ps.ParseFillOrder(log)
	if err != nil {
		return rfqTypes.RFQTrade{}, err
	}
	res := rfqTypes.RFQTrade{
		OrderHash:        common.Hash(o.OrderHash).String(),
		Maker:            strings.ToLower(o.MakerAddr.Hex()),
		Taker:            strings.ToLower(o.ReceiverAddr.Hex()),
		MakerToken:       strings.ToLower(o.MakerAssetAddr.Hex()),
		TakerToken:       strings.ToLower(o.TakerAssetAddr.Hex()),
		MakerTokenAmount: o.MakerAssetAmount.String(),
		TakerTokenAmount: o.TakerAssetAmount.String(),
		ContractAddress:  strings.ToLower(o.Raw.Address.String()),
		BlockNumber:      o.Raw.BlockNumber,
		TxHash:           o.Raw.TxHash.String(),
		LogIndex:         uint64(o.Raw.Index),
		Timestamp:        blockTime * 1000,
		EventHash:        p.eventHash,
	}
	return res, nil
}
