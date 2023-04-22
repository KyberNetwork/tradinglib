package httpsign_test

import (
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/KyberNetwork/tradinglib/pkg/httpsign"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRequest(t *testing.T) {
	t.Skip()
	client := http.Client{
		Transport: httpsign.NewTransport(http.DefaultTransport,
			"..",
			[]byte("..")),
	}
	resp, err := client.Get("https://test.knstats.com/authdata")
	require.NoError(t, err)
	t.Log(resp.StatusCode)
	out, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	require.NoError(t, err)
	t.Log(string(out))
}

func TestOption(t *testing.T) {
	timeout := time.Second * 3

	client := httpsign.NewClient(
		http.DefaultTransport, "..", []byte(".."), httpsign.WithTimeout(timeout),
	)
	assert.Equal(t, client.Timeout, timeout)
}
