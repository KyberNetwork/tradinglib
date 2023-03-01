package binanceutil

import "sync"

type bookTickerConn struct {
	l       sync.RWMutex
	symbols []string
	doneC   chan struct{}
}

func (b *bookTickerConn) SearchSymbol(symbol string) bool {
	b.l.RLock()
	defer b.l.RUnlock()

	for i := range b.symbols {
		if b.symbols[i] == symbol {
			return true
		}
	}

	return false
}

func (b *bookTickerConn) CheckAndAppendSymbol(maxLength int, symbol string) bool {
	b.l.Lock()
	defer b.l.Unlock()

	if len(b.symbols) >= maxLength {
		return false
	}

	b.symbols = append(b.symbols, symbol)

	return true
}

func (b *bookTickerConn) CheckAndReniveSymbol(symbol string) bool {
	b.l.Lock()
	defer b.l.Unlock()

	for i := range b.symbols {
		if b.symbols[i] != symbol {
			continue
		}

		b.symbols = append(b.symbols[:i], b.symbols[i+1:]...) // remove elem i
		return true
	}

	return false
}
