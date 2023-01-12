package stack_test

import (
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/ds/stack"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	testLen := 10

	s := stack.New[int](testLen)
	require.NotNil(t, s)

	assert.True(t, s.Empty())
	for i := 0; i < testLen; i++ {
		s.Push(i + 1)
		assert.False(t, s.Empty())
		assert.Equal(t, i+1, s.Len())
		assert.Equal(t, i+1, s.Peek())
	}

	for i := 0; i < testLen; i++ {
		assert.Equal(t, testLen-i, s.Len())
		elem := s.Pop()
		assert.Equal(t, testLen-i, elem)
	}

	assert.True(t, s.Empty())
}
