package requests

import "mime/multipart"

type AuthRegister struct {
	Name            string `validate:"required,alpha,min=2,max=50"`
	Email           string `validate:"required,email,max=50"`
	Password        string `validate:"required,ascii,min=6,max=50"`
	ConfirmPassword string `validate:"required,eqfield=Password"`
	Avatar          *multipart.FileHeader
}
