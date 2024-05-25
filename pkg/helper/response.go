package helper

import (
	"github.com/gofiber/fiber/v2"
)

type Paginate struct {
	Page  int `json:"page" example:"0"`
	Limit int `json:"limit" example:"0"`
	Total int `json:"total" example:"0"`
}

func SendCustom(c *fiber.Ctx, data interface{}, statusCode ...int) error {
	code := fiber.StatusOK // Default status code
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	return c.Status(code).JSON(data)
}

func SendSuccess(c *fiber.Ctx, msg string) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": msg,
	})
}

func SendData(c *fiber.Ctx, msg string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    data,
		"message": msg,
	})
}

func SendDatas(c *fiber.Ctx, msg string, data interface{}, paginate Paginate) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":       data,
		"pagination": paginate,
		"message":    msg,
	})
}

func SendError(c *fiber.Ctx, msg string, statusCode ...int) error {
	code := fiber.StatusBadRequest // Default status code
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	return c.Status(code).JSON(fiber.Map{
		"message": msg,
	})
}

func SendErrors(c *fiber.Ctx, msg string, err interface{}) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"errors":  err,
		"message": msg,
	})
}
