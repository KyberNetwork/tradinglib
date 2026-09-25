package kipselipmm

import "errors"

var (
	ErrEmptyPriceLevels      = errors.New("empty price levels")
	ErrAmountInTooSmall      = errors.New("amount in is too small")
	ErrAmountOutTooSmall     = errors.New("amount out too small")
	ErrInsufficientLiquidity = errors.New("insufficient liquidity")
	ErrInvalidSwapInfo       = errors.New("invalid swap info")
	ErrTokenNotFound         = errors.New("token not found")
	ErrNoPathFound           = errors.New("no path found")
)
