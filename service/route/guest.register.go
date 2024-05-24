package route

import (
	"go-fiber/lang"
	"go-fiber/pkg/db"
	"go-fiber/pkg/hash"
	"go-fiber/pkg/helper"
	"go-fiber/service/dto"

	"github.com/gofiber/fiber/v2"
)

// @Summary 	Register
// @Description Register
// @Tags 		Guest
// @Accept 		json
// @Produce 	json
// @Param		payload body dto.GuestRegisterRequest true "payload"
// @Success 	200 {object} dto.GuestRegisterDataResponse "ok"
// @Router 		/auth/register [post]
func GuestRegister(c *fiber.Ctx) error {
	validated := &dto.GuestRegisterRequest{}
	if check, err := helper.Validate(c, validated); check {
		return err
	}

	user := &db.User{}
	db.G.Where(&db.User{Username: validated.Username}).First(&user)
	if user.Id != 0 {
		return helper.SendError(
			c,
			lang.L(lang.SetL().ALREADY_EXIST, fiber.Map{"operator": "Username"}),
		)
	}

	newPassword, err := hash.Create(validated.Password)
	if err != nil {
		return helper.SendError(c, err.Error())
	}

	user.Username = validated.Username
	user.Name = validated.Name
	user.Password = newPassword
	user.Create()

	id, _ := hash.Encode(user.Id)
	return helper.SendData(
		c,
		lang.L(lang.SetL().SAVED_SUCCESSFULLY, fiber.Map{"operator": lang.SetL().USER}),
		dto.GuestRegisterDataResponse{
			Id:        id,
			Username:  user.Username,
			Name:      user.Name,
			CreatedAt: helper.Time2Str(user.CreatedAt),
			UpdatedAt: helper.Time2Str(user.UpdatedAt),
		},
	)
}
