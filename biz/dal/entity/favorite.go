package entity

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	User   User
	UserID uint `gorm:"not null"`

	Product   Product
	ProductID uint `gorm:"not null"`

	Boss   User
	BossID uint `gorm:"not null"`
}
