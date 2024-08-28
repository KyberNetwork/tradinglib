package limitorder_test

import (
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/limitorder"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtension(t *testing.T) {
	t.Run("should encode/decode", func(t *testing.T) {
		extension := limitorder.Extension{
			MakerAssetSuffix: []byte{0x01},
			TakerAssetSuffix: []byte{0x02},
			MakingAmountData: []byte{0x03},
			TakingAmountData: []byte{0x04},
			Predicate:        []byte{0x05},
			MakerPermit:      []byte{0x06},
			PreInteraction:   []byte{0x07},
			PostInteraction:  []byte{0x08},
			CustomData:       []byte{0xff},
		}

		encodedExtension := extension.Encode()

		decodedExtension, err := limitorder.DecodeExtension(hexutil.Encode(encodedExtension))
		require.NoError(t, err)

		require.Equal(t, extension, decodedExtension)
	})

	t.Run("decode empty", func(t *testing.T) {
		encodedExtension := "0x"
		e, err := limitorder.DecodeExtension(encodedExtension)
		require.NoError(t, err)

		assert.True(t, e.IsEmpty())
	})

	t.Run("decode", func(t *testing.T) {
		// nolint: lll
		encodedExtension := "0x000000e5000000540000005400000054000000540000002a0000000000000000fb2809a5314473e1165f6b58018e20ed8f07b84000f1b8000005e566bb30120000b401def800f1b800b4fb2809a5314473e1165f6b58018e20ed8f07b84000f1b8000005e566bb30120000b401def800f1b800b4fb2809a5314473e1165f6b58018e20ed8f07b84066bb2ffab09498030ae3416b66dc0000db05a6a504f04d92e79d0000339fb574bdc56763f9950000d18bd45f0b94f54a968f0000d61b892b2ad6249011850000f7f4f96b98e102b56b0400000000000000000000000000006de5e0e428ac771d77b5000006455390207c1d485be90000ade19567bb538035ed36000050"
		e, err := limitorder.DecodeExtension(encodedExtension)
		require.NoError(t, err)

		t.Logf("Extension: %+v", e)

		assert.False(t, e.IsEmpty())
		assert.NotEmpty(t, e.MakingAmountData)
		assert.NotEmpty(t, e.TakingAmountData)
		assert.NotEmpty(t, e.PostInteraction)
	})
}
