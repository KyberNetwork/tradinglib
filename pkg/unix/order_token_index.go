package unix

import (
	"fmt"
	"math/big"
)

const (
	// there are 248 bits for order index
	// 1 order is 4 bits, 1 token is 3 bits, then there are 248/7 = 35 orders
	maxOrderLength = 35
	maxOrderIndex  = 15 // 4 bits
	maxTokenIndex  = 7  // 3 bits
	zippedSize     = 256
)

var (
	maskMaxOrder      = big.NewInt(0xF)
	maskMaxToken      = big.NewInt(0x7)
	maskMaxOrderToken = big.NewInt(0x7F)
)

type OrderTokenIndex struct {
	OrderIndex int
	TokenIndex int
}

func (o OrderTokenIndex) Validate() error {
	if o.OrderIndex < 0 || o.OrderIndex > maxOrderIndex {
		return fmt.Errorf("%w, orderIndex %v", ErrOrderIndexInvalid, o.OrderIndex)
	}
	if o.TokenIndex < 0 || o.TokenIndex > maxTokenIndex {
		return fmt.Errorf("%w, interactionIndex %v", ErrTokenIndexInvalid, o.TokenIndex)
	}

	return nil
}

func PackTokenOutputIndex(orderTokenIndexes []OrderTokenIndex) (*big.Int, error) {
	result := big.NewInt(0)
	if len(orderTokenIndexes) == 0 {
		return result, nil
	}

	orderLength := len(orderTokenIndexes)
	if uint64(orderLength) > maxOrderLength {
		return nil, fmt.Errorf("%w, actual %v", ErrOrderLimitExceeded, orderLength)
	}

	// Write the first 8 bits as the number of orders
	result.Lsh(big.NewInt(int64(orderLength)), 248)

	for i, orderTokenIndex := range orderTokenIndexes {
		if err := orderTokenIndex.Validate(); err != nil {
			return nil, err
		}

		// 4 bits for order index, 3 bits for token index
		value := (orderTokenIndex.OrderIndex << 3) | orderTokenIndex.TokenIndex
		valBig := big.NewInt(int64(value))

		shifted := new(big.Int).Lsh(valBig, uint(i*7+8))
		result.Or(result, shifted)
	}

	return result, nil
}

// testing purpose only
func unpackTokenOutputIndex(input *big.Int) ([]OrderTokenIndex, error) {
	// Get the first 8 bits as the number of orders
	orders := int(new(big.Int).Rsh(input, 248).Uint64())

	var result []OrderTokenIndex
	for order := 0; order < orders; order++ {
		// Process totalToken token indices
		orderTokenIndex := getOrderTokenIndex(input, order)
		if err := orderTokenIndex.Validate(); err != nil {
			return nil, err
		}
		result = append(result, orderTokenIndex)
	}
	return result, nil
}

func getOrderTokenIndex(input *big.Int, i int) (result OrderTokenIndex) {
	shift := uint(i)*7 + 8

	// shifted = input >> shift
	shifted := new(big.Int).Rsh(input, shift)

	value := new(big.Int).And(shifted, maskMaxOrderToken)

	orderBits := new(big.Int).Rsh(value, 3)
	result.OrderIndex = int(new(big.Int).And(orderBits, maskMaxOrder).Int64())
	result.TokenIndex = int(value.And(value, maskMaxToken).Int64())

	return
}
