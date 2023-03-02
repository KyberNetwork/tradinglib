package syncmap_test

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/KyberNetwork/tradinglib/x/syncmap"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

func TestSyncMap(t *testing.T) {
	type testData struct {
		k string
		v int
	}
	var (
		tests     []testData
		testRange = 1_000_000
	)
	for i := 0; i < testRange; i++ {
		tests = append(tests, testData{
			k: strconv.Itoa(i),
			v: i,
		})
	}

	sm := syncmap.New[string, int]()
	errGr := errgroup.Group{}

	errGr.Go(func() error {
		for i := range tests {
			sm.Store(tests[i].k, tests[i].v)
		}
		return nil
	})
	errGr.Go(func() error {
		for i := range tests {
			sm.Store(tests[i].k, tests[i].v)
		}
		return nil
	})
	errGr.Go(func() error {
		sm.Load(strconv.Itoa(rand.Intn(testRange))) // nolint: gosec
		return nil
	})
	errGr.Go(func() error {
		sm.Delete(strconv.Itoa(rand.Intn(testRange))) // nolint: gosec
		return nil
	})
	errGr.Go(func() error {
		for i := range tests {
			_, err := sm.Update(tests[i].k, func(v int) (int, error) {
				if v < 5 {
					v *= 5
				}
				return v, nil
			})
			require.NoError(t, err)
		}
		return nil
	})
	_ = errGr.Wait()
}
