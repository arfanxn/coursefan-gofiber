package bootstrap

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// NewValidator bootstraps application validator
func NewValidator() (*validator.Validate, error) {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}
		return name
	})

	return validate, nil
}
