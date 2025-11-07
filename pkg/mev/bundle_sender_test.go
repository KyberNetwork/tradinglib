package mev_test

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/convert"
	"github.com/KyberNetwork/tradinglib/pkg/mev"
	"github.com/duoxehyon/mev-share-go/rpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/flashbots/mev-share-node/mevshare"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSendBundle(t *testing.T) {
	t.Skip()

	// Generate a new private key
	privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		t.Error("Failed to generate private key:", err)
		return
	}
	var (
		endpoint       = "https://rpc.titanbuilder.xyz"
		ctx            = context.Background()
		client         = http.DefaultClient
		sepoliaChainID = 1
	)
	require.NoError(t, err)
	address := crypto.PubkeyToAddress(privateKey.PublicKey)

	ethClient, err := ethclient.Dial("wss://ethereum.kyberengineering.io")
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

	uuid := uuid.NewString()
	require.NoError(t, err)
	sender, err := mev.NewClient(client, endpoint, privateKey, mev.BundleSenderTypeFlashbot, false, false)
	require.NoError(t, err)

	resp, err := sender.SendBundle(ctx, &uuid, blockNumber+12, signedTx)
	require.NoError(t, err) // sepolia: code: [-32000], message: [internal server error]
	t.Log("send bundle response", resp)

	signedTxBin, err := signedTx.MarshalBinary()
	require.NoError(t, err)

	resp, err = sender.SendBundleHex(ctx, &uuid, blockNumber+12, hexutil.Encode(signedTxBin))
	require.NoError(t, err) // sepolia: code: [-32000], message: [internal server error]
	t.Log("send bundle response hex", resp)

	require.NoError(t, sender.CancelBundle(ctx, uuid))
}

func Ptr[T any](v T) *T {
	return &v
}

func TestSendBundleV2(t *testing.T) {
	t.Skip()

	// Generate a new private key
	privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		t.Error("Failed to generate private key:", err)
		return
	}
	var (
		endpoint   = "https://bsc.blinklabs.xyz/v1/<API_KEY>"
		ctx        = context.Background()
		client     = http.DefaultClient
		bscChainID = 56
	)
	require.NoError(t, err)
	address := crypto.PubkeyToAddress(privateKey.PublicKey)

	ethClient, err := ethclient.Dial("wss://bsc.kyberengineering.io")
	require.NoError(t, err)

	head, err := ethClient.HeaderByNumber(ctx, nil)
	require.NoError(t, err)
	t.Log("blockNumber", head.Number)

	nonce, err := ethClient.PendingNonceAt(ctx, address)
	require.NoError(t, err)

	tip, err := convert.FloatToWei(0.3, 18)
	require.NoError(t, err)
	fee, err := convert.FloatToWei(1, 18)
	require.NoError(t, err)

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(int64(bscChainID)),
		Nonce:     nonce,
		To:        &address,
		GasTipCap: tip,
		GasFeeCap: fee,
		Value:     big.NewInt(1),
	})
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(big.NewInt(int64(bscChainID))), privateKey)
	require.NoError(t, err)

	t.Log("new tx", signedTx.Hash().String())

	uuid := uuid.NewString()
	require.NoError(t, err)
	sender, err := mev.NewClient(client, endpoint, privateKey, mev.BundleSenderTypeBlink, false, false)
	require.NoError(t, err)

	resp, err := sender.SendBundleV2(ctx, mev.SendBundleV2Request{
		MinTimestamp: Ptr(head.Time + 1),
		MaxTimestamp: Ptr(head.Time + 5),
		UUID:         Ptr(uuid),
	})
	require.NoError(t, err) // sepolia: code: [-32000], message: [internal server error]
	t.Log("send bundle response", resp)

	// require.NoError(t, sender.CancelBundle(ctx, uuid))
}

func TestUnmarshal(t *testing.T) {
	var (
		data = "{\"id\":\"1\",\"result\":{\"bundleHash\":\"0xe0e0592348830d57fac820a6bec9fdbf0ac20a2b503351c63217cf8c274b70a8\"},\"jsonrpc\":\"2.0\"}\n" // nolint:lll
		resp mev.BLXRSubmitBundleResponse
	)

	require.NoError(t, json.Unmarshal([]byte(data), &resp))

	t.Logf("%+v", resp)
}

