package list_test

import (
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/ds/list"
	"github.com/stretchr/testify/assert"
)

func TestCycleBuffer(t *testing.T) {
	c := list.NewRing[int](5)
	for i := 0; i < 5; i++ {
		c.Append(i)
	}
	assert.Equal(t, list.List[int]{0, 1, 2, 3, 4}, c.Filter(func(e int) bool { return true }))
	for i := 0; i < 3; i++ {
		c.Append(i + 5)
	}
	assert.Equal(t, list.List[int]{0, 1, 2, 3, 4, 5, 6, 7}, c.Filter(func(e int) bool { return true }))
	c.Expire(2)
	assert.Equal(t, list.List[int]{2, 3, 4, 5, 6, 7}, c.Filter(func(e int) bool { return true }))
	c.ExpireCond(func(item int) bool {
		return item < 5
	})
	assert.Equal(t, list.List[int]{5, 6, 7}, c.Filter(func(e int) bool { return true }))
	c.Log()
}

func TestList(t *testing.T) {
	ll := []int{1, 2, 3, 4, 5}
	var list2 list.List[int] = ll
	t.Log(list2.First())
}
