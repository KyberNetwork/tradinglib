package fusionextention

import (
	"fmt"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/bps"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/constants"
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

	whitelistDiscountSub, err := iter.NextUint8()
	if err != nil {
		return InteractionData{}, fmt.Errorf("get whitelist discount: %w", err)
	}

	whitelistDiscount := int(constants.FeeBase1e2.Int64()) - int(whitelistDiscountSub)
	return InteractionData{
		IntegratorFee:     bps.FromFraction(int(integratorFee), constants.FeeBase1e5),
		IntegratorShare:   bps.FromFraction(int(integratorShare), constants.FeeBase1e2),
		ResolverFee:       bps.FromFraction(int(resolverFee), constants.FeeBase1e5),
		WhitelistDiscount: bps.FromFraction(whitelistDiscount, constants.FeeBase1e2),
	}, nil
}
