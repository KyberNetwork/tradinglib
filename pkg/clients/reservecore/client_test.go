package reservecore_test

import (
	"net/http"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/clients/reservecore"
	"github.com/KyberNetwork/tradinglib/pkg/httpsign"
	"github.com/KyberNetwork/tradinglib/pkg/testutil"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	t.Skip()
	hc := httpsign.NewClient(http.DefaultTransport,
		"..",
		[]byte(".."))
	client, err := reservecore.New("https://staging-core-v3-gateway.knstats.com", true, hc)
	require.NoError(t, err)
	resp, err := client.GetAuthData(0)
	require.NoError(t, err)
	t.Log(testutil.MustJsonify(resp))

	m, err := client.GetMarginConfig()
	require.NoError(t, err)
	t.Log(testutil.MustJsonify(m))
}
