package sb_test

import (
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/sb"
	"github.com/stretchr/testify/require"
)

func TestConcat(t *testing.T) {
	ss := sb.Concat("one", "two")
	require.Equal(t, "onetwo", ss)
}
