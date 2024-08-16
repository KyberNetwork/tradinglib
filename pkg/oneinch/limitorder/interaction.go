package limitorder

import (
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Interaction struct {
	Target common.Address
	Data   []byte
}

func (i Interaction) IsZero() bool {
	return i.Target.String() == common.Address{}.String() && len(i.Data) == 0
}

func (i Interaction) Encode() string {
	return i.Target.String() + utils.Trim0x(hexutil.Encode(i.Data))
}

func DecodeInteraction(encoded []byte) Interaction {
	return Interaction{
		Target: common.BytesToAddress(encoded[:common.AddressLength]),
		Data:   encoded[common.AddressLength:],
	}
}
