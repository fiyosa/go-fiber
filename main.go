package main

import (
	"go-fiber/cmd"
	"go-fiber/pkg/db"
	"go-fiber/pkg/logger"
	"go-fiber/pkg/secret"
	"go-fiber/pkg/validation"
	"go-fiber/router"

	"github.com/gofiber/fiber/v2/log"
)

// @title		Service API
// @version		1.0
// @description Service API in Go using Fiber Framework

// @securityDefinitions.apikey	BearerAuth
// @in 							header
// @name 						Authorization
// @description 				Type "Bearer" followed by a space and JWT token.
func main() {
	secret.Setup()

	if logger.Setup() {
		return
	}

	db.Setup()

	if cmd.Setup() {
		return
	}

	if err := validation.Setup(); err != nil {
		log.Error(err.Error())
		return
	}

	r := router.Config()
	r = router.Setup(r)

	r.Listen(":" + secret.PORT)
}
