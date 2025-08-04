package limitorder

import (
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/util"
	"github.com/ethereum/go-ethereum/common"
)

type Whitelist struct {
	Addresses []util.AddressHalf
}

func (wl Whitelist) IsWhitelisted(taker common.Address) bool {
	addressHalf := util.HalfAddressFromAddress(taker)
	for _, item := range wl.Addresses {
		if addressHalf == item {
			return true
		}
	}

	return false
}
