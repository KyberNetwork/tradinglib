package decode

import (
	"errors"
	"fmt"
)

var ErrInvalidDataLength = errors.New("invalid data length")

func ValidateDataLength(data []byte, size int) error {
	if len(data) < size {
		return fmt.Errorf("%w: expected minimum size %d, got %d", ErrInvalidDataLength, size, len(data))
	}
	return nil
}
