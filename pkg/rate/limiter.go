package rate

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Limiter struct {
	limit  int
	period time.Duration

	mu               sync.Mutex
	start, end, size int
	requests         []time.Time
}

func NewLimiter(limit int, period time.Duration) *Limiter {
	return &Limiter{
		limit:    limit,
		period:   period,
		start:    0,
		end:      0,
		requests: make([]time.Time, limit),
	}
}

func (l *Limiter) Wait(ctx context.Context) error {
	return l.WaitN(ctx, 1)
}

func (l *Limiter) WaitN(ctx context.Context, n int) error {
	if n <= 0 {
		return nil
	}

	if n > l.limit {
		return fmt.Errorf("rate: Wait (n=%d) exceed limiter %d", n, l.limit)
	}

	now := time.Now()

	// Get the oldest request in queue for waiting and add new requests to queue.
	var (
		shouldWait bool
		oldest     time.Time
	)

	l.mu.Lock()

	{
		if l.requestsSize()+n > l.limit {
			shouldWait = true
			oldest, _ = l.requestAt(l.requestsSize() + n - l.limit - 1)
		}

		for range n {
			l.addRequest(now)
		}
	}

	l.mu.Unlock()

	// Wait if rate limit is reached.
	if shouldWait {
		waitDuration := l.period - now.Sub(oldest)
		if waitDuration > 0 {
			timer := time.NewTimer(waitDuration)
			defer timer.Stop()

			select {
			case <-timer.C:
				return nil
			case <-ctx.Done():
				// Context was canceled before we could proceed.
				return ctx.Err()
			}
		}
	}

	return nil
}

func (l *Limiter) requestsSize() int {
	return l.size
}

func (l *Limiter) requestAt(i int) (time.Time, bool) {
	if l.size == 0 {
		return time.Now(), false
	}

	return l.requests[(l.start+i)%l.limit], true
}

func (l *Limiter) addRequest(t time.Time) {
	if l.size == l.limit {
		l.start++
		if l.start >= l.limit {
			l.start = 0
		}
	} else {
		l.size++
	}

	l.requests[l.end] = t

	l.end++
	if l.end >= l.limit {
		l.end = 0
	}
}
