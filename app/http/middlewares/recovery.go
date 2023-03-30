package middlewares

import (
	"github.com/arfanxn/coursefan-gofiber/app/helpers/responseh"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func Recovery() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		defer func() {
			errAny := recover()
			if errAny != nil {
				logrus.Error(errAny) // logs any error
				err = responseh.Write(c, resources.NewResponseError(fiber.ErrInternalServerError))
			}
		}()
		err = c.Next()
		return err
	}
}
