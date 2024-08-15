package limitorder_test

import (
	"math/big"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/limitorder"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInteraction(t *testing.T) {
	t.Run("should encode/decode", func(t *testing.T) {
		interaction := limitorder.Interaction{
			Target: common.BigToAddress(big.NewInt(1337)),
			Data:   "0xdeadbeef",
		}

		encodedInteraction := interaction.Encode()
		decodedInteraction, err := limitorder.DecodeInteraction(encodedInteraction)
		require.NoError(t, err)

		assert.Equal(t, interaction, decodedInteraction)
	})
}
