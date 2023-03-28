package resources

type AuthLogin struct {
	User
	AccessToken string `json:"access_token"`
}
