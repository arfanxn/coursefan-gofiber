package responseh

import (
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/gofiber/fiber/v2"
)

// Write writes response to client
func Write(c *fiber.Ctx, response *resources.Response) error {
	return c.Status(response.Code).Send(response.Bytes())
}
