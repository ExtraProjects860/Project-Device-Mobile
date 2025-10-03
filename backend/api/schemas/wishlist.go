package schemas

import (
	"time"
)

type WishList struct {
	UserID    uint      `gorm:"primaryKey"`
	ProductID uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	User    User
	Product Product
}

func (s *WishList) validateWishList() error {
	return nil
}

func (s *WishList) formatWishList() {

}
