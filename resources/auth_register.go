package resources

type AuthRegister struct {
	User
	Avatar Media `json:"avatar"`
}
