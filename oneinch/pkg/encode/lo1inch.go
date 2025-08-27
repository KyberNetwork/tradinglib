package encode

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/lo1inch"
	helper1inch "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/lo1inch/helper"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/util/bignumber"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/valueobject"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/limitorder"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

const (
	MethodIDLength              = 4
	MethodFillOrderArgs         = "fillOrderArgs"
	MethodFillContractOrderArgs = "fillContractOrderArgs"
)

//nolint:gochecknoglobals
var reserveFundCallbackArguments = abi.Arguments{
	{Name: "minMakingAmount", Type: Uint256},
}

type Interaction struct {
	Target   common.Address `json:"target"`
	CallData []byte         `json:"call_data"`
}

func EncodeReserveFundCallback(minMakingAmount *big.Int) ([]byte, error) {
	return reserveFundCallbackArguments.Pack(minMakingAmount)
}

func DecodeReserveFundCallback(data []byte) (*big.Int, error) {
	unpacked, err := reserveFundCallbackArguments.Unpack(data)
	if err != nil {
		return nil, fmt.Errorf("unpack reserve fund callback: %w", err)
	}

	if len(unpacked) == 0 {
		return nil, fmt.Errorf("no data unpacked from reserve fund callback")
	}

	minMakingAmount, ok := unpacked[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("failed to convert unpacked data to *big.Int")
	}

	return minMakingAmount, nil
}

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

		orderRemainingTakingAmount := filledOrder.FilledTakingAmount
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

		// minAmountOutWei = max(0, filledOrder.FilledMakingAmount - 1)
		// we sub 1 wei for rounding issue prevent
		minAmountOutWei := new(big.Int).Set(filledOrder.FilledMakingAmount.ToBig())
		minAmountOutWei.Sub(minAmountOutWei, bignumber.One)
		if minAmountOutWei.Sign() < 0 {
			minAmountOutWei = bignumber.ZeroBI
		}

		encodedReserveFundCallback, err := EncodeReserveFundCallback(minAmountOutWei)
		if err != nil {
			return nil, nil, fmt.Errorf("encode reserve fund callback: %w", err)
		}

		// init interaction to check min amount out returned.
		interaction := helper1inch.Interaction{
			Target: common.HexToAddress(poolMeta.TakerTargetInteraction),
			Data:   encodedReserveFundCallback,
		}

		takerTraitsEncoded, args := helper1inch.NewTakerTraits(big.NewInt(0), &receiver, &extension, &interaction).
			// as we are taker side, we need check makingAmount >= threshold
			SetAmountThreshold(minAmountOutWei).
			Encode()

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
				order, signature, takingAmount, takerTraitsEncoded, args,
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

type UnpackLO1inchResult struct {
	FillOrderArgs              *FillOrderArgs
	FillContractOrderArgs      *FillContractOrderArgs
	TakerTraitsIsMakingAmount  bool
	AmountThreshold            *big.Int
	Receiver                   *common.Address
	Extension                  *limitorder.Extension
	InteractionTarget          *common.Address
	InteractionMinMakingAmount *big.Int
}

// UnpackLO1inch
// nolint:funlen,cyclop
func UnpackLO1inch(decoded []byte) (*UnpackLO1inchResult, error) {
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
		fillOrderArgs, err := UnpackFillOrderArgs(decoded)
		if err != nil {
			return nil, err
		}

		takerTraits := helper1inch.NewTakerTraits(
			fillOrderArgs.TakerTraits,
			nil,
			nil,
			nil,
		)

		receiver, extension, interaction, err := limitorder.DecodeArgs(
			fillOrderArgs.TakerTraits,
			fillOrderArgs.Args,
		)
		if err != nil {
			return nil, err
		}

		minMakingAmount, err := DecodeReserveFundCallback(interaction.Data)
		if err != nil {
			return nil, err
		}

		if interaction == nil {
			return nil, ErrInteractionIsNil
		}

		return &UnpackLO1inchResult{
			FillOrderArgs:              &fillOrderArgs,
			AmountThreshold:            takerTraits.AmountThreshold(),
			Receiver:                   receiver,
			Extension:                  extension,
			InteractionTarget:          &interaction.Target,
			InteractionMinMakingAmount: minMakingAmount,
			// expect to be false as VT is the taker side
			TakerTraitsIsMakingAmount: takerTraits.IsMakingAmount(),
		}, nil

	case MethodFillContractOrderArgs:
		fillContractOrderArgs, err := UnpackFillContractOrderArgs(decoded)
		if err != nil {
			return nil, err
		}

		takerTraits := helper1inch.NewTakerTraits(
			fillContractOrderArgs.TakerTraits,
			nil,
			nil,
			nil,
		)

		receiver, extension, interaction, err := limitorder.DecodeArgs(
			fillContractOrderArgs.TakerTraits,
			fillContractOrderArgs.Args,
		)
		if err != nil {
			return nil, err
		}

		minMakingAmount, err := DecodeReserveFundCallback(interaction.Data)
		if err != nil {
			return nil, err
		}

		if interaction == nil {
			return nil, ErrInteractionIsNil
		}

		return &UnpackLO1inchResult{
			FillContractOrderArgs:      &fillContractOrderArgs,
			AmountThreshold:            takerTraits.AmountThreshold(),
			Receiver:                   receiver,
			Extension:                  extension,
			InteractionTarget:          &interaction.Target,
			InteractionMinMakingAmount: minMakingAmount,
			// expect to be false as VT is the taker side
			TakerTraitsIsMakingAmount: takerTraits.IsMakingAmount(),
		}, nil

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
