package paraswap

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

func TestFetchEvent(t *testing.T) {
	t.Skip()
	p := MustNewParser()
	require.Equal(t, p.abi.Events[OrderFilledEvent].ID, common.HexToHash("0x6621486d9c28838df4a87d2cca5007bc2aaf6a5b5de083b1db8faf709302c473"))
	client, err := ethclient.Dial("https://eth.llamarpc.com")
	require.NoError(t, err)
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(16872690),
		ToBlock:   big.NewInt(16872690),
		Addresses: nil,
		Topics: [][]common.Hash{
			{
				common.HexToHash("0x6621486d9c28838df4a87d2cca5007bc2aaf6a5b5de083b1db8faf709302c473"),
			},
		},
	})
	require.NoError(t, err)
	d, _ := json.Marshal(logs)
	t.Log(string(d))
}

func TestParseEvent(t *testing.T) {
	eventRaw := `{
		"address":"0xe92b586627cca7a83dc919cc7127196d70f55a06",
		"topics":["0x6621486d9c28838df4a87d2cca5007bc2aaf6a5b5de083b1db8faf709302c473","0xd3c7eed1fe6c2c0684383c9f4a4f87d7136a96e60e9a5db381d324e9198133f0","0x000000000000000000000000bb289bc97591f70d8216462df40ed713011b968a","0x000000000000000000000000def171fe48cf0115b1d80b88dc8eab59176fee57"],
		"data":"0x000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000000000000000000000000000006f285be082e6dc410000000000000000000000004d224452801aced8b2f0aebe155379bb5d5943810000000000000000000000000000000000000000000000bad91fdd52aac724d2",
		"blockNumber":"0x10174f2",
		"transactionHash":"0xbed8812c5b929036c8b6c20c834482afcee3bc5a7100e2926a90604413c35e5f",
		"transactionIndex":"0x57",
		"blockHash":"0x66f5cae466d1a24aab266717d06df7234b9b826b2fe2891bdfa343ac42c8605f",
		"logIndex":"0xe7",
		"removed":false
		}`
	event := types.Log{}
	err := json.Unmarshal([]byte(eventRaw), &event)
	require.NoError(t, err)
	p := MustNewParser()
	log, err := p.Parse(event, uint64(time.Now().Unix()))
	require.NoError(t, err)
	require.Equal(t, log.EventHash, p.eventHash)
	t.Log(log)
}
