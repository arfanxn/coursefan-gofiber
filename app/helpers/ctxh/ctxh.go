package ctxh

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

// GetAcceptLang returns the client's Accept language from request headers or query
func GetAcceptLang(c *fiber.Ctx) string {
	headerKey := "Accept-Language"
	queryKey := "lang"
	headers := c.GetReqHeaders()
	lang, ok := headers[headerKey]
	if !ok || lang == "" {
		defaultAppLang := os.Getenv("APP_LANGUAGE")
		lang = c.Query(queryKey, defaultAppLang)
	}
	return lang
}
