package reservecore

import (
	"math/big"
	"time"

	ethereum "github.com/ethereum/go-ethereum/common"
)

//go:generate enumer -type=SetRate -linecomment -json=true -sql=true
type SetRate int

//go:generate stringer -type=ExchangeID -linecomment
type ExchangeID uint64

type (
	AssetID         uint64
	TradingPairID   uint64
	TradingByID     uint64
	AssetExchangeID uint64
	AssetAddressID  uint64
	SettingChangeID uint64
	FeedWeightID    uint64
)

type PWIEquation struct {
	A                   float64 `json:"a"`
	B                   float64 `json:"b"`
	C                   float64 `json:"c"`
	MinMinSpread        float64 `json:"min_min_spread"`
	PriceMultiplyFactor float64 `json:"price_multiply_factor"`
}

type AssetPWI struct {
	Ask PWIEquation `json:"ask"`
	Bid PWIEquation `json:"bid"`
}

// RebalanceQuadratic is params of quadratic equation.
type RebalanceQuadratic struct {
	SizeA       float64 `json:"size_a"`
	SizeB       float64 `json:"size_b"`
	SizeC       float64 `json:"size_c"`
	PriceA      float64 `json:"price_a"`
	PriceB      float64 `json:"price_b"`
	PriceC      float64 `json:"price_c"`
	PriceOffset float64 `json:"price_offset"`
}

// AssetTarget is the target setting of an asset.
type AssetTarget struct {
	Total                float64 `json:"total"`
	Reserve              float64 `json:"reserve"`
	RebalanceThreshold   float64 `json:"rebalance_threshold"`
	TransferThreshold    float64 `json:"transfer_threshold"`
	MinWithdrawThreshold float64 `json:"min_withdraw_threshold"`
	TriggerRebalanceTS   int64   `json:"trigger_rebalance_timestamp"`
}

// AssetExchange is the configuration of an asset for a specific exchange.
type AssetExchange struct {
	ID                    AssetExchangeID  `json:"id"`
	AssetID               AssetID          `json:"asset_id"`
	ExchangeID            ExchangeID       `json:"exchange_id"`
	Symbol                string           `json:"symbol" binding:"required"`
	DepositAddress        ethereum.Address `json:"deposit_address"`
	MinDeposit            float64          `json:"min_deposit"`
	WithdrawFee           float64          `json:"withdraw_fee"`
	TargetRecommended     float64          `json:"target_recommended"`
	TargetRatio           float64          `json:"target_ratio"`
	TradingPairs          []TradingPair    `json:"trading_pairs,omitempty" binding:"dive"`
	DepositWithdrawSymbol string           `json:"deposit_withdraw_symbol"`
}

// TradingPair is a trading in an exchange.
type TradingPair struct {
	ID              TradingPairID `json:"id"`
	Base            AssetID       `json:"base"`
	Quote           AssetID       `json:"quote"`
	PricePrecision  uint64        `json:"price_precision"`
	AmountPrecision uint64        `json:"amount_precision"`
	AmountLimitMin  float64       `json:"amount_limit_min"`
	AmountLimitMax  float64       `json:"amount_limit_max"`
	PriceLimitMin   float64       `json:"price_limit_min"`
	PriceLimitMax   float64       `json:"price_limit_max"`
	MinNotional     float64       `json:"min_notional"`
	ExchangeID      ExchangeID    `json:"-"`
	BaseSymbol      string        `json:"base_symbol"`
	QuoteSymbol     string        `json:"quote_symbol"`
	StaleThreshold  float64       `json:"stale_threshold"`
}

// StableParam is params of stablize action.
type StableParam struct {
	PriceUpdateThreshold float64 `json:"price_update_threshold"`
	AskSpread            float64 `json:"ask_spread"`
	BidSpread            float64 `json:"bid_spread"`
	SingleFeedMaxSpread  float64 `json:"single_feed_max_spread"`
	MultipleFeedsMaxDiff float64 `json:"multiple_feeds_max_diff"`
}

type FeedWeight map[string]float64

type SanityInfo struct {
	Provider  string   `json:"provider"`
	Threshold float64  `json:"threshold"`
	Path      []string `json:"path"`
}

// StableCoin is params for stable coins like BUSD, DAI...
type StableCoin struct {
	IsStable           bool    `json:"is_stable"`
	BinanceConvert     bool    `json:"binance_convert"`
	RateAmount         float64 `json:"rate_amount"`
	PriceDiffThreshold float64 `json:"price_diff_threshold"`
}

