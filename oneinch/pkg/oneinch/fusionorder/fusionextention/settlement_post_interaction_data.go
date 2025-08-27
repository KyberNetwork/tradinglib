package fusionextention

import (
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/limitorder"
	"github.com/ethereum/go-ethereum/common"
)

type SettlementPostInteractionData struct {
	IntegratorFeeRecipient common.Address
	ProtocolFeeRecipient   common.Address
	CustomReceiver         common.Address
	Fee                    limitorder.Fee
	Whitelist              Whitelist
	SurplusParam           SurplusParam
}

func (s SettlementPostInteractionData) HasFees() bool {
	return s.IntegratorFeeRecipient != (common.Address{}) || s.ProtocolFeeRecipient != (common.Address{})
}
