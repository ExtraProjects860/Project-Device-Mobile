package model

import "gorm.io/gorm"

type TypeUser struct {
	gorm.Model
	Name  string `gorm:"unique;not null"`
	Users []User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
