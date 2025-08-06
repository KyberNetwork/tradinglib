package encode

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/KyberNetwork/blockchain-toolkit/number"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/lo1inch"
	helper1inch "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/lo1inch/helper"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/util/bignumber"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/valueobject"
	"github.com/ethereum/go-ethereum/common"
)

const (
	MethodIDLength              = 4
	MethodFillOrderArgs         = "fillOrderArgs"
	MethodFillContractOrderArgs = "fillContractOrderArgs"
)

// https://github.com/KyberNetwork/aggregator-encoding/blob/v0.37.6/pkg/encode/l1encode/executor/swapdata/lo1inch.go#L19
func PackLO1inch(_ valueobject.ChainID, encodingSwap EncodingSwap) ([][]byte, *big.Int, error) { //nolint:funlen,cyclop
	// get contract address for LO.
	if encodingSwap.PoolExtra == nil {
		return nil, nil, fmt.Errorf("[PackLO1inch] PoolExtra is nil")
	}

	byteData, err := json.Marshal(encodingSwap.Extra)
	if err != nil {
		return nil, nil, fmt.Errorf("[buildLO1inch] ErrMarshalFailed err :[%w]", err)
	}

	var swapInfo lo1inch.SwapInfo
	if err = json.Unmarshal(byteData, &swapInfo); err != nil {
		return nil, nil, fmt.Errorf("[buildLO1inch] ErrUnmarshalFailed err :[%w]", err)
	}

	poolExtraBytes, err := json.Marshal(encodingSwap.PoolExtra)
	if err != nil {
		return nil, nil, fmt.Errorf("[buildLO1inch] ErrMarshalFailed err :[%w]", err)
	}

	var poolMeta lo1inch.MetaInfo
	if err = json.Unmarshal(poolExtraBytes, &poolMeta); err != nil {
		return nil, nil, fmt.Errorf("[buildLO1inch] ErrUnmarshalFailed err :[%w]", err)
	}

	if len(swapInfo.FilledOrders) == 0 {
		return nil, nil, fmt.Errorf("[buildLO1inch] cause by filledOrders is empty")
	}

	encodeds := make([][]byte, 0, len(swapInfo.FilledOrders))

	amountIn := new(big.Int).Set(encodingSwap.SwapAmount)
	for _, filledOrder := range swapInfo.FilledOrders {
		if amountIn.Sign() <= 0 {
			break
		}

		// calculate order's remaining taking amount
		// orderRemainingTakingAmount = order.TakingAmount * orderRemainingMakingAmount / order.MakingAmount
		orderRemainingTakingAmount := number.Set(filledOrder.TakingAmount)
		orderRemainingTakingAmount.Mul(orderRemainingTakingAmount, filledOrder.RemainingMakerAmount)
		orderRemainingTakingAmount.Div(orderRemainingTakingAmount, filledOrder.MakingAmount)

		takingAmount := orderRemainingTakingAmount.ToBig()
		if takingAmount.Sign() == 0 {
			continue
		}
		switch amountIn.Cmp(takingAmount) {
		case -1:
			takingAmount.Set(amountIn)
			amountIn.SetInt64(0)

		case 0:
			amountIn.SetInt64(0)

		case 1:
			amountIn.Sub(amountIn, takingAmount)
		}

		extension, err := helper1inch.DecodeExtension(filledOrder.Extension)
		if err != nil {
			return nil, nil, fmt.Errorf("decode extension: %w", err)
		}

		receiver := common.HexToAddress(encodingSwap.Recipient)

		// init interaction to check min amount out returned.
		interaction := helper1inch.Interaction{
			Target: common.HexToAddress(poolMeta.TakerTargetInteraction),
			Data:   filledOrder.MakingAmount.Bytes(),
		}

		takerTraitsEncoded, args := helper1inch.NewTakerTraits(big.NewInt(0), &receiver, &extension, &interaction).Encode()

		order := OneInchV6Order{
			Salt:         bignumber.NewBig(filledOrder.Salt),
			Maker:        bignumber.NewBig(filledOrder.Maker),
			Receiver:     bignumber.NewBig(filledOrder.Receiver),
			MakerAsset:   bignumber.NewBig(filledOrder.MakerAsset),
			TakerAsset:   bignumber.NewBig(filledOrder.TakerAsset),
			MakingAmount: filledOrder.MakingAmount.ToBig(),
			TakingAmount: filledOrder.TakingAmount.ToBig(),
			MakerTraits:  bignumber.NewBig(filledOrder.MakerTraits),
		}

		switch filledOrder.IsMakerContract {
		case true:
			signature := common.FromHex(filledOrder.Signature)

			/*
				function fillContractOrderArgs(
					LimitOrder calldata order,
					bytes calldata signature,
					uint256 amount,
					TakerTraits takerTraits,
					bytes calldata args
				) external returns(uint256 makingAmount, uint256 takingAmount, bytes32 orderHash);
			*/
			packed, err := OneInchAggregationRouterV6ABI.Pack("fillContractOrderArgs",
				order, signature, filledOrder, takingAmount, takerTraitsEncoded, args,
			)
			if err != nil {
				return nil, nil, fmt.Errorf("pack fillContractOrderArgs error: %w", err)
			}
			encodeds = append(encodeds, packed)

		case false:
			bytesSignature, err := helper1inch.LO1inchParseSignature(filledOrder.Signature)
			if err != nil {
				return nil, nil, fmt.Errorf("parse lo1inch sig: %w", err)
			}
			r := bytesSignature.R
			vs := bytesSignature.GetCompactedSignatureBytes()[len(r):]

			/*
				function fillOrderArgs(
					LimitOrder calldata order,
					bytes32 r,
					bytes32 vs,
					uint256 amount,
					uint256 takerTraits,
					bytes calldata args
				)
				external payable returns (uint256 makingAmount, uint256 takingAmount, bytes32 orderHash);
			*/

			var rArray, vsArray [32]byte
			copy(rArray[:], r)
			copy(vsArray[:], vs)
			packed, err := OneInchAggregationRouterV6ABI.Pack("fillOrderArgs",
				order, rArray, vsArray, takingAmount, takerTraitsEncoded, args,
			)
			if err != nil {
				return nil, nil, fmt.Errorf("pack fillOrderArgs error: %w", err)
			}
			encodeds = append(encodeds, packed)
		}
	}

	return encodeds, amountIn, nil
}

