package fusionextention

import (
	"encoding/json"
	"math/big"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/decode"
)

// nolint:recvcheck
type SurplusParam struct {
	EstimatedTakerAmount *big.Int `json:"estimated_taker_amount"`
	ProtocolFee          int64    `json:"protocol_fee"`
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

func (s SurplusParam) MarshalJSON() ([]byte, error) {
	type surplusParam struct {
		EstimatedTakerAmount string `json:"estimated_taker_amount"`
		ProtocolFee          int64  `json:"protocol_fee"`
	}
	data := surplusParam{
		ProtocolFee: s.ProtocolFee,
	}
	if s.EstimatedTakerAmount != nil {
		data.EstimatedTakerAmount = s.EstimatedTakerAmount.String()
	}
	return json.Marshal(data)
}

func (s *SurplusParam) UnmarshalJSON(data []byte) error {
	type surplusParam struct {
		EstimatedTakerAmount string `json:"estimated_taker_amount"`
		ProtocolFee          int64  `json:"protocol_fee"`
	}

	var dto surplusParam
	if err := json.Unmarshal(data, &dto); err != nil {
		return err
	}

	s.ProtocolFee = dto.ProtocolFee
	if dto.EstimatedTakerAmount != "" {
		x, ok := new(big.Int).SetString(dto.EstimatedTakerAmount, 10)
		if !ok {
			return nil
		}
		s.EstimatedTakerAmount = x
	}
	return nil
}
