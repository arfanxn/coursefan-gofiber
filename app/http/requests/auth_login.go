package requests

type AuthLogin struct {
	Email    string `validate:"required,email,max=50"`
	Password string `validate:"required,ascii,min=6,max=50"`
}
