package fusionextention_test

import (
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/fusionorder/fusionextention"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestWhitelist_CanExecuteAt(t *testing.T) {
	first := common.HexToAddress("0x00000000000000000000000000000000000000a1")
	second := common.HexToAddress("0x00000000000000000000000000000000000000b2")
	outsider := common.HexToAddress("0x00000000000000000000000000000000000000c3")

	// first may fill from 1000, second from 1036.
	wl := fusionextention.Whitelist{
		ResolvingStartTime: 1000,
		Whitelist: []fusionextention.WhitelistItem{
			{AddressHalf: util.HalfAddressFromAddress(first), Delay: 0},
			{AddressHalf: util.HalfAddressFromAddress(second), Delay: 36},
		},
	}

	assert.True(t, wl.CanExecuteAt(first, 1000))
	assert.False(t, wl.CanExecuteAt(second, 1035))
	assert.True(t, wl.CanExecuteAt(second, 1036))

	// An access-NFT holder outside the whitelist may fill only after the last slot opens.
	assert.False(t, wl.CanExecuteAt(outsider, 1035))
	assert.True(t, wl.CanExecuteAt(outsider, 1036))
}
