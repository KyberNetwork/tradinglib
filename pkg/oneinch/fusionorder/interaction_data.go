package fusionorder

import (
	"fmt"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/decode"
)

type InteractionData struct {
	IntegratorFee     uint16 // In bps
	IntegratorShare   uint16 // In bps
	ResolverFee       uint16 // In bps
	WhitelistDiscount uint16 // In bps
}

func ParseAmountData(iter *decode.BytesIterator) (InteractionData, error) {
	integratorFee, err := iter.NextUint16()
	if err != nil {
		return InteractionData{}, fmt.Errorf("get intergator fee: %w", err)
	}

	integratorShare, err := iter.NextUint8()
	if err != nil {
		return InteractionData{}, fmt.Errorf("get intergator share: %w", err)
	}

	resolverFee, err := iter.NextUint16()
	if err != nil {
		return InteractionData{}, fmt.Errorf("get resolver fee: %w", err)
	}

	whitelistDiscount, err := iter.NextUint8()
	if err != nil {
		return InteractionData{}, fmt.Errorf("get whitelist discount: %w", err)
	}

	return InteractionData{
		IntegratorFee:     integratorFee / 10,
		IntegratorShare:   uint16(integratorShare) * 100,
		ResolverFee:       resolverFee / 10,
		WhitelistDiscount: (100 - uint16(whitelistDiscount)) * 100, // contract uses 1 - discount.
	}, nil
}
