package goopensea

type Asset struct {
	ID           int64         `json:"id"`
	NumSales     int64         `json:"num_sales"`
	ImageURL     string        `json:"image_url"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	ExternalLink string        `json:"external_link"`
	Contract     AssetContract `json:"asset_contract"`
	SellOrders   []AssetOrder  `json:"sell_orders,omitempty"`
	Permalink    string        `json:"permalink"`
	Owner        Account       `json:"owner"`
	Creator      Account       `json:"creator"`
	Traits       []AssetTrait  `json:"traits"`
}

type AssetOrder struct {
	CreatedDate  string `json:"created_date"`
	CurrentPrice string `json:"current_price"`
	HowToCall    int    `json:"how_to_call"`
}

type AssetContract struct {
	Address              string  `json:"address"`
	AssetContractType    string  `json:"asset_contract_type"`
	BuyerFeeBasisPoints  float64 `json:"buyer_fee_basis_points"`
	SellerFeeBasisPoints float64 `json:"eller_fee_basis_points"`
	CreatedDate          string  `json:"created_date"`
	Description          string  `json:"description"`
	ExternalLink         string  `json:"external_link"`
	ImageURL             string  `json:"image_url"`
	Name                 string  `json:"name"`
	PayoutAddress        string  `json:"payout_address"`
	Owner                int64   `json:"owner"`
	SchemaName           string  `json:"schema_name"`
	Symbol               string  `json:"symbol"`
}

type AssetTrait struct {
	Type  string `json:"trait_type"`
	Value string `json:"trait_value"`
	Count int64  `json:"trait_count"`
}
