package encode

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/math"
)

func EncodeInt64ToBytes(n int64, size int) []byte {
	return math.PaddedBigBytes(big.NewInt(n), size)
}
