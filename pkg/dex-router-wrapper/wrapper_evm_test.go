//nolint:testpackage
package dexrouterwrapper

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	_ "embed"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// This file exercises DexRouterWrapper against a real, in-process EVM (via
// go-ethereum's simulated backend), instead of only round-tripping Go-side
// ABI encoding. It deploys the wrapper both for real and via a state
// override (through eth_simulateV1, since the simulated backend
// deliberately hides the raw rpc.Client that gethclient's override-aware
// eth_call needs), so a bug that mixed up creation vs. deployed bytecode
// would fail here.

//go:embed testdata/MockRouter.abi.json
var mockRouterABIJSON []byte

//go:embed testdata/MockRouter.bin
var mockRouterCreationBytecodeHex string

type evmTestEnv struct {
	backend *simulated.Backend
	client  simulated.Client
	key     *ecdsa.PrivateKey
	from    common.Address
	chainID *big.Int
}

func newEVMTestEnv(t *testing.T) *evmTestEnv {
	t.Helper()

	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	from := crypto.PubkeyToAddress(key.PublicKey)

	balance := new(big.Int).Mul(big.NewInt(1000), big.NewInt(params.Ether))
	backend := simulated.NewBackend(types.GenesisAlloc{
		from: {Balance: balance},
	})
	t.Cleanup(func() { require.NoError(t, backend.Close()) })

	client := backend.Client()

	chainID, err := client.ChainID(context.Background())
	require.NoError(t, err)

	return &evmTestEnv{
		backend: backend,
		client:  client,
		key:     key,
		from:    from,
		chainID: chainID,
	}
}

// deploy sends creationBytecode as a contract-creation transaction and
// returns the resulting contract address.
func (env *evmTestEnv) deploy(ctx context.Context, t *testing.T, creationBytecode []byte) common.Address {
	t.Helper()

	nonce, err := env.client.PendingNonceAt(ctx, env.from)
	require.NoError(t, err)

	head, err := env.client.HeaderByNumber(ctx, nil)
	require.NoError(t, err)

	gasFeeCap := new(big.Int).Add(head.BaseFee, big.NewInt(params.GWei))
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   env.chainID,
		Nonce:     nonce,
		GasTipCap: big.NewInt(params.GWei),
		GasFeeCap: gasFeeCap,
		Gas:       3_000_000,
		Data:      creationBytecode,
	})

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(env.chainID), env.key)
	require.NoError(t, err)

	require.NoError(t, env.client.SendTransaction(ctx, signedTx))
	env.backend.Commit()

	receipt, err := env.client.TransactionReceipt(ctx, signedTx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status, "contract deployment must succeed")

	return receipt.ContractAddress
}

// sendValue sends amount of native ETH from env.from to addr as a plain
// transaction, e.g. to pre-fund a contract before a test calls it.
func (env *evmTestEnv) sendValue(ctx context.Context, t *testing.T, addr common.Address, amount *big.Int) {
	t.Helper()

	nonce, err := env.client.PendingNonceAt(ctx, env.from)
	require.NoError(t, err)

	head, err := env.client.HeaderByNumber(ctx, nil)
	require.NoError(t, err)

	gasFeeCap := new(big.Int).Add(head.BaseFee, big.NewInt(params.GWei))
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   env.chainID,
		Nonce:     nonce,
		GasTipCap: big.NewInt(params.GWei),
		GasFeeCap: gasFeeCap,
		Gas:       100_000,
		To:        &addr,
		Value:     amount,
	})

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(env.chainID), env.key)
	require.NoError(t, err)

	require.NoError(t, env.client.SendTransaction(ctx, signedTx))
	env.backend.Commit()

	receipt, err := env.client.TransactionReceipt(ctx, signedTx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status, "pre-funding transfer must succeed")
}

// simulateV1Caller is the subset of *ethclient.Client's SimulateV1 method
// that simulated.Client's interface doesn't expose statically, even though
// the backend's concrete client implements it.
type simulateV1Caller interface {
	SimulateV1(
		ctx context.Context, opts ethclient.SimulateOptions, blockNrOrHash *rpc.BlockNumberOrHash,
	) ([]ethclient.SimulateBlockResult, error)
}

// callWithOverride runs msg through eth_simulateV1 with the given state
// overrides applied and returns the call's raw return data.
func (env *evmTestEnv) callWithOverride(
	ctx context.Context, t *testing.T, msg ethereum.CallMsg, overrides map[common.Address]ethereum.OverrideAccount,
) ([]byte, error) {
	t.Helper()

	caller, ok := env.client.(simulateV1Caller)
	require.True(t, ok, "simulated client must support eth_simulateV1")

	results, err := caller.SimulateV1(ctx, ethclient.SimulateOptions{
		BlockStateCalls: []ethclient.SimulateBlock{
			{
				StateOverrides: overrides,
				Calls:          []ethereum.CallMsg{msg},
			},
		},
	}, nil)
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Len(t, results[0].Calls, 1)

	call := results[0].Calls[0]
	if call.Error != nil {
		return nil, fmt.Errorf("simulated call reverted: %s", call.Error.Message)
	}

	return call.ReturnValue, nil
}

func mustParseABI(t *testing.T, data []byte) abi.ABI {
	t.Helper()
	parsed, err := abi.JSON(bytes.NewReader(data))
	require.NoError(t, err)

	return parsed
}

