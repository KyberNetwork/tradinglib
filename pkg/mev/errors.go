package mev

import (
	"fmt"
)

// nolint: gochecknoglobals
var (
	ErrMethodNotSupport         = fmt.Errorf("method not support")
	ErrMevShareClientNil        = fmt.Errorf("mev share client is nil")
	ErrInvalidLenTx             = fmt.Errorf("only one tx is allowed")
	ErrMissingPrivKey           = fmt.Errorf("missing private key")
	ErrInvalidMaxBlock          = fmt.Errorf("max block number must be greater than block number")
	ErrInvalidLenPendingTx      = fmt.Errorf("only one pending tx is allowed")
	ErrInvalidRefundPercent     = fmt.Errorf("refund percent must be between 0 and 99")
	ErrInvalidLenRefundTxHashes = fmt.Errorf("only one refund tx hash is allowed")
)
