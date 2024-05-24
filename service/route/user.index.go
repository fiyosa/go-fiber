package route

import (
	"go-fiber/lang"
	"go-fiber/pkg/db"
	"go-fiber/pkg/hash"
	"go-fiber/pkg/helper"
	"go-fiber/service/dto"

	"github.com/gofiber/fiber/v2"
)

// @Summary 	Get users
// @Description Get users
// @Tags 		User
// @Accept 		json
// @Produce 	json
// @Param 		page query int false "page"
// @Param 		limit query int false "limit"
// @Param 		keyword query string false "keyword"
// @Param 		orderBy query string false "orderBy"
// @Param 		sortedBy query string false "sortedBy" Enums(asc, desc)
// @Success 	200 {object} dto.UserIndexResponse "ok"
// @Security	BearerAuth
// @Router 		/user [get]
func UserIndex(c *fiber.Ctx) error {
	user := &db.User{}
	if ok, err := user.GetUser(c); !ok {
		return err
	}

	query := helper.QueryStr(c)

	users := []db.User{}
	db.G.
		Offset(helper.Offset(query.Page, query.Limit)).
		Limit(query.Limit).
		Order(query.OrderBy+" "+query.SortedBy).
		Where("username like ?", "%"+query.Keyword+"%").
		Or("name like ?", "%"+query.Keyword+"%").
		Find(&users)

	var countUsers int64
	db.G.
		Model(&db.User{}).
		Where("username like ?", "%"+query.Keyword+"%").
		Or("name like ?", "%"+query.Keyword+"%").
		Count(&countUsers)

	newUsers := []dto.UserIndexDataResponse{}
	for _, v := range users {
		id, _ := hash.Encode(v.Id)
		newUsers = append(newUsers, dto.UserIndexDataResponse{
			Id:        id,
			Username:  v.Username,
			Name:      v.Name,
			CreatedAt: helper.Time2Str(v.CreatedAt),
			UpdatedAt: helper.Time2Str(v.UpdatedAt),
		})
	}

	return helper.SendDatas(
		c,
		lang.L(lang.SetL().RETRIEVED_SUCCESSFULLY, fiber.Map{"operator": lang.SetL().USER}),
		newUsers,
		helper.Paginate{
			Page:  query.Page,
			Limit: query.Limit,
			Total: int(countUsers),
		},
	)
}
