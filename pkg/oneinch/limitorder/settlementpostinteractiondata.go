package limitorder

import (
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type WhitelistItem struct {
	/**
	 * last 10 bytes of address, no 0x prefix
	 */
	AddressHalf string
	/**
	 * Delay from previous resolver in seconds
	 * For first resolver delay from `resolvingStartTime`
	 */
	Delay *big.Int
}

type IntegratorFee struct {
	Ratio    *big.Int
	Receiver common.Address
}

type SettlementPostInteractionData struct {
	Whitelist          []WhitelistItem
	IntegratorFee      *IntegratorFee
	BankFee            *big.Int
	ResolvingStartTime *big.Int
	CustomReceiver     common.Address
}

func DecodeSettlementPostInteractionData(data []byte) (SettlementPostInteractionData, error) {
	flags := big.NewInt(int64(data[len(data)-1]))
	bytesWithoutFlags := data[:len(data)-1]

	iter := NewBytesIter(bytesWithoutFlags)
	var (
		bankFee        *big.Int
		integratorFee  *IntegratorFee
		customReceiver common.Address
	)

	if flags.Bit(0) == 1 {
		bankFee = iter.NextUint32()
	}

	if flags.Bit(1) == 1 {
		integratorFee = &IntegratorFee{
			Ratio:    iter.NextUint16(),
			Receiver: common.HexToAddress(iter.NextUint160().Text(16)),
		}

		if flags.Bit(2) == 1 {
			customReceiver = common.HexToAddress(iter.NextUint160().Text(16))
		}
	}

	resolvingStartTime := iter.NextUint32()
	var whitelist []WhitelistItem

	for !iter.IsEmpty() {
		addressHalf := hex.EncodeToString(iter.NextBytes(10))
		delay := iter.NextUint16()
		whitelist = append(whitelist, WhitelistItem{
			AddressHalf: addressHalf,
			Delay:       delay,
		})
	}

	return SettlementPostInteractionData{
		IntegratorFee:      integratorFee,
		BankFee:            bankFee,
		ResolvingStartTime: resolvingStartTime,
		Whitelist:          whitelist,
		CustomReceiver:     customReceiver,
	}, nil
}

func (s SettlementPostInteractionData) IsExclusiveResolver(resolver common.Address) bool {
	addressHalf := resolver.Hex()[len(resolver.Hex())-20:]

	if len(s.Whitelist) == 1 {
		return addressHalf == s.Whitelist[0].AddressHalf
	}

	if s.Whitelist[0].Delay.Cmp(s.Whitelist[1].Delay) == 0 {
		return false
	}

	return addressHalf == s.Whitelist[0].AddressHalf
}
