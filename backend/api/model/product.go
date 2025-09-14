package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `gorm:"unique;not null"`
	Description string  `gorm:"not null"`
	Value       float64 `gorm:"type:decimal(10,2);not null"`
	Quantity    int     `gorm:"not null"`
	IsAvaible   bool    `gorm:"default:true;not null"`

	Promotions []Promotion `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
