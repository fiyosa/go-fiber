package route

import (
	"go-fiber/lang"
	"go-fiber/pkg/db"
	"go-fiber/pkg/hash"
	"go-fiber/pkg/helper"
	"go-fiber/service/dto"

	"github.com/gofiber/fiber/v2"
)

// @Summary 	Get user by auth
// @Description Get user by auth
// @Tags 		User
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} dto.UserAuthDataResponse "ok"
// @Security	BearerAuth
// @Router 		/auth/user [get]
func UserAuth(c *fiber.Ctx) error {
	user := &db.User{}
	if ok, err := user.GetUser(c); !ok {
		return err
	}

	id, _ := hash.Encode(user.Id)
	return helper.SendData(
		c,
		lang.L(lang.SetL().RETRIEVED_SUCCESSFULLY, fiber.Map{"operator": lang.SetL().USER}),
		dto.UserAuthDataResponse{
			Id:        id,
			Username:  user.Username,
			Name:      user.Name,
			CreatedAt: helper.Time2Str(user.CreatedAt),
			UpdatedAt: helper.Time2Str(user.UpdatedAt),
		},
	)
}
