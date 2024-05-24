package db

import (
	"go-fiber/pkg/secret"
	"time"

	"gorm.io/gorm"
)

type Role struct {
	Id        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"type:varchar(255);unique"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp(0);not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp(0);not null;default:CURRENT_TIMESTAMP"`

	RoleHashPermission []RoleHashPermission `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Role) TableName() string {
	return secret.DB_SCHEMA + "." + "roles"
}

func (r *Role) GetRoles(user_id int, roles *[]string) *gorm.DB {
	return G.
		Table("users AS u").
		Select("r.name").
		Joins("LEFT JOIN user_has_roles AS uhr ON uhr.user_id = u.id").
		Joins("LEFT JOIN roles AS r ON r.id = uhr.role_id").
		Where("u.id = ?", user_id).
		Scan(&roles)
}
