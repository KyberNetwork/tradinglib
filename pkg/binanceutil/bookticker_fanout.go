package binanceutil

import (
	"fmt"
	"sync"

	"github.com/KyberNetwork/tradinglib/pkg/ds/mutexmap"
	"github.com/adshao/go-binance/v2"
)

type BookTickerFanOut struct {
	MaxSymbolPerConn int
	MaxEventPerChan  int
	fanOutMap        mutexmap.MutexMap[string, []chan binance.WsMarketStatEvent]
	connsLock        sync.RWMutex
	conns            []*bookTickerConn
}

func NewBookTickerFanOut(maxSymbolPerConn, maxEventPerChan int) *BookTickerFanOut {
	bt := &BookTickerFanOut{
		MaxSymbolPerConn: maxEventPerChan,
		MaxEventPerChan:  maxEventPerChan,
		fanOutMap:        mutexmap.New[string, []chan binance.WsMarketStatEvent](),
	}
	go bt.autoResubcribe()

	return bt
}

func (bt *BookTickerFanOut) Subscribe(symbol string) <-chan binance.WsMarketStatEvent {
	bt.connsLock.Lock()
	defer bt.connsLock.Unlock()

	c := make(chan binance.WsMarketStatEvent, bt.MaxEventPerChan)

	chans, ok := bt.fanOutMap.Load(symbol)
	if ok {
		chans = append(chans, c)
		bt.fanOutMap.Store(symbol, chans)
		return c
	}
	bt.fanOutMap.Store(symbol, []chan binance.WsMarketStatEvent{c})

	added := false
	for i := range bt.conns {
		if !bt.conns[i].CheckAndAppendSymbol(bt.MaxSymbolPerConn, symbol) {
			continue
		}

		if err := bt.subscribe(bt.conns[i]); err != nil {
			panic(err)
		}
		added = true
	}

	if added {
		return c
	}

	conn := &bookTickerConn{symbols: []string{symbol}}
	bt.conns = append(bt.conns, conn)

	if err := bt.subscribe(conn); err != nil {
		panic(err)
	}

	return c
}

func (bt *BookTickerFanOut) autoResubcribe() {
	bt.connsLock.RLock()
	defer bt.connsLock.RUnlock()

	for i := range bt.conns {
		select {
		case <-bt.conns[i].doneC:
			if err := bt.subscribe(bt.conns[i]); err != nil {
				panic(err)
			}
		default:
		}
	}
}

func (bt *BookTickerFanOut) subscribe(conn *bookTickerConn) error {
	done, _, err := binance.WsCombinedMarketStatServe(
		conn.symbols,
		bt.wsMarketStatHandler,
		func(err error) {
			panic("subscribe error")
		})
	if err != nil {
		return fmt.Errorf("WsCombinedMarketStatServe error: %w", err)
	}
	conn.doneC = done

	return nil
}

func (bt *BookTickerFanOut) wsMarketStatHandler(event *binance.WsMarketStatEvent) {
	chans, ok := bt.fanOutMap.Load(event.Symbol)
	if !ok {
		return
	}

	// safer to send a copy
	for _, c := range chans {
		select {
		case c <- *event:
		default:
			<-c // discard oldest event
			c <- *event
		}
	}
}
