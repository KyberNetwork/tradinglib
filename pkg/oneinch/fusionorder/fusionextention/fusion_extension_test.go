package fusionextention_test

import (
	"math/big"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/fusionorder/auctiondetail"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/fusionorder/fusionextention"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/limitorder"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewFusionExtension(t *testing.T) {
	extension, err := limitorder.DecodeExtension("0x000002eb000002160000021600000122000001220000009100000000000000002ad5004c60e16e54d5007c80ce329adde5b51ef50021c000000a81687086770000b4019b520200b03600900021c0002400000000006409d61b892b2ad624901185b09498030ae3416b66dc74db31d09524fa87b1f700000000000000000000d18bd45f0b94f54a968f2fb3a8220b328962674b339fb574bdc56763f99595770895ad27ad6b0d95972ad4499f120902631a2ad5004c60e16e54d5007c80ce329adde5b51ef50021c000000a81687086770000b4019b520200b03600900021c0002400000000006409d61b892b2ad624901185b09498030ae3416b66dc74db31d09524fa87b1f700000000000000000000d18bd45f0b94f54a968f2fb3a8220b328962674b339fb574bdc56763f99595770895ad27ad6b0d95972ad4499f120902631aa0b86991c6218b36c1d19d4a2e9eb0ce3606eb480000000000000000000000009a143ac36894a6d5d10245819df72c8dfef5c29b000000000000000000000000111111125421ca6dc452d289314280a0f8842a65ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000000000000000000000000000000000000000000000000006871d7bf000000000000000000000000000000000000000000000000000000000000001c0a37d13a9907785dcc324c4b41c3e944591edff7bb8cb9857ed9fbd1d66fda5d0ab4a0012fa9aa43a929ece5ce33104b7f898283455fe5463a8fe4771ba4351b2ad5004c60e16e54d5007c80ce329adde5b51ef500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000646870865309d61b892b2ad6249011850000b09498030ae3416b66dc002474db31d09524fa87b1f70000000000000000000000000000d18bd45f0b94f54a968f00002fb3a8220b328962674b0000339fb574bdc56763f995000095770895ad27ad6b0d950000972ad4499f120902631a00000000000000000000000000000000000000000000000000f73644bc1869ba1f2900")
	require.NoError(t, err)

	fusionExtension, err := fusionextention.NewFusionExtensionFromExtension(extension)
	require.NoError(t, err)

	opt := cmp.Comparer(func(x, y *big.Int) bool {
		if x == nil && y == nil {
			return true
		}
		if x == nil || y == nil {
			return false
		}
		return x.Cmp(y) == 0
	})

	assert.True(t, cmp.Equal(
		fusionextention.FusionExtension{
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
			SurplusParam: fusionextention.SurplusParam{
				EstimatedTakerAmount: new(big.Int).SetBytes(hexutil.MustDecode("0xf73644bc1869ba1f29")),
				ProtocolFee:          0,
			},
		}, fusionExtension, opt))
}

func TestDecodeFusionExtension(t *testing.T) {
	t.Skip("debug only")
	extension, err := limitorder.DecodeExtension("0x000002570000014a0000014a0000014a0000014a000000a500000000000000002ad5004c60e16e54d5007c80ce329adde5b51ef50000000000000068b0f91d00383d0003e8000000000000640cb09498030ae3416b66dc74db31d09524fa87b1f7972ad4499f120902631ad18bd45f0b94f54a968fc90ed87a54c23dc480b32fb3a8220b328962674bd61b892b2ad624901185b8394f2220fac7e6ade6200b94fa1a7585930e2195770895ad27ad6b0d956ea9a11ae13b29f5c555339fb574bdc56763f9952ad5004c60e16e54d5007c80ce329adde5b51ef50000000000000068b0f91d00383d0003e8000000000000640cb09498030ae3416b66dc74db31d09524fa87b1f7972ad4499f120902631ad18bd45f0b94f54a968fc90ed87a54c23dc480b32fb3a8220b328962674bd61b892b2ad624901185b8394f2220fac7e6ade6200b94fa1a7585930e2195770895ad27ad6b0d956ea9a11ae13b29f5c555339fb574bdc56763f9952ad5004c60e16e54d5007c80ce329adde5b51ef501000000000000000000000000000000000000000000000000000000000000000000000000000000005866984ca08e121f3d97138ce69909d84fc8261a00000000006468b0f9110cb09498030ae3416b66dc000074db31d09524fa87b1f70000972ad4499f120902631a0000d18bd45f0b94f54a968f0000c90ed87a54c23dc480b300002fb3a8220b328962674b0000d61b892b2ad6249011850000b8394f2220fac7e6ade60000200b94fa1a7585930e21000095770895ad27ad6b0d9500006ea9a11ae13b29f5c5550000339fb574bdc56763f9950000000000000000000000000000000000000000000000000000dc9e8459f70ee1dc00")
	require.NoError(t, err)

	fusionExtension, err := fusionextention.NewFusionExtensionFromExtension(extension)
	require.NoError(t, err)

	t.Log(fusionExtension)
}
