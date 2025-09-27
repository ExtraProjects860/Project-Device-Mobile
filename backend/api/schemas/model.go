package schemas

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	RoleID         uint `gorm:"not null"`
	EnterpriseID   *uint
	Name           string `gorm:"not null"`
	Email          string `gorm:"uniqueIndex;not null"`
	Password       string `gorm:"not null"`
	Cpf            string `gorm:"uniqueIndex;not null"`
	RegisterNumber uint   `gorm:"not null"`
	PhotoUrl       *string

	Role            Role          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Enterprise      Enterprise    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TokenPassword   TokenPassword `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	WishListEntries []WishList    `gorm:"foreignKey:UserID"`
}

type Enterprise struct {
	gorm.Model
	Name string `gorm:"uniqueIndex;not null"`
}

type TokenPassword struct {
	gorm.Model
	UserID uint `gorm:"unique;not null"`
	Code   *string
	TimeUp *time.Time
	User   *User
}

type WishList struct {
	UserID    uint      `gorm:"primaryKey"`
	ProductID uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	User    User
	Product Product
}

type Role struct {
	gorm.Model
	Name  string `gorm:"unique;not null"`
	Users []User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

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

func AllModelsSlice() []any {
	return []any{
		&User{},
		&TokenPassword{},
		&Enterprise{},
		&Role{},
		&Product{},
		&WishList{},
	}
}
