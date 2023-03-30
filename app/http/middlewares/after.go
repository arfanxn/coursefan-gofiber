package middlewares

import (
	"github.com/arfanxn/coursefan-gofiber/app/exceptions"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func After() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		err = c.Next()

		if err != nil {
			var resBody *resources.Response
			logrus.Error(err)
			switch err.(type) {
			default: // sends internal server error if error is default error
				resBody = resources.NewResponseError(fiber.ErrInternalServerError)
			case *fiber.Error:
				resBody = resources.NewResponseError(err.(*fiber.Error))
			case *exceptions.ValidationError:
				resBody = resources.NewResponseValidationErrs([]*exceptions.ValidationError{
					err.(*exceptions.ValidationError),
				})
			}
			err = c.Status(resBody.Code).Send(resBody.Bytes())
		}

		return err
	}
}
