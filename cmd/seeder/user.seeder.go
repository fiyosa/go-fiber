package seeder

import (
	"go-fiber/pkg/db"
	"go-fiber/pkg/hash"

	"gorm.io/gorm"
)

func UserSeeder(g *gorm.DB) {
	password, _ := hash.Create("Password")

	users := []*db.User{
		{Username: "admin", Name: "Admin", Password: password},
		{Username: "user", Name: "User", Password: password},
	}

	g.Create(&users)

	roles := []*db.Role{}
	db.G.Find(&roles)

	userHasRoles := []*db.UserHasRole{}
	for _, u := range users {
		for _, r := range roles {
			if u.Username == "admin" && r.Name == "admin" {
				userHasRoles = append(userHasRoles, &db.UserHasRole{UserId: u.Id, RoleId: r.Id})
			}
			if u.Username == "user" && r.Name == "user" {
				userHasRoles = append(userHasRoles, &db.UserHasRole{UserId: u.Id, RoleId: r.Id})
			}
		}
	}
	db.G.Create(&userHasRoles)
}
