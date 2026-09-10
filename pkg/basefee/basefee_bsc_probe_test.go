package basefee_test

import (
	"context"
	"math/big"
	"os"
	"strconv"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/basefee"
	bscChainEIP1559 "github.com/KyberNetwork/tradinglib/pkg/basefee/op-geth/bsc-eip1559"
	bscChainParams "github.com/KyberNetwork/tradinglib/pkg/basefee/op-geth/bsc-params"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

// TestBSCNextBaseFee_AgainstTheChain checks CalcNextBaseFee for BSC against what the chain actually
// produced, and shows which ChainConfig the BSC branch should be passing.
//
// It runs only with BASEFEE_BSC_RPC set, rather than the unconditional t.Skip on the table test
// beside it: a skip that cannot be turned on is a test that never ran, which is why the BSC branch
// went unverified.
//
//	BASEFEE_BSC_RPC=https://bsc-rpc.publicnode.com go test ./pkg/basefee/ -run BSCNextBaseFee -v
func TestBSCNextBaseFee_AgainstTheChain(t *testing.T) {
	rpcURL := os.Getenv("BASEFEE_BSC_RPC")
	if rpcURL == "" {
		t.Skip("set BASEFEE_BSC_RPC to run this against a live node")
	}

	ctx := context.Background()
	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)
	t.Cleanup(client.Close)

	head, err := client.HeaderByNumber(ctx, nil)
	require.NoError(t, err)

	// Step back from the head so the next block is certainly available, and sample a few parents so
	// the gasUsed-vs-target branches are not all the same one.
	for _, back := range []int64{8, 24, 72} {
		parentNum := new(big.Int).Sub(head.Number, big.NewInt(back))
		parent, err := client.HeaderByNumber(ctx, parentNum)
		require.NoError(t, err)
		next, err := client.HeaderByNumber(ctx, new(big.Int).Add(parentNum, big.NewInt(1)))
		require.NoError(t, err)

		actual := next.BaseFee
		viaAPI, err := basefee.CalcNextBaseFee(basefee.BscChainID, parent)
		require.NoError(t, err)
		withEthCfg := bscChainEIP1559.CalcBaseFee(bscChainParams.MainnetChainConfig, parent)
		withBSCCfg := bscChainEIP1559.CalcBaseFee(bscChainParams.BSCChainConfig, parent)

		gasTarget := parent.GasLimit / bscChainParams.MainnetChainConfig.ElasticityMultiplier()
		t.Logf("block %s: parentBaseFee=%v gasUsed=%d target=%d | actual next=%v | CalcNextBaseFee=%v | ethCfg=%v | bscCfg=%v",
			parentNum, parent.BaseFee, parent.GasUsed, gasTarget, actual, viaAPI, withEthCfg, withBSCCfg)

		require.Zerof(t, withBSCCfg.Cmp(actual),
			"BSCChainConfig must reproduce the chain: got %v want %v", withBSCCfg, actual)
	}
}

// TestNextBaseFee_AgainstTheChain is the same check for any chain CalcNextBaseFee claims to
// support: does it reproduce the base fee the chain actually produced?
//
//	BASEFEE_RPC=https://mainnet.base.org BASEFEE_CHAIN_ID=8453 \
//	  go test ./pkg/basefee/ -run TestNextBaseFee_AgainstTheChain -v
func TestNextBaseFee_AgainstTheChain(t *testing.T) {
	rpcURL, chainIDStr := os.Getenv("BASEFEE_RPC"), os.Getenv("BASEFEE_CHAIN_ID")
	if rpcURL == "" || chainIDStr == "" {
		t.Skip("set BASEFEE_RPC and BASEFEE_CHAIN_ID to run this against a live node")
	}
	chainID, err := strconv.ParseUint(chainIDStr, 10, 64)
	require.NoError(t, err)

	ctx := context.Background()
	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)
	t.Cleanup(client.Close)

	head, err := client.HeaderByNumber(ctx, nil)
	require.NoError(t, err)

	for _, back := range []int64{8, 24, 72} {
		parentNum := new(big.Int).Sub(head.Number, big.NewInt(back))
		parent, err := client.HeaderByNumber(ctx, parentNum)
		require.NoError(t, err)
		next, err := client.HeaderByNumber(ctx, new(big.Int).Add(parentNum, big.NewInt(1)))
		require.NoError(t, err)

		got, err := basefee.CalcNextBaseFee(chainID, parent)
		require.NoError(t, err)
		t.Logf("chain %d block %s: parentBaseFee=%v gasUsed=%d gasLimit=%d | actual next=%v | calculated=%v",
			chainID, parentNum, parent.BaseFee, parent.GasUsed, parent.GasLimit, next.BaseFee, got)
		require.Zerof(t, got.Cmp(next.BaseFee),
			"CalcNextBaseFee must reproduce the chain: got %v want %v", got, next.BaseFee)
	}
}
