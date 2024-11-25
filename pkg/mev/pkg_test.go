package mev_test

import (
	"encoding/json"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/mev"
	"github.com/test-go/testify/require"
)

func TestUnmarshalSendBundleResponse1(t *testing.T) {
	raws := []string{
		"{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":\"nil\"}",
		"{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":{}}",
		"{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":{\"bundleHash\": \"0x0\"}}",
		"{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":\"0x1ae69e3198840607d9946c52b0564624ab421d0678402e8696d08f9e5bc93a01\"}",
	}

	for _, raw := range raws {
		var resp mev.SendBundleResponse
		require.NoError(t, json.Unmarshal([]byte(raw), &resp))
		t.Logf("%+v\n", resp)
	}
}
