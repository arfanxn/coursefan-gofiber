package validators

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// EnglishTranslator returns validator english translator
func EnglishTranslator(validate *validator.Validate) (trans ut.Translator) {
	english := en.New()
	uni := ut.New(english, english)
	trans, _ = uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, trans)
	return
}
