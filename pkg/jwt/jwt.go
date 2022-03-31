package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const TokenExpireDuration = time.Hour * 2

var mySecret = []byte("hello world")

type MyClaims struct {
	UserID   int64  `json:"userId"`
	Username string `json:"username"`

	jwt.StandardClaims
}

func GenToken(userId int64, username string) (string, error) {
	c := MyClaims{
		userId,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "my-project",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, c)
	return token.SignedString(mySecret)
}

// 解析token数据
func ParseToken(tokenString string) (*MyClaims, error) {
	var mc = new(MyClaims)

	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("无效的token")
	}
	return mc, nil
}
