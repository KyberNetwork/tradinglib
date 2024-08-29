package limitorder

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
)

func TestEncodeTakerTraits(t *testing.T) {
	extension := Extension{
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

	takerTraits := NewDefaultTakerTraits()
	takerTraits.SetExtension(extension).SetAmountMode(AmountModeMaker).SetAmountThreshold(big.NewInt(1))

	encodedTakerTraits, args := takerTraits.Encode()
	assert.Equal(t, common.HexToHash("0x8000002900000000000000000000000000000000000000000000000000000001"), encodedTakerTraits)
	assert.Equal(t, hexutil.Encode(extension.Encode()), hexutil.Encode(args))
}
