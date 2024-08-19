package fusionorder

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/decode"
	"github.com/ethereum/go-ethereum/common"
)

const (
	resolverFeeFlag    = 0x01
	integratorFeeFlag  = 0x02
	customReceiverFlag = 0x04
	whitelistShift     = 3
)

func resolverFeeEnabled(flags byte) bool {
	return flags&resolverFeeFlag == resolverFeeFlag
}

func integratorFeeEnabled(flags byte) bool {
	return flags&integratorFeeFlag == integratorFeeFlag
}

func hasCustomReceiver(flags byte) bool {
	return flags&customReceiverFlag == customReceiverFlag
}

func resolversCount(flags byte) byte {
	return flags >> whitelistShift
}

var ErrDataTooShort = errors.New("data is too short")

// DecodeSettlementPostInteractionData decodes SettlementPostInteractionData from bytes
// nolint: gomnd
func DecodeSettlementPostInteractionData(data []byte) (SettlementPostInteractionData, error) {
	// must have at least 1 byte for flags
	if len(data) < 1 {
		return SettlementPostInteractionData{}, ErrDataTooShort
	}

	flags := data[len(data)-1]
	bi := decode.NewBytesIterator(data[:len(data)-1])

	var bankFee uint32
	var integratorFee IntegratorFee
	var customReceiver common.Address
	var err error

	if resolverFeeEnabled(flags) {
		bankFee, err = bi.NextUint32()
		if err != nil {
			return SettlementPostInteractionData{}, fmt.Errorf("get bank fee: %w", err)
		}
	}

	integratorFee, customReceiver, err = decodeIntegratorFee(flags, bi)
	if err != nil {
		return SettlementPostInteractionData{}, fmt.Errorf("get integrator fee: %w", err)
	}

	resolvingStartTime, err := bi.NextUint32()
	if err != nil {
		return SettlementPostInteractionData{}, fmt.Errorf("get resolving start time: %w", err)
	}

	whitelistCount := resolversCount(flags)
	whitelist := make([]WhitelistItem, 0, whitelistCount)

	for i := byte(0); i < whitelistCount; i++ {
		addressHalfBytes, err := bi.NextBytes(addressHalfLength)
		if err != nil {
			return SettlementPostInteractionData{}, fmt.Errorf("get whitelist item address half: %w", err)
		}

		var address AddressHalf
		copy(address[:], addressHalfBytes)

		delay, err := bi.NextUint16()
		if err != nil {
			return SettlementPostInteractionData{}, fmt.Errorf("get whitelist item delay: %w", err)
		}

		whitelist = append(whitelist, WhitelistItem{
			AddressHalf: address,
			Delay:       int64(delay),
		})
	}

	return SettlementPostInteractionData{
		Whitelist:          whitelist,
		IntegratorFee:      integratorFee,
		BankFee:            int64(bankFee),
		ResolvingStartTime: int64(resolvingStartTime),
		CustomReceiver:     customReceiver,
	}, nil
}

func decodeIntegratorFee(
	flags byte, bi *decode.BytesIterator,
) (integratorFee IntegratorFee, customReceiver common.Address, err error) {
	if !integratorFeeEnabled(flags) {
		return integratorFee, customReceiver, nil
	}

	integratorFeeRatio, err := bi.NextUint16()
	if err != nil {
		return integratorFee, customReceiver, fmt.Errorf("get integrator fee ratio: %w", err)
	}

	integratorAddress, err := bi.NextBytes(common.AddressLength)
	if err != nil {
		return integratorFee, customReceiver, fmt.Errorf("get integrator fee address: %w", err)
	}

	integratorFee = IntegratorFee{
		Ratio:    int64(integratorFeeRatio),
		Receiver: common.BytesToAddress(integratorAddress),
	}

	if hasCustomReceiver(flags) {
		customReceiverBytes, err := bi.NextBytes(common.AddressLength)
		if err != nil {
			return integratorFee, customReceiver, fmt.Errorf("get custom receiver: %w", err)
		}
		customReceiver = common.BytesToAddress(customReceiverBytes)
	}

	return integratorFee, customReceiver, nil
}

// Encode encodes SettlementPostInteractionData to bytes
// nolint: gomnd
func (s SettlementPostInteractionData) Encode() []byte {
	buf := new(bytes.Buffer)
	var flags byte
	if s.BankFee != 0 {
		flags |= resolverFeeFlag
		buf.Write(encodeInt64ToBytes(s.BankFee, 4))
	}
	if s.IntegratorFee.Ratio != 0 {
		flags |= integratorFeeFlag
		buf.Write(encodeInt64ToBytes(s.IntegratorFee.Ratio, 2))
		buf.Write(s.IntegratorFee.Receiver.Bytes())
		if s.CustomReceiver != (common.Address{}) {
			flags |= customReceiverFlag
			buf.Write(s.CustomReceiver.Bytes())
		}
	}
	buf.Write(encodeInt64ToBytes(s.ResolvingStartTime, 4))

	for _, wl := range s.Whitelist {
		buf.Write(wl.AddressHalf[:])
		buf.Write(encodeInt64ToBytes(wl.Delay, 2))
	}

	flags |= byte(len(s.Whitelist)) << whitelistShift

	buf.WriteByte(flags)

	return buf.Bytes()
}
