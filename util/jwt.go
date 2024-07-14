package util

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func (u *utilImpl) CreateSession(username string) *jwt.Token {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	return token
}

func (u *utilImpl) SignToken(token *jwt.Token) (string, error) {
	tokenString, err := token.SignedString([]byte(u.config.JWT.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
