package strh

import (
	"fmt"
	"strings"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/iancoleman/strcase"
)

// StrToDelimetedCamel converts a string to Delimeted{delimeter}CamelCase
func StrToDelimetedCamel(s, delimeter string) string {
	return strings.Join(
		sliceh.Map(
			strings.Split(s, delimeter), func(s string) string {
				return strcase.ToCamel(s)
			},
		),
		delimeter,
	)
}

// StrToDelimetedSnake converts a string to delimeted{delimeter}snake_case
func StrToDelimetedSnake(s, delimeter string) string {
	return strings.Join(
		sliceh.Map(
			strings.Split(s, delimeter), func(s string) string {
				return strcase.ToSnake(s)
			},
		),
		delimeter,
	)
}

// AnyToStr converts any to string
func AnyToStr(value any) string { return fmt.Sprintf("%v", value) }

// StrToAny converts string to any
func StrToAny(value string) any { return any(value) }
