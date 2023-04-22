package httpsign

import (
	"net/http"
	"time"
)

// becareful, this client doesn't have timeout.
func NewClient(transport http.RoundTripper, key string, secret []byte) *http.Client {
	return &http.Client{
		Transport: NewTransport(transport, key, secret),
	}
}

func NewClientWithOptions(transport http.RoundTripper, key string, secret []byte, timeout time.Duration) *http.Client {
	return &http.Client{
		Transport: NewTransport(transport, key, secret),
		Timeout:   timeout,
	}
}

func NewTransport(inner http.RoundTripper, key string, secret []byte) HTTPSign {
	return HTTPSign{
		inner:  inner,
		keyID:  key,
		secret: secret,
	}
}

type HTTPSign struct {
	inner  http.RoundTripper
	keyID  string
	secret []byte
}

func (s HTTPSign) RoundTrip(request *http.Request) (*http.Response, error) {
	if len(s.keyID) == 0 {
		return s.inner.RoundTrip(request)
	}
	signed, err := Sign(request, s.keyID, s.secret)
	if err != nil {
		return nil, err
	}
	return s.inner.RoundTrip(signed)
}
