package requests

import (
	"github.com/gofiber/fiber/v2"
)

type AuthLogin struct {
	Email    string `json:"email" validate:"required,email,max=50"`
	Password string `json:"password" validate:"required,ascii,min=6,max=50"`
}

func (input *AuthLogin) FromContext(c *fiber.Ctx) {
	c.BodyParser(input)
}
