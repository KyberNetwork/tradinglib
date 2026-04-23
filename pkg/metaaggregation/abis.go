package metaaggregation

import (
	"bytes"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

var RouterV2ABI abi.ABI

//nolint:gochecknoinits
func init() {
	builder := []struct {
		ABI  *abi.ABI
		data []byte
	}{
		{&RouterV2ABI, metaAggregationRouterV2JSON},
	}

	for _, b := range builder {
		var err error
		*b.ABI, err = abi.JSON(bytes.NewReader(b.data))
		if err != nil {
			panic(err)
		}
	}
}
