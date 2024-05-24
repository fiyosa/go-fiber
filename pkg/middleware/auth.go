package middleware

import (
	"go-fiber/lang"
	"go-fiber/pkg/db"
	"go-fiber/pkg/helper"
	"go-fiber/pkg/jwt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Auth(permission ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := db.User{}
		if ok, err := authentication(c, &user); !ok {
			return err
		}
		if ok, err := authorization(c, &user, permission...); !ok {
			return err
		}
		return c.Next()
	}
}

func authorization(c *fiber.Ctx, user *db.User, permissions ...string) (bool, error) {
	if len(permissions) == 0 {
		return true, nil
	}

	getRoles := []string{}
	r := &db.Role{}
	r.GetRoles(user.Id, &getRoles)

	getPermissions := []string{}
	p := &db.Permission{}
	p.GetPermissions(getRoles, &getPermissions)

	check := false
	for _, v := range getPermissions {
		if v == permissions[0] {
			check = true
		}
	}

	if !check {
		return false, helper.SendError(c, lang.L(lang.SetL().PERMISSION_FAILED, nil))
	}

	return true, nil
}

func authentication(c *fiber.Ctx, user *db.User) (bool, error) {
	getToken := c.Get("Authorization")

	if getToken == "" {
		return false, helper.SendError(c, lang.L(lang.SetL().UNAUTHORIZED_ACCESS, nil))
	}

	tokenParts := strings.Split(getToken, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return false, helper.SendError(c, lang.L(lang.SetL().UNAUTHORIZED_ACCESS, nil))
	}

	token := tokenParts[1]

	_, err := jwt.Verify(c, token)

	if err != nil {
		return false, helper.SendError(c, lang.L(lang.SetL().UNAUTHORIZED_ACCESS, nil))
	}

	auth := db.Auth{}
	db.G.Preload("User").Where(&db.Auth{Token: token}).First(&auth)
	if auth.Id == 0 {
		return false, helper.SendError(c, lang.L(lang.SetL().UNAUTHORIZED_ACCESS, nil))
	}

	*user = auth.User
	c.Locals("user", auth.User)
	return true, nil
}
