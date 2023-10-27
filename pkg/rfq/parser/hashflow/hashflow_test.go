package hashflow

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
	require.Equal(t, p.abi.Events[TradeEvent].ID, common.HexToHash("0x8cf3dec1929508e5677d7db003124e74802bfba7250a572205a9986d86ca9f1e"))
	client, err := ethclient.Dial("https://eth.llamarpc.com")
	require.NoError(t, err)
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(17363812),
		ToBlock:   big.NewInt(17363812),
		Addresses: nil,
		Topics:    [][]common.Hash{{common.HexToHash("0x8cf3dec1929508e5677d7db003124e74802bfba7250a572205a9986d86ca9f1e")}},
	})
	require.NoError(t, err)
	d, _ := json.Marshal(logs)
	t.Log(string(d))
}

func TestParseEvent(t *testing.T) {
	eventRaw := `{
		"address":"0x4306aefe6c0e605f22207a43203ae7d7e4fb42c9",
		"topics":[
			"0x8cf3dec1929508e5677d7db003124e74802bfba7250a572205a9986d86ca9f1e"
			],
		"data":"0x0000000000000000000000009863470bd3549563d5499c179ede3c0fcbcf3eb40067616e64616c6674686562726f776e67786d786e6900190878e8b3c77600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7000000000000000000000000000000000000000000000000002386f26fc10000000000000000000000000000000000000000000000000000000000000121a0f6",
		"blockNumber":"0x108f364",
		"transactionHash":"0x2345ecbe3c5f79d11b632bbd606ed74d7ada809992945bb4ed67f826660229af",
		"transactionIndex":"0x20",
		"blockHash":"0x1057d2b4e14ff582aa1571c4d26c28f662825f2cfb01cfed4009dd125248e5fa",
		"logIndex":"0xe0",
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
