package eth_test

import (
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/convert"
	"github.com/KyberNetwork/tradinglib/pkg/eth"
)

func TestCalcNextBaseFee(t *testing.T) {
	// block: 23135915
	baseFee := eth.CalcNextBaseFee(
		convert.MustFloatToWei(0.887536596, 9), // 0.887536596 Gwei
		43718352,
		45043901,
	)

	t.Log(baseFee) // 991949082
}
