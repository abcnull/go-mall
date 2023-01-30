package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Num                int `gorm:"not null"`
	Name               string
	CategoryId         uint
	Title              string
	Info               string
	ImgPath            string
	Price              string
	DiscountPrice      string
	Locate             string
	Freight            uint
	OnSale             bool `gorm:"default:false"`
	BossID             uint `gorm:"not null"`
	BossName           string
	BossAvatar         string
	ExposureTimes      int64
	ClickTimes         int64
	CommunicationTimes int64
}
