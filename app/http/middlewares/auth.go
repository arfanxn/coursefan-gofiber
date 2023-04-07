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
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/skip"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func Auth() fiber.Handler {
	middleware := func(c *fiber.Ctx) error {
		response := resources.Response{}
		authCookieName := os.Getenv("AUTH_COOKIE_NAME")
		token := c.Cookies(authCookieName)
		if token == "" {
			response.FromError(fiber.ErrUnauthorized)
			return c.Send(response.Bytes())
		}

		signature := os.Getenv("APP_KEY")
		tokenizer, err := jwth.Decode(signature, token) // parse jwt token from cookie access token
		claims, ok := tokenizer.Claims.(jwt.MapClaims)  // get jwt claims
		userMap := claims["user"].(map[string]any)
		c.Locals("user", models.User{
			Id:    uuid.MustParse(userMap["id"].(string)),
			Name:  userMap["name"].(string),
			Email: userMap["email"].(string),
		}) // set claims user into context

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				response.FromError(fiber.ErrUnauthorized)
				return c.Send(response.Bytes())
			case jwt.ValidationErrorExpired:
				response.FromError(exceptions.AuthSessionExpired)
				return c.Send(response.Bytes())
			default:
				response.FromError(fiber.ErrInternalServerError)
				return c.Send(response.Bytes())
			}
		} else if !ok || !tokenizer.Valid {
			response.FromError(fiber.ErrInternalServerError)
			return c.Send(response.Bytes())
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
		excludedPaths := []string{
			// Auth routes
			"login", "register", "forgot-password", "reset-password",
			// General routes
			"public/",
		}

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
