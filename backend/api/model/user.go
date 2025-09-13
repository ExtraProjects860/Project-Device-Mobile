package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	TypeUserID     uint   `gorm:"not null"`
	Name           string `gorm:"not null"`
	Email          string `gorm:"unique;not null"`
	Password       string `gorm:"not null"`
	Cpf            string `gorm:"unique;not null"`
	RegisterNumber uint   `gorm:"not null"`
	PhotoUrl       *string

	TypeUser      TypeUser      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TokenPassword TokenPassword `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	WishLists     []WishList    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
