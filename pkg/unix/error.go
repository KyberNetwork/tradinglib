package blockchain

import "errors"

var (
	ErrTokenIndexInvalid  = errors.New("token index invalid")
	ErrOrderIndexInvalid  = errors.New("order index invalid")
	ErrOrderLimitExceeded = errors.New("order limit exceeded")
)
