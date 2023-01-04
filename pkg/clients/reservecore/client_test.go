package reservecore

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/httpsign"
	"github.com/stretchr/testify/require"
)

func jsonify(data interface{}) string {
	d, _ := json.MarshalIndent(data, "", " ")
	return string(d)
}
func TestNewClient(t *testing.T) {
	t.Skip()
	hc := httpsign.NewClient(http.DefaultTransport,
		"..",
		[]byte(".."))
	client, err := New("https://staging-core-v3-gateway.knstats.com", true, hc)
	require.NoError(t, err)
	resp, err := client.GetAuthData(0)
	require.NoError(t, err)
	t.Log(jsonify(resp))

	m, err := client.GetMarginConfig()
	require.NoError(t, err)
	t.Log(jsonify(m))
}
