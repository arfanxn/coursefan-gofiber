package routes

import "github.com/gofiber/fiber/v2"

func registerSandboxRouter(router fiber.Router) {
	router.Get("sandbox", func(c *fiber.Ctx) (err error) {

		return
	})
}
