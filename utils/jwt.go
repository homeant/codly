package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/homeanter/codly/config"
	"time"
)

type Claims struct {
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(userID uint) (string, error) {
	claims := &Claims{
		UserId: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token过期时间为24小时
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Config.JwtSecretKey))
}
