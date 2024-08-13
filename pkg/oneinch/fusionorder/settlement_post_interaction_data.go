package fusionorder

import (
	"errors"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/exp/slices"
)

const (
	addressHalfLength = common.AddressLength / 2
)

type AddressHalf [addressHalfLength]byte

var (
	ErrEmptyWhitelist              = errors.New("white list cannot be empty")
	ErrResolvingStartTimeNil       = errors.New("resolving start time can not be nil")
	ErrFeeReceiverZero             = errors.New("fee receiver can not be zero when fee set")
	ErrTooBigDiffBetweenTimestamps = errors.New("too big diff between timestamps")
)

type SettlementPostInteractionData struct {
	Whitelist          []WhitelistItem
	IntegratorFee      IntegratorFee
	BankFee            *big.Int
	ResolvingStartTime *big.Int
	CustomReceiver     common.Address
}

func NewSettlementPostInteractionData(
	whitelist []WhitelistItem,
	integratorFee IntegratorFee,
	bankFee *big.Int,
	resolvingStartTime *big.Int,
	customReceiver common.Address,
) (SettlementPostInteractionData, error) {
	// set default
	if bankFee == nil {
		bankFee = big.NewInt(0)
	}

	// assert
	if !integratorFee.IsZero() && integratorFee.Ratio.Cmp(big.NewInt(0)) != 0 {
		if integratorFee.Receiver.Cmp(common.Address{}) == 0 { // integrator fee receiver is empty
			return SettlementPostInteractionData{}, ErrFeeReceiverZero
		}
	}
	if resolvingStartTime == nil {
		return SettlementPostInteractionData{}, ErrResolvingStartTimeNil
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
		if allowFrom.Cmp(data.ResolvingStartTime) == -1 { // allowFrom < resolvingStartTime
			allowFrom = data.ResolvingStartTime
		}
		auctionWhitelist = append(auctionWhitelist, AuctionWhitelistItem{
			Address:   item.Address,
			AllowFrom: allowFrom,
		})
	}

	slices.SortFunc(auctionWhitelist, func(a, b AuctionWhitelistItem) int {
		return a.AllowFrom.Cmp(b.AllowFrom) // sort by AllowFrom in ascending order
	})

	whitelist := make([]WhitelistItem, 0, len(data.Whitelist))
	sumDelay := big.NewInt(0)
	for _, item := range auctionWhitelist {
		delay := new(big.Int).Sub(new(big.Int).Sub(item.AllowFrom, data.ResolvingStartTime), sumDelay)
		sumDelay = new(big.Int).Add(sumDelay, delay)

		if delay.Cmp(new(big.Int).SetUint64(math.MaxUint16)) >= 0 { // delay >= math.MaxUint16
			return SettlementPostInteractionData{}, ErrTooBigDiffBetweenTimestamps
		}

		var addressHalf AddressHalf
		copy(addressHalf[:], item.Address.Bytes()[common.AddressLength-addressHalfLength:]) // take the last 10 bytes
		whitelist = append(whitelist, WhitelistItem{
			AddressHalf: addressHalf,
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
	Delay       *big.Int
}

type IntegratorFee struct {
	Ratio    *big.Int
	Receiver common.Address
}

func (f IntegratorFee) IsZero() bool {
	return f == IntegratorFee{}
}

type AuctionWhitelistItem struct {
	Address   common.Address
	AllowFrom *big.Int
}

type SettlementSuffixData struct {
	Whitelist          []AuctionWhitelistItem
	IntegratorFee      IntegratorFee
	BankFee            *big.Int
	ResolvingStartTime *big.Int
	CustomReceiver     common.Address
}