func mustDecodeHexBytecode(t *testing.T, hexStr string) []byte {
	t.Helper()
	code, err := hex.DecodeString(strings.TrimSpace(hexStr))
	require.NoError(t, err)

	return code
}

func TestWrapEVM(t *testing.T) {
	env := newEVMTestEnv(t)
	ctx := context.Background()

	mockRouterABI := mustParseABI(t, mockRouterABIJSON)
	mockRouterCreation := mustDecodeHexBytecode(t, mockRouterCreationBytecodeHex)

	mockRouterAddr := env.deploy(ctx, t, mockRouterCreation)
	wrapperAddr := env.deploy(ctx, t, CreationBytecode())

	recipient := common.HexToAddress("0x00000000000000000000000000000000000c0de")
	amount := big.NewInt(1_000_000_000_000_000) // 0.001 ETH

	swapData, err := mockRouterABI.Pack("swap", recipient, amount)
	require.NoError(t, err)

	wrapCalldata, err := EncodeWrapCalldata(mockRouterAddr, swapData, NativeTokenAddress, recipient)
	require.NoError(t, err)

	t.Run("deployed for real", func(t *testing.T) {
		msg := ethereum.CallMsg{From: env.from, To: &wrapperAddr, Data: wrapCalldata, Value: amount}

		result, err := env.client.CallContract(ctx, msg, nil)
		require.NoError(t, err)

		returnAmount, gasUsed, err := DecodeWrapOutput(result)
		require.NoError(t, err)
		assert.Equal(t, 0, amount.Cmp(returnAmount))
		assert.Positive(t, gasUsed)
	})

	t.Run("injected via state override", func(t *testing.T) {
		// overrideAddr never receives a real deployment; only DeployedBytecode
		// (via StateOverride) makes it behave like DexRouterWrapper.
		overrideAddr := common.HexToAddress("0x00000000000000000000000000000000badc0d")
		msg := ethereum.CallMsg{From: env.from, To: &overrideAddr, Data: wrapCalldata, Value: amount}

		result, err := env.callWithOverride(ctx, t, msg, StateOverride(overrideAddr))
		require.NoError(t, err)

		returnAmount, gasUsed, err := DecodeWrapOutput(result)
		require.NoError(t, err)
		assert.Equal(t, 0, amount.Cmp(returnAmount))
		assert.Positive(t, gasUsed)
	})

	t.Run("bubbles target revert", func(t *testing.T) {
		failData, err := mockRouterABI.Pack("fail")
		require.NoError(t, err)

		failCalldata, err := EncodeWrapCalldata(mockRouterAddr, failData, NativeTokenAddress, recipient)
		require.NoError(t, err)

		msg := ethereum.CallMsg{From: env.from, To: &wrapperAddr, Data: failCalldata}

		_, err = env.client.CallContract(ctx, msg, nil)
		assert.Error(t, err)
	})

	t.Run("reverts when output balance decreases", func(t *testing.T) {
		// Pre-fund mockRouterAddr, then have it pay part of its own balance
		// out to a third party. With recipient == mockRouterAddr, wrap's
		// balance-diff goes negative, which must be a clean revert rather
		// than an underflow panic.
		env.sendValue(ctx, t, mockRouterAddr, big.NewInt(1_000_000_000_000_000))

		bystander := common.HexToAddress("0x0000000000000000000000000000000000b7a5")
		payOutData, err := mockRouterABI.Pack("payOut", bystander, big.NewInt(1))
		require.NoError(t, err)

		payOutCalldata, err := EncodeWrapCalldata(mockRouterAddr, payOutData, NativeTokenAddress, mockRouterAddr)
		require.NoError(t, err)

		msg := ethereum.CallMsg{From: env.from, To: &wrapperAddr, Data: payOutCalldata}

		_, err = env.client.CallContract(ctx, msg, nil)
		assert.Error(t, err)
	})

	t.Run("wrapper itself receives native swap output", func(t *testing.T) {
		// Some routers pay swap output to msg.sender rather than an explicit
		// recipient argument. Since the wrapper is msg.sender of the call to
		// target, that means the payout lands on the wrapper contract
		// itself; without a receive() function the plain ETH transfer would
		// revert and the whole simulated call would fail.
		//
		// mockRouterAddr pays out of its own pre-funded balance (rather than
		// forwarding wrap's msg.value) so that wrapperAddr's balanceBefore
		// snapshot isn't itself inflated by the value of this very call.
		env.sendValue(ctx, t, mockRouterAddr, amount)

		payOutToWrapperData, err := mockRouterABI.Pack("payOut", wrapperAddr, amount)
		require.NoError(t, err)

		selfWrapCalldata, err := EncodeWrapCalldata(
			mockRouterAddr, payOutToWrapperData, NativeTokenAddress, wrapperAddr,
		)
		require.NoError(t, err)

		msg := ethereum.CallMsg{From: env.from, To: &wrapperAddr, Data: selfWrapCalldata}

		result, err := env.client.CallContract(ctx, msg, nil)
		require.NoError(t, err)

		returnAmount, gasUsed, err := DecodeWrapOutput(result)
		require.NoError(t, err)
		assert.Equal(t, 0, amount.Cmp(returnAmount))
		assert.Positive(t, gasUsed)
	})
}
