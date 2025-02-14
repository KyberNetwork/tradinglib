package encode

import "errors"

var (
	ErrInvalidSwapLength              = errors.New("invalid swap length")
	ErrInvalidPoolsAndPoolTypesLength = errors.New("invalid pools and pool types lengths")
	ErrInvalidPoolsAndTokensLength    = errors.New("invalid pools and tokens lengths")
	ErrNotSupportedDex                = errors.New("not supported dex")
)
