package middlewares

import (
	"github.com/arfanxn/coursefan-gofiber/app/exceptions"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/gofiber/fiber/v2"
)

func After() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		err = c.Next()

		if err != nil {
			switch err.(type) {
			default: // sends internal server error if error is default error
				return c.Send(resources.NewResponseError(fiber.ErrInternalServerError).Bytes())
			case *fiber.Error:
				return c.Send(resources.NewResponseError(err.(*fiber.Error)).Bytes())
			case *exceptions.ValidationError:
				return c.Send(resources.NewResponseValidationErrs([]*exceptions.ValidationError{
					err.(*exceptions.ValidationError),
				}).Bytes())
			}
		}

		return c.Next()
	}
}
