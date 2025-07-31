//nolint:testpackage
package unix

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func mustFromString(s string) *big.Int {
	v, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic("invalid big integer number: " + s)
	}
	return v
}

func TestPackNativeTokenOutputIndex(t *testing.T) {
	tests := []struct {
		name           string
		tokenOutputIdx []OrderTokenIndex
		expected       *big.Int
		expectedErr    error
	}{
		{
			name: "normal case",
			tokenOutputIdx: []OrderTokenIndex{
				{
					OrderIndex: 0,
					TokenIndex: 0,
				},
				{
					OrderIndex: 0,
					TokenIndex: 1,
				},
				{
					OrderIndex: 1,
					TokenIndex: 2,
				},
			},
			expected: mustFromString("1356938545749799165119972480570561420155507632800475359837393562592773963776"),
		},
		{
			name: "normal case for tokenOutputIdx with 0",
			tokenOutputIdx: []OrderTokenIndex{
				{
					OrderIndex: 0,
					TokenIndex: 0,
				},
			},
			expected: mustFromString("452312848583266388373324160190187140051835877600158453279131187530910662656"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PackTokenOutputIndex(tt.tokenOutputIdx)
			if tt.expectedErr != nil {
				require.ErrorIs(t, tt.expectedErr, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, got)
			}
		})
	}
}

func TestUnpackNativeTokenOutputIndex(t *testing.T) {
	tests := []struct {
		name     string
		packed   *big.Int
		expected []OrderTokenIndex
		hasError bool
	}{
		{
			name:   "normal case",
			packed: mustFromString("1356938545749799165119972480570561420155507632800475359837393562592773963776"),
			expected: []OrderTokenIndex{
				{
					OrderIndex: 0,
					TokenIndex: 0,
				},
				{
					OrderIndex: 0,
					TokenIndex: 1,
				},
				{
					OrderIndex: 1,
					TokenIndex: 2,
				},
			},
		},
		{
			name:   "normal case for tokenOutputIdx with 0",
			packed: mustFromString("452312848583266388373324160190187140051835877600158453279131187530910662656"),
			expected: []OrderTokenIndex{
				{
					OrderIndex: 0,
					TokenIndex: 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unpackTokenOutputIndex(tt.packed)
			if tt.hasError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, got)
			}
		})
	}
}
