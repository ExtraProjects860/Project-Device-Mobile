package schemas

import (
	"time"

	"gorm.io/gorm"
)

type WishList struct {
	gorm.Model
	UserID    uint           `gorm:"primaryKey"`
	ProductID uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"default:current_timestamp"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`

	User    User
	Product Product
}
