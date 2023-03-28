package middlewares

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/exceptions"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/jwth"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(next http.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authCookieName := "Authorization"
		token := c.Cookies(authCookieName)
		if token == "" {
			return c.Send(resources.NewResponseError(fiber.ErrUnauthorized).Bytes())
		}

		signature := os.Getenv("APP_KEY")
		tokenizer, err := jwth.Decode(signature, token) // parse jwt token from cookie access token
		claims, ok := tokenizer.Claims.(jwt.MapClaims)  // get jwt claims
		c.Locals("user", claims["user"])                // set claims user into context

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				return c.Send(resources.NewResponseError(fiber.ErrUnauthorized).Bytes())
			case jwt.ValidationErrorExpired:
				return c.Send(resources.NewResponseError(exceptions.JWTExpired).Bytes())
			default:
				return c.Send(resources.NewResponseError(fiber.ErrInternalServerError).Bytes())
			}
		} else if !ok || !tokenizer.Valid {
			return c.Send(resources.NewResponseError(fiber.ErrInternalServerError).Bytes())
		}

		// Update token expiration
		authExpSec, err := strconv.ParseInt(os.Getenv("AUTH_EXP"), 10, 64)
		authExpTime := time.Now().Add(time.Minute * time.Duration(authExpSec))
		claims["exp"] = authExpTime.Unix()
		// Regenerate new token every request
		token, err = jwth.Encode(signature, claims)
		if err != nil {
			return err
		}

		// Refresh authentication cookie with new token
		c.Cookie(&fiber.Cookie{
			Name:     authCookieName,
			Path:     "/",
			Value:    token,
			HTTPOnly: true,
			MaxAge:   int(authExpSec),
			Expires:  authExpTime,
		})

		return c.Next()
	}
}
