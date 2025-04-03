package fusionorder_test

import (
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/fusionorder"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/limitorder"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewFusionExtension(t *testing.T) {
	extension, err := limitorder.DecodeExtension("0x0000024c00000168000001680000016800000168000000b40000000000000000abd4e5fb590aa132749bbf2a04ea57efbaac399e00ca070000023367ed11940000b401934d0100ca0700b40000000000640db09498030ae3416b66dc5dcd8578ca14eec99e63972ad4499f120902631ad18bd45f0b94f54a968fd61b892b2ad62490118595770895ad27ad6b0d95339fb574bdc56763f995617556ed277ab32233786de5e0e428ac771d77b55b7d1434eae4a48b2c8626813bd1b091ea6bedbd00000000000000000000b8394f2220fac7e6ade6abd4e5fb590aa132749bbf2a04ea57efbaac399e00ca070000023367ed11940000b401934d0100ca0700b40000000000640db09498030ae3416b66dc5dcd8578ca14eec99e63972ad4499f120902631ad18bd45f0b94f54a968fd61b892b2ad62490118595770895ad27ad6b0d95339fb574bdc56763f995617556ed277ab32233786de5e0e428ac771d77b55b7d1434eae4a48b2c8626813bd1b091ea6bedbd00000000000000000000b8394f2220fac7e6ade6abd4e5fb590aa132749bbf2a04ea57efbaac399e000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006467ed117c0db09498030ae3416b66dc00005dcd8578ca14eec99e630000972ad4499f120902631a0000d18bd45f0b94f54a968f0000d61b892b2ad624901185000095770895ad27ad6b0d950000339fb574bdc56763f9950000617556ed277ab322337800006de5e0e428ac771d77b500005b7d1434eae4a48b2c86000026813bd1b091ea6bedbd0000000000000000000000000000b8394f2220fac7e6ade60000")
	require.NoError(t, err)

	fusionExtension, err := fusionorder.NewFusionExtensionFromExtension(extension)
	require.NoError(t, err)
	assert.Equal(t, fusionorder.FusionExtension{
		Address: common.HexToAddress("0xAbD4e5fB590Aa132749bbF2A04eA57EFbaAC399E"),
		AuctionDetails: fusionorder.AuctionDetails{
			StartTime:       1743589780,
			Duration:        180,
			InitialRateBump: 103245,
			Points: []fusionorder.AuctionPoint{
				{
					Delay:       180,
					Coefficient: 51719,
				},
			},
			GasCost: fusionorder.AuctionGasCostInfo{
				GasBumpEstimate:  51719,
				GasPriceEstimate: 563,
			},
		},
		Whitelist: fusionorder.Whitelist{
			ResolvingStartTime: 1743589756,
			Whitelist: []fusionorder.WhitelistItem{
				{
					AddressHalf: fusionorder.HexToAddressHalf("0xb09498030ae3416b66dc"),
					Delay:       0,
				},
				{
					AddressHalf: fusionorder.HexToAddressHalf("0x5dcd8578ca14eec99e63"),
					Delay:       0,
				},
				{
					AddressHalf: fusionorder.HexToAddressHalf("0x972ad4499f120902631a"),
					Delay:       0,
				},
				{
					AddressHalf: fusionorder.HexToAddressHalf("0xd18bd45f0b94f54a968f"),
					Delay:       0,
				},
				{
					AddressHalf: fusionorder.HexToAddressHalf("0xd61b892b2ad624901185"),
					Delay:       0,
				},
				{
					AddressHalf: fusionorder.HexToAddressHalf("0x95770895ad27ad6b0d95"),
					Delay:       0,
				},
				{
					AddressHalf: fusionorder.HexToAddressHalf("0x339fb574bdc56763f995"),
					Delay:       0,
				},
				{
					AddressHalf: fusionorder.HexToAddressHalf("0x617556ed277ab3223378"),
					Delay:       0,
				},
				{
					AddressHalf: fusionorder.HexToAddressHalf("0x6de5e0e428ac771d77b5"),
					Delay:       0,
				},
				{
					AddressHalf: fusionorder.HexToAddressHalf("0x5b7d1434eae4a48b2c86"),
					Delay:       0,
				},
				{
					AddressHalf: fusionorder.HexToAddressHalf("0x26813bd1b091ea6bedbd"),
					Delay:       0,
				},
				{
					AddressHalf: fusionorder.HexToAddressHalf("0x00000000000000000000"),
					Delay:       0,
				},
				{
					AddressHalf: fusionorder.HexToAddressHalf("0xb8394f2220fac7e6ade6"),
					Delay:       0,
				},
			},
		},
	}, fusionExtension)
}
