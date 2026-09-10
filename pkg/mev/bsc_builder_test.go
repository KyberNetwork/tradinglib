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
