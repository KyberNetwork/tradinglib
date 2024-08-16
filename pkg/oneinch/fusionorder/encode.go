package fusionorder

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/math"
)

func encodeInt64ToBytes(n int64, size int) []byte {
	return math.PaddedBigBytes(big.NewInt(n), size)
}
