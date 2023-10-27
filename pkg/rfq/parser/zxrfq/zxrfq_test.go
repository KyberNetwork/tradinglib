package zxrfq

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
	require.Equal(t, p.abi.Events[RfqOrderFilledEvent].ID, common.HexToHash("0x829fa99d94dc4636925b38632e625736a614c154d55006b7ab6bea979c210c32"))
	client, err := ethclient.Dial("https://eth.llamarpc.com")
	require.NoError(t, err)
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(16376824),
		ToBlock:   big.NewInt(16376824),
		Addresses: nil,
		Topics: [][]common.Hash{
			{
				common.HexToHash("0x829fa99d94dc4636925b38632e625736a614c154d55006b7ab6bea979c210c32"),
			},
		},
	})
	require.NoError(t, err)
	d, _ := json.Marshal(logs)
	t.Log(string(d))
}

func TestParseEvent(t *testing.T) {
	eventRaw := `{
		"address":"0xdef1c0ded9bec7f1a1670819833240f027b25eff",
		"topics":["0x829fa99d94dc4636925b38632e625736a614c154d55006b7ab6bea979c210c32"],
		"data":"0x0bcf3d2e6e7c55d641843ce75fadb99eca10ed21d56d34c6bdfcf2492e2f12e300000000000000000000000056178a0d5f301baf6cf3e1cd53d9863437345bf900000000000000000000000074de5d4fcbf63e00296fd95d33236b9794016631000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc200000000000000000000000000000000000000000000000039c8e65eb874cc00000000000000000000000000000000000000000000000000000000014903f8810000000000000000000000000000000000000000000000000000000000000000",
		"blockNumber":"0xf9e3f8",
		"transactionHash":"0x09a4736a17994a57b91356c2dfce87af8b46e78279a0ef2d70eeacd0729e11a7",
		"transactionIndex":"0x4",
		"blockHash":"0x06b9e6a0fc43e123cbf624aef470b77afe260e67613a5ff0323d82a39acdc88e",
		"logIndex":"0x16",
		"removed":false}`
	event := types.Log{}
	err := json.Unmarshal([]byte(eventRaw), &event)
	require.NoError(t, err)
	p := MustNewParser()
	log, err := p.Parse(event, uint64(time.Now().Unix()))
	require.NoError(t, err)
	require.Equal(t, log.EventHash, p.eventHash)
	t.Log(log)
}
