package helper

import (
	"encoding/json"
	"go-fiber/lang"
	"go-fiber/pkg/secret"
	"go-fiber/pkg/validation"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Validate(c *fiber.Ctx, input interface{}) (bool, error) {
	if err := c.BodyParser(input); err != nil {
		return true, generateError(c, err)
	}

	if err := validation.Validate.Struct(input); err != nil {
		return true, generateError(c, err)
	}

	return false, nil
}

func generateError(c *fiber.Ctx, err error) error {
	newErrors := map[string]string{}
	msg := ""

	switch v := err.(type) {
	case *json.UnmarshalTypeError:
		newMsg := "Json binding error: " + v.Field + " type error"
		newErrors[v.Field] = newMsg
		msg = newMsg

	case validator.ValidationErrors:
		for _, e := range v {
			newMsg := e.Translate(validation.Translator)
			newErrors[e.Field()] = newMsg
			if msg == "" {
				msg = newMsg
			}
		}

	default:
		if secret.APP_ENV == "development" {
			msg = v.Error()
		} else {
			msg = lang.L(lang.SetL().SOMETHING_WENT_WRONG, nil)
		}
	}

	return SendErrors(c, msg, newErrors)
}
