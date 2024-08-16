package limitorder

import (
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/utils"
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

func (e Extension) validate() error {
	if !utils.IsHexString(e.MakerAssetSuffix) {
		return fmt.Errorf("invalid maker asset suffix: %s", e.MakerAssetSuffix)
	}
	if !utils.IsHexString(e.TakerAssetSuffix) {
		return fmt.Errorf("invalid taker asset suffix: %s", e.TakerAssetSuffix)
	}
	if !utils.IsHexString(e.MakingAmountData) {
		return fmt.Errorf("invalid making amount data: %s", e.MakingAmountData)
	}
	if !utils.IsHexString(e.TakingAmountData) {
		return fmt.Errorf("invalid taking amount data: %s", e.TakingAmountData)
	}
	if !utils.IsHexString(e.Predicate) {
		return fmt.Errorf("invalid predicate: %s", e.Predicate)
	}
	if !utils.IsHexString(e.MakerPermit) {
		return fmt.Errorf("invalid maker permit: %s", e.MakerPermit)
	}
	if !utils.IsHexString(e.PreInteraction) {
		return fmt.Errorf("invalid pre interaction: %s", e.PreInteraction)
	}
	if !utils.IsHexString(e.PostInteraction) {
		return fmt.Errorf("invalid post interaction: %s", e.PostInteraction)
	}
	if !utils.IsHexString(e.CustomData) {
		return fmt.Errorf("invalid custom data: %s", e.CustomData)
	}
	return nil
}

func (e Extension) HasMakerPermit() bool {
	return e.MakerPermit != ZX
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
	return ZX + paddedOffsetHex + interactionsConcatenated + utils.Trim0x(e.CustomData)
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
		interaction = utils.Trim0x(interaction)
		builder.WriteString(interaction)
	}
	return builder.String()
}

func (e Extension) getOffsets() *big.Int {
	var lengthMap [totalOffsetSlots]int
	for i, interaction := range e.interactionsArray() {
		// nolint: gomnd
		lengthMap[i] = len(utils.Trim0x(interaction)) / 2 // divide by 2 because each byte is represented by 2 hex characters
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

	encodedExtension = utils.Trim0x(encodedExtension)

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

	e := Extension{
		MakerAssetSuffix: utils.Add0x(data[0]),
		TakerAssetSuffix: utils.Add0x(data[1]),
		MakingAmountData: utils.Add0x(data[2]),
		TakingAmountData: utils.Add0x(data[3]),
		Predicate:        utils.Add0x(data[4]),
		MakerPermit:      utils.Add0x(data[5]),
		PreInteraction:   utils.Add0x(data[6]),
		PostInteraction:  utils.Add0x(data[7]),
		CustomData:       utils.Add0x(customData),
	}

	return e, e.validate()
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
