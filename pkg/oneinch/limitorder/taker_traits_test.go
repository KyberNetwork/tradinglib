package limitorder

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecodeArgs(t *testing.T) {
	// Test with receiver, extension, and interaction
	t.Run("full decode", func(t *testing.T) {
		// Create test data
		receiver := common.HexToAddress("0x1234567890123456789012345678901234567890")
		extension := Extension{
			MakerAssetSuffix: []byte("maker"),
			TakerAssetSuffix: []byte("taker"),
		}
		interaction := Interaction{
			Target: common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"),
			Data:   []byte("test data"),
		}

		// Create TakerTraits and encode
		takerTraits := NewTakerTraits(big.NewInt(0), &receiver, &extension, &interaction)
		flags, args := takerTraits.Encode()

		// Decode back
		decodedReceiver, decodedExtension, decodedInteraction, err := DecodeArgs(flags, args)
		require.NoError(t, err)

		// Verify receiver
		assert.NotNil(t, decodedReceiver)
		assert.Equal(t, receiver, *decodedReceiver)

		// Verify extension
		assert.NotNil(t, decodedExtension)
		assert.Equal(t, extension.MakerAssetSuffix, decodedExtension.MakerAssetSuffix)
		assert.Equal(t, extension.TakerAssetSuffix, decodedExtension.TakerAssetSuffix)

		// Verify interaction
		assert.NotNil(t, decodedInteraction)
		assert.Equal(t, interaction.Target, decodedInteraction.Target)
		assert.Equal(t, interaction.Data, decodedInteraction.Data)
	})

	// Test with only receiver
	t.Run("receiver only", func(t *testing.T) {
		receiver := common.HexToAddress("0x1234567890123456789012345678901234567890")

		takerTraits := NewTakerTraits(big.NewInt(0), &receiver, nil, nil)
		flags, args := takerTraits.Encode()

		decodedReceiver, decodedExtension, decodedInteraction, err := DecodeArgs(flags, args)
		require.NoError(t, err)

		assert.NotNil(t, decodedReceiver)
		assert.Equal(t, receiver, *decodedReceiver)
		assert.Nil(t, decodedExtension)
		assert.Nil(t, decodedInteraction)
	})

	// Test with empty args
	t.Run("empty args", func(t *testing.T) {
		flags := big.NewInt(0)
		args := []byte{}

		decodedReceiver, decodedExtension, decodedInteraction, err := DecodeArgs(flags, args)
		require.NoError(t, err)

		assert.Nil(t, decodedReceiver)
		assert.Nil(t, decodedExtension)
		assert.Nil(t, decodedInteraction)
	})

	// Test with only extension
	t.Run("extension only", func(t *testing.T) {
		extension := Extension{
			MakerAssetSuffix: []byte("test"),
		}

		takerTraits := NewTakerTraits(big.NewInt(0), nil, &extension, nil)
		flags, args := takerTraits.Encode()

		decodedReceiver, decodedExtension, decodedInteraction, err := DecodeArgs(flags, args)
		require.NoError(t, err)

		assert.Nil(t, decodedReceiver)
		assert.NotNil(t, decodedExtension)
		assert.Equal(t, extension.MakerAssetSuffix, decodedExtension.MakerAssetSuffix)
		assert.Nil(t, decodedInteraction)
	})

	// Test with only interaction
	t.Run("interaction only", func(t *testing.T) {
		interaction := Interaction{
			Target: common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"),
			Data:   []byte("test data"),
		}

		takerTraits := NewTakerTraits(big.NewInt(0), nil, nil, &interaction)
		flags, args := takerTraits.Encode()

		decodedReceiver, decodedExtension, decodedInteraction, err := DecodeArgs(flags, args)
		require.NoError(t, err)

		assert.Nil(t, decodedReceiver)
		assert.Nil(t, decodedExtension)
		assert.NotNil(t, decodedInteraction)
		assert.Equal(t, interaction.Target, decodedInteraction.Target)
		assert.Equal(t, interaction.Data, decodedInteraction.Data)
	})
}
