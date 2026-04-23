//nolint:testpackage
package nativev2

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// quoteInput mirrors the RFQTQuote ABI tuple layout for packing purposes.
type quoteInput struct {
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
	WidgetFee                 widgetFeeInput
	WidgetFeeSignature        []byte
}

type widgetFeeInput struct {
	FeeRecipient common.Address
	FeeRate      *big.Int
}

func packTradeRFQT(t *testing.T, quote quoteInput, actualSellerAmount, actualMinOutputAmount *big.Int) []byte {
	t.Helper()
	packed, err := nativeABI.Methods[methodTradeRFQT].Inputs.Pack(quote, actualSellerAmount, actualMinOutputAmount)
	require.NoError(t, err)
	return append(nativeABI.Methods[methodTradeRFQT].ID, packed...)
}

// nolint: gocritic
func TestDecodeTradeRFQT(t *testing.T) {
	t.Run("decodes all fields correctly", func(t *testing.T) {
		quote := quoteInput{
			Pool:                      common.HexToAddress("0x1111111111111111111111111111111111111111"),
			Signer:                    common.HexToAddress("0x2222222222222222222222222222222222222222"),
			Recipient:                 common.HexToAddress("0x3333333333333333333333333333333333333333"),
			SellerToken:               common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
			BuyerToken:                common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
			SellerTokenAmount:         big.NewInt(1e18),
			BuyerTokenAmount:          big.NewInt(3000e6),
			AmountOutMinimum:          big.NewInt(2990e6),
			DeadlineTimestamp:         big.NewInt(1700000000),
			Nonce:                     big.NewInt(42),
			ConfidenceExtractedValueT: big.NewInt(100),
			ConfidenceExtractedValueN: big.NewInt(200),
			ConfidenceExtractedValueE: big.NewInt(300),
			ConfidenceExtractedValueM: big.NewInt(400),
			QuoteId:                   [16]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10},
			MultiHop:                  false,
			Signature:                 []byte{0xde, 0xad, 0xbe, 0xef},
			WidgetFee: widgetFeeInput{
				FeeRecipient: common.HexToAddress("0x4444444444444444444444444444444444444444"),
				FeeRate:      big.NewInt(50),
			},
			WidgetFeeSignature: []byte{0xca, 0xfe, 0xba, 0xbe},
		}
		actualSellerAmount := big.NewInt(1e18)
		actualMinOutputAmount := big.NewInt(2985e6)

		args, err := DecodeTradeRFQT(packTradeRFQT(t, quote, actualSellerAmount, actualMinOutputAmount))
		require.NoError(t, err)

		assert.Equal(t, quote.Pool, args.Quote.Pool)
		assert.Equal(t, quote.Signer, args.Quote.Signer)
		assert.Equal(t, quote.Recipient, args.Quote.Recipient)
		assert.Equal(t, quote.SellerToken, args.Quote.SellerToken)
		assert.Equal(t, quote.BuyerToken, args.Quote.BuyerToken)
		assert.Equal(t, quote.SellerTokenAmount, args.Quote.SellerTokenAmount)
		assert.Equal(t, quote.BuyerTokenAmount, args.Quote.BuyerTokenAmount)
		assert.Equal(t, quote.AmountOutMinimum, args.Quote.AmountOutMinimum)
		assert.Equal(t, quote.DeadlineTimestamp, args.Quote.DeadlineTimestamp)
		assert.Equal(t, quote.Nonce, args.Quote.Nonce)
		assert.Equal(t, quote.ConfidenceExtractedValueT, args.Quote.ConfidenceExtractedValueT)
		assert.Equal(t, quote.ConfidenceExtractedValueN, args.Quote.ConfidenceExtractedValueN)
		assert.Equal(t, quote.ConfidenceExtractedValueE, args.Quote.ConfidenceExtractedValueE)
		assert.Equal(t, quote.ConfidenceExtractedValueM, args.Quote.ConfidenceExtractedValueM)
		assert.Equal(t, quote.QuoteId, args.Quote.QuoteId)
		assert.Equal(t, quote.MultiHop, args.Quote.MultiHop)
		assert.Equal(t, quote.Signature, args.Quote.Signature)
		assert.Equal(t, quote.WidgetFee.FeeRecipient, args.Quote.WidgetFee.FeeRecipient)
		assert.Equal(t, quote.WidgetFee.FeeRate, args.Quote.WidgetFee.FeeRate)
		assert.Equal(t, quote.WidgetFeeSignature, args.Quote.WidgetFeeSignature)
		assert.Equal(t, actualSellerAmount, args.ActualSellerAmount)
		assert.Equal(t, actualMinOutputAmount, args.ActualMinOutputAmount)
	})

	t.Run("multihop true", func(t *testing.T) {
		quote := quoteInput{
			Pool:                      common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"),
			Signer:                    common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"),
			Recipient:                 common.HexToAddress("0xCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC"),
			SellerToken:               common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
			BuyerToken:                common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"),
			SellerTokenAmount:         big.NewInt(5000e6),
			BuyerTokenAmount:          new(big.Int).Mul(big.NewInt(5000), new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)),
			AmountOutMinimum:          new(big.Int).Mul(big.NewInt(4990), new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)),
			DeadlineTimestamp:         big.NewInt(1800000000),
			Nonce:                     big.NewInt(99),
			ConfidenceExtractedValueT: big.NewInt(0),
			ConfidenceExtractedValueN: big.NewInt(0),
			ConfidenceExtractedValueE: big.NewInt(0),
			ConfidenceExtractedValueM: big.NewInt(0),
			QuoteId:                   [16]byte{0xff, 0xee, 0xdd, 0xcc, 0xbb, 0xaa, 0x99, 0x88, 0x77, 0x66, 0x55, 0x44, 0x33, 0x22, 0x11, 0x00},
			MultiHop:                  true,
			Signature:                 []byte{0x01, 0x02, 0x03},
			WidgetFee: widgetFeeInput{
				FeeRecipient: common.HexToAddress("0x0000000000000000000000000000000000000000"),
				FeeRate:      big.NewInt(0),
			},
			WidgetFeeSignature: []byte{},
		}
		actualSellerAmount := big.NewInt(5000e6)
		actualMinOutputAmount := new(big.Int).Mul(big.NewInt(4980), new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))

		args, err := DecodeTradeRFQT(packTradeRFQT(t, quote, actualSellerAmount, actualMinOutputAmount))
		require.NoError(t, err)

		assert.True(t, args.Quote.MultiHop)
		assert.Equal(t, quote.QuoteId, args.Quote.QuoteId)
		assert.Equal(t, actualSellerAmount, args.ActualSellerAmount)
		assert.Equal(t, actualMinOutputAmount, args.ActualMinOutputAmount)
	})

	t.Run("only selector no payload", func(t *testing.T) {
		data := nativeABI.Methods[methodTradeRFQT].ID
		_, err := DecodeTradeRFQT(data)
		require.Error(t, err)
	})

	t.Run("garbage payload after selector", func(t *testing.T) {
		data := append(nativeABI.Methods[methodTradeRFQT].ID, []byte{0xde, 0xad, 0xbe, 0xef}...)
		_, err := DecodeTradeRFQT(data)
		require.Error(t, err)
	})
}