func UnpackLO1inch(decoded []byte) (any, error) {
	if len(decoded) < MethodIDLength {
		return nil, fmt.Errorf("invalid data length: %d", len(decoded))
	}

	methodID := decoded[:MethodIDLength]
	method, err := OneInchAggregationRouterV6ABI.MethodById(methodID)
	if err != nil {
		return nil, fmt.Errorf("get method: %w", err)
	}

	switch method.Name {
	case MethodFillOrderArgs:
		return UnpackFillOrderArgs(decoded)

	case MethodFillContractOrderArgs:
		return UnpackFillContractOrderArgs(decoded)

	default:
		return nil, fmt.Errorf("method not support: %s", method.Name)
	}
}

func UnpackFillOrderArgs(decoded []byte) (FillOrderArgs, error) {
	if len(decoded) < MethodIDLength {
		return FillOrderArgs{}, fmt.Errorf("invalid data length: %d", len(decoded))
	}

	method, err := OneInchAggregationRouterV6ABI.MethodById(decoded[:MethodIDLength])
	if err != nil {
		return FillOrderArgs{}, fmt.Errorf("get method: %w", err)
	}

	if method.Name != MethodFillOrderArgs {
		return FillOrderArgs{}, fmt.Errorf("invalid method: %s", method.Name)
	}

	unpacked, err := method.Inputs.Unpack(decoded[MethodIDLength:])
	if err != nil {
		return FillOrderArgs{}, fmt.Errorf("unpack fillOrderArgs: %w", err)
	}

	var args FillOrderArgs
	if err := method.Inputs.Copy(&args, unpacked); err != nil {
		return FillOrderArgs{}, fmt.Errorf("copy FillOrderArgs: %w", err)
	}

	return args, nil
}

func UnpackFillContractOrderArgs(decoded []byte) (FillContractOrderArgs, error) {
	if len(decoded) < MethodIDLength {
		return FillContractOrderArgs{}, fmt.Errorf("invalid data length: %d", len(decoded))
	}

	method, err := OneInchAggregationRouterV6ABI.MethodById(decoded[:MethodIDLength])
	if err != nil {
		return FillContractOrderArgs{}, fmt.Errorf("get method: %w", err)
	}

	if method.Name != MethodFillContractOrderArgs {
		return FillContractOrderArgs{}, fmt.Errorf("invalid method: %s", method.Name)
	}

	unpacked, err := method.Inputs.Unpack(decoded[MethodIDLength:])
	if err != nil {
		return FillContractOrderArgs{}, fmt.Errorf("unpack fillContractOrderArgs: %w", err)
	}

	var args FillContractOrderArgs
	if err := method.Inputs.Copy(&args, unpacked); err != nil {
		return FillContractOrderArgs{}, fmt.Errorf("copy FillOrderArgs: %w", err)
	}

	return args, nil
}
