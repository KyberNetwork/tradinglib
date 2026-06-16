package types_test

import (
	"errors"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/types"
	"github.com/stretchr/testify/require"
)

func TestParseBloxrouteBackrunConfig(t *testing.T) {
	const (
		targetAddress = "0x435D8cB64ba4B68f6A28405D3C19E7C169B01917"
		blxrAddress   = "0xF1Ce037Cc1d02c046e800C3265D2Efa91940864E"
	)

	cfg, err := types.ParseBloxrouteBackrunConfig([]string{"50", targetAddress, "15", blxrAddress, "20"})
	require.NoError(t, err)
	require.Equal(t, 50.0, cfg.ContractSplit)
	require.Equal(t, targetAddress, cfg.TargetRewardAddress)
	require.Equal(t, 15.0, cfg.TargetSplit)
	require.Equal(t, blxrAddress, cfg.BlxrRewardAddress)
	require.Equal(t, 20.0, cfg.SearcherSplit)

	blxr, err := cfg.BlxrSplit()
	require.NoError(t, err)
	require.Equal(t, 15.0, blxr) // 100 - 50 - 15 - 20
}

func TestParseBloxrouteBackrunConfigErrors(t *testing.T) {
	const (
		targetAddress = "0x435D8cB64ba4B68f6A28405D3C19E7C169B01917"
		blxrAddress   = "0xF1Ce037Cc1d02c046e800C3265D2Efa91940864E"
	)

	tests := []struct {
		name  string
		parts []string
	}{
		{"wrong length", []string{"50", targetAddress, "15", blxrAddress}},
		{"non-numeric split", []string{"x", targetAddress, "15", blxrAddress, "20"}},
		{"sum exceeds 100", []string{"60", targetAddress, "30", blxrAddress, "20"}},
		{"zero target address", []string{"50", "0x0000000000000000000000000000000000000000", "15", blxrAddress, "20"}},
		{"invalid blxr address", []string{"50", targetAddress, "15", "not-an-address", "20"}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := types.ParseBloxrouteBackrunConfig(tc.parts)
			require.Error(t, err)
			require.True(t, errors.Is(err, types.ErrInvalidBloxrouteBackrunConfig))
		})
	}
}
