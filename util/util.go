package util

import (
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

type Util interface {
	WriteJSON(http.ResponseWriter, int, interface{})
	Input(*http.Request, interface{}) error
	CreateSession(string) *jwt.Token
}
