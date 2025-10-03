package schemas

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name  string `gorm:"unique;not null"`
	Users []User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (s *Role) validateRole() error {
	return nil
}

func (s *Role) formatRole() {

}
