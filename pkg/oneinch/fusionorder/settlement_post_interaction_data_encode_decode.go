package fusionorder

import (
	"bytes"
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

// DecodeSettlementPostInteractionData decodes SettlementPostInteractionData from bytes
// nolint: gomnd
func DecodeSettlementPostInteractionData(data []byte) (SettlementPostInteractionData, error) {
	// must have at least 1 byte for flags
	if err := decode.ValidateDataLength(data, 1); err != nil {
		return SettlementPostInteractionData{}, err
	}

	flags := data[len(data)-1]

	var bankFee int64
	var integratorFee IntegratorFee
	var customReceiver common.Address
	var resolvingStartTime int64
	var err error

	if resolverFeeEnabled(flags) {
		const lengthBankFee = 4
		bankFee, data, err = decode.NextInt64(data, lengthBankFee)
		if err != nil {
			return SettlementPostInteractionData{}, fmt.Errorf("get bank fee: %w", err)
		}
	}

	integratorFee, customReceiver, data, err = decodeIntegratorFee(flags, data)
	if err != nil {
		return SettlementPostInteractionData{}, fmt.Errorf("get integrator fee: %w", err)
	}

	resolvingStartTime, data, err = decode.NextInt64(data, 4)
	if err != nil {
		return SettlementPostInteractionData{}, fmt.Errorf("get resolving start time: %w", err)
	}

	whitelistCount := resolversCount(flags)
	whitelist := make([]WhitelistItem, 0, whitelistCount)

	for i := byte(0); i < whitelistCount; i++ {
		var addressHalfBytes []byte
		addressHalfBytes, data, err = decode.Next(data, addressHalfLength)
		if err != nil {
			return SettlementPostInteractionData{}, fmt.Errorf("get whitelist item address half: %w", err)
		}

		var address AddressHalf
		copy(address[:], addressHalfBytes)

		var delay int64
		delay, data, err = decode.NextInt64(data, 2)
		if err != nil {
			return SettlementPostInteractionData{}, fmt.Errorf("get whitelist item delay: %w", err)
		}

		whitelist = append(whitelist, WhitelistItem{
			AddressHalf: address,
			Delay:       delay,
		})
	}

	return SettlementPostInteractionData{
		Whitelist:          whitelist,
		IntegratorFee:      integratorFee,
		BankFee:            bankFee,
		ResolvingStartTime: resolvingStartTime,
		CustomReceiver:     customReceiver,
	}, nil
}

func decodeIntegratorFee(
	flags byte, data []byte,
) (integratorFee IntegratorFee, customReceiver common.Address, remainingData []byte, err error) {
	if !integratorFeeEnabled(flags) {
		return integratorFee, customReceiver, data, nil
	}

	var integratorFeeRatio int64
	integratorFeeRatio, data, err = decode.NextInt64(data, 2) // nolint: gomnd
	if err != nil {
		return integratorFee, customReceiver, data, fmt.Errorf("get integrator fee ratio: %w", err)
	}
	var integratorAddress []byte
	integratorAddress, data, err = decode.Next(data, common.AddressLength)
	if err != nil {
		return integratorFee, customReceiver, data, fmt.Errorf("get integrator fee address: %w", err)
	}

	integratorFee = IntegratorFee{
		Ratio:    integratorFeeRatio,
		Receiver: common.BytesToAddress(integratorAddress),
	}

	if hasCustomReceiver(flags) {
		var customReceiverBytes []byte
		customReceiverBytes, data, err = decode.Next(data, common.AddressLength)
		if err != nil {
			return integratorFee, customReceiver, data, fmt.Errorf("get custom receiver: %w", err)
		}
		customReceiver = common.BytesToAddress(customReceiverBytes)
	}

	return integratorFee, customReceiver, data, nil
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
