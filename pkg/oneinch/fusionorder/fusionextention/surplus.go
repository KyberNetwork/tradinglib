package fusionextention

import (
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/decode"
	"math/big"
)

type SurplusParam struct {
	estimatedTakerAmount *big.Int
	protocolFee          int64
}

func (s SurplusParam) IsZero() bool {
	return s.protocolFee == 0
}

// https://github.com/1inch/fusion-sdk/blob/bcdb9d67f5c528bbbeb20d6a6b1355e3fcafcb2a/src/fusion-order/surplus-params.ts#L6
func GetNoFeeSurplusParam() SurplusParam {
	maxUint256 := new(big.Int).Sub(
		new(big.Int).Lsh(big.NewInt(1), 256),
		big.NewInt(1),
	)
	return SurplusParam{
		estimatedTakerAmount: maxUint256,
	}
}

func DecodeSurplusParam(iter *decode.BytesIterator) (SurplusParam, error) {
	estimatedTakerAmount, err := iter.NextUint256()
	if err != nil {
		return SurplusParam{}, err
	}

	protocolFee, err := iter.NextUint8()
	if err != nil {
		return SurplusParam{}, err
	}

	return SurplusParam{
		estimatedTakerAmount: estimatedTakerAmount,
		protocolFee:          int64(protocolFee * 100),
	}, nil
}
