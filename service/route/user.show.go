package route

import (
	"go-fiber/lang"
	"go-fiber/pkg/db"
	"go-fiber/pkg/hash"
	"go-fiber/pkg/helper"
	"go-fiber/service/dto"

	"github.com/gofiber/fiber/v2"
)

// @Summary 	Get user by id
// @Description Get user by id
// @Tags 		User
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "id"
// @Success 	200 {object} dto.UserAuthResponse "ok"
// @Security	BearerAuth
// @Router 		/user/{id} [get]
func UserShow(c *fiber.Ctx) error {
	get_id := c.Params("id")
	user_id, err := hash.Decode(get_id)
	if err != nil {
		return helper.SendError(c, lang.L(lang.SetL().NOT_FOUND, fiber.Map{"operator": lang.SetL().USER}))
	}

	user := &db.User{}
	user.Show(user_id)
	if user.Id == 0 {
		return helper.SendError(c, lang.L(lang.SetL().NOT_FOUND, fiber.Map{"operator": lang.SetL().USER}))
	}

	id, _ := hash.Encode(user.Id)
	res := dto.UserAuthResponse{
		Data: dto.UserAuthDataResponse{
			Id:        id,
			Username:  user.Username,
			Name:      user.Name,
			CreatedAt: helper.Time2Str(user.CreatedAt),
			UpdatedAt: helper.Time2Str(user.UpdatedAt),
		},
		Message: lang.L(lang.SetL().RETRIEVED_SUCCESSFULLY, fiber.Map{"operator": lang.SetL().USER}),
	}

	return helper.SendCustom(c, res)
}
