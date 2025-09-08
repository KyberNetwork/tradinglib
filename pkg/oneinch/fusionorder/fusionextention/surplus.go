package fusionextention

import (
	"math/big"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/decode"
)

type SurplusParam struct {
	EstimatedTakerAmount *big.Int `json:"estimatedTakerAmount"`
	ProtocolFee          int64    `json:"protocolFee"`
}

func (s SurplusParam) IsZero() bool {
	return s.ProtocolFee == 0
}

// https://github.com/1inch/fusion-sdk/blob/bcdb9d67f5c528bbbeb20d6a6b1355e3fcafcb2a/src/fusion-order/surplus-params.ts#L6
func GetNoFeeSurplusParam() SurplusParam {
	maxUint256 := new(big.Int).Sub(
		new(big.Int).Lsh(big.NewInt(1), 256),
		big.NewInt(1),
	)
	return SurplusParam{
		EstimatedTakerAmount: maxUint256,
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
		EstimatedTakerAmount: estimatedTakerAmount,
		ProtocolFee:          int64(protocolFee * 100),
	}, nil
}
