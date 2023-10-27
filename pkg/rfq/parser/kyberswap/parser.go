package kyberswap

import (
	"errors"
	"strings"

	rfqTypes "github.com/KyberNetwork/tradinglib/pkg/rfq/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	SwappedEvent = "Swapped"
)

var ErrInvalidKSSwappedTopic = errors.New("invalid KS Swapped topic")

type Parser struct {
	abi       *abi.ABI
	ps        *SwappedFilterer
	eventHash string
}

func MustNewParser() *Parser {
	ps, err := NewSwappedFilterer(common.Address{}, nil)
	if err != nil {
		panic(err)
	}
	ab, err := SwappedMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	swapEvent, ok := ab.Events[SwappedEvent]
	if !ok {
		panic("no such event: Swapped")
	}
	return &Parser{
		ps:        ps,
		abi:       ab,
		eventHash: swapEvent.ID.String(),
	}
}

func (p *Parser) Topics() []string {
	return []string{
		p.eventHash,
	}
}

func (p *Parser) Parse(log types.Log, blockTime uint64) (rfqTypes.RFQTrade, error) {
	if len(log.Topics) > 0 && log.Topics[0].Hex() != p.eventHash {
		return rfqTypes.RFQTrade{}, ErrInvalidKSSwappedTopic
	}
	e, err := p.ps.ParseSwapped(log)
	if err != nil {
		return rfqTypes.RFQTrade{}, err
	}
	res := rfqTypes.RFQTrade{
		Taker:            strings.ToLower(e.DstReceiver.String()),
		MakerToken:       strings.ToLower(e.DstToken.String()),
		TakerToken:       strings.ToLower(e.SrcToken.String()),
		MakerTokenAmount: e.ReturnAmount.String(),
		TakerTokenAmount: e.SpentAmount.String(),
		ContractAddress:  strings.ToLower(e.Raw.Address.String()),
		BlockNumber:      e.Raw.BlockNumber,
		TxHash:           e.Raw.TxHash.String(),
		LogIndex:         uint64(e.Raw.Index),
		Timestamp:        blockTime * 1000,
		EventHash:        p.eventHash,
	}
	return res, nil
}
