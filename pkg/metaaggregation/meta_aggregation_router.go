package metaaggregation

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/KyberNetwork/aggregator-encoding/pkg/abis"
	"github.com/KyberNetwork/aggregator-encoding/pkg/encode/v3/router"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

var (
	ErrInvalidMethodName = errors.New("invalid method name")
	ErrUnsupportedType   = errors.New("unsupported type")
)

type InputsTypes string

// nolint:staticcheck
const (
	Swap       InputsTypes = "swap"
	methodSwap             = "swap"
)

type MetaAggregationRouterSwapInputs interface {
	GetMinReturnAmount() *big.Int
	SetMinReturnAmount(amount *big.Int)
	Type() InputsTypes
}
type metaAggregationSwapInputs struct {
	router.SwapInputs
}

func (m *metaAggregationSwapInputs) GetMinReturnAmount() *big.Int {
	return new(big.Int).Set(m.Execution.Desc.MinReturnAmount)
}

func (m *metaAggregationSwapInputs) SetMinReturnAmount(amount *big.Int) {
	m.Execution.Desc.MinReturnAmount = amount
}

func (m *metaAggregationSwapInputs) Type() InputsTypes {
	return Swap
}

func DecodeMetaAggregationRouterSwapData(data []byte) (MetaAggregationRouterSwapInputs, error) {
	method, err := abis.MetaAggregationRouterV2.MethodById(data)
	if err != nil {
		return nil, err
	}

	switch method.Name {
	case methodSwap:
		swapInputs, err := UnpackSwapInputs(data)
		if err != nil {
			return nil, err
		}
		input := &metaAggregationSwapInputs{}
		input.SwapInputs = swapInputs
		return input, nil
	default:
		return nil, ErrInvalidMethodName
	}
}

// nolint: forcetypeassert
func EncodeMetaAggregationRouterSwapData(
	input MetaAggregationRouterSwapInputs,
) ([]byte, error) {
	switch input.Type() {
	case Swap:
		return router.PackSwapInputs(input.(*metaAggregationSwapInputs).SwapInputs)
	default:
		return nil, ErrUnsupportedType
	}
}

func UnpackSwapInputs(data []byte) (router.SwapInputs, error) {
	method, err := abis.MetaAggregationRouterV2.MethodById(data)
	if err != nil {
		return router.SwapInputs{}, err
	}

	if len(data) < 4 {
		return router.SwapInputs{}, fmt.Errorf("invalid data length: %d", len(data))
	}

	unpacked, err := method.Inputs.Unpack(data[4:])
	if err != nil {
		return router.SwapInputs{}, err
	}

	var inputs router.SwapInputs
	if err = method.Inputs.Copy(&inputs, unpacked); err != nil {
		return router.SwapInputs{}, err
	}

	return inputs, nil
}

func DecodeMetaAggregationSwapOutput(data []byte) (*big.Int, uint64, error) {
	var res [2]interface{}
	err := RouterV2ABI.UnpackIntoInterface(&res, methodSwap, data)
	if err != nil {
		return nil, 0, fmt.Errorf("unpack data: %w", err)
	}

	returnAmount := *abi.ConvertType(res[0], new(*big.Int)).(**big.Int)   //nolint:forcetypeassert
	gas := (*abi.ConvertType(res[1], new(*big.Int)).(**big.Int)).Uint64() //nolint:forcetypeassert
	if gas == 0 {
		return nil, 0, errors.New("gas is zero")
	}

	return returnAmount, gas, nil
}
