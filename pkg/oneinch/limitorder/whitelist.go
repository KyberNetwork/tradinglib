package limitorder

import (
	"slices"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/util"
	"github.com/ethereum/go-ethereum/common"
)

type Whitelist struct {
	Addresses []util.AddressHalf
}

func (wl Whitelist) IsWhitelisted(taker common.Address) bool {
	addressHalf := util.HalfAddressFromAddress(taker)
	return slices.Contains(wl.Addresses, addressHalf)
}
