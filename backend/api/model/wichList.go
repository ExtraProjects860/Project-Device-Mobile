package model

import "gorm.io/gorm"

// dps vou ter que usar aqui um unique de junção
type WishList struct {
	gorm.Model
	UserID    uint `gorm:"not null"`
	ProductID uint `gorm:"not null"`

	User    User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Product Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (WishList) TableName() string {
	return "wish_lists"
}