type RFQParams struct {
	Asset          AssetID `json:"asset_id"`
	Integration    string  `json:"integration"`
	RefETHAmount   float64 `json:"ref_eth_amount"`
	ETHStep        float64 `json:"eth_step"`
	MaxETHSizeBuy  float64 `json:"max_eth_size_buy"`
	MaxETHSizeSell float64 `json:"max_eth_size_sell"`
	MinSlippage    float64 `json:"min_slippage"`
	MaxSlippage    float64 `json:"max_slippage"`
	Enabled        bool    `json:"enabled"`
	SwapEnabled    bool    `json:"swap_enabled"`
	AskOffset      float64 `json:"ask_offset"`
	BidOffset      float64 `json:"bid_offset"`
	A              float64 `json:"a"`
	B              float64 `json:"b"`
	C              float64 `json:"c"`
	MultiplyFactor float64 `json:"multiply_factor"`
	MinMin         float64 `json:"min_min"`
	MaxImb         float64 `json:"max_imb"`
}

// AssetMarginBalance ...
type AssetMarginBalance struct {
	Borrowed float64 `json:"borrowed"`
	Free     float64 `json:"free"`
	Interest float64 `json:"interest"`
	Locked   float64 `json:"locked"`
	NetAsset float64 `json:"net_asset"`
}

type Asset struct {
	ID                    AssetID             `json:"id"`
	Symbol                string              `json:"symbol" binding:"required"`
	Name                  string              `json:"name"`
	Address               ethereum.Address    `json:"address"`
	OldAddresses          []ethereum.Address  `json:"old_addresses,omitempty"`
	Decimals              uint64              `json:"decimals"`
	Transferable          bool                `json:"transferable"`
	MarginEnable          bool                `json:"margin_enable"`
	PerpetualEnable       bool                `json:"perpetual_enable"`
	SetRate               SetRate             `json:"set_rate"`
	Rebalance             bool                `json:"rebalance"`
	IsQuote               bool                `json:"is_quote"`
	PWI                   *AssetPWI           `json:"pwi,omitempty"`
	RebalanceQuadratic    *RebalanceQuadratic `json:"rebalance_quadratic,omitempty"`
	Exchanges             []AssetExchange     `json:"exchanges,omitempty" binding:"dive"`
	Target                *AssetTarget        `json:"target,omitempty"`
	StableParam           StableParam         `json:"stable_param"`
	Created               time.Time           `json:"created"`
	Updated               time.Time           `json:"updated"`
	FeedWeight            *FeedWeight         `json:"feed_weight,omitempty"`
	NormalUpdatePerPeriod float64             `json:"normal_update_per_period"`
	MaxImbalanceRatio     float64             `json:"max_imbalance_ratio"`
	OrderDurationMillis   uint64              `json:"order_duration_millis"`
	PriceETHAmount        float64             `json:"price_eth_amount"`
	ExchangeETHAmount     float64             `json:"exchange_eth_amount"`
	SanityInfo            SanityInfo          `json:"sanity_info"`
	RFQParams             []RFQParams         `json:"rfq_params"`
	StableCoin            StableCoin          `json:"stablecoin"`
}

// ExchangeBalance is balance of a token of an exchange.
type ExchangeBalance struct {
	ExchangeID    ExchangeID         `json:"exchange_id"`
	Available     float64            `json:"available"`
	Locked        float64            `json:"locked"`
	Name          string             `json:"name"`
	MarginBalance AssetMarginBalance `json:"margin_balance"`
	Error         string             `json:"error"`
}

// Total return total amount.
func (eb *ExchangeBalance) Total() float64 {
	return eb.Available + eb.Locked
}

// Lock amount coin.
func (eb *ExchangeBalance) Lock(amount float64) {
	if amount > eb.Available {
		return
	}
	eb.Available -= amount
	eb.Locked += amount
}

// Version indicate fetched data version.
type Version uint64

// AuthdataBalance is balance for a token in reservesetting authata.
type AuthdataBalance struct {
	Valid        bool              `json:"valid"`
	AssetID      AssetID           `json:"asset_id"`
	Exchanges    []ExchangeBalance `json:"exchanges"`
	Reserve      float64           `json:"reserve"`
	ReserveError string            `json:"reserve_error"`
	Symbol       string            `json:"symbol"`
}

// PendingActivities is pending activities for authdata.
type PendingActivities struct {
	SetRates []ActivityRecord `json:"set_rates"`
	Withdraw []ActivityRecord `json:"withdraw"`
	Deposit  []ActivityRecord `json:"deposit"`
	Trades   []ActivityRecord `json:"trades"`
}

// ActivityID specify an activity.
type ActivityID struct {
	Timepoint uint64
	EID       string
}

