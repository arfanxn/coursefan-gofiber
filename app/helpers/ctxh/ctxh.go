package ctxh

import (
	"mime/multipart"
	"os"
	"regexp"

	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
)

// MustGetUser returns current logged in user in from the given context, or panic if not in the context
func MustGetUser(c *fiber.Ctx) models.User {
	userAny := c.Locals("user")
	if userAny == nil {
		panic(fiber.ErrNotFound)
	}
	return userAny.(models.User)
}

// GetAcceptLang returns the client's Accept language from request headers or query
func GetAcceptLang(c *fiber.Ctx) string {
	headerKey := "Accept-Language"
	queryKey := "lang"
	headers := c.GetReqHeaders()
	lang, ok := headers[headerKey]
	if lang != "" && ok {
		lang = regexp.MustCompile("[a-z]{2}").FindString(lang)
	}
	if lang == "" {
		defaultAppLang := os.Getenv("APP_LANGUAGE")
		lang = c.Query(queryKey, defaultAppLang)
	}
	return lang
}

// GetFileHeader returns file header from the given ctx with the given key or nil if no file header provided
func GetFileHeader(c *fiber.Ctx, key string) (fileHeader *multipart.FileHeader) {
	fileHeader, err := c.FormFile(key)
	if err != nil {
		fileHeader = nil
		return
	}
	return
}

// GetFullURIString get current URI string
func GetFullURIString(c *fiber.Ctx) string {
	return string(c.Request().URI().FullURI())
}
