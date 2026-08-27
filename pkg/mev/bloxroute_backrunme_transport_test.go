// The test reaches sender.httpClient to make the sender's OWN transport issue the request;
// an external test package could only swap in a different client, which would not exercise
// what the constructor built.
//
//nolint:testpackage
package mev

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

// TestBackrunmeSenderNegotiatesHTTP2 pins the protocol the sender actually speaks.
//
// The constructor used to hand its Transport a TLSClientConfig (for InsecureSkipVerify). net/http
// only enables HTTP/2 when it is free to append "h2" to TLSClientConfig.NextProtos, and it refuses
// to touch a caller-supplied config unless ForceAttemptHTTP2 is set (Transport.protocols, Go issue
// 14275) — so that one field silently pinned every submit to HTTP/1.1, where the default
// MaxIdleConnsPerHost of 2 costs a TCP+TLS handshake per concurrent submit past the second.
//
// The sender's OWN transport makes the request; only RootCAs is added so it trusts the test
// server's self-signed cert. Swapping in httptest's client instead would test that client's
// transport rather than the sender's, and the assertion would hold no matter what the constructor
// did.
func TestBackrunmeSenderNegotiatesHTTP2(t *testing.T) {
	t.Parallel()

	var gotProto string
	srv := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotProto = r.Proto
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":"1","result":{"bundleHash":"0xabc"}}`))
	}))
	srv.EnableHTTP2 = true
	srv.StartTLS()
	defer srv.Close()

	sender, err := NewBloxrouteBackrunmeSender("token", srv.URL)
	require.NoError(t, err)

	tr, ok := sender.httpClient.Transport.(*http.Transport)
	require.True(t, ok, "sender must use an *http.Transport")
	pool := x509.NewCertPool()
	pool.AddCert(srv.Certificate())
	if tr.TLSClientConfig == nil {
		tr.TLSClientConfig = &tls.Config{MinVersion: tls.VersionTLS12}
	}
	tr.TLSClientConfig.RootCAs = pool

	tx := types.NewTx(&types.LegacyTx{Nonce: 1, Gas: 21000, GasPrice: big.NewInt(1)})
	_, err = sender.SendBackrunBundle(context.Background(), nil, 1, 1,
		[]common.Hash{common.HexToHash("0x01")}, nil, big.NewInt(0), tx)
	require.NoError(t, err)

	require.Equal(t, "HTTP/2.0", gotProto,
		"backrunme submits must negotiate HTTP/2; a caller-supplied TLSClientConfig silently disables it")

	require.Positive(t, tr.IdleConnTimeout,
		"idle conns must expire; a Transport built from scratch never expires them")
}
