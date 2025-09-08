package fusionextention

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestMarshalUnMarshal(t *testing.T) {
	data := SurplusParam{
		EstimatedTakerAmount: big.NewInt(1),
		ProtocolFee:          100,
	}

	jsonData, err := json.Marshal(data)
	assert.NoError(t, err, "marshal should not error")
	var decoded SurplusParam
	err = json.Unmarshal(jsonData, &decoded)
	assert.NoError(t, err, "unmarshal should not error")
	assert.Equal(t, data, decoded, "decoded data should match original")
}
