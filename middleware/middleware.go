package middleware

import "net/http"

type Middleware interface {
	Authenticate(http.Handler) http.Handler
	IsAdmin(http.Handler) http.Handler
	IsUser(http.Handler) http.Handler
}
