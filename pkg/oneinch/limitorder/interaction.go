package limitorder

import (
	"fmt"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/utils"
	"github.com/ethereum/go-ethereum/common"
)

type Interaction struct {
	Target common.Address
	Data   string
}

func (i Interaction) validate() error {
	if !utils.IsHexString(i.Data) {
		return fmt.Errorf("invalid data: %s", i.Data)
	}
	return nil
}

func (i Interaction) IsZero() bool {
	return i.Target == common.Address{} && i.Data == ""
}

func (i Interaction) Encode() string {
	return i.Target.String() + utils.Trim0x(i.Data)
}

func DecodeInteraction(encoded string) (Interaction, error) {
	const addressLength = 42 // 42 is the length of an address len(0x) + 20 bytes
	i := Interaction{
		Target: common.HexToAddress(encoded[:addressLength]),
		Data:   utils.Add0x(encoded[addressLength:]),
	}
	return i, i.validate()
}
