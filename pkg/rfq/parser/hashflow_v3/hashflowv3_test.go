package hashflowv3

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
	require.Equal(t, p.abi.Events[TradeEvent].ID, common.HexToHash("0x34f57786fb01682fb4eec88d340387ef01a168fe345ea5b76f709d4e560c10eb"))
	client, err := ethclient.Dial("https://eth.llamarpc.com")
	require.NoError(t, err)
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(18267792),
		ToBlock:   big.NewInt(18268792),
		Addresses: nil,
		Topics:    [][]common.Hash{{common.HexToHash("0x34f57786fb01682fb4eec88d340387ef01a168fe345ea5b76f709d4e560c10eb")}},
	})
	require.NoError(t, err)
	d, _ := json.Marshal(logs)
	t.Log(string(d))
}

func TestParseEvent(t *testing.T) {
	eventRaw := `{
		"address":"0x403a1a99095d95501f57269d0f137e5790d40406",
		"topics":[
			"0x34f57786fb01682fb4eec88d340387ef01a168fe345ea5b76f709d4e560c10eb"
			],
		"data":"0x00000000000000000000000013414b047539298d5aed429722211681eaab43b700000000000000000000000013414b047539298d5aed429722211681eaab43b7100100000000000000001748090b00ffffffffffffff001b95ca55a556b000000000000000000000000000006b175474e89094c44da98b954eedeac495271d0f000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000000000000000000000000000004563918244f40000000000000000000000000000000000000000000000000000000aa888b55c36e5",
		"blockNumber":"0x116c013",
		"transactionHash":"0x092681c0f57a996361aa0e6634a5f8ccd0c59cafeb9ceb746f2c5ba0692b39da",
		"transactionIndex":"0x41",
		"blockHash":"0xb77e1bdfe3f03464b4a70cb81e40b6b89704e773d4daac8328ac9bbb000735a0",
		"logIndex":"0x5d",
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
