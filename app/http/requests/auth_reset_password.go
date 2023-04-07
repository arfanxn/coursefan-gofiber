package requests

import "github.com/gofiber/fiber/v2"

type AuthResetPassword struct {
	Email           string `json:"email" validate:"required,email,max=50"`
	Otp             string `json:"otp" validate:"required,min=6,max=6"`
	Password        string `json:"password" validate:"required,ascii,min=6,max=50"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" validate:"required,eqfield=Password"`
}

func (input *AuthResetPassword) FromContext(c *fiber.Ctx) {
	c.BodyParser(input)
}
