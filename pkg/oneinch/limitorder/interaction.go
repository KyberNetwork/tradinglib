package limitorder

import (
	"encoding/hex"

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

func DecodeInteraction(encoded []byte) Interaction {
	return Interaction{
		Target: common.BytesToAddress(encoded[:common.AddressLength]),
		Data:   encoded[common.AddressLength:],
	}
}
