package limitorder

import "github.com/ethereum/go-ethereum/common"

type Interaction struct {
	Target common.Address
	Data   string
}

func (i Interaction) Encode() string {
	return i.Target.String() + trim0x(i.Data)
}

func DecodeInteraction(encoded string) Interaction {
	const addressLength = 42 // 42 is the length of an address len(0x) + 20 bytes
	return Interaction{
		Target: common.HexToAddress(encoded[:addressLength]),
		Data:   add0x(encoded[addressLength:]),
	}
}
