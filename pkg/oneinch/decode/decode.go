package decode

import (
	"errors"
	"fmt"
	"math/big"
)

var ErrInvalidDataLength = errors.New("invalid data length")

func ValidateDataLength(data []byte, size int) error {
	if len(data) < size {
		return fmt.Errorf("%w: expected minimum size %d, got %d", ErrInvalidDataLength, size, len(data))
	}
	return nil
}

func Next(data []byte, size int) ([]byte, []byte, error) {
	if err := ValidateDataLength(data, size); err != nil {
		return nil, nil, err
	}
	return data[:size], data[size:], nil
}

func NextInt64(data []byte, size int) (int64, []byte, error) {
	d, remainingData, err := Next(data, size)
	if err != nil {
		return 0, nil, err
	}

	return new(big.Int).SetBytes(d).Int64(), remainingData, nil
}
