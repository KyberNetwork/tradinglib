package httpclient_test

import (
	"testing"
	"time"

	"github.com/KyberNetwork/tradinglib/pkg/httpclient"
	"github.com/shopspring/decimal"
)

func strPtr(s string) *string {
	return &s
}

func TestClient(t *testing.T) {
	opt := struct {
		unexported int
		Name       *string         `json:"name"`
		Age        int64           `json:"age"`
		Balance    decimal.Decimal `json:"balance"`
		Rate       float64         `json:"rate,omitempty"`
		Delta32    float32         `json:"delta32"`
		Delta      *float64        `json:"delta,omitempty"`
		BoolValue  bool
		IntValue   int
		Int8       int8
		Int16      int16
		Int32      int32
		Int64      int64

		UIntValue uint
		UInt8     uint8
		UInt16    uint16
		UInt32    uint32
		UInt64    uint64

		Created time.Time `json:"created,unixMilli"` // nolint: staticcheck
	}{
		Name:    strPtr("booss"),
		Age:     120,
		Balance: decimal.NewFromInt(1000000),

		Created: time.Now(),
	}

	query := httpclient.NewQuery().
		SetString("param1", "value1").
		Float("price", 100.4512).
		Uint64("nonce", 1001).
		Int64("ts", 12345).
		Unix("unixKey", time.Now()).
		UnixMillis("millis", time.Now()).
		Bool("saveGas", true).
		Bool("skipEstimate", false).
		Struct(opt).
		String()
	t.Log("param:", query)
}
