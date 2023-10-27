package okx

import (
	"errors"
	"fmt"
	"strings"

	rfqTypes "github.com/KyberNetwork/tradinglib/pkg/rfq/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	FilledEvent = "OrderFilledRFQ"
)

var ErrInvalidOKXFilledTopic = errors.New("invalid KS order filled topic")

type Parser struct {
	abi       *abi.ABI
	ps        *OkxFilterer
	eventHash string
}

func MustNewParser() *Parser {
	ps, err := NewOkxFilterer(common.Address{}, nil)
	if err != nil {
		panic(err)
	}
	ab, err := OkxMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	event, ok := ab.Events[FilledEvent]
	if !ok {
		panic(fmt.Sprintf("no such event: %s", FilledEvent))
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
		return rfqTypes.RFQTrade{}, ErrInvalidOKXFilledTopic
	}
	e, err := p.ps.ParseOrderFilledRFQ(log)
	if err != nil {
		return rfqTypes.RFQTrade{}, err
	}
	res := rfqTypes.RFQTrade{
		OrderHash:        common.Hash(e.OrderHash).String(),
		MakerTokenAmount: e.MakingAmount.String(),
		ContractAddress:  strings.ToLower(e.Raw.Address.String()),
		BlockNumber:      e.Raw.BlockNumber,
		TxHash:           e.Raw.TxHash.String(),
		LogIndex:         uint64(e.Raw.Index),
		Timestamp:        blockTime * 1000,
		EventHash:        p.eventHash,
	}
	return res, nil
}
