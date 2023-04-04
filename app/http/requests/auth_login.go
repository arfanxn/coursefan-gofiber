package requests

type AuthLogin struct {
	Email    string `json:"email" validate:"required,email,max=50"`
	Password string `json:"password" validate:"required,ascii,min=6,max=50"`
}
