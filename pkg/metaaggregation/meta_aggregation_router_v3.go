package metaaggregation

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// ErrNotSingleOutput is returned for KSAggregationRouterV3 swaps that do not have exactly one output token.
var ErrNotSingleOutput = errors.New("swap must have exactly one output token")

const SwapV3 InputsTypes = "swapV3"

// SwapV3InputTokenData mirrors IKSAggregationRouterV3.InputTokenData.
type SwapV3InputTokenData struct {
	PermitData    []byte           `json:"permitData"`
	FeeRecipients []common.Address `json:"feeRecipients"`
	Fees          []*big.Int       `json:"fees"`
	Targets       []common.Address `json:"targets"`
	Amounts       []*big.Int       `json:"amounts"`
}

// SwapV3OutputTokenData mirrors IKSAggregationRouterV3.OutputTokenData.
type SwapV3OutputTokenData struct {
	MinAmount     *big.Int         `json:"minAmount"`
	FeeRecipients []common.Address `json:"feeRecipients"`
	Fees          []*big.Int       `json:"fees"`
}

// SwapV3Params mirrors IKSAggregationRouterV3.SwapParams.
type SwapV3Params struct {
	InputTokens  []common.Address        `json:"inputTokens"`
	InputAmounts []*big.Int              `json:"inputAmounts"`
	InputData    []SwapV3InputTokenData  `json:"inputData"`
	OutputTokens []common.Address        `json:"outputTokens"`
	OutputData   []SwapV3OutputTokenData `json:"outputData"`
	Permit2Data  []byte                  `json:"permit2Data"`
	Executor     common.Address          `json:"executor"`
	ExecutorData []byte                  `json:"executorData"`
	Recipient    common.Address          `json:"recipient"`
	Deadline     *big.Int                `json:"deadline"`
	ClientData   []byte                  `json:"clientData"`
}

// metaAggregationSwapV3Inputs holds a KSAggregationRouterV3 swap with exactly one output token.
type metaAggregationSwapV3Inputs struct {
	SwapV3Params
}

func (m *metaAggregationSwapV3Inputs) GetMinReturnAmount() *big.Int {
	return new(big.Int).Set(m.OutputData[0].MinAmount)
}

func (m *metaAggregationSwapV3Inputs) SetMinReturnAmount(amount *big.Int) {
	m.OutputData[0].MinAmount = amount
}

func (m *metaAggregationSwapV3Inputs) Type() InputsTypes {
	return SwapV3
}

func isSwapV3Calldata(data []byte) bool {
	selector := RouterV3ABI.Methods[methodSwap].ID
	return len(data) >= len(selector) && bytes.Equal(data[:len(selector)], selector)
}

func decodeSwapV3Inputs(data []byte) (MetaAggregationRouterSwapInputs, error) {
	params, err := UnpackSwapV3Inputs(data)
	if err != nil {
		return nil, err
	}
	return &metaAggregationSwapV3Inputs{SwapV3Params: params}, nil
}

// UnpackSwapV3Inputs decodes KSAggregationRouterV3 swap calldata and rejects swaps
// that do not have exactly one output token.
func UnpackSwapV3Inputs(data []byte) (SwapV3Params, error) {
	if !isSwapV3Calldata(data) {
		return SwapV3Params{}, ErrInvalidMethodName
	}

	method := RouterV3ABI.Methods[methodSwap]
	unpacked, err := method.Inputs.Unpack(data[len(method.ID):])
	if err != nil {
		return SwapV3Params{}, fmt.Errorf("unpack swap params: %w", err)
	}

	// A method whose only input is a tuple is copied atomically, so it needs a containing struct.
	var call struct {
		Params SwapV3Params
	}
	if err = method.Inputs.Copy(&call, unpacked); err != nil {
		return SwapV3Params{}, fmt.Errorf("copy swap params: %w", err)
	}

	if err = validateSingleOutput(call.Params); err != nil {
		return SwapV3Params{}, err
	}
	return call.Params, nil
}

func packSwapV3Inputs(params SwapV3Params) ([]byte, error) {
	if err := validateSingleOutput(params); err != nil {
		return nil, err
	}
	return RouterV3ABI.Pack(methodSwap, params)
}

func validateSingleOutput(params SwapV3Params) error {
	if len(params.OutputTokens) != 1 || len(params.OutputData) != 1 {
		return fmt.Errorf("%w: got %d output tokens and %d output data",
			ErrNotSingleOutput, len(params.OutputTokens), len(params.OutputData))
	}
	if params.OutputData[0].MinAmount == nil {
		return fmt.Errorf("%w: missing min amount", ErrNotSingleOutput)
	}
	return nil
}

// DecodeMetaAggregationSwapV3Output decodes the return data of a KSAggregationRouterV3 swap
// with exactly one output token.
func DecodeMetaAggregationSwapV3Output(data []byte) (*big.Int, uint64, error) {
	var res struct {
		OutputAmounts []*big.Int
		GasUsed       *big.Int
	}
	if err := RouterV3ABI.UnpackIntoInterface(&res, methodSwap, data); err != nil {
		return nil, 0, fmt.Errorf("unpack data: %w", err)
	}

	if len(res.OutputAmounts) != 1 {
		return nil, 0, fmt.Errorf("%w: got %d output amounts", ErrNotSingleOutput, len(res.OutputAmounts))
	}
	if res.GasUsed.Sign() == 0 {
		return nil, 0, errors.New("gas is zero")
	}

	return res.OutputAmounts[0], res.GasUsed.Uint64(), nil
}
