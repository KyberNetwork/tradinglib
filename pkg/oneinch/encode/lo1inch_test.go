package encode_test

import (
	_ "embed"
	"encoding/json"
	"log"
	"math/big"
	"testing"

	ksent "github.com/KyberNetwork/kyberswap-dex-lib/pkg/entity"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/valueobject"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/encode"
	"github.com/KyberNetwork/tradinglib/pkg/poolsimulators"
	"github.com/KyberNetwork/tradinglib/pkg/testutil"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"
)

//go:embed lo1inch_test_pool_data.json
var testLO1inchPoolData string

func TestPackLO1inch(t *testing.T) {
	var (
		poolEnt ksent.Pool
		chainID = 1
	)
	require.NoError(t, json.Unmarshal([]byte(testLO1inchPoolData), &poolEnt))

	pSim, err := poolsimulators.PoolSimulatorFromPool(poolEnt, uint(chainID))
	require.NoError(t, err)

	out, err := pSim.CalcAmountOut(
		pool.CalcAmountOutParams{
			TokenAmountIn: pool.TokenAmount{
				Token:  "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
				Amount: testutil.NewBig10("271133267321"),
			},
			TokenOut: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
		},
	)
	require.NoError(t, err)

	log.Printf("%+v", out.SwapInfo)

	encodingSwap := encode.EncodingSwap{
		Pool:              poolEnt.Address,
		TokenIn:           "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
		TokenOut:          "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
		SwapAmount:        testutil.NewBig10("271133267321"),
		AmountOut:         out.TokenAmountOut.Amount,
		LimitReturnAmount: big.NewInt(0),
		Exchange:          valueobject.Exchange(poolEnt.Exchange),
		PoolLength:        len(poolEnt.Tokens),
		PoolType:          poolEnt.Type,
		PoolExtra:         poolEnt.Extra,
		Extra:             out.SwapInfo,
		Recipient:         "0x807cf9a772d5a3f9cefbc1192e939d62f0d9bd38 ",
	}
	data, remain, err := encode.PackLO1inch(valueobject.ChainID(chainID), encodingSwap)
	require.NoError(t, err)
	t.Log("remain", remain.String())

	for _, b := range data {
		t.Log(hexutil.Encode(b))
	}
	// Test result: https://www.tdly.co/shared/simulation/ce0ee2f6-3d5e-4015-894a-700793bd8e84
}

func TestUnpackLO1inch(t *testing.T) {
	encoded := "0xf497df7544dd7f63a495d4c57c7b5445eea1735f63a2298f36ce4f48477d58e522e8a3ca000000000000000000000000e22259232b3cf5c74104cf2ded7f878f0201b1980000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48000000000000000000000000000000000000000000000028a61ed32d6ae800000000000000000000000000000000000000000000000000000000020bc1d72c0044000000000000000000000000000000000067c7e7530000000000000000000068963745db680beeeede005d48280a7a767b21196999dd17e4d29de8da474715246588864a0628dc0810d84c22a31035a4f79e62ab6163ec0c7eb84a954669c70000000000000000000000000000000000000000000000000000003f20cd5579080000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001a000000000000000000000000000000000000000000000000000000000000000140807cf9a772d5a3f9cefbc1192e939d62f0d9bd3000000000000000000000000"

	unpacked, err := encode.UnpackLO1inch(hexutil.MustDecode(encoded))
	require.NoError(t, err)

	t.Logf("%+v", unpacked)
}
