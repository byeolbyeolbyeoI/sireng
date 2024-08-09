package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

type MiddlewareService interface {
	GetMethod(string) (string, error)
	ExportJWTString(string) (string, error)
	ExtractClaims(*http.Request) (jwt.MapClaims, error)
	IsAuthenticated(jwt.MapClaims) bool
	IsAdmin(jwt.MapClaims) bool
	IsUser(jwt.MapClaims) bool
}
