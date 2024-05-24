package cmd

import (
	"flag"
	"go-fiber/pkg/db"
	"go-fiber/pkg/secret"

	"github.com/gofiber/fiber/v2/log"
)

func Setup() bool {
	if secret.APP_ENV != "development" {
		log.Error("Cannot access cmd while mode development")
		return true
	}

	dropFlag := flag.Bool("drop", false, "Drop the database tables")
	seedFlag := flag.Bool("seed", false, "Seed the database with initial data")
	migrateFlag := flag.Bool("migrate", false, "Run database migrations")

	flag.Parse()
	status := false

	if *dropFlag {
		RunDropAllTable(db.G)
		status = true
	}

	if *migrateFlag {
		RunMigrate(db.G)
		status = true
	}

	if *seedFlag {
		RunSeeder(db.G)
		status = true
	}

	return status
}
