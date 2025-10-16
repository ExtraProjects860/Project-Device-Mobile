package schemas

import (
	"time"

	"gorm.io/gorm"
)

type TokenPassword struct {
	gorm.Model
	UserID uint `gorm:"unique;not null"`
	Code   *string
	TimeUp *time.Time
	User   *User
}
