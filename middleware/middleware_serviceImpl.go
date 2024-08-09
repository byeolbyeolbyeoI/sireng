package middleware

import (
	"errors"
	"fmt"
	"github.com/chaaaeeee/sireng/config"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

type middlewareServiceImpl struct {
	config *config.Config
}

func NewMiddlewareService(c *config.Config) MiddlewareService {
	return &middlewareServiceImpl{config: c}
}

func (m *middlewareServiceImpl) GetMethod(authorizationHeader string) (string, error) {
	if authorizationHeader == "" {
		return "", errors.New("authorization header is empty")
	}

	parts := strings.Split(authorizationHeader, " ")
	if len(parts) != 2 {
		return "", errors.New("authorization method unknown :" + authorizationHeader)
	}

	return parts[0], nil // returning the Bearer
}

func (m *middlewareServiceImpl) ExportJWTString(authorizationHeader string) (string, error) {
	parts := strings.Split(authorizationHeader, " ")
	if len(parts) != 2 {
		return "", errors.New("authorization header unknown :" + authorizationHeader)
	}

	return parts[1], nil
}

// fix err.Error()
func (m *middlewareServiceImpl) ExtractClaims(r *http.Request) (jwt.MapClaims, error) {
	authorizationHeader := r.Header.Get("Authorization") // get
	method, err := m.GetMethod(authorizationHeader)
	if err != nil {
		return nil, err
	}

	if method != "Bearer" {
		return nil, errors.New("Method not allowed :" + authorizationHeader)
	}

	tokenString, err := m.ExportJWTString(authorizationHeader)
	if err != nil {
		return nil, err
	}

	// parse it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.config.JWT.Secret), nil // return the secret, it is used to parse the tokenString
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("can't retrieve claims")
	}

	return claims, nil
}

func (m *middlewareServiceImpl) IsAuthenticated(claims jwt.MapClaims) bool {
	if claims["iss"] != "Sireng" {
		return false
	}

	return true
}

// is admin is called only after is authenticated is called and it returns true, if not then ur app sucks
func (m *middlewareServiceImpl) IsAdmin(claims jwt.MapClaims) bool {
	if claims["role"] != "admin" {
		return false
	}

	return true
}
func (m *middlewareServiceImpl) IsUser(claims jwt.MapClaims) bool {
	fmt.Println(claims["role"])
	if claims["role"] != "user" {
		return false
	}

	return true
}