// ActivityParams is params for activity.
type ActivityParams struct {
	// deposit, withdraw params
	Exchange  ExchangeID `json:"exchange,omitempty"`
	Asset     AssetID    `json:"asset,omitempty"`
	Amount    float64    `json:"amount,omitempty"`
	Timepoint uint64     `json:"timepoint,omitempty"`
	// SetRates params
	Assets []AssetID  `json:"assets,omitempty"` // list of asset id
	Buys   []*big.Int `json:"buys,omitempty"`
	Sells  []*big.Int `json:"sells,omitempty"`
	Block  *big.Int   `json:"block,omitempty"`
	AFPMid []*big.Int `json:"afpMid,omitempty"`
	Msgs   []string   `json:"msgs,omitempty"`
	// Trade params
	Type          string        `json:"type,omitempty"`
	Base          string        `json:"base,omitempty"`
	Quote         string        `json:"quote,omitempty"`
	Rate          float64       `json:"rate,omitempty"`
	FilledPrice   float64       `json:"filled_price,omitempty"`
	Triggers      []bool        `json:"triggers,omitempty"`
	TradingPairID TradingPairID `json:"trading_pair_id,omitempty"`
}

// ActivityResult is result of an activity.
type ActivityResult struct {
	Tx       string `json:"tx,omitempty"`
	Nonce    uint64 `json:"nonce,omitempty"`
	GasPrice string `json:"gasPrice,omitempty"`
	Error    string `json:"error,omitempty"`
	// ID of withdraw
	ID string `json:"id,omitempty"`
	// params of trade
	Done        float64 `json:"done,omitempty"`
	Remaining   float64 `json:"remaining,omitempty"`
	FilledPrice float64 `json:"filled_price,omitempty"`
	Finished    bool    `json:"finished,omitempty"`
	//
	StatusError string `json:"status_error,omitempty"`
	BlockNumber uint64 `json:"blockNumber,omitempty"`
	//
	WithdrawFee float64 `json:"withdraw_fee,omitempty"`
	TxTime      uint64  `json:"tx_time"` // when the tx was sent, will be update when tx get override to speed up
	IsReplaced  bool    `json:"is_replaced,omitempty"`
}

type Timestamp string

// ActivityRecord object.
type ActivityRecord struct {
	OrgTime uint64 `json:"org_time"` // origin timestamp - timestamp of the first activity
	// (in case it has many activities like override or replace)
	Action         string          `json:"action,omitempty"`
	ID             ActivityID      `json:"id,omitempty"`
	EID            string          `json:"eid"`
	Destination    string          `json:"destination,omitempty"`
	Params         *ActivityParams `json:"params,omitempty"`
	Result         *ActivityResult `json:"result,omitempty"`
	ExchangeStatus string          `json:"exchange_status,omitempty"`
	MiningStatus   string          `json:"mining_status,omitempty"`
	Timestamp      Timestamp       `json:"timestamp,omitempty"` // created
	LastTime       uint64          `json:"last_time"`           // activity finished time
}

// AuthDataResponseV3 is auth data format for reservesetting.
type AuthDataResponseV3 struct {
	Balances          []AuthdataBalance `json:"balances"`
	PendingActivities PendingActivities `json:"pending_activities"`
	Version           Version           `json:"version"`
}

type BinanceMainAccountBalance struct {
	AssetID AssetID `json:"asset_id"`
	Symbol  string  `json:"symbol"`
	Free    string  `json:"free"`
	Locked  string  `json:"locked"`
}

// FeedConfiguration feed configuration.
type FeedConfiguration struct {
	Name                 string  `json:"name" db:"name"`
	SetRate              SetRate `json:"set_rate" db:"set_rate"`
	Enabled              bool    `json:"enabled" db:"enabled"`
	BaseVolatilitySpread float64 `json:"base_volatility_spread" db:"base_volatility_spread"`
	NormalSpread         float64 `json:"normal_spread" db:"normal_spread"`
}

// Exchange represents a centralized exchange in database.
type Exchange struct {
	ID              ExchangeID `json:"id"`
	Name            string     `json:"name"`
	TradingFeeMaker float64    `json:"trading_fee_maker"`
	TradingFeeTaker float64    `json:"trading_fee_taker"`
	Disable         bool       `json:"disable"`
}

