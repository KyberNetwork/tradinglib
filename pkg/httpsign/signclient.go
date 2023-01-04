package httpsign

import (
	"net/http"
)

func NewClient(transport http.RoundTripper, key string, secret []byte) *http.Client {
	return &http.Client{
		Transport: NewTransport(transport, key, secret),
	}
}

func NewTransport(inner http.RoundTripper, key string, secret []byte) HttpSign {
	return HttpSign{
		inner:  inner,
		keyID:  key,
		secret: secret,
	}
}

type HttpSign struct {
	inner  http.RoundTripper
	keyID  string
	secret []byte
}

func (s HttpSign) RoundTrip(request *http.Request) (*http.Response, error) {
	if len(s.keyID) == 0 {
		return s.inner.RoundTrip(request)
	}
	signed, err := Sign(request, s.keyID, s.secret)
	if err != nil {
		return nil, err
	}
	return s.inner.RoundTrip(signed)
}
