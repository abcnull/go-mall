package entity

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID    uint `gorm:"not null"`
	ProductID uint `gorm:"not null"`
	BossID    uint `gorm:"not null"`

	Num    int  `gorm:"not null"`
	MaxNum int  `gorm:"not null"`
	Check  bool `gorm:"default:false"`
}
