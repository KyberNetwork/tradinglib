package mutexmap_test

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/ds/mutexmap"
	"golang.org/x/sync/errgroup"
)

func TestMutextMap(t *testing.T) {
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

	mm := mutexmap.New[string, int]()
	errGr := errgroup.Group{}

	errGr.Go(func() error {
		for i := range tests {
			mm.Store(tests[i].k, tests[i].v)
		}
		return nil
	})
	errGr.Go(func() error {
		for i := range tests {
			mm.Store(tests[i].k, tests[i].v)
		}
		return nil
	})
	errGr.Go(func() error {
		mm.Load(strconv.Itoa(rand.Intn(testRange))) // nolint: gosec
		return nil
	})
	errGr.Go(func() error {
		mm.Delete(strconv.Itoa(rand.Intn(testRange))) // nolint: gosec
		return nil
	})
	_ = errGr.Wait()
}
