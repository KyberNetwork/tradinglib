package fusionorder

import (
	"errors"
	"math"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/exp/slices"
)

var (
	ErrEmptyWhitelist              = errors.New("white list cannot be empty")
	ErrResolvingStartTimeZero      = errors.New("resolving start time can not be 0")
	ErrFeeReceiverZero             = errors.New("fee receiver can not be zero when fee set")
	ErrTooBigDiffBetweenTimestamps = errors.New("too big diff between timestamps")
)

type SettlementPostInteractionData struct {
	Whitelist          []WhitelistItem
	IntegratorFee      IntegratorFee
	BankFee            int64
	ResolvingStartTime int64
	CustomReceiver     common.Address
}

func NewSettlementPostInteractionData(
	whitelist []WhitelistItem,
	integratorFee IntegratorFee,
	bankFee int64,
	resolvingStartTime int64,
	customReceiver common.Address,
) (SettlementPostInteractionData, error) {
	// assert
	if !integratorFee.IsZero() && integratorFee.Ratio > 0 {
		if integratorFee.Receiver.Cmp(common.Address{}) == 0 { // integrator fee receiver is empty
			return SettlementPostInteractionData{}, ErrFeeReceiverZero
		}
	}
	if resolvingStartTime == 0 {
		return SettlementPostInteractionData{}, ErrResolvingStartTimeZero
	}
	return SettlementPostInteractionData{
		Whitelist:          whitelist,
		IntegratorFee:      integratorFee,
		BankFee:            bankFee,
		ResolvingStartTime: resolvingStartTime,
		CustomReceiver:     customReceiver,
	}, nil
}

func NewSettlementPostInteractionDataFromSettlementSuffixData(
	data SettlementSuffixData,
) (SettlementPostInteractionData, error) {
	if len(data.Whitelist) == 0 {
		return SettlementPostInteractionData{}, ErrEmptyWhitelist
	}

	auctionWhitelist := make([]AuctionWhitelistItem, 0, len(data.Whitelist))
	for _, item := range data.Whitelist {
		allowFrom := item.AllowFrom
		if allowFrom < data.ResolvingStartTime {
			allowFrom = data.ResolvingStartTime
		}
		auctionWhitelist = append(auctionWhitelist, AuctionWhitelistItem{
			Address:   item.Address,
			AllowFrom: allowFrom,
		})
	}

	slices.SortFunc(auctionWhitelist, func(a, b AuctionWhitelistItem) int {
		// sort by AllowFrom in ascending order
		return int(a.AllowFrom - b.AllowFrom)
	})

	whitelist := make([]WhitelistItem, 0, len(data.Whitelist))
	sumDelay := int64(0)
	for _, item := range auctionWhitelist {
		delay := item.AllowFrom - data.ResolvingStartTime - sumDelay
		sumDelay += delay

		if delay > math.MaxUint16 {
			return SettlementPostInteractionData{}, ErrTooBigDiffBetweenTimestamps
		}

		whitelist = append(whitelist, WhitelistItem{
			AddressHalf: HalfAddressFromAddress(item.Address),
			Delay:       delay,
		})
	}

	return NewSettlementPostInteractionData(
		whitelist,
		data.IntegratorFee,
		data.BankFee,
		data.ResolvingStartTime,
		data.CustomReceiver,
	)
}

type WhitelistItem struct {
	AddressHalf AddressHalf
	Delay       int64
}

type IntegratorFee struct {
	Ratio    int64
	Receiver common.Address
}

func (f IntegratorFee) IsZero() bool {
	return f == IntegratorFee{}
}

type AuctionWhitelistItem struct {
	Address   common.Address
	AllowFrom int64
}

type SettlementSuffixData struct {
	Whitelist          []AuctionWhitelistItem
	IntegratorFee      IntegratorFee
	BankFee            int64
	ResolvingStartTime int64
	CustomReceiver     common.Address
}

func (s SettlementPostInteractionData) CanExecuteAt(executor common.Address, executionTime int64) bool {
	addressHalf := HalfAddressFromAddress(executor)

	allowedFrom := s.ResolvingStartTime

	for _, item := range s.Whitelist {
		allowedFrom += item.Delay

		if addressHalf == item.AddressHalf {
			return executionTime >= allowedFrom
		} else if executionTime < allowedFrom {
			return false
		}
	}

	return false
}
