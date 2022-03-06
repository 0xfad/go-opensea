package goopensea

type PaymentToken struct {
	ID       int64   `json:"id"`
	Symbol   string  `json:"symbol"`
	Address  string  `json:"address"`
	ImageURL string  `json:"image_url"`
	Name     string  `json:"name"`
	ETHPrice float64 `json:"eth_price"`
	USDPrice float64 `json:"usd_price"`
}
