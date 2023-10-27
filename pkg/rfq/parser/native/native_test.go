package native

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
	require.Equal(t, p.abi.Events[SwapEvent].ID, common.HexToHash("0xe3a54b69726c85299f4e794bac96150af56af801be76cafd11947a1103b6308a"))
	client, err := ethclient.Dial("https://eth.llamarpc.com")
	require.NoError(t, err)
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(18176063),
		ToBlock:   big.NewInt(18176063),
		Addresses: nil,
		Topics:    [][]common.Hash{{common.HexToHash("0xe3a54b69726c85299f4e794bac96150af56af801be76cafd11947a1103b6308a")}},
	})
	require.NoError(t, err)
	d, _ := json.Marshal(logs)
	t.Log(string(d))
}

func TestParseEvent(t *testing.T) {
	eventRaw := `{
		"address":"0x3d130bf4686b3d4b6eb91a8e26ac629c5bea6082",
		"topics":["0xe3a54b69726c85299f4e794bac96150af56af801be76cafd11947a1103b6308a","0x0000000000000000000000003208684f96458c540eb08f6f01b9e9afb2b7d4f0","0x0000000000000000000000003208684f96458c540eb08f6f01b9e9afb2b7d4f0"],
		"data":"0x000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7000000000000000000000000000000000000000000000000191b9c77998fd61cffffffffffffffffffffffffffffffffffffffffffffffffffffffff4f6cd327000000000000000000000000000000000000000000000000000000000000000077bc10edff9b41b1b9fddc1965b20b9500000000000000000000000000000000",
		"blockNumber":"0x115583f",
		"transactionHash":"0x9e452a5d1c537a2383fd25adf038785a7104cf3763e7537f0fcc61422746b15f",
		"transactionIndex":"0x75",
		"blockHash":"0xd50211ae01767f277610a36c2c515e76ebb13775e29bf404196fe8a6d76602a4",
		"logIndex":"0x124",
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
