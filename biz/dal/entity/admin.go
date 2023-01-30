package entity

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	Avatar         string
}
