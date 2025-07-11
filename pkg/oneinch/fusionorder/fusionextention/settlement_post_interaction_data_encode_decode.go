package fusionextention

import (
	"fmt"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/decode"
	"github.com/ethereum/go-ethereum/common"
)

const customReceiverFlag = 0x01

func hasCustomReceiver(flags byte) bool {
	return flags&customReceiverFlag == customReceiverFlag
}

// DecodeSettlementPostInteractionData decodes SettlementPostInteractionData from bytes
// nolint: gomnd
// https://github.com/1inch/fusion-sdk/blob/6d40f680a2f1cd0148c314d4c8608a004fffdc09/src/fusion-order/fusion-extension.ts#L73-L89
func DecodeSettlementPostInteractionData(data []byte) (SettlementPostInteractionData, error) {
	iter := decode.NewBytesIterator(data)
	if _, err := iter.NextUint160(); err != nil {
		return SettlementPostInteractionData{}, fmt.Errorf("skip address of extension: %w", err)
	}
	flags, err := iter.NextUint8()
	if err != nil {
		return SettlementPostInteractionData{}, fmt.Errorf(
			"get settlement post interaction data flags: %w", err)
	}

	integratorFeeRecipientBytes, err := iter.NextBytes(common.AddressLength)
	if err != nil {
		return SettlementPostInteractionData{}, fmt.Errorf("get integrator fee recipient: %w", err)
	}
	integratorFeeRecipient := common.BytesToAddress(integratorFeeRecipientBytes)

	protocolFeeRecipientBytes, err := iter.NextBytes(common.AddressLength)
	if err != nil {
		return SettlementPostInteractionData{}, fmt.Errorf("get protocol fee recipient: %w", err)
	}
	protocolFeeRecipient := common.BytesToAddress(protocolFeeRecipientBytes)

	var customReceiver common.Address
	if hasCustomReceiver(flags) {
		customReceiverBytes, err := iter.NextBytes(common.AddressLength)
		if err != nil {
			return SettlementPostInteractionData{}, fmt.Errorf("get custom receiver: %w", err)
		}
		customReceiver = common.BytesToAddress(customReceiverBytes)
	}

	interactionData, err := ParseAmountData(iter)
	if err != nil {
		return SettlementPostInteractionData{}, fmt.Errorf("parse amount data: %w", err)
	}

	whitelist, err := DecodeWhitelist(iter)
	if err != nil {
		return SettlementPostInteractionData{}, fmt.Errorf("decode whitelist: %w", err)
	}

	//surplusParam, err := DecodeSurplusParam(iter)
	//if err != nil {
	//	return SettlementPostInteractionData{}, fmt.Errorf("decode surplus param: %w", err)
	//}
	return SettlementPostInteractionData{
		IntegratorFeeRecipient: integratorFeeRecipient,
		ProtocolFeeRecipient:   protocolFeeRecipient,
		CustomReceiver:         customReceiver,
		InteractionData:        interactionData,
		Whitelist:              whitelist,
		//SurplusParam:           surplusParam,
	}, nil
}
