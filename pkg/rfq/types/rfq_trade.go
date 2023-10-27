package types

type RFQTrade struct {
	OrderHash        string `json:"order_hash"`
	Maker            string `json:"maker"`
	Taker            string `json:"taker"`
	MakerToken       string `json:"maker_token"`
	TakerToken       string `json:"taker_token"`
	MakerTokenAmount string `json:"maker_token_amount"`
	TakerTokenAmount string `json:"taker_token_amount"`
	ContractAddress  string `json:"contract_address"`
	BlockNumber      uint64 `json:"block_number"`
	TxHash           string `json:"tx_hash"`
	LogIndex         uint64 `json:"log_index"`
	Timestamp        uint64 `json:"timestamp"`
	EventHash        string `json:"event_hash"`
}
