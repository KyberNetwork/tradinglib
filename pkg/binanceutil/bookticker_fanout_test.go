package binanceutil_test

import (
	"testing"
	"time"

	"github.com/KyberNetwork/tradinglib/pkg/binanceutil"
)

func TestBookTickerFanout(t *testing.T) {
	var (
		symbols = []string{"KNCUSDT", "LINKUSDT", "C98USDT"}
		tick    = time.Tick(time.Second * 5)
	)

	bt := binanceutil.NewBookTickerFanOut(10, 10)
	for i := range symbols {
		go func(idx int) {
			c := bt.Subscribe(symbols[idx])
			for e := range c {
				t.Log(e)
			}
		}(i)
	}
	<-tick
}