// Order across multiple exchanges.
type Order struct {
	ID            string        `json:"id,omitempty"` // standard id across multiple exchanges
	Symbol        string        `json:"symbol,omitempty"`
	Base          string        `json:"base"`
	Quote         string        `json:"quote"`
	OrderID       string        `json:"order_id"`
	Price         float64       `json:"price"`
	OrigQty       float64       `json:"orig_qty"`     // original quantity
	ExecutedQty   float64       `json:"executed_qty"` // matched quantity
	TimeInForce   string        `json:"time_in_force,omitempty"`
	Type          string        `json:"type"` // market or limit
	Side          string        `json:"side"` // buy or sell
	StopPrice     string        `json:"stop_price,omitempty"`
	IcebergQty    string        `json:"iceberg_qty,omitempty"`
	Time          uint64        `json:"time,omitempty"`
	TradingPairID TradingPairID `json:"trading_pair_id,omitempty"`
}

type RequestOrder struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
}

// CancelOrderRequest type.
type CancelOrderRequest struct {
	ExchangeID ExchangeID     `json:"exchange_id"`
	Orders     []RequestOrder `json:"orders"`
}

// CancelOrderResult is response when calling cancel an order.
type CancelOrderResult struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// AssetPriceFactorResponse is on element in asset price factor list output in getPriceFactors.
type AssetPriceFactorResponse struct {
	Timestamp uint64  `json:"timestamp"`
	AfpMid    float64 `json:"afp_mid"`
	Spread    float64 `json:"spread"`
}

// AssetPriceFactorListResponse present for price factor list of an asset.
type AssetPriceFactorListResponse struct {
	AssetID AssetID                    `json:"id"`
	Data    []AssetPriceFactorResponse `json:"data"`
}

type commonResponse struct {
	Success bool   `json:"success"`
	Reason  string `json:"reason,omitempty"`
}

// PriceFactorResponse present for out getPriceFactors result.
type PriceFactorResponse struct {
	commonResponse
	Timestamp  uint64                          `json:"timestamp"`
	ReturnTime uint64                          `json:"returnTime"`
	Data       []*AssetPriceFactorListResponse `json:"data"`
}

// DepositRequest type.
type DepositRequest struct {
	ExchangeID ExchangeID `json:"exchange"`
	Amount     *big.Int   `json:"amount"`
	Asset      AssetID    `json:"asset"`
}

// WithdrawRequest type.
type WithdrawRequest struct {
	ExchangeID ExchangeID `json:"exchange"`
	Asset      AssetID    `json:"asset"`
	Amount     *big.Int   `json:"amount"`
}

type CEXDEXWithdrawRequest struct {
	WithdrawOrderID string `json:"withdraw_order_id"`
	Amount          string `json:"amount"`
	WithdrawAsset   string `json:"withdraw_asset"`
	Asset           string `json:"asset"`
	Address         string `json:"address"`
	Name            string `json:"name"`
	Network         string `json:"network"` // default = ethereum
}

// BorrowTransferRequest type.
type BorrowTransferRequest struct {
	ExchangeID ExchangeID `json:"exchange"`
	Asset      AssetID    `json:"asset"`
	Amount     float64    `json:"amount"`
}

// TransferRepayRequest type.
type TransferRepayRequest struct {
	ExchangeID ExchangeID `json:"exchange"`
	Asset      AssetID    `json:"asset"`
	Amount     float64    `json:"amount"`
}

// RawAssetMarginBalance raw margin balance data response from binance.
type RawAssetMarginBalance struct {
	Asset    string `json:"asset"`
	Borrowed string `json:"borrowed"`
	Free     string `json:"free"`
	Interest string `json:"interest"`
	Locked   string `json:"locked"`
	NetAsset string `json:"netAsset"`
}

// CrossMarginAccountDetails binance margin account info.
type CrossMarginAccountDetails struct {
	BorrowEnabled       bool                    `json:"borrowEnabled"`
	MarginLevel         string                  `json:"marginLevel"`
	TotalAssetOfBtc     string                  `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc string                  `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc  string                  `json:"totalNetAssetOfBtc"`
	TradeEnabled        bool                    `json:"tradeEnabled"`
	TransferEnabled     bool                    `json:"transferEnabled"`
	UserAssets          []RawAssetMarginBalance `json:"userAssets"`
}

// CrossMarginData ..
type CrossMarginData struct {
	VipLevel        int      `json:"vipLevel"`
	Coin            string   `json:"coin"`
	TransferIn      bool     `json:"transferIn"`
	Borrowable      bool     `json:"borrowable"`
	DailyInterest   string   `json:"dailyInterest"`
	YearlyInterest  string   `json:"yearlyInterest"`
	BorrowLimit     string   `json:"borrowLimit"`
	MarginablePairs []string `json:"marginablePairs"`
}

type idResponse struct {
	commonResponse
	ID string `json:"id"`
}
