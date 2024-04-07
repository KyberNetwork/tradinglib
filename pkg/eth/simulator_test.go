package eth_test

import (
	"context"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/convert"
	"github.com/KyberNetwork/tradinglib/pkg/eth"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/require"
)

func TestEstimateGasWithOverrides(t *testing.T) {
	t.Skip()

	var (
		url         = "https://ethereum.kyberengineering.io"
		wallet      = common.HexToAddress("0x72CE0F4a9dbB974D3B1a9bF7ca857fD381260e97")
		ethValue, _ = convert.FloatToWei(10, 18)
	)

	c, err := rpc.Dial(url)
	require.NoError(t, err)

	s := eth.NewSimulator(c)

	gasUnit, err := s.EstimateGasWithOverrides(context.Background(), ethereum.CallMsg{
		From:  wallet,
		To:    &wallet,
		Value: ethValue,
	}, nil, &map[common.Address]gethclient.OverrideAccount{
		wallet: {
			Code:    []byte("0x6"),
			Balance: ethValue,
		},
	},
	)
	require.NoError(t, err)
	t.Log(gasUnit)
}
