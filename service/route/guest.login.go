package route

import (
	"errors"
	"go-fiber/lang"
	"go-fiber/pkg/db"
	"go-fiber/pkg/hash"
	"go-fiber/pkg/helper"
	"go-fiber/pkg/jwt"
	"go-fiber/service/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary 	Login
// @Description Login
// @Tags 		Guest
// @Accept 		json
// @Produce 	json
// @Param		payload body dto.GuestLoginRequest true "payload"
// @Success 	200 {object} dto.GuestLoginResponse "ok"
// @Router 		/auth/login [post]
func GuestLogin(c *fiber.Ctx) error {
	validated := &dto.GuestLoginRequest{}
	if check, err := helper.Validate(c, validated); check {
		return err
	}

	user := db.User{}
	if err := db.G.Where(&db.User{Username: validated.Username}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helper.SendError(c, lang.L(lang.SetL().AUTH_NOT_FOUND, nil), http.StatusNotFound)
		}
		return helper.SendError(c, lang.L(lang.SetL().AUTH_FAILED, nil))
	}
	if !hash.Verify(validated.Password, user.Password) {
		return helper.SendError(c, lang.L(lang.SetL().AUTH_FAILED, nil))
	}

	token, err := jwt.Create(helper.Int2Str(user.Id))
	if err != nil {
		return helper.SendError(c, lang.L(err.Error(), nil))
	}

	auth := &db.Auth{
		UserId: user.Id,
		Token:  token,
	}

	auth.Create()

	return c.Status(fiber.StatusOK).JSON(dto.GuestLoginResponse{
		Token: token,
	})
}