func TestDecodeTradeRFQTFromRawData(t *testing.T) {
	testCases := []struct {
		name    string
		rawData string
	}{
		{
			name:    "normal case",
			rawData: "0x0947c2d90000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005d1a34369686ae59ac97ae4e1df5635ffda9ee7c000000000000000000000000129b3d9a0a6e4beab88f5cb1e57995d72a6e24f100000000000000000000000063242a4ea82847b20e506b63b0e2e2eff0cc6cb0000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec70000000000000000000000000000000000000000000000000de0b6b3a7640000000000000000000000000000000000000000000000000000000000008c1facfb000000000000000000000000000000000000000000000000000000008c1facfb0000000000000000000000000000000000000000000000000000000069e9b91f0000000000000000000000000000000000000000000000000c2f2bf24ca3f6250000000000000000000000000000000000000000000000000000000069e9b9090000000000000000000000000000000000000000000000000000000069e9b91d000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000643017e1b4cf0843b4a1cd27c0e319fc6500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002800000000000000000000000006044eef7179034319e2c8636ea885b37cbfa9aba0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000000000414e7c025d42643738711cd05f9665d95a0848bfb3b5057d66f2fb28510c11fa443647b1a81967af6f90905ceb95dc9d1b69802082f0b3f1705afd3b0a7a849f4b1b0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000417afc296a94f4db8af9a549853e2cb5f08ae3128e9d8fe626c5436d6d7ad81702325d2cbfff1999675d7ef4d747e288c24fa7e4e6720f3d74c7a096c48286b5171b00000000000000000000000000000000000000000000000000000000000000",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := hexutil.Decode(tc.rawData)
			require.NoError(t, err)
			result, err := DecodeTradeRFQT(data)
			require.NoError(t, err)
			t.Log(hexutil.Encode(result.Quote.QuoteId[:]))
		})
	}

}
