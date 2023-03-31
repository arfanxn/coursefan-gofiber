package bootstrap

import (
	"os"
	"strconv"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/numh"
	"github.com/gofiber/fiber/v2"
)

// NewApp bootstraps a new App with configured config from environment variables
func NewApp() (app *fiber.App, err error) {
	bodyLimit, err := strconv.ParseInt(os.Getenv("REQUEST_BODY_LIMIT"), 10, 64)
	if err != nil {
		return
	}
	bodyLimit = numh.MegabyteToByte(bodyLimit)

	app = fiber.New(fiber.Config{
		AppName:   os.Getenv("APP_NAME"),
		BodyLimit: int(bodyLimit),
	})
	return
}
