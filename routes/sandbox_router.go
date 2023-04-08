package routes

import (
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/davecgh/go-spew/spew"
	"github.com/gofiber/fiber/v2"
)

func registerSandboxRouter(router fiber.Router) {
	router.Get("/sandbox", func(c *fiber.Ctx) (err error) {
		input := requests.QueryExp{}
		err = input.FromContext(c)
		spew.Dump(input)
		if err != nil {
			return
		}
		return
	})
}
