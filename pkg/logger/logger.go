package logger

import (
	"go-fiber/pkg/secret"
	"io"
	"os"

	"github.com/gofiber/fiber/v2/log"
)

const out = "file"

func Setup() bool {
	if secret.APP_ENV == "development" {
		_, err := os.Stat("logs")
		if os.IsNotExist(err) {
			err := os.MkdirAll("logs", 0755)
			if err != nil {
				log.Fatalf("error creating directory logs: %v", err)
				return true
			}
		}

		file, err := os.OpenFile("logs/fiber.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
			return true
		}

		if out == "file" {
			log.SetOutput(file)
		} else {
			iw := io.MultiWriter(os.Stdout, file)
			log.SetOutput(iw)
		}
	}
	return false
}
