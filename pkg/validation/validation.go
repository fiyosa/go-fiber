package validation

import (
	"fmt"
	"go-fiber/pkg/secret"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	id_translations "github.com/go-playground/validator/v10/translations/id"
)

var (
	Translator ut.Translator
	Validate   *validator.Validate
)

func Setup() error {
	locale := secret.APP_LOCALE
	en := en.New()
	uni := ut.New(en, en, id.New())

	var found bool
	Translator, found = uni.GetTranslator(locale)
	if !found {
		return fmt.Errorf("translator for locale %s not found", locale)
	}

	Validate = validator.New()

	switch locale {
	case "en":
		return en_translations.RegisterDefaultTranslations(Validate, Translator)
	case "id":
		return id_translations.RegisterDefaultTranslations(Validate, Translator)
	default:
		return en_translations.RegisterDefaultTranslations(Validate, Translator)
	}
}
