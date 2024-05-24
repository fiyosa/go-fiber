package router

import (
	"go-fiber/pkg/middleware"
	"go-fiber/service/route"

	"github.com/gofiber/fiber/v2"
)

func Setup(r *fiber.App) *fiber.App {
	auth := middleware.Auth
	p := r.Group("/api")

	p.Post("/auth/register", route.GuestRegister)
	p.Post("/auth/login", route.AuthLogin)
	p.Get("/auth/secret", route.AuthSecret)

	p.Get("/auth/user", auth(), route.UserAuth)
	p.Get("/user", auth("user_index"), route.UserIndex)

	r.Use(route.AuthNotFound)

	return r
}
