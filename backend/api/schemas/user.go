package schemas

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	RoleID         uint `gorm:"not null"`
	EnterpriseID   *uint
	Name           string `gorm:"not null"`
	Email          string `gorm:"uniqueIndex;not null"`
	Password       string `gorm:"not null"`
	Cpf            string `gorm:"uniqueIndex;not null"`
	RegisterNumber uint   `gorm:"not null"`
	PhotoUrl       *string

	Role            Role          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Enterprise      Enterprise    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TokenPassword   TokenPassword `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	WishListEntries []WishList    `gorm:"foreignKey:UserID"`
}
