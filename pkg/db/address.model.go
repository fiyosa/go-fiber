package db

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	Id         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	ContactId  int       `json:"contact_id" `
	Street     string    `json:"street" gorm:"type:varchar(255)"`
	City       string    `json:"city" gorm:"type:varchar(255)"`
	Province   string    `json:"province" gorm:"type:varchar(255)"`
	Country    string    `json:"country" gorm:"type:varchar(255);not null"`
	PostalCode string    `json:"postal_code" gorm:"type:varchar(255);not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"type:timestamp(0);not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"type:timestamp(0);not null;default:CURRENT_TIMESTAMP"`

	Contact Contact `gorm:"foreignKey:ContactId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Address) TableName() string {
	return "addresses"
}

func (a *Address) Create() *gorm.DB {
	return G.Create(&a)
}

func (a *Address) Show(id int) *gorm.DB {
	return G.First(&a, id)
}

func (a *Address) Update() *gorm.DB {
	return G.Save(&a)
}

func (a *Address) Delete(id int) *gorm.DB {
	return G.Delete(&a, id)
}
