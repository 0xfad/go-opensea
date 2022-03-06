package goopensea

type User struct {
	Username string `json:"username"`
}

type Account struct {
	User            User   `json:"user"`
	ProfileImageURL string `json:"profile_img_url"`
	Address         string `json:"address"`
}
