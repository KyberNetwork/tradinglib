package kyberswap

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
	require.Equal(t, p.abi.Events[SwappedEvent].ID, common.HexToHash("0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8"))
	client, err := ethclient.Dial("https://eth.llamarpc.com")
	require.NoError(t, err)
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(16282225),
		ToBlock:   big.NewInt(16282225),
		Addresses: nil,
		Topics:    [][]common.Hash{{common.HexToHash("0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8")}},
	})
	require.NoError(t, err)
	d, _ := json.Marshal(logs)
	t.Log(string(d))
}

func TestParseEvent(t *testing.T) {
	eventRaw := `{
    "address": "0x617dee16b86534a5d792a4d7a62fb491b544111e",
    "topics": [
        "0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8"
    ],
    "data": "0x000000000000000000000000db306e5c24cd28a02b50c6f893d46a3572835195000000000000000000000000eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee000000000000000000000000952193fcf398550b6caa0ecf5e706a6ad0ab8d8a000000000000000000000000db306e5c24cd28a02b50c6f893d46a357283519500000000000000000000000000000000000000000000000006f05b59d3b2000000000000000000000000000000000000000000000000000000005a91a0920233",
    "blockNumber": "0xf87271",
    "transactionHash": "0x6461e9058f2779725aa97bf5a80d0170967d36c5eddc387396359bb0caafe89c",
    "transactionIndex": "0x1",
    "blockHash": "0x8e4418b9ed9de011aaf66f8122488e6bc209f19b2a52653cd9e06780ac2f11c2",
    "logIndex": "0xa",
    "removed": false
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
