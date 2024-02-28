package bsync_test

import (
	"github.com/KyberNetwork/tradinglib/pkg/bsync"
	"testing"
	"time"
)

func TestWorker(t *testing.T) {
	v := 0
	w := bsync.New(time.Millisecond, func() (int, error) {
		v++
		return v, nil
	})
	go w.Start()
	for i := 0; i < 10; i++ {
		c, at := w.Get()
		t.Log(c, at)
		time.Sleep(time.Millisecond)
	}
	w.Stop()
}
