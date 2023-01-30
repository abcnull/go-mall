package util

import (
	"github.com/dgrijalva/jwt-go"
	"go-mall/biz/model"
	"time"
)

var jwtSecret = []byte("jwt_jiamichuan")

type Claims struct { // todo
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

// GenerateToken 签发 token
func GenerateToken(id uint, username string, auth int) (string, error) {
	nowTime := time.Now()                          // 签发时间
	expireTime := nowTime.Add(time.Hour * 24 * 30) // 过期时间 1M

	claims := &Claims{
		ID:        id,
		UserName:  username,
		Authority: auth,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // todo
		},
	}
	claimsWithMethod := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 先设定 hash 加密 method
	token, err := claimsWithMethod.SignedString(jwtSecret)                // 拿密钥通过 hs256 加密得到 token
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	claimsToken, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || claimsToken == nil {
		return nil, err
	}

	claims, ok := claimsToken.Claims.(*Claims)
	if ok && claimsToken.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

type EmailClaims struct {
	UserId        uint   `json:"user_id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	OperationType uint   `json:"operation_type"`
	jwt.StandardClaims
}

// GenerateEmailToken 签发 emailToken
func GenerateEmailToken(userId uint, email string, password string, operationType model.SendEmailType) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour * 24 * 30)
	emailClaims := EmailClaims{
		UserId:        userId,
		Email:         email,
		Password:      password,
		OperationType: uint(operationType),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	emailClaimsToken := jwt.NewWithClaims(jwt.SigningMethodHS256, emailClaims) // 设置签名方式
	emailToken, err := emailClaimsToken.SignedString(jwtSecret)                // 签名获取 token
	return emailToken, err
}

// ParseEmailToken 解密 emailToken
func ParseEmailToken(token string) (*EmailClaims, error) {
	claimsToken, err := jwt.ParseWithClaims(token, &EmailClaims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || claimsToken == nil {
		return nil, err
	}

	claims, ok := claimsToken.Claims.(*EmailClaims)
	if ok && claimsToken.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
