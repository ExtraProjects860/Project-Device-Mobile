package schemas

import (
	"gorm.io/gorm"
)

type Enterprise struct {
	gorm.Model
	Name string `gorm:"uniqueIndex;not null" validate:"required,min=3"`
}
