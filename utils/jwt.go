package utils

import (
	"messageserver/utils/configure"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(configure.JwtSecret)

type Claims struct {
	Phone string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(phone string) (string, error) {
	nowTime := time.Now()

	expireTime := nowTime.Add(time.Duration(configure.JwtExpireTime) * time.Hour) // 过期时间

	claims := Claims{
		phone,
		jwt.StandardClaims {
			ExpiresAt : expireTime.Unix(),
			Issuer : "ai-message",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}