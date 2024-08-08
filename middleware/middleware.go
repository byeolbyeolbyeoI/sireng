package middleware

import "net/http"

type Middleware interface {
	Authenticate(http.Handler) http.Handler
}
