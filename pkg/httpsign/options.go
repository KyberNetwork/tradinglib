package httpsign

import (
	"net/http"
	"time"
)

type Option func(c *http.Client)

func WithHTTPTimeout(timeout time.Duration) Option {
	return func(c *http.Client) {
		if timeout == 0 {
			return
		}
		c.Timeout = timeout
	}
}
