package db

import (
	"time"

	"gorm.io/gorm"
)

type Contact struct {
	Id        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId    int       `json:"user_id"`
	FirstName string    `json:"first_name" gorm:"type:varchar(255);not null"`
	LastName  string    `json:"last_name" gorm:"type:varchar(255)"`
	Email     string    `json:"password" gorm:"type:varchar(255);not null"`
	Phone     string    `json:"phone" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp(0);not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp(0);not null;default:CURRENT_TIMESTAMP"`

	User      User      `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Addresses []Address `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Contact) TableName() string {
	return "contacts"
}

func (c *Contact) Create() *gorm.DB {
	return G.Create(&c)
}

func (c *Contact) Show(id int) *gorm.DB {
	return G.First(&c, id)
}

func (c *Contact) Update() *gorm.DB {
	return G.Save(&c)
}

func (c *Contact) Delete(id int) *gorm.DB {
	return G.Delete(&c, id)
}
