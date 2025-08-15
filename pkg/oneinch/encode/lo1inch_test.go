package encode_test

import (
	_ "embed"
	"encoding/json"
	"log"
	"math/big"
	"testing"

	ksent "github.com/KyberNetwork/kyberswap-dex-lib/pkg/entity"
	_ "github.com/KyberNetwork/kyberswap-dex-lib/pkg/msgpack" // make sure that every init registerFactory function is ran
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pool"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/valueobject"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/encode"
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

	factoryFn := pool.Factory(poolEnt.Type)
	require.NotNil(t, factoryFn)

	pSim, err := factoryFn(pool.FactoryParams{
		EntityPool:  poolEnt,
		ChainID:     valueobject.ChainID(chainID),
		BasePoolMap: nil,
		EthClient:   nil,
	})
	require.NoError(t, err)

	out, err := pSim.CalcAmountOut(
		pool.CalcAmountOutParams{
			TokenAmountIn: pool.TokenAmount{
				Token:  "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
				Amount: testutil.NewBig10("101133267321"),
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
		PoolExtra:         pSim.GetMetaInfo("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
		Extra:             out.SwapInfo,
		Recipient:         "0x807cf9a772d5a3f9cefbc1192e939d62f0d9bd38 ",
	}
	data, remain, err := encode.PackLO1inch(valueobject.ChainID(chainID), encodingSwap)
	require.NoError(t, err)
	t.Log("remain", remain.String())

	for _, b := range data {
		t.Log(hexutil.Encode(b))

		unpacked, err := encode.UnpackLO1inch(hexutil.MustDecode(hexutil.Encode(b)))
		require.NoError(t, err)
		t.Logf("%+v", unpacked)
	}
	// Test result: https://www.tdly.co/shared/simulation/ce0ee2f6-3d5e-4015-894a-700793bd8e84
}

func TestUnpackLO1inch(t *testing.T) {
	encoded := "0xf497df75238aed66c26048cb6f769c930000000000000000000000000000000000000000000000000000000000000000e67a89fe03e493cdfbdcedcd925235349d76d0010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb4800000000000000000000000000000000000000000000000009935f581f0505000000000000000000000000000000000000000000000000000000000321fb01c04400000000000000000000000000000001006d76c7a800000000000000000000911885c6a0e6ea35e900d03b0d1f276ce244ce25c3e847360bfc1ef311c22c182ee8aab950585031672398dfbab9101ea65142b5c430144dc52da32705c016430000000000000000000000000000000000000000000000000000000262201b5f0800000000003400000000000000000000000000000000000748f05c8c65b48500000000000000000000000000000000000000000000000000000000000001a000000000000000000000000000000000000000000000000000000000000000480807cf9a772d5a3f9cefbc1192e939d62f0d9bd300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000748f05c8c65b485000000000000000000000000000000000000000000000000"

	unpacked, err := encode.UnpackLO1inch(hexutil.MustDecode(encoded))
	require.NoError(t, err)

	t.Logf("%+v", unpacked)
}
