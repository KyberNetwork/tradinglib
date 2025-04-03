package fusionorder

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
func DecodeSettlementPostInteractionData(data []byte) (SettlementPostInteractionData, error) {
	iter := decode.NewBytesIterator(data)
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

	return SettlementPostInteractionData{
		IntegratorFeeRecipient: integratorFeeRecipient,
		ProtocolFeeRecipient:   protocolFeeRecipient,
		CustomReceiver:         customReceiver,
		InteractionData:        interactionData,
		Whitelist:              whitelist,
	}, nil
}
