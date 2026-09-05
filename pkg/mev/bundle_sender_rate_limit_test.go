package mev_test

import (
	"context"
	"math/big"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/mev"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func rateLimitTestTx(t *testing.T) *types.Transaction {
	t.Helper()
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	tx, err := types.SignNewTx(key, types.NewCancunSigner(big.NewInt(1)), &types.DynamicFeeTx{
		ChainID:   big.NewInt(1),
		Nonce:     1,
		Gas:       21000,
		GasFeeCap: big.NewInt(100),
		GasTipCap: big.NewInt(10),
		To:        &common.Address{},
	})
	require.NoError(t, err)

	return tx
}

// newRateLimitServer returns a server that answers every request with one fixed
// bundle hash, and the counter of the requests that reached it.
func newRateLimitServer(t *testing.T) (*httptest.Server, *atomic.Int64) {
	t.Helper()
	var hits atomic.Int64
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		hits.Add(1)
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":{"bundleHash":"0xabc"}}`))
	}))
	t.Cleanup(srv.Close)

	return srv, &hits
}

func TestSendBundleRateLimitRejectsOverBurst(t *testing.T) {
	srv, hits := newRateLimitServer(t)
	client, err := mev.NewClient(
		srv.Client(), srv.URL, nil, mev.BundleSenderTypeUltrasound, false,
		mev.WithSendBundleRateLimit(2, 2),
	)
	require.NoError(t, err)

	ctx := context.Background()
	tx := rateLimitTestTx(t)
	for range 2 {
		resp, err := client.SendBundle(ctx, nil, 100, tx)
		require.NoError(t, err)
		require.Equal(t, "0xabc", resp.Result.BundleHash)
	}

	_, err = client.SendBundle(ctx, nil, 100, tx)
	require.ErrorIs(t, err, mev.ErrSendBundleRateLimited)
	require.Equal(t, int64(2), hits.Load())
}

func TestSendBundleRateLimitIgnoresOtherMethods(t *testing.T) {
	srv, hits := newRateLimitServer(t)
	client, err := mev.NewClient(
		srv.Client(), srv.URL, nil, mev.BundleSenderTypeUltrasound, false,
		mev.WithSendBundleRateLimit(1, 1),
	)
	require.NoError(t, err)

	ctx := context.Background()
	tx := rateLimitTestTx(t)
	_, err = client.SendBundle(ctx, nil, 100, tx)
	require.NoError(t, err)
	_, err = client.SendBundle(ctx, nil, 100, tx)
	require.ErrorIs(t, err, mev.ErrSendBundleRateLimited)

	resp, err := client.SimulateBundle(ctx, 100, tx)
	require.NoError(t, err)
	require.Equal(t, "0xabc", resp.Result.BundleHash)
	require.Equal(t, int64(2), hits.Load())
}

func TestSendBundleNoRateLimitByDefault(t *testing.T) {
	srv, hits := newRateLimitServer(t)
	client, err := mev.NewClient(srv.Client(), srv.URL, nil, mev.BundleSenderTypeUltrasound, false)
	require.NoError(t, err)

	ctx := context.Background()
	tx := rateLimitTestTx(t)
	for range 5 {
		_, err := client.SendBundle(ctx, nil, 100, tx)
		require.NoError(t, err)
	}
	require.Equal(t, int64(5), hits.Load())
}

func TestSendBundleRateLimitIgnoresNonPositive(t *testing.T) {
	srv, hits := newRateLimitServer(t)
	client, err := mev.NewClient(
		srv.Client(), srv.URL, nil, mev.BundleSenderTypeUltrasound, false,
		mev.WithSendBundleRateLimit(0, 0),
	)
	require.NoError(t, err)

	ctx := context.Background()
	tx := rateLimitTestTx(t)
	for range 3 {
		_, err := client.SendBundle(ctx, nil, 100, tx)
		require.NoError(t, err)
	}
	require.Equal(t, int64(3), hits.Load())
}

func TestUltrasoundSendBundleRateLimitValues(t *testing.T) {
	require.Equal(t, 50, mev.UltrasoundSendBundleRPS)
	require.Equal(t, 50, mev.UltrasoundSendBundleBurst)
}
