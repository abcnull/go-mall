package entity

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID    uint `gorm:"not null"`
	ProductID uint `gorm:"not null"`
	BossID    uint `gorm:"not null"`
	AddressID uint `gorm:"not null"`
	Num       int  `gorm:"not null"` // todo: 后续调整数量，没有数量
	OrderNum  int  // todo: 后续改下表里头的 schema
	Type      uint // 1 未支付 2 已支付 // todo: 后续调整成 int
	Money     float64
}
