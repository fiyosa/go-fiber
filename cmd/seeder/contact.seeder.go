package seeder

import (
	"go-fiber/pkg/db"

	"gorm.io/gorm"
)

func ContactSeeder(g *gorm.DB) {
	var users []*db.User

	contacts := []*db.Contact{
		{FirstName: "admin", LastName: "admin", Email: "admin@gmail.com", Phone: "081234567890"},
		{FirstName: "user", LastName: "user", Email: "user@gmail.com", Phone: "081234567890"},
	}

	g.Find(&users)

	for i, user := range users {
		contacts[i].UserId = user.Id
	}

	g.Create(&contacts)
}
