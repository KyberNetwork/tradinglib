package mev

import (
	"fmt"
)

// nolint: gochecknoglobals
var (
	ErrMethodNotSupport  = fmt.Errorf("method not support")
	ErrMevShareClientNil = fmt.Errorf("mev share client is nil")
	ErrInvalidLenTx      = fmt.Errorf("only one tx is allowed")
	ErrMissingPrivKey    = fmt.Errorf("missing private key")
)
