package mev_test

import (
	"context"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/mev"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

func TestBloxrouteBackrunmeSender_SendBackrunBundle(t *testing.T) {
	t.Skip()

	sender, err := mev.NewBloxrouteBackrunmeSender("YzJjNTM5MDAtMmZiYy00M2Q0LTkyNGMtYTU3YTk4NzUwZDJlOmZjMDFmOTIxMjZhYmFlOGE0MjRmYmJlNzU1ZGYwMzBh", "https://backrunme.blxrbdn.com")
	require.NoError(t, err)

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

	pendingTXhash := common.HexToHash("0x79d48b1a25d7af0d815997d2ce3a127560080971c5ea98ca5a32424f604e09fb")

	resp, err := sender.SendBackrunBundle(context.Background(), nil, 1, 1, []common.Hash{pendingTXhash}, []string{}, tx)
	require.NoError(t, err)
	t.Log("resp", resp.Result.BundleHash)
}
