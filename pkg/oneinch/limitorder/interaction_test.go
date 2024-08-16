package limitorder_test

import (
	"math/big"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/limitorder"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInteraction(t *testing.T) {
	t.Run("should encode/decode", func(t *testing.T) {
		data, err := hexutil.Decode("0xdeadbeef")
		require.NoError(t, err)
		interaction := limitorder.Interaction{
			Target: common.BigToAddress(big.NewInt(1337)),
			Data:   data,
		}

		encodedInteraction, err := hexutil.Decode(interaction.Encode())
		require.NoError(t, err)

		decodedInteraction := limitorder.DecodeInteraction(encodedInteraction)
		require.NoError(t, err)

		assert.Equal(t, interaction, decodedInteraction)
	})
}
