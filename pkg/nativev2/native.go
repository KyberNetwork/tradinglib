package nativev2

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type WidgetFee struct {
	FeeRecipient common.Address
	FeeRate      *big.Int
}
type RFQTQuote struct {
	Pool                      common.Address
	Signer                    common.Address
	Recipient                 common.Address
	SellerToken               common.Address
	BuyerToken                common.Address
	SellerTokenAmount         *big.Int
	BuyerTokenAmount          *big.Int
	AmountOutMinimum          *big.Int
	DeadlineTimestamp         *big.Int
	Nonce                     *big.Int
	ConfidenceExtractedValueT *big.Int
	ConfidenceExtractedValueN *big.Int
	ConfidenceExtractedValueE *big.Int
	ConfidenceExtractedValueM *big.Int
	QuoteId                   [16]byte
	MultiHop                  bool
	Signature                 []byte
	WidgetFee                 WidgetFee
	WidgetFeeSignature        []byte
}

type TradeRFQTArguments struct {
	Quote                 RFQTQuote
	ActualSellerAmount    *big.Int
	ActualMinOutputAmount *big.Int
}

func DecodeTradeRFQT(data []byte) (TradeRFQTArguments, error) {
	if len(data) < 4 {
		return TradeRFQTArguments{}, fmt.Errorf("invalid data length: %d", len(data))
	}
	inputs, err := nativeABI.Methods[methodTradeRFQT].Inputs.Unpack(data[4:])
	if err != nil {
		return TradeRFQTArguments{}, err
	}

	var args TradeRFQTArguments
	if err = nativeABI.Methods[methodTradeRFQT].Inputs.Copy(&args, inputs); err != nil {
		return TradeRFQTArguments{}, err
	}

	return args, nil
}
