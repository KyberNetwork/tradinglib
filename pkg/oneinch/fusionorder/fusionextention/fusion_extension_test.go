package fusionextention_test

import (
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/fusionorder/auctiondetail"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/fusionorder/fusionextention"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/limitorder"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewFusionExtension(t *testing.T) {
	extension, err := limitorder.DecodeExtension("0x000002eb000002160000021600000122000001220000009100000000000000002ad5004c60e16e54d5007c80ce329adde5b51ef50021c000000a81687086770000b4019b520200b03600900021c0002400000000006409d61b892b2ad624901185b09498030ae3416b66dc74db31d09524fa87b1f700000000000000000000d18bd45f0b94f54a968f2fb3a8220b328962674b339fb574bdc56763f99595770895ad27ad6b0d95972ad4499f120902631a2ad5004c60e16e54d5007c80ce329adde5b51ef50021c000000a81687086770000b4019b520200b03600900021c0002400000000006409d61b892b2ad624901185b09498030ae3416b66dc74db31d09524fa87b1f700000000000000000000d18bd45f0b94f54a968f2fb3a8220b328962674b339fb574bdc56763f99595770895ad27ad6b0d95972ad4499f120902631aa0b86991c6218b36c1d19d4a2e9eb0ce3606eb480000000000000000000000009a143ac36894a6d5d10245819df72c8dfef5c29b000000000000000000000000111111125421ca6dc452d289314280a0f8842a65ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000000000000000000000000000000000000000000000000006871d7bf000000000000000000000000000000000000000000000000000000000000001c0a37d13a9907785dcc324c4b41c3e944591edff7bb8cb9857ed9fbd1d66fda5d0ab4a0012fa9aa43a929ece5ce33104b7f898283455fe5463a8fe4771ba4351b2ad5004c60e16e54d5007c80ce329adde5b51ef500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000646870865309d61b892b2ad6249011850000b09498030ae3416b66dc002474db31d09524fa87b1f70000000000000000000000000000d18bd45f0b94f54a968f00002fb3a8220b328962674b0000339fb574bdc56763f995000095770895ad27ad6b0d950000972ad4499f120902631a00000000000000000000000000000000000000000000000000f73644bc1869ba1f2900")
	require.NoError(t, err)

	fusionExtension, err := fusionextention.NewFusionExtensionFromExtension(extension)
	require.NoError(t, err)
	assert.Equal(t, fusionextention.FusionExtension{
		Address: common.HexToAddress("0x2Ad5004c60e16E54d5007C80CE329Adde5B51Ef5"),
		AuctionDetails: auctiondetail.AuctionDetails{
			StartTime:       1752204919,
			Duration:        180,
			InitialRateBump: 105298,
			Points: []auctiondetail.AuctionPoint{
				{
					Delay:       144,
					Coefficient: 45110,
				},
				{
					Delay:       36,
					Coefficient: 8640,
				},
			},
			GasCost: auctiondetail.AuctionGasCostInfo{
				GasBumpEstimate:  8640,
				GasPriceEstimate: 2689,
			},
		},
		Whitelist: fusionextention.Whitelist{
			ResolvingStartTime: 1752204883,
			Whitelist: []fusionextention.WhitelistItem{
				{
					AddressHalf: util.HexToAddressHalf("0xd61b892b2ad624901185"),
					Delay:       0,
				},
				{
					AddressHalf: util.HexToAddressHalf("0xb09498030ae3416b66dc"),
					Delay:       36,
				},
				{
					AddressHalf: util.HexToAddressHalf("0x74db31d09524fa87b1f7"),
					Delay:       0,
				},
				{
					AddressHalf: util.HexToAddressHalf("0x00000000000000000000"),
					Delay:       0,
				},
				{
					AddressHalf: util.HexToAddressHalf("0xd18bd45f0b94f54a968f"),
					Delay:       0,
				},
				{
					AddressHalf: util.HexToAddressHalf("0x2fb3a8220b328962674b"),
					Delay:       0,
				},
				{
					AddressHalf: util.HexToAddressHalf("0x339fb574bdc56763f995"),
					Delay:       0,
				},
				{
					AddressHalf: util.HexToAddressHalf("0x95770895ad27ad6b0d95"),
					Delay:       0,
				},
				{
					AddressHalf: util.HexToAddressHalf("0x972ad4499f120902631a"),
					Delay:       0,
				},
			},
		},
		Extra: fusionextention.Extra{
			MakerPermit: limitorder.Interaction{
				Target: common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
				Data:   hexutil.MustDecode("0x0000000000000000000000009a143ac36894a6d5d10245819df72c8dfef5c29b000000000000000000000000111111125421ca6dc452d289314280a0f8842a65ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000000000000000000000000000000000000000000000000006871d7bf000000000000000000000000000000000000000000000000000000000000001c0a37d13a9907785dcc324c4b41c3e944591edff7bb8cb9857ed9fbd1d66fda5d0ab4a0012fa9aa43a929ece5ce33104b7f898283455fe5463a8fe4771ba4351b"),
			},
		},
		// SurplusParam: SurplusParam{
		//	estimatedTakerAmount: new(big.Int).SetBytes(hexutil.MustDecode("0xf73644bc1869ba1f29")),
		// ,
	}, fusionExtension)
}
