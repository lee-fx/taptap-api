package utils

import (
	"api/admin/defs"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Id       int64  `json:"id"`
	UserName string `json:"userName"`
	Iphone   string `json:"iphone"`
	jwt.StandardClaims
}

// 生成jwt-token
func GenerateToken(user defs.User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(100 * time.Second) // 失效 24小时
	issuer := "lee-fx"
	claims := Claims{
		Id:       user.Id,
		UserName: user.Name,
		Iphone:   user.Iphone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("lfx-pupa"))
	return token, err
}

// 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("lfx-pupa"), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
