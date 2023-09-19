package queue_test

import (
	"slices"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/ds/queue"
	"github.com/stretchr/testify/assert"
)

var vals = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //nolint: gochecknoglobals

func TestPushBack(t *testing.T) {
	q := queue.New[int]()
	for _, v := range vals {
		q.PushBack(v)
	}
	qVals := q.List()
	assert.Equal(t, vals, qVals)
}

func TestPushFront(t *testing.T) {
	q := queue.New[int]()
	for _, v := range vals {
		q.PushFront(v)
	}
	qVals := q.List()

	reverse := make([]int, len(vals))
	copy(reverse, vals)
	slices.Reverse(reverse)
	assert.Equal(t, reverse, qVals)
}

func TestPopBack(t *testing.T) {
	testVals := make([]int, len(vals))
	copy(testVals, vals)
	q := queue.New[int]()
	for _, v := range vals {
		q.PushBack(v)
	}

	for {
		v, ok := q.PopBack()
		if !ok {
			break
		}
		testVal := testVals[len(testVals)-1]
		assert.Equal(t, testVal, v)

		testVals = testVals[:len(testVals)-1]
		assert.Equal(t, len(testVals), int((q.Size())))
		if len(testVals) == 0 {
			assert.Equal(t, 0, len(q.List()))
			break
		}
		assert.Equal(t, testVals, q.List())
	}
}

func TestPopFront(t *testing.T) {
	testVals := make([]int, len(vals))
	copy(testVals, vals)

	q := queue.New[int]()
	for _, v := range vals {
		q.PushBack(v)
	}

	for {
		v, ok := q.PopFront()
		if !ok {
			break
		}
		testVal := testVals[0]
		assert.Equal(t, testVal, v)

		testVals = testVals[1:]
		assert.Equal(t, len(testVals), int((q.Size())))
		if len(testVals) == 0 {
			assert.Equal(t, 0, len(q.List()))
			break
		}
		assert.Equal(t, q.List(), testVals)
	}
}

func TestQueue(t *testing.T) {
	// random
	q := queue.New[int]()

	q.PushBack(1)
	assert.Equal(t, []int{1}, q.List())

	q.PushFront(2)
	assert.Equal(t, []int{2, 1}, q.List())

	q.PushBack(3)
	assert.Equal(t, []int{2, 1, 3}, q.List())

	q.PushBack(4)
	assert.Equal(t, []int{2, 1, 3, 4}, q.List())

	q.PushFront(5)
	assert.Equal(t, []int{5, 2, 1, 3, 4}, q.List())

	assert.Equal(t, 5, int(q.Size()))

	v, ok := q.PopBack()
	assert.True(t, ok)
	assert.Equal(t, 4, v)
	assert.Equal(t, []int{5, 2, 1, 3}, q.List())
	assert.Equal(t, 4, int(q.Size()))
	v, ok = q.PeekFront()
	assert.True(t, ok)
	assert.Equal(t, 5, v)
	v, ok = q.PeekBack()
	assert.True(t, ok)
	assert.Equal(t, 3, v)

	v, ok = q.PopFront()
	assert.True(t, ok)
	assert.Equal(t, 5, v)
	assert.Equal(t, []int{2, 1, 3}, q.List())
	assert.Equal(t, 3, int(q.Size()))
	v, ok = q.PeekFront()
	assert.True(t, ok)
	assert.Equal(t, 2, v)
	v, ok = q.PeekBack()
	assert.True(t, ok)
	assert.Equal(t, 3, v)

	q.PushFront(6)
	assert.Equal(t, []int{6, 2, 1, 3}, q.List())

	for i := 0; i < 20; i++ {
		q.PopBack()
	}
	v, ok = q.PopBack()
	assert.False(t, ok)
	assert.Equal(t, 0, v) // default int
	assert.True(t, q.IsEmpty())

	v, ok = q.PopFront()
	assert.False(t, ok)
	assert.Equal(t, 0, v) // default int
	assert.True(t, q.IsEmpty())

	_, ok = q.PeekFront()
	assert.False(t, ok)
	_, ok = q.PeekBack()
	assert.False(t, ok)
}
