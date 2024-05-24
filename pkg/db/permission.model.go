package db

import (
	"go-fiber/pkg/secret"
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	Id        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"type:varchar(255);unique"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp(0);not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp(0);not null;default:CURRENT_TIMESTAMP"`

	RoleHashPermission []RoleHashPermission `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Permission) TableName() string {
	return secret.DB_SCHEMA + "." + "permissions"
}

func (p *Permission) GetPermissions(roles []string, permissions *[]string) *gorm.DB {
	return G.
		Table("permissions AS p").
		Select("p.name").
		Joins("LEFT JOIN role_has_permissions AS rhp ON rhp.permission_id = p.id").
		Joins("LEFT JOIN roles AS r ON r.id = rhp.role_id").
		Where("r.name IN ?", roles).
		Scan(&permissions)
}
