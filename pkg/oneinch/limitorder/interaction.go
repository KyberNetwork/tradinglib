package limitorder

import (
	"encoding/hex"
	"fmt"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/decode"
	"github.com/ethereum/go-ethereum/common"
)

type Interaction struct {
	Target common.Address
	Data   []byte
}

func (i Interaction) IsZero() bool {
	return i.Target.String() == common.Address{}.String() && len(i.Data) == 0
}

func (i Interaction) Encode() string {
	return i.Target.String() + hex.EncodeToString(i.Data)
}

func DecodeInteraction(data []byte) (Interaction, error) {
	bi := decode.NewBytesIterator(data)
	target, err := bi.NextBytes(common.AddressLength)
	if err != nil {
		return Interaction{}, fmt.Errorf("get target: %w", err)
	}
	return Interaction{
		Target: common.BytesToAddress(target),
		Data:   bi.RemainingData(),
	}, nil
}
