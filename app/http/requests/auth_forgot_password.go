package requests

type AuthForgotPassword struct {
	Email string `validate:"required,email,max=50"`
}
