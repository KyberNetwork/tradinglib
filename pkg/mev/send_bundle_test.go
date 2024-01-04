package mev_test

import (
	"context"
	"math/big"
	"net/http"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/convert"
	"github.com/KyberNetwork/tradinglib/pkg/mev"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

func TestSendBundle(t *testing.T) {
	t.Skip()
	var (
		rawKey         = "...."
		endpoint       = "https://relay-sepolia.flashbots.net"
		ctx            = context.Background()
		client         = http.DefaultClient
		sepoliaChainID = 11155111
	)
	privateKey, err := crypto.HexToECDSA(rawKey)
	require.NoError(t, err)
	address := crypto.PubkeyToAddress(privateKey.PublicKey)

	ethClient, err := ethclient.Dial("https://ethereum-sepolia.blockpi.network/v1/rpc/public")
	require.NoError(t, err)

	blockNumber, err := ethClient.BlockNumber(ctx)
	require.NoError(t, err)
	t.Log("blockNumber", blockNumber)

	nonce, err := ethClient.PendingNonceAt(ctx, address)
	require.NoError(t, err)

	tip, err := convert.FloatToWei(0.3, 18)
	require.NoError(t, err)
	fee, err := convert.FloatToWei(1, 18)
	require.NoError(t, err)

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(int64(sepoliaChainID)),
		Nonce:     nonce,
		To:        &address,
		GasTipCap: tip,
		GasFeeCap: fee,
		Value:     big.NewInt(1),
	})
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(big.NewInt(int64(sepoliaChainID))), privateKey)
	require.NoError(t, err)

	t.Log("new tx", signedTx.Hash().String())

	sender := mev.NewClient(client, endpoint, privateKey)
	resp, err := sender.SendBundle(ctx, blockNumber+12, signedTx)
	require.NoError(t, err) // sepolia: code: [-32000], message: [internal server error]

	t.Log("send bundle response", resp)
}
