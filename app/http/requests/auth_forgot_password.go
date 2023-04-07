package requests

import "github.com/gofiber/fiber/v2"

type AuthForgotPassword struct {
	Email string `validate:"required,email,max=50"`
}

func (input *AuthForgotPassword) FromContext(c *fiber.Ctx) {
	c.BodyParser(input)
}
