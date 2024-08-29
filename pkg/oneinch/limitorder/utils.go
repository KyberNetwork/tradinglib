package limitorder

import "math/big"

func newBitMask(start uint, end uint) *big.Int {
	mask := big.NewInt(1)
	mask.Lsh(mask, end)
	mask.Sub(mask, big.NewInt(1))
	if start == 0 {
		return mask
	}

	notMask := newBitMask(0, start)
	notMask.Not(notMask)
	mask.And(mask, notMask)

	return mask
}

func setMask(n *big.Int, mask *big.Int, value *big.Int) {
	// Clear bits in range.
	n.And(n, new(big.Int).Not(mask))

	// Shift value to correct position and ensure value fits in mask.
	value = new(big.Int).Lsh(value, mask.TrailingZeroBits())
	value.And(value, mask)

	// Set the bits in range.
	n.Or(n, value)
}
