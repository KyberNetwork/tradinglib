package mev_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/mev"
	"github.com/stretchr/testify/require"
)

// captured is one request as it left the client: the only view that proves what a builder
// actually receives, since every field here is optional and a dropped one is silent.
type captured struct {
	method string
	params map[string]any
}

func sendOne(
	t *testing.T, senderType mev.BundleSenderType, req mev.SendBundleV2Request,
	opts ...mev.NewBundleSendleClientOption,
) *captured {
	t.Helper()
	got := &captured{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)

		var parsed struct {
			Method string           `json:"method"`
			Params []map[string]any `json:"params"`
		}
		require.NoError(t, json.Unmarshal(body, &parsed))
		got.method = parsed.Method
		if len(parsed.Params) > 0 {
			got.params = parsed.Params[0]
		}
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0xdead"}`))
	}))
	t.Cleanup(srv.Close)

	c, err := mev.NewClient(srv.Client(), srv.URL, nil, senderType, false, opts...)
	require.NoError(t, err)
	_, err = c.SendBundleV2(context.Background(), req)
	require.NoError(t, err)

	return got
}

// TestSendBundleV2_HonoursTheBuilderNetRefundOption: the option is a property of the client, but
// only sendBundle applied it, so a caller on SendBundleV2 configured a refund recipient and
// silently never sent one — forgoing the refunds it had opted into.
func TestSendBundleV2_HonoursTheBuilderNetRefundOption(t *testing.T) {
	t.Parallel()
	blockNumber := uint64(1)
	addr := "0x0000000000000000000000000000000000001234"

	got := sendOne(t, mev.BundleSenderTypeBeaver, mev.SendBundleV2Request{BlockNumber: &blockNumber},
		mev.WithBuilderNetRefundAddress(addr))

	require.Equal(t, addr, got.params["builderNetRefundAddress"])
	require.Equal(t, true, got.params["allowBuilderNetRefunds"])
}

// TestSendBundleV2_LeavesTheRefundUnsetWithoutTheOption: a client that never opted in must not
// claim refunds for an empty address.
func TestSendBundleV2_LeavesTheRefundUnsetWithoutTheOption(t *testing.T) {
	t.Parallel()
	blockNumber := uint64(1)
	got := sendOne(t, mev.BundleSenderTypeBeaver, mev.SendBundleV2Request{BlockNumber: &blockNumber})

	require.Empty(t, got.params["builderNetRefundAddress"])
	require.Equal(t, false, got.params["allowBuilderNetRefunds"])
}

// TestSendBundleV2_48ClubSchedulingFieldsAreCarried pins the two fields 48Club's eth_sendBundle
// documents and this library did not send. Absent is not the same as false to a builder that
// treats the key's presence as the opt-in.
func TestSendBundleV2_48ClubSchedulingFieldsAreCarried(t *testing.T) {
	t.Parallel()
	blockNumber, noMerge, positionFirst := uint64(1), true, true

	got := sendOne(t, mev.BundleSenderType48Club, mev.SendBundleV2Request{
		BlockNumber:   &blockNumber,
		NoMerge:       &noMerge,
		PositionFirst: &positionFirst,
	})

	require.Equal(t, mev.ETHSendBundleMethod, got.method)
	require.Equal(t, true, got.params["noMerge"])
	require.Equal(t, true, got.params["positionFirst"])
}

// TestSendBundleV2_48ClubFieldsAreOmittedWhenUnset: both are optional, and a caller that never
// asked for one must not have a default chosen for it here.
func TestSendBundleV2_48ClubFieldsAreOmittedWhenUnset(t *testing.T) {
	t.Parallel()
	blockNumber := uint64(1)
	got := sendOne(t, mev.BundleSenderType48Club, mev.SendBundleV2Request{BlockNumber: &blockNumber})

	require.NotContains(t, got.params, "noMerge")
	require.NotContains(t, got.params, "positionFirst")
}

