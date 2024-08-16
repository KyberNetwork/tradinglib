package limitorder

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	ZX = "0x"

	totalOffsetSlots     = 8
	offsetSlotSizeInBits = 32
	offsetLength         = totalOffsetSlots * offsetSlotSizeInBits / 8
)

// Extension represents the extension data of a 1inch order.
// This is copied from
// nolint: lll
// https://github.com/1inch/limit-order-sdk/blob/999852bc3eb92fb75332b7e3e0300e74a51943c1/src/limit-order/extension.ts#L6
type Extension struct {
	MakerAssetSuffix []byte
	TakerAssetSuffix []byte
	MakingAmountData []byte
	TakingAmountData []byte
	Predicate        []byte
	MakerPermit      []byte
	PreInteraction   []byte
	PostInteraction  []byte
	CustomData       []byte
}

func (e Extension) HasMakerPermit() bool {
	return len(e.MakerPermit) > 0
}

func (e Extension) IsEmpty() bool {
	return len(e.getConcatenatedInteractions()) == 0
}

func (e Extension) Encode() string {
	interactionsConcatenated := e.getConcatenatedInteractions()
	if len(interactionsConcatenated) == 0 {
		return ZX
	}

	offsetsBytes := e.getOffsets()
	paddedOffsetHex := fmt.Sprintf("%064x", offsetsBytes)
	return ZX + paddedOffsetHex + hex.EncodeToString(interactionsConcatenated) + hex.EncodeToString(e.CustomData)
}

func (e Extension) interactionsArray() [totalOffsetSlots][]byte {
	return [totalOffsetSlots][]byte{
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

func (e Extension) getConcatenatedInteractions() []byte {
	builder := new(bytes.Buffer)
	for _, interaction := range e.interactionsArray() {
		builder.Write(interaction)
	}
	return builder.Bytes()
}

func (e Extension) getOffsets() *big.Int {
	var lengthMap [totalOffsetSlots]int
	for i, interaction := range e.interactionsArray() {
		lengthMap[i] = len(interaction)
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

	extensionDataBytes, err := hexutil.Decode(encodedExtension)
	if err != nil {
		return Extension{}, fmt.Errorf("decode extension data: %w", err)
	}

	offset := new(big.Int).SetBytes(extensionDataBytes[:offsetLength])

	maxInt32 := big.NewInt(math.MaxInt32)

	extensionData := extensionDataBytes[offsetLength:]

	data := [totalOffsetSlots][]byte{}
	prevLength := 0
	for i := 0; i < totalOffsetSlots; i++ {
		length := int(new(big.Int).And(
			new(big.Int).Rsh(
				offset, uint(i*offsetSlotSizeInBits),
			),
			maxInt32,
		).Int64())

		start := prevLength
		end := length

		data[i] = extensionData[start:end]

		prevLength = length
	}
	customData := extensionData[prevLength:]

	e := Extension{
		MakerAssetSuffix: data[0],
		TakerAssetSuffix: data[1],
		MakingAmountData: data[2],
		TakingAmountData: data[3],
		Predicate:        data[4],
		MakerPermit:      data[5],
		PreInteraction:   data[6],
		PostInteraction:  data[7],
		CustomData:       customData,
	}

	return e, nil
}

func defaultExtension() Extension {
	return Extension{}
}
