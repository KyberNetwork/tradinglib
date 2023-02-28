package types

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type APIError struct {
	Code    int64
	Message string
}

func NewAPIError(code int64, message string) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
	}
}

func (e *APIError) Error() string {
	return fmt.Sprintf("code %d, message: %s", e.Code, e.Message)
}

// BigInt is helper use to marshal,unmarhsal bigInt as json string (with double quote).
type BigInt big.Int

func (b BigInt) String() string {
	return b.Int().String()
}

func (b *BigInt) UnmarshalText(text []byte) error {
	return (*big.Int)(b).UnmarshalText(text)
}

func (b BigInt) MarshalText() ([]byte, error) {
	return b.Int().MarshalText()
}

func (b *BigInt) UnmarshalJSON(data []byte) error {
	data = unquote(data)
	return (*big.Int)(b).UnmarshalJSON(data)
}

func (b BigInt) MarshalJSON() ([]byte, error) {
	d := b.Int().String()
	return strconv.AppendQuote(nil, d), nil
}

func (b *BigInt) Int() *big.Int {
	return (*big.Int)(b)
}

func BigIntFromStd(b *big.Int) *BigInt {
	return (*BigInt)(b)
}

// Bytes is a helper to unmarshal hex encode string (without 0x prefix, ethereum common type require 0x).
type Bytes []byte

func (b Bytes) MarshalJSON() ([]byte, error) {
	return strconv.AppendQuote(nil, hexutil.Encode(b)), nil
}

func (b Bytes) Bytes() []byte {
	return b
}

func (b *Bytes) UnmarshalJSON(data []byte) error {
	data = unquote(data)
	if len(data) == 0 {
		return nil
	}
	dd, err := hex.DecodeString(string(data))
	*b = dd
	return err
}

func unquote(data []byte) []byte {
	if len(data) < 2 { // nolint: gomnd
		return data
	}
	if (data[0] == '\'' && data[len(data)-1] == '\'') ||
		data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}
	return data
}

func BytesFromRaw(b []byte) Bytes {
	return b
}

// Duration is a helper type to unmarshal json duration string.
type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(time.Duration(d).String())), nil
}

func (d *Duration) UnmarshalJSON(bytes []byte) error {
	bytes = unquote(bytes)
	v, err := time.ParseDuration(string(bytes))
	if err != nil {
		return err
	}
	*d = Duration(v)
	return nil
}

func (d Duration) Unwrap() time.Duration {
	return time.Duration(d)
}