// TestSendBundleV2_BlockRazorTakesNeitherFieldNorTheBuilderMethod: BlockRazor's two products are
// not interchangeable. This library targets the free bsc.blockrazor.xyz RPC, whose
// eth_sendMevBundle documents only txs, revertingTxHashes and maxBlockNumber; noMerge and
// positionFirst belong to the paid virginia.builder.blockrazor.io builder on eth_sendBundle. Both
// halves are pinned here because sending one product's fields under the other's method is the
// mix-up that produces bundles which never land and no error to see.
func TestSendBundleV2_BlockRazorTakesNeitherFieldNorTheBuilderMethod(t *testing.T) {
	t.Parallel()
	blockNumber, noMerge, positionFirst := uint64(1), true, true

	got := sendOne(t, mev.BundleSenderTypeBlockRazor, mev.SendBundleV2Request{
		BlockNumber:   &blockNumber,
		NoMerge:       &noMerge,
		PositionFirst: &positionFirst,
	})

	require.Equal(t, mev.ETHSendMevBundle, got.method)
	require.NotContains(t, got.params, "noMerge")
	require.NotContains(t, got.params, "positionFirst")
}

// TestSendBundleV2_48ClubFieldsStayOffOtherBuilders: an unknown parameter is a rejection risk on
// any builder that validates strictly, not only on BSC.
func TestSendBundleV2_48ClubFieldsStayOffOtherBuilders(t *testing.T) {
	t.Parallel()
	blockNumber, noMerge := uint64(1), true
	got := sendOne(t, mev.BundleSenderTypeFlashbot, mev.SendBundleV2Request{
		BlockNumber: &blockNumber,
		NoMerge:     &noMerge,
	})

	require.NotContains(t, got.params, "noMerge")
}

// captureMethod records the JSON-RPC method of whatever call send makes.
func captureMethod(
	t *testing.T, senderType mev.BundleSenderType, send func(*mev.Client) error,
) string {
	t.Helper()
	var method string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)

		var parsed struct {
			Method string `json:"method"`
		}
		require.NoError(t, json.Unmarshal(body, &parsed))
		method = parsed.Method
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0xdead"}`))
	}))
	t.Cleanup(srv.Close)

	c, err := mev.NewClient(srv.Client(), srv.URL, nil, senderType, false)
	require.NoError(t, err)
	require.NoError(t, send(c))

	return method
}

// TestSendBundle_BlockRazorTakesTheFreeRPCMethod is the regression this pair of entry points
// needed: SendBundleV2 picked eth_sendMevBundle for BlockRazor while SendBundle and SendBundleHex
// hardcoded eth_sendBundle. The free bsc.blockrazor.xyz RPC serves only the former, so a caller on
// the plain path reached a method that endpoint does not implement — the bundle never lands and the
// send itself looks fine.
func TestSendBundle_BlockRazorTakesTheFreeRPCMethod(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		send func(*mev.Client) error
	}{
		{"SendBundle", func(c *mev.Client) error {
			_, err := c.SendBundle(context.Background(), nil, 1)
			return err
		}},
		{"SendBundleHex", func(c *mev.Client) error {
			_, err := c.SendBundleHex(context.Background(), nil, 1, "0x01")
			return err
		}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, "eth_sendMevBundle",
				captureMethod(t, mev.BundleSenderTypeBlockRazor, tc.send),
				"BlockRazor's free RPC serves eth_sendMevBundle on every submit path")
			require.Equal(t, "eth_sendBundle",
				captureMethod(t, mev.BundleSenderTypeFlashbot, tc.send),
				"every other sender keeps eth_sendBundle")
		})
	}
}

// TestBlockRazor_NonSubmitPathsKeepTheirOwnMethod: the swap belongs to bundle SUBMISSION only.
// Simulation and cancellation pass a method of their own, and rewriting those would turn a
// simulate into a submit.
func TestBlockRazor_NonSubmitPathsKeepTheirOwnMethod(t *testing.T) {
	t.Parallel()

	require.Equal(t, "eth_callBundle", captureMethod(t, mev.BundleSenderTypeBlockRazor,
		func(c *mev.Client) error {
			_, err := c.SimulateBundle(context.Background(), 1)
			return err
		}), "a simulation must never be rewritten into a send")
}
