package goopensea

type SingleCollectionResponse struct {
	Collection Collection `json:"collection"`
}

type CollectionListResponse struct {
	Collections []Collection `json:"collections"`
}

type SingleAssetResponse struct {
	Asset Asset `json:"asset"`
}

type AssetListResponse struct {
	Assets []Asset `json:"assets"`
}
