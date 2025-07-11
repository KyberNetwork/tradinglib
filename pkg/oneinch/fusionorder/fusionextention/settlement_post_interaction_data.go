package fusionextention

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrEmptyWhitelist         = errors.New("white list cannot be empty")
	ErrResolvingStartTimeZero = errors.New("resolving start time can not be 0")
)

type SettlementPostInteractionData struct {
	IntegratorFeeRecipient common.Address
	ProtocolFeeRecipient   common.Address
	CustomReceiver         common.Address
	InteractionData        InteractionData
	Whitelist              Whitelist
	SurplusParam           SurplusParam
}

func (s SettlementPostInteractionData) HasFees() bool {
	return s.IntegratorFeeRecipient != (common.Address{}) || s.ProtocolFeeRecipient != (common.Address{})
}
