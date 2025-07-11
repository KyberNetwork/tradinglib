package auctioncalculator_test

import (
	"math/big"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/auctioncalculator"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/fusionorder/fusionextention"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/limitorder"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAmountCalculator(t *testing.T) {
	extension, err := limitorder.DecodeExtension("0x00000292000001ae000001ae000001ae000001ae000000d70000000000000000abd4e5fb590aa132749bbf2a04ea57efbaac399e0000d3000001c867efac06000168020ea50801cde80030018ae800300146b10030010638003000c5f2003000c44b001800c3a1003c0000d300240000000000640db09498030ae3416b66dc5dcd8578ca14eec99e63972ad4499f120902631ad18bd45f0b94f54a968fd61b892b2ad62490118595770895ad27ad6b0d95339fb574bdc56763f995617556ed277ab32233786de5e0e428ac771d77b55b7d1434eae4a48b2c8626813bd1b091ea6bedbd00000000000000000000b8394f2220fac7e6ade6abd4e5fb590aa132749bbf2a04ea57efbaac399e0000d3000001c867efac06000168020ea50801cde80030018ae800300146b10030010638003000c5f2003000c44b001800c3a1003c0000d300240000000000640db09498030ae3416b66dc5dcd8578ca14eec99e63972ad4499f120902631ad18bd45f0b94f54a968fd61b892b2ad62490118595770895ad27ad6b0d95339fb574bdc56763f995617556ed277ab32233786de5e0e428ac771d77b55b7d1434eae4a48b2c8626813bd1b091ea6bedbd00000000000000000000b8394f2220fac7e6ade6abd4e5fb590aa132749bbf2a04ea57efbaac399e000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006467efabee0db09498030ae3416b66dc00005dcd8578ca14eec99e630000972ad4499f120902631a0000d18bd45f0b94f54a968f0000d61b892b2ad624901185000095770895ad27ad6b0d950000339fb574bdc56763f9950000617556ed277ab322337800006de5e0e428ac771d77b500005b7d1434eae4a48b2c86000026813bd1b091ea6bedbd0000000000000000000000000000b8394f2220fac7e6ade60000")
	require.NoError(t, err)
	fusionExtension, err := fusionextention.NewFusionExtensionFromExtension(extension)
	require.NoError(t, err)

	takingAmount, _ := new(big.Int).SetString("140444897314051230680", 10)
	amountCalculator := auctioncalculator.NewAmountCalculatorFromExtension(fusionExtension)
	requiredTakingAmount := amountCalculator.GetRequiredTakingAmount(
		common.HexToAddress("0xad3b67bca8935cb510c8d18bd45f0b94f54a968f"),
		takingAmount,
		big.NewInt(1743760379),
		big.NewInt(411998030),
	)
	assert.Equal(t, "142335721011080033804", requiredTakingAmount.String())
}
