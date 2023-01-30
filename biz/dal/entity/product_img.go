package entity

import "gorm.io/gorm"

type ProductImg struct {
	gorm.Model
	ProductID uint `gorm:"not null"`
	ImgPath   string
}
