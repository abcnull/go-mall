package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string
	PasswordDigest string
	NickName       string
	Status         UserStatus // 0 正常状态; 1 临时封控; 2 永久黑名单用户; 3 账号停用注销用户
	Avatar         string
	Money          string
}

type UserStatus uint8

const (
	DefaultUserStatus     UserStatus = 0 // 常规用户
	SealingUserStatus     UserStatus = 1 // 封控用户
	BlackListUserStatus   UserStatus = 2 // 黑名单用户
	DeactivatedUserStatus UserStatus = 3 // 账号停用用户
)

func (user *User) SetPwd(pwd string) error { // todo
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 12) // 密码加密难度 12
	if err != nil {
		return err
	}

	user.PasswordDigest = string(hash)
	return nil
}

func (user *User) CheckPwd(pwd string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(pwd)) // todo
	if err != nil {
		return false, err
	}

	return true, nil
}
