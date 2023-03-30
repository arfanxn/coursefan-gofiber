package requests

import "mime/multipart"

type AuthRegister struct {
	Name            string                `json:"name" validate:"required,ascii,min=2,max=50"`
	Email           string                `json:"email" validate:"required,email,max=50"`
	Password        string                `json:"password" validate:"required,ascii,min=6,max=50"`
	ConfirmPassword string                `json:"confirm_password" form:"confirm_password" validate:"required,eqfield=Password"`
	Avatar          *multipart.FileHeader `json:"avatar" form:"avatar" fhlidate:"max=10,mimes=image/jpeg"`
}
