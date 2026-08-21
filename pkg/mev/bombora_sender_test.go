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
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func newBundleServer(tb testing.TB) (*httptest.Server, <-chan []byte) {
	tb.Helper()

	bodyCh := make(chan []byte, 1)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			tb.Errorf("read request body: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		bodyCh <- body

		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":{"bundleHash":"0xbombora"}}`)); err != nil {
			tb.Errorf("write response body: %v", err)
		}
	}))

	return server, bodyCh
}

func newBomboraRequest() mev.SendBundleV2Request {
	droppingTxs := []string{"0xdropping"}
	replacementSeqNumber := uint64(7)
	refundPercent := uint64(50)
	refundTxHashes := []string{"0xrefund"}
	uuid := "11111111-1111-1111-1111-111111111111"

	return mev.SendBundleV2Request{
		DroppingTxs:          &droppingTxs,
		ReplacementSeqNumber: &replacementSeqNumber,
		RefundPercent:        &refundPercent,
		RefundRecipient:      "0x000000000000000000000000000000000000dEaD",
		RefundTxHashes:       &refundTxHashes,
		UUID:                 &uuid,
	}
}

func newSignedLegacyTx(tb testing.TB) *types.Transaction {
	tb.Helper()

	key, err := crypto.GenerateKey()
	require.NoError(tb, err)

	to := common.HexToAddress("0x0000000000000000000000000000000000000001")
	tx, err := types.SignNewTx(key, types.NewEIP155Signer(big.NewInt(1)), &types.LegacyTx{
		Nonce:    0,
		To:       &to,
		Value:    big.NewInt(1),
		Gas:      21000,
		GasPrice: big.NewInt(1),
	})
	require.NoError(tb, err)

	return tx
}

func sendBundleV2Request(tb testing.TB, senderType mev.BundleSenderType) []byte {
	tb.Helper()

	server, bodyCh := newBundleServer(tb)
	tb.Cleanup(server.Close)

	client, err := mev.NewClient(server.Client(), server.URL, nil, senderType, false)
	require.NoError(tb, err)

	_, err = client.SendBundleV2(context.Background(), newBomboraRequest(), newSignedLegacyTx(tb))
	require.NoError(tb, err)

	return <-bodyCh
}

func decodeBundleRequest(tb testing.TB, body []byte) (method string, params map[string]json.RawMessage) {
	tb.Helper()

	var got struct {
		Method string                       `json:"method"`
		Params []map[string]json.RawMessage `json:"params"`
	}
	require.NoError(tb, json.Unmarshal(body, &got))
	require.Len(tb, got.Params, 1)

	return got.Method, got.Params[0]
}

func TestBundleSenderTypeBomboraEnumer(t *testing.T) {
	require.Equal(t, "BundleSenderTypeBombora", mev.BundleSenderTypeBombora.String())
	require.True(t, mev.BundleSenderTypeBombora.IsABundleSenderType())

	v, err := mev.BundleSenderTypeString("BundleSenderTypeBombora")
	require.NoError(t, err)
	require.Equal(t, mev.BundleSenderTypeBombora, v)
}

func TestSendBundleV2BomboraFields(t *testing.T) {
	method, params := decodeBundleRequest(t, sendBundleV2Request(t, mev.BundleSenderTypeBombora))

	require.Equal(t, "eth_sendBundle", method)
	for _, key := range []string{
		"droppingTxHashes",
		"replacementSeqNumber",
		"refundPercent",
		"refundRecipient",
		"refundTxHashes",
		"replacementUuid",
	} {
		require.Contains(t, params, key)
	}
	require.NotContains(t, params, "ReplacementUuid")

	var replacementSeqNumber uint64
	require.NoError(t, json.Unmarshal(params["replacementSeqNumber"], &replacementSeqNumber))
	require.Equal(t, uint64(7), replacementSeqNumber)

	var refundPercent uint64
	require.NoError(t, json.Unmarshal(params["refundPercent"], &refundPercent))
	require.Equal(t, uint64(50), refundPercent)

	var refundRecipient string
	require.NoError(t, json.Unmarshal(params["refundRecipient"], &refundRecipient))
	require.Equal(t, "0x000000000000000000000000000000000000dEaD", refundRecipient)
}

func TestSendBundleV2NonBomboraOmitsBomboraFields(t *testing.T) {
	_, params := decodeBundleRequest(t, sendBundleV2Request(t, mev.BundleSenderTypeBeaver))

	for _, key := range []string{
		"droppingTxHashes",
		"replacementSeqNumber",
		"refundPercent",
		"refundRecipient",
		"refundTxHashes",
		"replacementUuid",
	} {
		require.NotContains(t, params, key)
	}
}

func TestSetBomboraFieldsValidation(t *testing.T) {
	refundPercent100 := uint64(100)
	refundPercent99 := uint64(99)
	twoRefundTxHashes := []string{"0xrefund1", "0xrefund2"}
	oneRefundTxHash := []string{"0xrefund"}

	tests := []struct {
		name    string
		req     mev.SendBundleV2Request
		wantErr error
	}{
		{
			name:    "refund percent above maximum",
			req:     mev.SendBundleV2Request{RefundPercent: &refundPercent100},
			wantErr: mev.ErrInvalidRefundPercent,
		},
		{
			name:    "more than one refund transaction hash",
			req:     mev.SendBundleV2Request{RefundTxHashes: &twoRefundTxHashes},
			wantErr: mev.ErrInvalidLenRefundTxHashes,
		},
		{
			name: "valid refund fields",
			req: mev.SendBundleV2Request{
				RefundPercent:  &refundPercent99,
				RefundTxHashes: &oneRefundTxHash,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := new(mev.SendBundleParams).SetBomboraFields(test.req).Err()
			if test.wantErr == nil {
				require.NoError(t, err)
				return
			}

			require.ErrorIs(t, err, test.wantErr)
		})
	}
}
