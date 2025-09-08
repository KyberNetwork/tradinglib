package fusionextention_test

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/fusionorder/fusionextention"
	"github.com/stretchr/testify/assert"
)

func TestMarshalUnMarshal(t *testing.T) {
	data := fusionextention.SurplusParam{
		EstimatedTakerAmount: big.NewInt(1),
		ProtocolFee:          100,
	}

	jsonData, err := json.Marshal(data)
	assert.NoError(t, err, "marshal should not error")
	var decoded fusionextention.SurplusParam
	err = json.Unmarshal(jsonData, &decoded)
	assert.NoError(t, err, "unmarshal should not error")
	assert.Equal(t, data, decoded, "decoded data should match original")
}
