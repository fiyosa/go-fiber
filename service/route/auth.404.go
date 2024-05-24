package route

import (
	"go-fiber/lang"

	"github.com/gofiber/fiber/v2"
)

func AuthNotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": lang.L(lang.SetL().API_NOT_FOUND, nil),
	})
}
