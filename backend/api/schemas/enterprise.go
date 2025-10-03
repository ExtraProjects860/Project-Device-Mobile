package schemas

import "gorm.io/gorm"

type Enterprise struct {
	gorm.Model
	Name string `gorm:"uniqueIndex;not null"`
}

func (s *Enterprise) validateEnterprise() error {
	return nil
}

func (s *Enterprise) formatEnterprise() {

}
