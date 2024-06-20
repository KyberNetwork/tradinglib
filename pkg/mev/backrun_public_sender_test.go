package mev_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/KyberNetwork/tradinglib/pkg/mev"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

func TestNewBackrunPublicClient(t *testing.T) {
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

	httpCl := http.Client{Timeout: time.Second * 5}
	// Initialize the client
	senderClient := mev.NewBackrunPublicClient(&httpCl, "https://rpc.mevblocker.io", nil, mev.BundleSenderTypeMevBlocker)

	pendingTXhash := common.HexToHash("0x79d48b1a25d7af0d815997d2ce3a127560080971c5ea98ca5a32424f604e09fb")
	resp, err := senderClient.SendBackrunBundle(context.Background(), nil, blockNumber, pendingTXhash, []string{}, tx)
	t.Log("resp", resp)
	t.Log("err", err)
}
