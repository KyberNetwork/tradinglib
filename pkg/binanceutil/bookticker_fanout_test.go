package binanceutil_test

import (
	"testing"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/stretchr/testify/require"
)

// func TestBookTickerFanout(t *testing.T) {
// 	var (
// 		symbols = []string{"KNCUSDT", "LINKUSDT", "C98USDT"}
// 		tick    = time.Tick(time.Second * 5)
// 	)

// 	bt := binanceutil.NewBookTickerFanOut(10, 10)
// 	for i := range symbols {
// 		go func(idx int) {
// 			c := bt.Subscribe(symbols[idx])
// 			for e := range c {
// 				t.Log(e)
// 			}
// 		}(i)
// 	}
// 	<-tick
// }

func TestWs(t *testing.T) {
	tick := time.Tick(time.Second * 5)
	symbols := []string{"KNCUSDT", "LINKUSDT", "C98USDT"}
	go func() {
		done, _, err := binance.WsCombinedMarketStatServe(
			symbols,
			func(event *binance.WsMarketStatEvent) {
				t.Log(event)
			},
			func(err error) {
				panic("subscribe error")
			})
		require.NoError(t, err)
		<-done
	}()

	<-tick
}
