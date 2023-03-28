package requests

type AuthLogin struct {
	Email    string `validate:"required,email,min=3,max=50"`
	Password string `validate:"required,ascii,min=6,max=50"`
}
