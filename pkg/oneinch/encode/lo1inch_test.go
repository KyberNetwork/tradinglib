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
	data, err := encode.PackLO1inch(valueobject.ChainID(chainID), encodingSwap)
	require.NoError(t, err)

	for _, b := range data {
		t.Log(hexutil.Encode(b))
	}
	// Test result: https://www.tdly.co/shared/simulation/ce0ee2f6-3d5e-4015-894a-700793bd8e84
}
