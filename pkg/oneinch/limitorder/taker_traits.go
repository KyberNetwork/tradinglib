package limitorder

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type AmountMode uint

const (
	makerAmountFlag = 255
	argsHasReceiver = 251

	// 224-247 bits `ARGS_EXTENSION_LENGTH`   - The length of the extension calldata in the args.
	argsExtensionLenStart = 224
	argsExtensionLenEnd   = 247

	// 200-223 bits `ARGS_INTERACTION_LENGTH` - The length of the interaction calldata in the args.
	argsInteractionLenStart = 200
	argsInteractionLenEnd   = 223

	// 0-184 bits                             - The threshold amount.
	numAmountThresholdBits = 185

	AmountModeTaker AmountMode = 0
	AmountModeMaker AmountMode = 1
)

type TakerTraits struct {
	flags       *big.Int
	receiver    *common.Address
	extension   *Extension
	interaction *Interaction
}

func NewTakerTraits(
	flags *big.Int, receiver *common.Address, extension *Extension, interaction *Interaction,
) *TakerTraits {
	return &TakerTraits{
		flags:       flags,
		receiver:    receiver,
		extension:   extension,
		interaction: interaction,
	}
}

func NewDefaultTakerTraits() *TakerTraits {
	return &TakerTraits{
		flags: new(big.Int),
	}
}

func (t *TakerTraits) SetAmountMode(mode AmountMode) *TakerTraits {
	t.flags.SetBit(t.flags, makerAmountFlag, uint(mode))
	return t
}

// SetAmountThreshold sets threshold amount.
//
// In taker amount mode: the minimum amount a taker agrees to receive in exchange for a taking amount.
// In maker amount mode: the maximum amount a taker agrees to give in exchange for a making amount.
func (t *TakerTraits) SetAmountThreshold(threshold *big.Int) *TakerTraits {
	clearAndSetLowerBits(t.flags, numAmountThresholdBits, threshold)
	return t
}

// SetExtension sets extension, it is required to provide same extension as in order creation (if any).
func (t *TakerTraits) SetExtension(ext Extension) *TakerTraits {
	t.extension = &ext
	return t
}

func (t *TakerTraits) Encode() (common.Hash, []byte) {
	var extension, interaction []byte
	if t.extension != nil {
		extension = t.extension.Encode()
	}
	if t.interaction != nil {
		interaction = t.interaction.Encode()
	}

	flags := new(big.Int).Set(t.flags)
	if t.receiver != nil {
		flags.SetBit(flags, argsHasReceiver, 1)
	}

	// Set length for extension and interaction.
	setBitsRange(flags, argsExtensionLenStart, argsExtensionLenEnd, big.NewInt(int64(len(extension))))
	setBitsRange(flags, argsInteractionLenStart, argsInteractionLenEnd, big.NewInt(int64(len(interaction))))

	var args []byte
	if t.receiver == nil {
		args = make([]byte, 0, len(extension)+len(interaction))
	} else {
		args = make([]byte, 0, len(t.receiver)+len(extension)+len(interaction))
		args = append(args, t.receiver.Bytes()...)
	}
	args = append(append(args, extension...), interaction...)

	return common.BigToHash(flags), args
}

func clearAndSetLowerBits(x *big.Int, n int, value *big.Int) {
	// Clear the lower n bits.
	mask := new(big.Int).Lsh(big.NewInt(1), uint(n))
	mask.Sub(mask, big.NewInt(1))
	mask.Not(mask)

	x = x.And(x, mask)

	// Mask and shift the provided value.
	value = new(big.Int).And(value, mask.Not(mask)) // Ensure value fits in n bits

	// Combine the results
	x.Or(x, value)
}

func setBitsRange(n *big.Int, start, end int, value *big.Int) {
	// Create a mask with bits set from start to end
	mask := new(big.Int)
	for i := start; i <= end; i++ {
		mask.SetBit(mask, i, 1)
	}

	// Shift the value to the correct position
	value = new(big.Int).Lsh(value, uint(start))

	// Clear the bits in the range (AND with the negated mask)
	n.And(n, mask.Not(mask))

	// Set the bits in the range (OR with the value)
	n.Or(n, value)
}
