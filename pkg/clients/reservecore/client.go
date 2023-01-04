package reservecore

import (
	"net/http"
	"strings"

	"github.com/KyberNetwork/tradinglib/pkg/httpclient"
	"github.com/KyberNetwork/tradinglib/pkg/sb"
	"github.com/KyberNetwork/tradinglib/pkg/types"
)

// Client implements a http client for core api.
type Client struct {
	baseURL    string
	httpClient *http.Client
	useGateway bool
}

// New returns a new Client object.
func New(baseURL string, useGateway bool, httpClient *http.Client) (*Client, error) {

	return &Client{
		baseURL:    strings.TrimRight(baseURL, "/"),
		httpClient: httpClient,
		useGateway: useGateway,
	}, nil
}

type response struct {
	Reason  string      `json:"reason,omitempty"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func (c *Client) shouldSuccessRequest(req *http.Request, out interface{}) error {
	res := response{Data: out}
	err := httpclient.DoHTTPRequest(c.httpClient, req, &res)
	if err != nil {
		return err
	}
	if !res.Success {
		return types.NewAPIError(0, res.Reason)
	}
	return nil
}

func (c *Client) doRequest(req *http.Request, out interface{}) error {
	return httpclient.DoHTTPRequest(c.httpClient, req, out)
}

// ListAssets returns a list of assets.
func (c *Client) ListAssets() (assets []Asset, err error) {
	req, err := httpclient.NewRequest(http.MethodGet, c.baseURL, "/v3/asset", nil, nil)
	if err != nil {
		return nil, err
	}
	err = c.shouldSuccessRequest(req, &assets)
	if err != nil {
		return
	}
	return
}

// GetAuthData returns auth data snapshot for given `timestamp`.
func (c *Client) GetAuthData(timestamp int64) (authData AuthDataResponseV3, err error) {
	req, err := httpclient.NewRequest(http.MethodGet, c.baseURL, "/v3/authdata",
		httpclient.NewQuery().Int64("timestamp", timestamp), nil)
	if err != nil {
		return authData, err
	}
	err = c.shouldSuccessRequest(req, &authData)
	return authData, err
}

// GetSetRateStatus returns set-rate status.
func (c *Client) GetSetRateStatus() (status bool, err error) {
	req, err := httpclient.NewRequest(http.MethodGet, c.baseURL, "/v3/set-rate-status", nil, nil)
	if err != nil {
		return false, err
	}
	err = c.shouldSuccessRequest(req, &status)
	return status, err
}

// GetRebalanceStatus returns rebalance status.
func (c *Client) GetRebalanceStatus() (status bool, err error) {
	req, err := httpclient.NewRequest(http.MethodGet, c.baseURL, "/v3/rebalance-status", nil, nil)
	if err != nil {
		return false, err
	}
	err = c.shouldSuccessRequest(req, &status)
	return status, err
}

// GetMainBalance returns balance of binance account.
func (c *Client) GetMainBalance() (balance []BinanceMainAccountBalance, err error) {
	req, err := httpclient.NewRequest(http.MethodGet, c.baseURL, "/v3/binance/main", nil, nil)
	if err != nil {
		return nil, err
	}
	err = c.shouldSuccessRequest(req, &balance)
	return balance, err
}

// GetFeedConfiguration returns configuration for feed data.
func (c *Client) GetFeedConfiguration() (config FeedConfiguration, err error) {
	req, err := httpclient.NewRequest(http.MethodGet, c.baseURL, "/v3/feed-configurations", nil, nil)
	if err != nil {
		return config, err
	}
	err = c.shouldSuccessRequest(req, &config)
	return config, err
}

// ListExchanges returns a list of supported exchanges.
func (c *Client) ListExchanges() (out []Exchange, err error) {
	req, err := httpclient.NewRequest(http.MethodGet, c.baseURL, "/v3/exchange", nil, nil)
	if err != nil {
		return out, err
	}
	err = c.shouldSuccessRequest(req, &out)
	return out, err
}

// GetOpenOrders returns a list of open orders.
func (c *Client) GetOpenOrders() (out map[ExchangeID][]Order, err error) {
	req, err := httpclient.NewRequest(http.MethodGet, c.baseURL, "/v3/open-orders", nil, nil)
	if err != nil {
		return out, err
	}
	err = c.shouldSuccessRequest(req, &out)
	return out, err
}

// CancelOrders cancel orders.
func (c *Client) CancelOrders(data CancelOrderRequest) (out map[string]CancelOrderResult, err error) {
	req, err := httpclient.NewPostJSON(c.baseURL, "/v3/cancel-orders", nil, data)
	if err != nil {
		return out, err
	}
	err = c.shouldSuccessRequest(req, &out)
	return out, err
}

// GetPriceFactor returns a list of price factors.
func (c *Client) GetPriceFactor(from int64, to int64) (out PriceFactorResponse, err error) {
	req, err := httpclient.NewRequest(http.MethodGet, c.baseURL, "/v3/price-factor",
		httpclient.NewQuery().Int64("from", from).Int64("to", to),
		nil)
	if err != nil {
		return out, err
	}
	err = c.doRequest(req, &out)
	if err != nil {
		return out, err
	}
	if !out.Success {
		return out, types.NewAPIError(0, out.Reason)
	}
	return out, nil
}

// Deposit deposits fund into the account on a CEX.
func (c *Client) Deposit(data DepositRequest) (string, error) {
	req, err := httpclient.NewPostJSON(c.baseURL, "/v3/deposit", nil, data)
	if err != nil {
		return "", err
	}
	var res idResponse
	err = httpclient.DoHTTPRequest(c.httpClient, req, &res)
	if err != nil {
		return "", err
	}
	if !res.Success {
		return "", types.NewAPIError(0, res.Reason)
	}
	return res.ID, nil
}

// Withdraw funds from the account on a CEX.
func (c *Client) Withdraw(data WithdrawRequest) (string, error) {
	req, err := httpclient.NewPostJSON(c.baseURL, "/v3/withdraw", nil, data)
	if err != nil {
		return "", err
	}
	var res idResponse
	err = httpclient.DoHTTPRequest(c.httpClient, req, &res)
	if err != nil {
		return "", err
	}
	if !res.Success {
		return "", types.NewAPIError(0, res.Reason)
	}
	return res.ID, nil
}

func (c *Client) cexDataPrefix() string {
	if c.useGateway {
		return "/cex-data"
	}
	return ""
}

// WithdrawWithLimitedPermission withdraw funds from the account on a CEX.
func (c *Client) WithdrawWithLimitedPermission(accountID string, data CEXDEXWithdrawRequest) (string, error) {
	req, err := httpclient.NewPostJSON(c.baseURL,
		sb.Concat(c.cexDataPrefix(), "/sapi/v1/capital/withdraw/apply/", accountID),
		httpclient.NewQuery().Struct(data), nil)
	if err != nil {
		return "", err
	}
	var res idResponse
	err = httpclient.DoHTTPRequest(c.httpClient, req, &res)
	if err != nil {
		return "", err
	}
	return res.ID, nil
}

// GetWithdrawActivityStatus get withdraw activity status
func (c *Client) GetWithdrawActivityStatus(eid string) (out ActivityRecord, err error) {
	req, err := httpclient.NewRequest(http.MethodGet, c.baseURL, sb.Concat(c.cexDataPrefix(), "/activity"),
		httpclient.NewQuery("eid", eid),
		nil)
	if err != nil {
		return out, err
	}
	err = c.shouldSuccessRequest(req, &out)
	return out, err
}

// BorrowTransferMargin borrow and transfer from margin to spot account
func (c *Client) BorrowTransferMargin(data BorrowTransferRequest) (string, error) {
	req, err := httpclient.NewPostJSON(c.baseURL, "/v3/margin/borrow-and-transfer", nil, data)
	if err != nil {
		return "", err
	}
	var res idResponse
	err = httpclient.DoHTTPRequest(c.httpClient, req, &res)
	if err != nil {
		return "", err
	}
	if !res.Success {
		return "", types.NewAPIError(0, res.Reason)
	}
	return res.ID, nil
}

// TransferRepayMargin transfer and repay from spot to margin account
func (c *Client) TransferRepayMargin(data TransferRepayRequest) (string, error) {
	req, err := httpclient.NewPostJSON(c.baseURL, "/v3/margin/transfer-and-repay", nil, data)
	if err != nil {
		return "", err
	}
	var res idResponse
	err = httpclient.DoHTTPRequest(c.httpClient, req, &res)
	if err != nil {
		return "", err
	}
	if !res.Success {
		return "", types.NewAPIError(0, res.Reason)
	}
	return res.ID, nil
}

// GetMarginAccountInfo get margin account info
func (c *Client) GetMarginAccountInfo(exchange ExchangeID) (out CrossMarginAccountDetails, err error) {
	req, err := httpclient.NewRequest(http.MethodGet, c.baseURL, "/v3/margin/account",
		httpclient.NewQuery("exchange", int64(exchange)),
		nil)
	if err != nil {
		return out, err
	}
	err = c.shouldSuccessRequest(req, &out)
	return out, err
}

// GetCrossMarginData get cross margin data
func (c *Client) GetCrossMarginData(exchange ExchangeID) (out []CrossMarginData, err error) {
	req, err := httpclient.NewRequest(http.MethodGet, c.baseURL, "/v3/margin/account",
		httpclient.NewQuery("exchange", int64(exchange)),
		nil)
	if err != nil {
		return out, err
	}
	err = c.shouldSuccessRequest(req, &out)
	return out, err
}

// MarginConfig margin config.
type MarginConfig struct {
	MarginEnable bool `json:"margin_enable"`
}

// GetMarginConfig returns margin config.
func (c *Client) GetMarginConfig() (out MarginConfig, err error) {
	req, err := httpclient.NewRequest(http.MethodGet, c.baseURL, "/v3/margin-config", nil, nil)
	if err != nil {
		return out, err
	}
	err = c.shouldSuccessRequest(req, &out)
	return out, err
}

// PerpetualConfig perpetual config.
type PerpetualConfig struct {
	PerpetualEnable bool `json:"perpetual_enable"`
}

// GetPerpetualConfig returns perpetual config.
func (c *Client) GetPerpetualConfig() (out PerpetualConfig, err error) {
	req, err := httpclient.NewRequest(http.MethodGet, c.baseURL, "/v3/perpetual-config", nil, nil)
	if err != nil {
		return out, err
	}
	err = c.shouldSuccessRequest(req, &out)
	return out, err
}

// MarginLevelThresholdResponse margin level threshold config.
type MarginLevelThresholdResponse struct {
	MarginLevelThreshold float64 `json:"margin_level_threshold"`
}

// GetMarginLevelThreshold returns margin level threshold config.
func (c *Client) GetMarginLevelThreshold() (out MarginLevelThresholdResponse, err error) {
	req, err := httpclient.NewRequest(http.MethodGet, c.baseURL, "/v3/margin-level-threshold",
		nil, nil)
	if err != nil {
		return out, err
	}
	err = c.shouldSuccessRequest(req, &out)
	return out, err
}

// TradeRequest form
type TradeRequest struct {
	Pair   TradingPairID `json:"pair"`
	Amount float64       `json:"amount"`
	Rate   float64       `json:"rate"`
	Type   string        `json:"type"`
}

// TradeResponse ...
type TradeResponse struct {
	ID        string  `json:"id"`
	Done      string  `json:"done"`
	Remaining float64 `json:"remaining"`
	Finished  float64 `json:"finished"`
	Reason    string  `json:"reason"`
	Success   bool    `json:"success"`
}

// Trade makes a trade on a CEX.
func (c *Client) Trade(data TradeRequest) (trade TradeResponse, err error) {
	req, err := httpclient.NewPostJSON(c.baseURL, "/v3/trade", nil, data)
	if err != nil {
		return trade, err
	}
	err = httpclient.DoHTTPRequest(c.httpClient, req, &trade)
	if err != nil {
		return trade, err
	}
	if !trade.Success {
		return trade, types.NewAPIError(0, trade.Reason)
	}
	return trade, nil
}
