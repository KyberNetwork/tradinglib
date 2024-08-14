package limitorder_test

import (
	"math/big"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/limitorder"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestInteraction(t *testing.T) {
	t.Run("should encode/decode", func(t *testing.T) {
		interaction := limitorder.Interaction{
			Target: common.BigToAddress(big.NewInt(1337)),
			Data:   "0xdeadbeef",
		}

		encodedInteraction := interaction.Encode()
		decodedInteraction := limitorder.DecodeInteraction(encodedInteraction)

		assert.Equal(t, interaction, decodedInteraction)
	})
}
