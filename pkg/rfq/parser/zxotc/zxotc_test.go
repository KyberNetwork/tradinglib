package zxotc

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
	require.Equal(t, p.abi.Events[OtcOrderFilledEvent].ID, common.HexToHash("0xac75f773e3a92f1a02b12134d65e1f47f8a14eabe4eaf1e24624918e6a8b269f"))
	client, err := ethclient.Dial("https://eth.llamarpc.com")
	require.NoError(t, err)
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: big.NewInt(16282125),
		ToBlock:   big.NewInt(16282255),
		Addresses: nil,
		Topics: [][]common.Hash{
			{
				common.HexToHash("0xac75f773e3a92f1a02b12134d65e1f47f8a14eabe4eaf1e24624918e6a8b269f"),
			},
		},
	})
	require.NoError(t, err)
	d, _ := json.Marshal(logs)
	t.Log(string(d))
}

func TestParseEvent(t *testing.T) {
	eventRaw := `{
        "address": "0xdef1c0ded9bec7f1a1670819833240f027b25eff",
        "topics": [
            "0xac75f773e3a92f1a02b12134d65e1f47f8a14eabe4eaf1e24624918e6a8b269f"
        ],
        "data": "0xb1db529b00f4a98a4c0dbf8fc7c888be9cd20da5b95136c0887aef7a6a96637d000000000000000000000000af0b0000f0210d0f421f0009c72406703b50506b00000000000000000000000074de5d4fcbf63e00296fd95d33236b9794016631000000000000000000000000c944e90c64b2c07662a292be6244bdf05cda44a7000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb4800000000000000000000000000000000000000000000124272b1b960ca00000000000000000000000000000000000000000000000000000000000001276a5f90",
        "blockNumber": "0xf87219",
        "transactionHash": "0x4153c12bbef7fcae36bed6a1b27b42f72b15302afea70e2c5aa52022f92a7eca",
        "transactionIndex": "0x15",
        "blockHash": "0x15db0c1f7d0fcf132f6598e4420bf45de265f4b9d8b168383a049f62bdcd0485",
        "logIndex": "0x3b",
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
