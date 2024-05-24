package router

import (
	"go-fiber/docs"
	"go-fiber/pkg/middleware"
	"go-fiber/pkg/secret"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/gofiber/template/html/v2"
)

func Config() *fiber.App {
	r := fiber.New(fiber.Config{
		Views: html.New("./service/view", ".html"),
	})
	r.Use(logger.New())
	r.Use(recover.New())

	r.Use(middleware.Cors())

	r.Get("/", func(c *fiber.Ctx) error { return c.Render("welcome", nil) })

	docs.SwaggerInfo.Host = secret.APP_URL
	docs.SwaggerInfo.BasePath = "/api"
	r.Get("/swagger/*", swagger.New(swagger.Config{
		DocExpansion:             "none",
		DefaultModelsExpandDepth: -1,
	}))

	return r
}
