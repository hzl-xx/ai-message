package services

import (
	"messageserver/grpc/message/configure"
	"github.com/dgrijalva/jwt-go"
	"messageserver/grpc/message/protos"
	"time"
)

var (
	jwt_key = []byte(configure.JWT_KEY)
)

type CustomClaims struct {
	Key *protos.Key
	jwt.StandardClaims
}

type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *protos.Key) (string, error)
}

type TokenService struct {
}

func (srv *TokenService) Decode(tokenString string) (*CustomClaims, error) {

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwt_key, nil
	})

	// Validate the token and return the custom claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func (srv *TokenService) Encode(key *protos.Key) (string, error) {

	expireToken := time.Now().Add(time.Hour * 72).Unix()

	// Create the Claims
	claims := CustomClaims{
		key,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "srv.wechat.message",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	return token.SignedString(jwt_key)
}

