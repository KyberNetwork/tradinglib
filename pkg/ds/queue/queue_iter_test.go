package queue_test

import (
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/ds/queue"
	"github.com/stretchr/testify/assert"
)

func TestQueueIterReadOnly(t *testing.T) {
	q := queue.New[int]()
	for _, v := range vals {
		q.PushBack(v)
	}

	iter := queue.NewIter[int](q)
	collect := make([]int, 0)
	for iter.Next() {
		v, ok := iter.Val()
		assert.True(t, ok)
		collect = append(collect, v)
	}

	assert.Equal(t, vals, collect)
}

func TestQueueIterRemoveEven(t *testing.T) {
	testVals := make([]int, 0, len(vals))
	for i := 0; i < len(vals); i++ {
		if vals[i]%2 == 0 {
			continue
		}

		testVals = append(testVals, vals[i])
	}

	q := queue.New[int]()
	for _, v := range vals {
		q.PushBack(v)
	}

	iter := queue.NewIter[int](q)
	collect := make([]int, 0)
	for iter.Next() {
		v, _ := iter.Val()
		if v%2 == 0 {
			iter.RemoveCurrent()
		}
	}

	iter.Reset()
	for iter.Next() {
		v, _ := iter.Val()
		collect = append(collect, v)
	}

	assert.Equal(t, testVals, collect)
}

func TestQueueIterRemoveOdd(t *testing.T) {
	testVals := make([]int, 0, len(vals))
	for i := 0; i < len(vals); i++ {
		if vals[i]%2 != 0 {
			continue
		}

		testVals = append(testVals, vals[i])
	}

	q := queue.New[int]()
	for _, v := range vals {
		q.PushBack(v)
	}

	iter := queue.NewIter[int](q)
	collect := make([]int, 0)
	for iter.Next() {
		v, _ := iter.Val()
		if v%2 != 0 {
			iter.RemoveCurrent()
		}
	}

	iter.Reset()
	for iter.Next() {
		v, _ := iter.Val()
		collect = append(collect, v)
	}

	assert.Equal(t, testVals, collect)
}

func TestQueueIterRemove456(t *testing.T) {
	skip := func(i int) bool {
		if i == 4 || i == 5 || i == 6 {
			return true
		}
		return false
	}

	testVals := make([]int, 0, len(vals))
	for i := 0; i < len(vals); i++ {
		if skip(vals[i]) {
			continue
		}

		testVals = append(testVals, vals[i])
	}

	q := queue.New[int]()
	for _, v := range vals {
		q.PushBack(v)
	}

	iter := queue.NewIter[int](q)
	collect := make([]int, 0)
	for iter.Next() {
		v, _ := iter.Val()
		if skip(v) {
			iter.RemoveCurrent()
		}
	}

	iter.Reset()
	for iter.Next() {
		v, _ := iter.Val()
		collect = append(collect, v)
	}

	assert.Equal(t, testVals, collect)
}
