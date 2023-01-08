package httpclient

import (
	"net/http"
)

type doRequestOption struct {
	expectedStatusCode int
}

type Option func(o *doRequestOption)

func defaultOption() doRequestOption {
	return doRequestOption{
		expectedStatusCode: http.StatusOK,
	}
}

func (x *doRequestOption) apply(options ...Option) {
	for _, o := range options {
		o(x)
	}
}

func WithStatusCode(status int) Option {
	return func(o *doRequestOption) {
		o.expectedStatusCode = status
	}
}
