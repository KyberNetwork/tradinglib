package nativev2

import (
	"bytes"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const (
	methodTradeRFQT = "tradeRFQT"
)

var nativeABI abi.ABI

//nolint:gochecknoinits
func init() {
	builder := []struct {
		ABI  *abi.ABI
		data []byte
	}{
		{&nativeABI, nativeJSON},
	}

	for _, b := range builder {
		var err error
		*b.ABI, err = abi.JSON(bytes.NewReader(b.data))
		if err != nil {
			panic(err)
		}
	}
}
