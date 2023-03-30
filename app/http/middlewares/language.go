package middlewares

import (
	"os"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/ctxh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/validateh"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/gofiber/fiber/v2"
)

// Language middleware validates "Accept-Language" from header or "lang" from qeury params if present
func Language() fiber.Handler {
	return func(c *fiber.Ctx) error {
		lang := ctxh.GetAcceptLang(c)
		input := struct {
			Lang string `json:"lang" validate:"alpha,oneof=en id"`
		}{
			Lang: lang,
		}

		if validationErrs := validateh.ValidateStruct(input, os.Getenv("APP_LANGUAGE")); validationErrs != nil {
			response := resources.NewResponseValidationErrs(validationErrs)
			return c.Send(response.Bytes())
		}

		return c.Next()
	}
}
