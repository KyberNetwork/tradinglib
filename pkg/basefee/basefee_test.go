package basefee_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/basefee"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

func TestCalculateNextBaseFee(t *testing.T) {
	t.Skip()
	testCases := []struct {
		name        string
		rpcURL      string
		chainID     uint64
		blockNumber *big.Int
	}{
		{
			name:        "Ethereum Mainnet",
			chainID:     1,
			rpcURL:      "https://ethereum.kyberengineering.io",
			blockNumber: big.NewInt(23637674),
		},
		{
			name:        "BSC Mainnet",
			chainID:     56,
			rpcURL:      "https://bsc-dataseed.binance.org/",
			blockNumber: big.NewInt(65588931),
		},
		{
			name:        "Base Mainnet",
			chainID:     8453,
			rpcURL:      "https://base.kyberengineering.io",
			blockNumber: big.NewInt(37084592),
		},
		{
			name:        "Polygon Mainnet",
			chainID:     137,
			rpcURL:      "https://polygon.kyberengineering.io",
			blockNumber: big.NewInt(81240460),
		},
	}

	ctx := context.Background()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ethClient, err := ethclient.Dial(tc.rpcURL)
			assert.NoError(t, err, "Failed to connect to Ethereum client")
			head, err := ethClient.HeaderByNumber(ctx, tc.blockNumber)
			assert.NoError(t, err, "Failed to get block header", tc.rpcURL)
			nextBaseFee, err := basefee.CalcNextBaseFee(tc.chainID, head)
			assert.NoError(t, err)
			nextHead, err := ethClient.HeaderByNumber(ctx, new(big.Int).Add(head.Number, big.NewInt(1)))
			assert.NoError(t, err, "Failed to get next block header")
			assert.Equal(t, nextBaseFee, nextHead.BaseFee)
			t.Log("Next base fee:", nextBaseFee.String())
		})
	}
}
