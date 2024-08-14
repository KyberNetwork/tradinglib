package limitorder

import (
	"fmt"
	"math"
	"math/big"
	"strings"
)

const (
	ZX = "0x"

	totalOffsetSlots     = 8
	offsetSlotSizeInBits = 32
	offsetLengthInBytes  = totalOffsetSlots * offsetSlotSizeInBits / 8
	offsetLengthInHex    = offsetLengthInBytes * 2
)

// Extension represents the extension data of a 1inch order.
// This is copied from
// nolint: lll
// https://github.com/1inch/limit-order-sdk/blob/999852bc3eb92fb75332b7e3e0300e74a51943c1/src/limit-order/extension.ts#L6
type Extension struct {
	MakerAssetSuffix string
	TakerAssetSuffix string
	MakingAmountData string
	TakingAmountData string
	Predicate        string
	MakerPermit      string
	PreInteraction   string
	PostInteraction  string
	CustomData       string
}

func (e Extension) IsEmpty() bool {
	return len(e.getConcatenatedInteractions()) == 0
}

func (e Extension) Encode() string {
	interactionsConcatenated := e.getConcatenatedInteractions()
	if interactionsConcatenated == "" {
		return ZX
	}

	offsetsBytes := e.getOffsets()
	paddedOffsetHex := fmt.Sprintf("%064x", offsetsBytes)
	return ZX + paddedOffsetHex + interactionsConcatenated + trim0x(e.CustomData)
}

func (e Extension) interactionsArray() [totalOffsetSlots]string {
	return [totalOffsetSlots]string{
		e.MakerAssetSuffix,
		e.TakerAssetSuffix,
		e.MakingAmountData,
		e.TakingAmountData,
		e.Predicate,
		e.MakerPermit,
		e.PreInteraction,
		e.PostInteraction,
	}
}

func (e Extension) getConcatenatedInteractions() string {
	var builder strings.Builder
	for _, interaction := range e.interactionsArray() {
		interaction = trim0x(interaction)
		builder.WriteString(interaction)
	}
	return builder.String()
}

func (e Extension) getOffsets() *big.Int {
	var lengthMap [totalOffsetSlots]int
	for i, interaction := range e.interactionsArray() {
		// nolint: gomnd
		lengthMap[i] = len(trim0x(interaction)) / 2 // divide by 2 because each byte is represented by 2 hex characters
	}

	cumulativeSum := 0
	bytesAccumulator := big.NewInt(0)
	var index uint64

	for _, length := range lengthMap {
		cumulativeSum += length
		shiftVal := big.NewInt(int64(cumulativeSum))
		shiftVal.Lsh(shiftVal, uint(offsetSlotSizeInBits*index)) // Shift left
		bytesAccumulator.Add(bytesAccumulator, shiftVal)         // Add to accumulator
		index++
	}

	return bytesAccumulator
}

// DecodeExtension decodes the encoded extension string into an Extension struct.
// The encoded extension string is expected to be in the format of "0x" followed by the hex-encoded extension data.
// The hex-encoded extension data is expected to be in
// the format of 32 bytes of offset data followed by the extension data.
func DecodeExtension(encodedExtension string) (Extension, error) {
	if encodedExtension == ZX {
		return defaultExtension(), nil
	}

	encodedExtension = trim0x(encodedExtension)

	// nolint: gomnd
	offset, ok := new(big.Int).SetString(encodedExtension[:offsetLengthInHex], 16)
	if !ok {
		return Extension{}, fmt.Errorf("decode offset from encoded extension")
	}

	maxInt32 := big.NewInt(math.MaxInt32)

	extensionData := encodedExtension[offsetLengthInHex:]

	data := [totalOffsetSlots]string{}
	prevLength := 0
	for i := 0; i < totalOffsetSlots; i++ {
		length := int(new(big.Int).And(
			new(big.Int).Rsh(
				offset, uint(i*offsetSlotSizeInBits),
			),
			maxInt32,
		).Int64())

		// multiply by 2 because each byte is represented by 2 hex characters
		start := prevLength * 2 // nolint: gomnd
		end := length * 2       // nolint: gomnd

		data[i] = extensionData[start:end]

		prevLength = length
	}
	customData := extensionData[prevLength*2:]

	return Extension{
		MakerAssetSuffix: add0x(data[0]),
		TakerAssetSuffix: add0x(data[1]),
		MakingAmountData: add0x(data[2]),
		TakingAmountData: add0x(data[3]),
		Predicate:        add0x(data[4]),
		MakerPermit:      add0x(data[5]),
		PreInteraction:   add0x(data[6]),
		PostInteraction:  add0x(data[7]),
		CustomData:       add0x(customData),
	}, nil
}

func defaultExtension() Extension {
	return Extension{
		MakerAssetSuffix: ZX,
		TakerAssetSuffix: ZX,
		MakingAmountData: ZX,
		TakingAmountData: ZX,
		Predicate:        ZX,
		MakerPermit:      ZX,
		PreInteraction:   ZX,
		PostInteraction:  ZX,
		CustomData:       ZX,
	}
}

func trim0x(s string) string {
	return strings.TrimPrefix(s, "0x")
}

func add0x(s string) string {
	return "0x" + s
}
