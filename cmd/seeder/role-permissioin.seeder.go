package seeder

import (
	"go-fiber/pkg/db"

	"gorm.io/gorm"
)

var roles = []string{
	"admin",
	"user",
}

var permissions = []string{
	"user_index",
	"user_show",
}

func RolePermissionSeeder(g *gorm.DB) {
	createRoles := []*db.Role{}
	for _, v := range roles {
		createRoles = append(createRoles, &db.Role{Name: v})
	}

	createPermissions := []*db.Permission{}
	for _, v := range permissions {
		createPermissions = append(createPermissions, &db.Permission{Name: v})
	}

	g.Create(&createRoles)
	g.Create(&createPermissions)

	createRoleHasPermissions := []*db.RoleHashPermission{}

	createRoleHasPermissions = append(createRoleHasPermissions, createAdmin(createRoles, createPermissions)...)
	createRoleHasPermissions = append(createRoleHasPermissions, createUser(createRoles, createPermissions)...)

	g.Create(&createRoleHasPermissions)
}

func createAdmin(cr []*db.Role, cp []*db.Permission) []*db.RoleHashPermission {
	roleName := "admin"

	var roleID int
	for _, v := range cr {
		if v.Name == roleName {
			roleID = v.Id
			break
		}
	}

	crhp := []*db.RoleHashPermission{}
	for _, v := range cp {
		crhp = append(crhp, &db.RoleHashPermission{
			RoleId:       roleID,
			PermissionId: v.Id,
		})
	}
	return crhp
}

func createUser(cr []*db.Role, cp []*db.Permission) []*db.RoleHashPermission {
	roleName := "user"
	permissions := []string{
		"user_show",
	}

	var roleID int
	for _, v := range cr {
		if v.Name == roleName {
			roleID = v.Id
			break
		}
	}

	crhp := []*db.RoleHashPermission{}
	for _, v := range cp {
		for _, p := range permissions {
			if p == v.Name {
				crhp = append(crhp, &db.RoleHashPermission{
					RoleId:       roleID,
					PermissionId: v.Id,
				})
			}
		}
	}
	return crhp
}
