package eth_test

import (
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/eth"
	"github.com/test-go/testify/require"
)

func TestRecover(t *testing.T) {
	hashHex := "0xef6777697377372751e5daa1772d3b2b03f6f50b85dd3d7a79b4f12381897ec8"
	signatureHex := "0xf89e6165aa337b8ebf74effa1a2b0980cc54e3780af077337b12438f30c901ab69ad655518007bd15073c919780ac61d3a4d5918c8966ffdb07256776ca6223f1c" // nolint: lll

	addr, err := eth.RecoverSignerAddress(hashHex, signatureHex)
	require.NoError(t, err)

	t.Log(addr.String())
}
