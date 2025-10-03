package schemas

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name               string   `gorm:"uniqueIndex;not null"`
	Description        string   `gorm:"not null"`
	Value              float64  `gorm:"type:decimal(10,2);not null"`
	Quantity           int      `gorm:"not null"`
	IsPromotionAvaible bool     `gorm:"not null"`
	Discount           *float64 `gorm:"type:decimal(10,2)"`
	PhotoUrl           *string
	IsAvaible          bool `gorm:"default:true;not null"`

	WishListEntries []WishList `gorm:"foreignKey:ProductID"`
}

func (s *Product) validateProduct() error {
	return nil
}

func (s *Product) formatProduct() {

}
