package validationh

import (
	"strings"

	"github.com/arfanxn/coursefan-gofiber/app/exceptions"
	validator_provider "github.com/arfanxn/coursefan-gofiber/app/providers/validators"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
)

func ValidateStruct[T any](structure T, lang string) []*exceptions.ValidationError {
	validate := validator.New()
	err := validate.Struct(structure)
	lang = strings.ToLower(lang)

	translators := map[string]func(*validator.Validate) ut.Translator{
		"en": validator_provider.EnglishTranslator,
	}
	translator := translators[lang]
	utTrans := translator(validate)
	return TranslateErrs(err, utTrans)
}

// TranslateErrs translates errors from validation errors
func TranslateErrs(errs error, trans ut.Translator) (translatedErrs []*exceptions.ValidationError) {
	if errs == nil {
		return nil
	}
	validationErrs := errs.(validator.ValidationErrors)
	for _, validationErr := range validationErrs {
		fieldName := strings.Join(strings.SplitAfter(validationErr.StructNamespace(), ".")[1:], ".")
		fieldName = strcase.ToSnake(fieldName)
		message := validationErr.Translate(trans)
		translatedErr := exceptions.NewValidationError(fieldName, message)
		translatedErrs = append(translatedErrs, translatedErr)
	}
	return
}
