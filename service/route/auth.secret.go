package route

import (
	"go-fiber/pkg/secret"

	"github.com/gofiber/fiber/v2"
)

func AuthSecret(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"PORT":       secret.PORT,
		"APP_ENV":    secret.APP_ENV,
		"APP_LOCALE": secret.APP_LOCALE,
		"APP_SECRET": secret.APP_SECRET,
		"APP_URL":    secret.APP_URL,

		"DB_HOST":    secret.DB_HOST,
		"DB_PORT":    secret.DB_PORT,
		"DB_NAME":    secret.DB_NAME,
		"DB_USER":    secret.DB_USER,
		"DB_PASS":    secret.DB_PASS,
		"DB_SCHEMA":  secret.DB_SCHEMA,
		"DB_SSLMODE": secret.DB_SSLMODE,
	})
}