func TestCancelBeaver(t *testing.T) {
	t.Skip()
	var (
		endpoint   = "https://rpc.beaverbuild.org"
		ctx        = context.Background()
		client     = http.DefaultClient
		bundleUUID = uuid.New().String()
	)

	sender, err := mev.NewClient(client, endpoint, nil, mev.BundleSenderTypeBeaver, true, false)
	require.NoError(t, err)

	require.NoError(t, sender.CancelBundle(ctx, bundleUUID))
}

func Test_UnmarshalSimulationResponse(t *testing.T) {
	response := "{\n    \"jsonrpc\": \"2.0\",\n    \"id\": 1,\n    \"result\": {\n        \"bundleGasPrice\": \"1\",\n        \"bundleHash\": \"0x4753e95178e232c1cd0436acbb2340d9fe3442331c4650379fb436c7ee8c8489\",\n        \"coinbaseDiff\": \"63000\",\n        \"ethSentToCoinbase\": \"0\",\n        \"gasFees\": \"63000\",\n        \"results\": [\n            {\n                \"coinbaseDiff\": \"21000\",\n                \"ethSentToCoinbase\": \"0\",\n                \"fromAddress\": \"0xf84bB4749ef5745258812243B65d6Ec06B52Cc4f\",\n                \"gasFees\": \"21000\",\n                \"gasPrice\": \"1\",\n                \"gasUsed\": 21000,\n                \"toAddress\": \"0x4592D8f8D7B001e72Cb26A73e4Fa1806a51aC79d\",\n                \"txHash\": \"0x31c0d14c4cf1dceaecad2b028472490fc7ed5a3b7f6cdcb78fa26673448b5665\",\n                \"value\": \"0x\"\n            },\n            {\n                \"coinbaseDiff\": \"21000\",\n                \"ethSentToCoinbase\": \"0\",\n                \"fromAddress\": \"0xf84bB4749ef5745258812243B65d6Ec06B52Cc4f\",\n                \"gasFees\": \"21000\",\n                \"gasPrice\": \"1\",\n                \"gasUsed\": 21000,\n                \"toAddress\": \"0x4592D8f8D7B001e72Cb26A73e4Fa1806a51aC79d\",\n                \"txHash\": \"0xe7e261a582b11be10ded10262e98a938230ecae1adc155e23d5cc805021d10f4\",\n                \"value\": \"0x\"\n            },\n            {\n                \"coinbaseDiff\": \"21000\",\n                \"ethSentToCoinbase\": \"0\",\n                \"fromAddress\": \"0xf84bB4749ef5745258812243B65d6Ec06B52Cc4f\",\n                \"gasFees\": \"21000\",\n                \"gasPrice\": \"1\",\n                \"gasUsed\": 21000,\n                \"toAddress\": \"0x4592D8f8D7B001e72Cb26A73e4Fa1806a51aC79d\",\n                \"txHash\": \"0x7ad464764e279a1849f517c83c459b0088b454f0928f61d0c3882ce09483e2d1\",\n                \"value\": \"0x\"\n            }\n        ],\n        \"stateBlockNumber\": 1,\n        \"totalGasUsed\": 63000\n    }\n}" // nolint:lll
	var submitResponse mev.SendBundleResponse

	require.NoError(t, json.Unmarshal([]byte(response), &submitResponse))

	t.Logf("%+v", submitResponse)
}

func Test_SimulateBundle(t *testing.T) {
	t.Skip()
	nodeEndpoint := "https://ethereum-rpc.publicnode.com"
	ethClient, err := ethclient.Dial(nodeEndpoint)
	require.NoError(t, err)

	// Transaction hashes you want to fetch from the node
	txHashes := []string{
		"0x5deec444557cb413fc483e517454eb2f7a717e922af60cd79a223ea9741299b3",
		"0x2e038916d175d9028c87d59e33f79ac96cb487e90aad6cd501dc9675b64d7245",
	}
	blockNumber := 19738428
	txs := make([]*types.Transaction, 0, len(txHashes))
	for _, hash := range txHashes {
		tx, isPending, err := ethClient.TransactionByHash(context.Background(), common.HexToHash(hash))
		require.NoError(t, err)
		require.False(t, isPending)
		txs = append(txs, tx)
	}

	simulationEndpoint := "https://relay.flashbots.net"
	require.NoError(t, err)

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	client, err := mev.NewClient(http.DefaultClient,
		simulationEndpoint, privateKey,
		mev.BundleSenderTypeFlashbot, false, false)
	require.NoError(t, err)

	simulationResponse, err := client.SimulateBundle(context.Background(), uint64(blockNumber), txs...) // nolint:gosec
	require.NoError(t, err)

	assert.Equal(t, "", simulationResponse.Error.Messange)
	assert.Equal(t, 0, simulationResponse.Error.Code)
	assert.Equal(
		t,
		"0x99872010193b755b7dfad328508c751173521ee9b5349eab111b33096bf9e19a",
		simulationResponse.Result.BundleHash,
	)
}

