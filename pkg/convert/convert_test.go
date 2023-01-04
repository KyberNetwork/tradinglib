package convert_test

import (
	"math/big"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/convert"
	"github.com/stretchr/testify/require"
)

func TestAddBPS(t *testing.T) {
	amount := big.NewInt(10000)
	add50bps := convert.AddBPS(amount, 50)
	require.Equal(t, big.NewInt(10050), add50bps)

	addn50bps := convert.AddBPS(amount, -50)
	require.Equal(t, big.NewInt(9950), addn50bps)
}
