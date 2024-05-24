package db

import (
	"go-fiber/lang"
	"go-fiber/pkg/helper"
	"go-fiber/pkg/secret"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type User struct {
	Id        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string    `json:"username" gorm:"type:varchar(255);not null;unique"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp(0);not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp(0);not null;default:CURRENT_TIMESTAMP"`

	UserHasRole []UserHasRole `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Auth        []Auth        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Contacts    []Contact     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (User) TableName() string {
	return secret.DB_SCHEMA + "." + "users"
}

func (u *User) GetUser(c *fiber.Ctx) (bool, error) {
	user := c.Locals("user")
	if user == "" {
		return false, helper.SendError(c, lang.L(lang.SetL().UNAUTHORIZED_ACCESS, nil))
	}
	userObj, ok := user.(User)
	if !ok {
		return false, helper.SendError(c, lang.L(lang.SetL().UNAUTHORIZED_ACCESS, nil))
	}
	*u = userObj
	return true, nil
}

func (u *User) Create() *gorm.DB {
	return G.Create(&u)
}

func (u *User) Show(id int) *gorm.DB {
	return G.First(&u, id)
}

func (u *User) Update() *gorm.DB {
	return G.Save(&u)
}

func (u *User) Delete(id int) *gorm.DB {
	return G.Delete(&u, id)
}
