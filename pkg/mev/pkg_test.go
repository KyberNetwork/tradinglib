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

func TestCleanBundleHash(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "hash with escaped quotes",
			input:    "\"0xf72e9e8afd22af2904857e03575eb6f125cabc0d18fe7fb89ee1f8c6861687ae\"",
			expected: "0xf72e9e8afd22af2904857e03575eb6f125cabc0d18fe7fb89ee1f8c6861687ae",
		},
		{
			name:     "hash with regular quotes",
			input:    "\"0xabc123\"",
			expected: "0xabc123",
		},
		{
			name:     "hash without quotes",
			input:    "0xdef456",
			expected: "0xdef456",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "only quotes",
			input:    "\"\"",
			expected: "",
		},
		{
			name:     "multiple escaped quotes",
			input:    "\"\\\"0xabc123\\\"\"",
			expected: "0xabc123",
		},
		{
			name:     "mixed quotes",
			input:    "\\\"0xabc123\"",
			expected: "0xabc123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mev.CleanBundleHash(tt.input)
			if got != tt.expected {
				t.Errorf("cleanBundleHash() = %v, want %v", got, tt.expected)
			}
		})
	}
}
