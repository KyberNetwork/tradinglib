package mev_test

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/json"
	"math/big"
	"net/http"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/convert"
	"github.com/KyberNetwork/tradinglib/pkg/mev"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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
		endpoint       = "https://relay-sepolia.flashbots.net"
		ctx            = context.Background()
		client         = http.DefaultClient
		sepoliaChainID = 11155111
	)
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

	uuid := uuid.NewString()
	ethClient, err = ethclient.Dial(endpoint)
	require.NoError(t, err)
	gasBundleEstimator := mev.NewGasBundleEstimator(ethClient)
	sender, err := mev.NewClient(client, endpoint, privateKey, false, mev.BundleSenderTypeFlashbot, gasBundleEstimator)
	require.NoError(t, err)

	resp, err := sender.SendBundle(ctx, &uuid, blockNumber+12, signedTx)
	require.NoError(t, err) // sepolia: code: [-32000], message: [internal server error]
	t.Log("send bundle response", resp)

	require.NoError(t, sender.CancelBundle(ctx, uuid))
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

	ethClient, err := ethclient.Dial(endpoint)
	require.NoError(t, err)
	gasBundleEstimator := mev.NewGasBundleEstimator(ethClient)

	sender, err := mev.NewClient(client, endpoint, nil, true, mev.BundleSenderTypeBeaver, gasBundleEstimator)
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
	rawTxs := []string{
		"0xf868808502540be400825208944592d8f8d7b001e72cb26a73e4fa1806a51ac79d82271080820a96a017691fd52972f52132d2db29c305afd89f50924b0cbc6e875ec8c6bcca14d287a015db7d1ab52dd19e40e7fe91018df63df9d5c5a5e21541dc2c8548a5ae9cee37", // nolint:lll
		"0xf868018502540be400825208944592d8f8d7b001e72cb26a73e4fa1806a51ac79d82271080820a96a0aba458c7c6d1feac6c414c8f8cf562251609a2fb6710fcfce0c7783b106e5f41a00c0e780a9ea923e177dcc52aa143887ec21c9697f305a54f41b648f733e98d3e", // nolint:lll
		"0xf868028502540be400825208944592d8f8d7b001e72cb26a73e4fa1806a51ac79d82271080820a95a05ddce4890eac66bfb20eba1493d12093c7551f1f5269a31487c718ee0ea2d12ca02ca061024dfad8d35cce7233eb3f75cdfc57ea80547f49ed887c1f257f4db719", // nolint:lll
	}

	blockNumber := 1

	txs := make([]*types.Transaction, 0, len(rawTxs))
	for _, rawTx := range rawTxs {
		var tx types.Transaction
		b, err := hexutil.Decode(rawTx)
		require.NoError(t, err)
		err = tx.UnmarshalBinary(b)
		require.NoError(t, err)
		txs = append(txs, &tx)
	}

	simulationEndpoint := "http://localhost:8545"
	ethClient, err := ethclient.Dial(simulationEndpoint)
	require.NoError(t, err)
	gasBundleEstimator := mev.NewGasBundleEstimator(ethClient)

	client, err := mev.NewClient(http.DefaultClient,
		simulationEndpoint, nil, false,
		mev.BundleSenderTypeFlashbot, gasBundleEstimator)
	require.NoError(t, err)

	simulationResponse, err := client.SimulateBundle(context.Background(), uint64(blockNumber), txs...)
	require.NoError(t, err)

	assert.Equal(t, "", simulationResponse.Error.Messange)
	assert.Equal(t, 0, simulationResponse.Error.Code)
	assert.Equal(
		t,
		"0x99872010193b755b7dfad328508c751173521ee9b5349eab111b33096bf9e19a",
		simulationResponse.Result.BundleHash,
	)
}
