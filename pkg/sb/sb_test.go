package sb

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConcat(t *testing.T) {
	ss := Concat("one", "two")
	require.Equal(t, "onetwo", ss)
}
