package middlewares

import (
	"github.com/arfanxn/coursefan-gofiber/app/exceptions"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/responseh"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func After() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		err = c.Next()

		if err != nil {
			var response resources.Response
			logrus.Error(err)
			switch err.(type) {
			default: // sends internal server error if error is default error
				response.FromError(fiber.ErrInternalServerError)
			case *fiber.Error:
				response.FromError(err.(*fiber.Error))
			case *exceptions.ValidationError:
				response.FromValidationErrs([]*exceptions.ValidationError{
					err.(*exceptions.ValidationError),
				})
			}
			err = responseh.Write(c, response)
		}

		return err
	}
}
