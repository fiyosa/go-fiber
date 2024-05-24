package main

import (
	"fmt"
	"go-fiber/cmd"
	"go-fiber/pkg/db"
	"go-fiber/pkg/secret"
	"go-fiber/pkg/validation"
	"go-fiber/router"
)

// @title		Service API
// @version		1.0
// @description Service API in Go useing Fiber Framework

// @securityDefinitions.apikey	BearerAuth
// @in 							header
// @name 						Authorization
// @description 				Type "Bearer" followed by a space and JWT token.
func main() {
	secret.Setup()

	db.Setup()

	if cmd.Setup() {
		return
	}

	if err := validation.Setup(); err != nil {
		fmt.Println(err.Error())
		return
	}

	r := router.Config()
	r = router.Setup(r)

	r.Listen(":" + secret.PORT)
}
