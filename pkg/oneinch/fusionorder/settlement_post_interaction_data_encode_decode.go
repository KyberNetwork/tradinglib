package fusionorder

import (
	"bytes"
	"fmt"
	"math/big"

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

func DecodeSettlementPostInteractionData(data []byte) (SettlementPostInteractionData, error) {
	// must have at least 1 byte for flags
	// nolint: gomnd
	if err := decode.ValidateDataLength(data, 1); err != nil {
		return SettlementPostInteractionData{}, err
	}

	flags := data[len(data)-1]

	bankFee := big.NewInt(0)
	var integratorFee IntegratorFee
	var customReceiver common.Address

	if resolverFeeEnabled(flags) {
		const lengthBankFee = 4
		if err := decode.ValidateDataLength(data, lengthBankFee); err != nil {
			return SettlementPostInteractionData{}, fmt.Errorf("get bank fee: %w", err)
		}
		bankFee.SetBytes(data[:lengthBankFee])
		data = data[lengthBankFee:]
	}

	if integratorFeeEnabled(flags) {
		const lengthIntegratorFee = 2 + common.AddressLength
		if err := decode.ValidateDataLength(data, lengthIntegratorFee); err != nil {
			return SettlementPostInteractionData{}, fmt.Errorf("get integrator fee: %w", err)
		}
		integratorFeeRatio := new(big.Int).SetBytes(data[:2]).Int64()
		integratorAddress := common.BytesToAddress(data[2:lengthIntegratorFee])
		integratorFee = IntegratorFee{
			Ratio:    integratorFeeRatio,
			Receiver: integratorAddress,
		}

		data = data[lengthIntegratorFee:]

		if hasCustomReceiver(flags) {
			if err := decode.ValidateDataLength(data, common.AddressLength); err != nil {
				return SettlementPostInteractionData{}, fmt.Errorf("get custom receiver: %w", err)
			}
			customReceiver = common.BytesToAddress(data[:common.AddressLength])
			data = data[common.AddressLength:]
		}
	}

	resolvingStartTime := new(big.Int).SetBytes(data[:4])
	data = data[4:]

	whitelistCount := resolversCount(flags)
	whitelist := make([]WhitelistItem, 0, whitelistCount)

	for i := byte(0); i < whitelistCount; i++ {
		const lengthWhitelistItem = addressHalfLength + 2
		if err := decode.ValidateDataLength(data, lengthWhitelistItem); err != nil {
			return SettlementPostInteractionData{}, fmt.Errorf("get whitelist item: %w", err)
		}

		var address AddressHalf
		copy(address[:], data[:addressHalfLength])
		data = data[addressHalfLength:]

		delay := new(big.Int).SetBytes(data[:2])
		data = data[2:]

		whitelist = append(whitelist, WhitelistItem{
			AddressHalf: address,
			Delay:       delay.Int64(),
		})
	}

	return SettlementPostInteractionData{
		Whitelist:          whitelist,
		IntegratorFee:      integratorFee,
		BankFee:            bankFee.Int64(),
		ResolvingStartTime: resolvingStartTime.Int64(),
		CustomReceiver:     customReceiver,
	}, nil
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
