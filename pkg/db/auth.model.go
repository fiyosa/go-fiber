package db

import (
	"go-fiber/pkg/secret"
	"time"

	"gorm.io/gorm"
)

type Auth struct {
	Id        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId    int       `json:"user_id"`
	Token     string    `json:"token" gorm:"type:varchar(255)"`
	Revoke    bool      `json:"revoke" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp(0);not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp(0);not null;default:CURRENT_TIMESTAMP"`

	User User `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Auth) TableName() string {
	return secret.DB_SCHEMA + "." + "auths"
}

func (a *Auth) Create() *gorm.DB {
	return G.Create(&a)
}