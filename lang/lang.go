package lang

import (
	"fmt"
	"go-fiber/pkg/secret"
	"strings"
)

func L(msg string, args map[string]any) string {
	newMsg := msg
	for key, value := range args {
		newMsg = strings.ReplaceAll(newMsg, ":"+key, fmt.Sprintf("%v", value))
	}
	return newMsg
}

func SetL() ILang {
	return locale[secret.APP_LOCALE]
}

var locale = map[string]ILang{
	"en": EN,
	"id": ID,
}

type ILang struct {
	USER   string
	POST   string
	ROLE   string
	LOGOUT string

	RETRIEVED_SUCCESSFULLY string
	SAVED_SUCCESSFULLY     string
	UPDATED_SUCCESSFULLY   string
	DELETED_SUCCESSFULLY   string
	ALREADY_EXIST          string

	UNAUTHORIZED_ACCESS  string
	PERMISSION_FAILED    string
	AUTH_NOT_FOUND       string
	AUTH_FAILED          string
	NOT_FOUND            string
	SOMETHING_WENT_WRONG string
	API_NOT_FOUND        string
	HASH_FAILED          string
}
