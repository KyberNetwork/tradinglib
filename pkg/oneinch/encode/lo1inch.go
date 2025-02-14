package encode

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/lo1inch"
	helper1inch "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/lo1inch/helper"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/util/bignumber"
	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/valueobject"
	"github.com/ethereum/go-ethereum/common"
)

// https://github.com/KyberNetwork/aggregator-encoding/blob/v0.37.6/pkg/encode/l1encode/executor/swapdata/lo1inch.go#L19
func PackLO1inch(_ valueobject.ChainID, encodingSwap EncodingSwap) ([][]byte, error) { //nolint:funlen,cyclop
	// get contract address for LO.
	if encodingSwap.PoolExtra == nil {
		return nil, fmt.Errorf("[PackLO1inch] PoolExtra is nil")
	}

	byteData, err := json.Marshal(encodingSwap.Extra)
	if err != nil {
		return nil, fmt.Errorf("[buildLO1inch] ErrMarshalFailed err :[%w]", err)
	}

	var swapInfo lo1inch.SwapInfo
	if err = json.Unmarshal(byteData, &swapInfo); err != nil {
		return nil, fmt.Errorf("[buildLO1inch] ErrUnmarshalFailed err :[%w]", err)
	}

	if len(swapInfo.FilledOrders) == 0 {
		return nil, fmt.Errorf("[buildLO1inch] cause by filledOrders is empty")
	}

	encodeds := make([][]byte, 0, len(swapInfo.FilledOrders))

	amountIn := new(big.Int).Set(encodingSwap.SwapAmount)
	for _, filledOrder := range swapInfo.FilledOrders {
		if amountIn.Sign() <= 0 {
			break
		}

		takingAmount := filledOrder.TakingAmount.ToBig()
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
			return nil, fmt.Errorf("decode extension: %w", err)
		}

		receiver := helper1inch.NewAddress(encodingSwap.Recipient)

		// In the orders response of 1inch limit order API, there is no interaction
		// so we only encode receiver + extension into takerTraits
		takerTraitsEncoded := helper1inch.NewTakerTraits(&receiver, extension, nil).Encode()

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
					Order calldata order,
					bytes calldata signature,
					uint256 amount,
					TakerTraits takerTraits,
					bytes calldata args
				) external returns(uint256 makingAmount, uint256 takingAmount, bytes32 orderHash);
			*/
			packed, err := OneInchAggregationRouterV6ABI.Pack("fillContractOrderArgs",
				order, signature, filledOrder, takingAmount, takerTraitsEncoded.TakerTraits, takerTraitsEncoded.Args,
			)
			if err != nil {
				return nil, fmt.Errorf("pack fillContractOrderArgs error: %w", err)
			}
			encodeds = append(encodeds, packed)

		case false:
			bytesSignature, err := helper1inch.LO1inchParseSignature(filledOrder.Signature)
			if err != nil {
				return nil, fmt.Errorf("parse lo1inch sig: %w", err)
			}
			r := bytesSignature.R
			vs := bytesSignature.GetCompactedSignatureBytes()[len(r):]

			/*
				function fillOrderArgs(
					Order calldata order,
					bytes32 r,
					bytes32 vs,
					uint256 amount,
					uint256 takerTraits,
					bytes calldata args
				)
				external payable returns (uint256 makingAmount, uint256 takingAmount, bytes32 orderHash);
			*/
			packed, err := OneInchAggregationRouterV6ABI.Pack("fillOrderArgs",
				order, r, vs, takingAmount, takerTraitsEncoded.TakerTraits, takerTraitsEncoded.Args,
			)
			if err != nil {
				return nil, fmt.Errorf("pack fillOrderArgs error: %w", err)
			}
			encodeds = append(encodeds, packed)
		}
	}

	return encodeds, nil
}
