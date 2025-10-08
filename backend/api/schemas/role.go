package schemas

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name  string `gorm:"unique;not null" validate:"required,min=3"`
	Users []User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