func TestMevSendBundle(t *testing.T) {
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

	// Serialize the transaction to RLP
	rlpEncodedTx, err := tx.MarshalBinary()
	if err != nil {
		log.Fatalf("Failed to encode transaction: %v", err)
	}

	// Flashbots header signing key
	fbSigningKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the client
	rpcClient := rpc.NewClient("https://relay.flashbots.net", fbSigningKey)

	txBytes := hexutil.Bytes(rlpEncodedTx)

	pendingTXhash := common.HexToHash("0x2e038916d175d9028c87d59e33f79ac96cb487e90aad6cd501dc9675b64d7245")
	// Define the bundle transactions
	txns := []mevshare.MevBundleBody{
		{
			Hash: &pendingTXhash,
		},
		{
			Tx: &txBytes,
		},
	}
	inclusion := mevshare.MevBundleInclusion{
		BlockNumber: hexutil.Uint64(blockNumber + 1),
	}
	// Make the bundle
	req := mevshare.SendMevBundleArgs{
		Body:      txns,
		Inclusion: inclusion,
	}

	// Send bundle
	res, err := rpcClient.SendBundle(req)
	assert.Nil(t, err)

	t.Log(res.BundleHash.String(), "bundleHash")
}

func TestClient_GetBundleStats(t *testing.T) {
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

	// Serialize the transaction to RLP
	rlpEncodedTx, err := tx.MarshalBinary()
	if err != nil {
		log.Fatalf("Failed to encode transaction: %v", err)
	}

	// Flashbots header signing key
	fbSigningKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	SimulationEndpoint := "https://relay.flashbots.net"
	// Initialize the client
	rpcClient := rpc.NewClient(SimulationEndpoint, fbSigningKey)

	txBytes := hexutil.Bytes(rlpEncodedTx)

	pendingTxHash := common.HexToHash("0x4ad277ae1dfba88e54bc68e81b345920691e6bf892f8799f2b0996ace875b1bf")
	// Define the bundle transactions
	txns := []mevshare.MevBundleBody{
		{
			Hash: &pendingTxHash,
		},
		{
			Tx: &txBytes,
		},
	}
	inclusion := mevshare.MevBundleInclusion{
		BlockNumber: hexutil.Uint64(blockNumber + 1),
	}
	// Make the bundle
	req := mevshare.SendMevBundleArgs{
		Body:      txns,
		Inclusion: inclusion,
		Privacy: &mevshare.MevBundlePrivacy{Builders: []string{
			mev.FlashbotBuilderRegistrationBobaBuilder,
			mev.FlashbotBuilderRegistrationFlashbot,
		}},
	}

	// Send bundle
	res, err := rpcClient.SendBundle(req)
	assert.Nil(t, err)

	t.Log(res.BundleHash.String(), "bundleHash")

	client, err := mev.NewClient(http.DefaultClient,
		SimulationEndpoint, fbSigningKey,
		mev.BundleSenderTypeFlashbot, false, false)
	require.NoError(t, err)
	// Get bundle stats
	stats, err := client.GetBundleStats(context.Background(), blockNumber+1, res.BundleHash)
	assert.NoError(t, err)

	t.Log(stats)
}

func TestGetUserStats(t *testing.T) {
	t.Skip()
	// Generate a new private key
	privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		t.Error("Failed to generate private key:", err)
		return
	}
	var (
		endpoint = "https://rpc.titanbuilder.xyz"
		ctx      = context.Background()
	)

	ethClient, err := ethclient.Dial("wss://ethereum.kyberengineering.io")
	require.NoError(t, err)

	blockNumber, err := ethClient.BlockNumber(ctx)
	require.NoError(t, err)
	t.Log("blockNumber", blockNumber)

	bundleSender, err := mev.NewClient(
		http.DefaultClient,
		endpoint,
		privateKey,
		mev.BundleSenderTypeTitan,
		false,
		false,
	)
	require.NoError(t, err)

	resp, err := bundleSender.GetUserStats(ctx, false, blockNumber)
	require.NoError(t, err)

	t.Log(resp)
}
