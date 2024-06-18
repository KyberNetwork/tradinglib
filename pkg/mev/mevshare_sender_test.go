package mev_test

import (
	"context"
	"log"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/mev"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSendBackrunBundle(t *testing.T) {
	t.Skip()

	ethClient, err := ethclient.Dial("https://ethereum-rpc.publicnode.com")
	require.NoError(t, err)

	blockNumber, err := ethClient.BlockNumber(context.Background())
	require.NoError(t, err)
	t.Log("blockNumber", blockNumber)

	// Transaction hashes you want to fetch from the node
	txHashes := []string{
		"0x2e038916d175d9028c87d59e33f79ac96cb487e90aad6cd501dc9675b64d7245",
	}
	tx, isPending, err := ethClient.TransactionByHash(context.Background(), common.HexToHash(txHashes[0]))
	require.NoError(t, err)
	require.False(t, isPending)

	// Flashbots header signing key
	fbSigningKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the client
	rpcClient, err := mev.NewMevShareSender("https://relay.flashbots.net", fbSigningKey)
	require.NoError(t, err)

	pendingTxHash := common.HexToHash("0x73767bdb9dd83040fa242887100bc460f1fdb56d7d7934ce2d21f2a1fa109e4f")

	// Send bundle
	res, err := rpcClient.SendBackrunBundle(context.Background(), nil, blockNumber+1, pendingTxHash, nil, tx)
	assert.Nil(t, err)

	t.Log(res, "result")
}
