package middlewares

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/exceptions"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/jwth"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/skip"
	"github.com/golang-jwt/jwt/v4"
)

func Auth() fiber.Handler {
	middleware := func(c *fiber.Ctx) error {
		authCookieName := os.Getenv("AUTH_COOKIE_NAME")
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
				return c.Send(resources.NewResponseError(exceptions.AuthSessionExpired).Bytes())
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

	// Skip authentication at some scenarios
	return skip.New(middleware, func(c *fiber.Ctx) bool {
		// Route names to be excluded from authentication middleware
		excludedPaths := []string{"login", "register", "forgot-password"}

		// Current accessed route path
		routePath := c.Path()

		// if current route path is contained in excluded paths then it will skip authentication middleware
		isContain := len(
			sliceh.Filter(excludedPaths, func(excludedPath string) bool {
				return regexp.
					MustCompile(fmt.Sprintf("(%s){1}", excludedPath)).
					MatchString(routePath)
			}),
		) > 0
		return isContain
	})
}
