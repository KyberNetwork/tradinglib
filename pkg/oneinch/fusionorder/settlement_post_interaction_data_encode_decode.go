package fusionorder

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/utils"
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

func DecodeSettlementPostInteractionData(extraData string) (SettlementPostInteractionData, error) {
	if !utils.IsHexString(extraData) {
		return SettlementPostInteractionData{}, errors.New("invalid auction details data")
	}

	data, err := hex.DecodeString(utils.Trim0x(extraData))
	if err != nil {
		return SettlementPostInteractionData{}, fmt.Errorf("decode settlement post interaction data: %w", err)
	}

	flags := data[len(data)-1]

	bankFee := big.NewInt(0)
	var integratorFee IntegratorFee
	var customReceiver common.Address
	resolvingStartTime := big.NewInt(0)

	if resolverFeeEnabled(flags) {
		bankFee.SetBytes(data[:4])
		data = data[4:]
	}

	if integratorFeeEnabled(flags) {
		integratorFeeRatio := new(big.Int).SetBytes(data[:2]).Int64()
		integratorAddress := common.BytesToAddress(data[2:22])
		integratorFee = IntegratorFee{
			Ratio:    integratorFeeRatio,
			Receiver: integratorAddress,
		}

		data = data[22:]

		if hasCustomReceiver(flags) {
			customReceiver = common.BytesToAddress(data[:20])
			data = data[20:]
		}
	}

	resolvingStartTime = new(big.Int).SetBytes(data[:4])
	data = data[4:]

	whitelistCount := resolversCount(flags)
	whitelist := make([]WhitelistItem, 0, whitelistCount)

	for i := byte(0); i < whitelistCount; i++ {
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

func (s SettlementPostInteractionData) Encode() string {
	buf := new(bytes.Buffer)
	var flags byte
	if s.BankFee != 0 {
		flags |= resolverFeeFlag
		buf.Write(utils.PadOrTrim(big.NewInt(s.BankFee).Bytes(), 4))
	}
	if s.IntegratorFee.Ratio != 0 {
		flags |= integratorFeeFlag
		buf.Write(utils.PadOrTrim(big.NewInt(s.IntegratorFee.Ratio).Bytes(), 2))
		buf.Write(s.IntegratorFee.Receiver.Bytes())
		if s.CustomReceiver != (common.Address{}) {
			flags |= customReceiverFlag
			buf.Write(s.CustomReceiver.Bytes())
		}
	}
	buf.Write(utils.PadOrTrim(big.NewInt(s.ResolvingStartTime).Bytes(), 4))

	for _, wl := range s.Whitelist {
		buf.Write(wl.AddressHalf[:])
		buf.Write(utils.PadOrTrim(big.NewInt(int64(int(wl.Delay))).Bytes(), 2))
	}

	flags |= byte(len(s.Whitelist)) << whitelistShift

	buf.WriteByte(flags)

	return utils.Add0x(hex.EncodeToString(buf.Bytes()))
}
