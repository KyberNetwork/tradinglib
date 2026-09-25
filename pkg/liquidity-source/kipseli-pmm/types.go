package kipselipmm

// SwapExtra represents extra information for a swap operation.
// it's compatible with kyberpmm.SwapExtra
type SwapExtra struct {
	TakerAsset   string `json:"t"`
	TakingAmount string `json:"i"`
	MakerAsset   string `json:"m"`
	MakingAmount string `json:"o"`

	// those fields are not used in kyberpmm, we used in our logic.
	IsExactOut bool      `json:"e"`
	SwapLogs   []SwapLog `json:"l"`
}

type SwapLog struct {
	Base             string  `json:"b"`
	Quote            string  `json:"q"`
	AmountIn         float64 `json:"i"`
	AmountOut        float64 `json:"o"`
	IsSell           bool    `json:"s"`
	SoldBaseAmount   float64 `json:"sba,omitempty"`
	BoughtBaseAmount float64 `json:"bba,omitempty"`
}

type PricingData struct {
	Tokens    map[string]Token
	Pairs     []Pair
	Timestamp int64
}

func (d PricingData) IsZero() bool {
	return len(d.Tokens) == 0 && len(d.Pairs) == 0
}

type Token struct {
	Symbol         string
	Address        string
	Decimals       int64
	MinTradeAmount float64
}

type Pair struct {
	Base  string
	Quote string
	Sell  []PriceLevel
	Buy   []PriceLevel
}

type PriceLevel struct {
	Rate     float64
	Quantity float64
}
