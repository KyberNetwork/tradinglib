package mev_test

import (
	"context"
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/mev"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func servoTestTx(t *testing.T) *types.Transaction {
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

func TestServoSenderBidShape(t *testing.T) {
	var got map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "key-123", r.Header.Get(mev.XAPIKeyHeader))
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		require.NoError(t, json.Unmarshal(body, &got))
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":{"bundleHash":"0xabc"}}`))
	}))
	defer srv.Close()

	s, err := mev.NewServoSender(srv.Client(), srv.URL, "key-123")
	require.NoError(t, err)

	tx := servoTestTx(t)
	victims := []common.Hash{common.HexToHash("0x01"), common.HexToHash("0x02")}
	uuid := "da48afaa-96ee-0000-0000-00000e80b778"
	resp, err := s.SendBackrunBundle(context.Background(), &uuid, 100, 102, victims, nil, nil, tx)
	require.NoError(t, err)
	require.Equal(t, "0xabc", resp.Result.BundleHash)

	require.Equal(t, mev.ETHSendBundleMethod, got["method"])
	params, ok := got["params"].([]any)
	require.True(t, ok)
	require.Len(t, params, 1)
	p, ok := params[0].(map[string]any)
	require.True(t, ok)

	// blockNumber is the DEADLINE (max), never the target, and is always present.
	require.Equal(t, hexutil.EncodeUint64(102), p["blockNumber"])
	require.Equal(t, uuid, p["replacementUuid"])

	raw, err := tx.MarshalBinary()
	require.NoError(t, err)
	// victim hashes first, in the order received; our raw backrun last.
	require.Equal(t, []any{victims[0].Hex(), victims[1].Hex(), hexutil.Encode(raw)}, p["txs"])
}

func TestServoSenderRejectsIncompleteBids(t *testing.T) {
	s, err := mev.NewServoSender(http.DefaultClient, "http://servo.invalid", "key")
	require.NoError(t, err)
	tx := servoTestTx(t)

	_, err = s.SendBackrunBundle(context.Background(), nil, 1, 1, nil, nil, nil, tx)
	require.ErrorIs(t, err, mev.ErrMissingPendingTxs)

	_, err = s.SendBackrunBundle(context.Background(), nil, 1, 1, []common.Hash{{}}, nil, nil)
	require.ErrorIs(t, err, mev.ErrInvalidLenTx)

	// A bid with no deadline would be retried by SERVO for 5 blocks on stale pricing.
	_, err = s.SendBackrunBundle(context.Background(), nil, 0, 0, []common.Hash{{}}, nil, nil, tx)
	require.ErrorIs(t, err, mev.ErrMissingBlockNumber)
}

func TestServoSenderRequiresCredentials(t *testing.T) {
	_, err := mev.NewServoSender(http.DefaultClient, "", "key")
	require.ErrorIs(t, err, mev.ErrMissingEndpoint)
	_, err = mev.NewServoSender(http.DefaultClient, "http://servo.invalid", "")
	require.ErrorIs(t, err, mev.ErrMissingAPIKey)
}
