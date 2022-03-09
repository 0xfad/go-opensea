package goopensea

type Collection struct {
	Name                  string          `json:"name"`
	Slug                  string          `json:"slug"`
	Editors               []string        `json:"editors"`
	Hidden                bool            `json:"hidden"`
	Featured              bool            `json:"featured"`
	CreatedDate           string          `json:"created_date"`
	Description           string          `json:"description"`
	ImageURL              string          `json:"image_url"`
	BuyerFeeBasisPoints   string          `json:"opensea_buyer_fee_basis_points"`
	SellerFeeBasisPoints  string          `json:"opensea_seller_fee_basis_points"`
	OwnedAssetCount       int             `json:"owned_asset_count"`
	PayoutAddress         string          `json:"payout_address"`
	PrimaryAssetContracts []AssetContract `json:"primary_asset_contracts"`
	Stats                 Stats           `json:"stats"`
	Traits                CollectionTrait `json:"traits"`
	PaymentTokens         []PaymentToken  `json:"payment_tokens"`
}

type CollectionTrait map[string]interface{}

type Stats struct {
	AveragePrice float64 `json:"average_price"`
	Count        float32 `json:"count"`
	FloorPrice   float64 `json:"floor_price"`
	MarketCap    float64 `json:"market_cap"`
	NumOwners    float32 `json:"num_owners"`
	NumReports   float32 `json:"num_reports"`
	TotalSales   float32 `json:"total_sales"`
	TotalSupply  float32 `json:"total_supply"`
	TotalVolume  float64 `json:"total_volume"`
}
