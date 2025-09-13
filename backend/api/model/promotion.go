package model

import "gorm.io/gorm"

type Promotion struct {
	gorm.Model
	ProductID uint    `gorm:"not null"`
	Product   Product `gorm:"not null"`
	Discount  *float64
	IsAvaible bool `gorm:"default:true;not null"`
}
