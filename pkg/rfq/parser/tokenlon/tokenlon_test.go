package tokenlon

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
	require.Equal(t, p.abi.Events[FillOrderEvent].ID, common.HexToHash("0x75d58426b26ab641a6a6a46f12fe35e17c570a1cd264c7248a73d90e3a8682ff"))
	client, err := ethclient.Dial("https://eth.llamarpc.com")
	require.NoError(t, err)
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(16381232),
		ToBlock:   big.NewInt(16381232),
		Addresses: nil,
		Topics: [][]common.Hash{
			{
				common.HexToHash("0x75d58426b26ab641a6a6a46f12fe35e17c570a1cd264c7248a73d90e3a8682ff"),
			},
		},
	})
	require.NoError(t, err)
	d, _ := json.Marshal(logs)
	t.Log(string(d))
}

func TestParseEvent(t *testing.T) {
	eventRaw := `{
		"address":"0xfd6c2d2499b1331101726a8ac68ccc9da3fab54f",
		"topics":["0x75d58426b26ab641a6a6a46f12fe35e17c570a1cd264c7248a73d90e3a8682ff","0x50cc97dc89796148c9975215413df35761ef27f9c59a222ec5eccde85bf35dd3","0xd2b210488c13475e08babd177b7eeaa5d5431edda7835fe9422f2085499a8454","0x000000000000000000000000ac06105579719ac8c22935be712ab17a3d51b415"],
		"data":"0x0000000000000000000000000000000000000000000000000000000000000120000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000000000000000000000000000016f5854c184970000000000000000000000000000b3c839dbde6b96d37c56ee4f9dad3390d49310aa000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7000000000000000000000000000000000000000000000000000000083081cc6b000000000000000000000000ac06105579719ac8c22935be712ab17a3d51b415000000000000000000000000000000000000000000000000000000082be50bb6000000000000000000000000000000000000000000000000000000000000001600000000000000000000000000000000000000000000000000000000000000065246512076310000000000000000000000000000000000000000000000000000",
		"blockNumber":"0xf9f530",
		"transactionHash":"0xf4423faa840a957508b0ddd031f001c69ad6453709711912e3d34b536acbaa89",
		"transactionIndex":"0x5d",
		"blockHash":"0xf09bfddb6c3bfa577c2ba95133fb0460faa474bbd5242bb816e181f7d662d79c",
		"logIndex":"0x118",
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
