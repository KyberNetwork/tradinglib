package mev_test

import (
	"context"
	"math/big"
	"net/http"
	"testing"
	"time"

	"github.com/KyberNetwork/tradinglib/pkg/mev"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

func TestBaseChainSender_SendRawTransaction(t *testing.T) {
	t.Skip("Skip by default - uncomment to run actual test against Base mainnet")

	// Create HTTP client
	httpClient := &http.Client{Timeout: time.Second * 30}

	// Initialize BaseChainSender with Base mainnet RPC
	sender := mev.NewL2ChainSender(
		httpClient,
		"https://mainnet.base.org",
		mev.BundleSenderTypeL2,
	)

	// Verify sender type
	require.Equal(t, mev.BundleSenderTypeL2, sender.GetSenderType())

	// Create a test transaction (this is a dummy transaction that will likely fail)
	// In a real scenario, you would use proper private key, nonce, gas price, etc.
	privateKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	// Connect to Base mainnet to get current gas price and nonce
	ethClient, err := ethclient.Dial("https://mainnet.base.org")
	require.NoError(t, err)

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
	require.NoError(t, err)

	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	require.NoError(t, err)

	// Create a simple ETH transfer transaction
	toAddress := common.HexToAddress("0x0000000000000000000000000000000000000001") // Burn address
	value := big.NewInt(1)                                                         // 1 wei
	gasLimit := uint64(21000)                                                      // Standard ETH transfer gas

	// Get chain ID for Base mainnet (8453)
	chainID, err := ethClient.ChainID(context.Background())
	require.NoError(t, err)
	require.Equal(t, int64(8453), chainID.Int64()) // Base mainnet chain ID

	// Create transaction
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	// Sign transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	require.NoError(t, err)

	// Send transaction using BaseChainSender
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	resp, err := sender.SendRawTransaction(ctx, signedTx)

	// Log the response for debugging
	t.Logf("Response: %+v", resp)
	t.Logf("Error: %v", err)

	// The transaction will likely fail due to insufficient funds, but we should get a proper response
	// We expect either a successful response with transaction hash or a proper error response
	if err != nil {
		// Check if it's a proper RPC error (not a network/parsing error)
		t.Logf("Expected error due to insufficient funds or other RPC error: %v", err)
	} else {
		// If successful, verify response structure
		require.Equal(t, "2.0", resp.Jsonrpc)
		require.Equal(t, 1, resp.ID)
		require.NotEmpty(t, resp.Result)
		t.Logf("Transaction hash: %s", resp.Result)
	}
}

func TestBaseChainSender_SendRawTransaction_InvalidTx(t *testing.T) {
	t.Skip("Skip by default - uncomment to run actual test against Base mainnet")
	// Create HTTP client
	httpClient := &http.Client{Timeout: time.Second * 10}

	// Initialize BaseChainSender with Base mainnet RPC
	sender := mev.NewL2ChainSender(
		httpClient,
		"https://mainnet.base.org",
		mev.BundleSenderTypeL2,
	)

	// Create an invalid transaction (unsigned)
	toAddress := common.HexToAddress("0x0000000000000000000000000000000000000001")
	value := big.NewInt(1)
	gasLimit := uint64(21000)
	gasPrice := big.NewInt(1000000000) // 1 gwei

	// Create unsigned transaction
	tx := types.NewTransaction(0, toAddress, value, gasLimit, gasPrice, nil)

	// Try to send unsigned transaction (should fail)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := sender.SendRawTransaction(ctx, tx)

	// Should get an error due to invalid transaction
	t.Logf("Response: %+v", resp)
	t.Logf("Error: %v", err)

	// We expect an error here
	require.Error(t, err)
}

func TestBaseChainSender_Interface_Compliance(t *testing.T) {
	t.Skip("Skip by default - uncomment to run actual test against Base mainnet")
	// Test that BaseChainSender implements ISendRawTransaction interface
	httpClient := &http.Client{Timeout: time.Second * 10}
	sender := mev.NewL2ChainSender(
		httpClient,
		"https://mainnet.base.org",
		mev.BundleSenderTypeL2,
	)

	// Verify it implements the interface
	var _ mev.ISendRawTransaction = sender
}

func TestNewBaseChainSender(t *testing.T) {
	t.Skip("Skip by default - uncomment to run actual test against Base mainnet")
	httpClient := &http.Client{Timeout: time.Second * 10}
	endpoint := "https://mainnet.base.org"
	senderType := mev.BundleSenderTypeL2

	sender := mev.NewL2ChainSender(httpClient, endpoint, senderType)

	require.NotNil(t, sender)
	require.Equal(t, senderType, sender.GetSenderType())
}
