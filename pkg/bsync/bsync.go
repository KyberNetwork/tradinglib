package bsync

import (
	"sync"
	"time"
)

type LocalReader[T any] interface {
	Get() (T, time.Time)
}

// Worker support for background sync local state
type Worker[T any] struct {
	interval   time.Duration
	stop       chan struct{}
	fetchFunc  func() (T, error)
	localValue T
	localTime  time.Time
	lock       sync.Mutex
}

func New[T any](updateInternal time.Duration, fetchFunc func() (T, error)) *Worker[T] {
	return &Worker[T]{
		interval:  updateInternal,
		fetchFunc: fetchFunc,
		stop:      make(chan struct{}, 1),
	}
}

func (w *Worker[T]) setValue(v T) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.localValue = v
	w.localTime = time.Now()
}

func (w *Worker[T]) Get() (T, time.Time) {
	w.lock.Lock()
	defer w.lock.Unlock()
	return w.localValue, w.localTime
}

func (w *Worker[T]) Start() {
	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()
	if v, err := w.fetchFunc(); err == nil {
		w.setValue(v)
	}
	for {
		select {
		case <-w.stop:
			break
		case <-ticker.C:
			v, err := w.fetchFunc()
			if err != nil { // user should handle error themself.
				continue
			}
			w.setValue(v)
		}
	}
}

func (w *Worker[T]) Stop() {
	select {
	case w.stop <- struct{}{}:
	}
}
