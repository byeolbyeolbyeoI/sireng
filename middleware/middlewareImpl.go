package middleware

import (
	"github.com/chaaaeeee/sireng/util"
	"net/http"
)

// the only struct that only has service and not [something]Service for one of it's variable name
type MiddlewareImpl struct {
	service MiddlewareService
	util    util.Util
}

func NewMiddleware(s MiddlewareService, util util.Util) *MiddlewareImpl {
	return &MiddlewareImpl{
		service: s,
		util:    util,
	}
}

func (m *MiddlewareImpl) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := m.service.ExtractClaims(r)
		if err != nil {
			m.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		ok := m.service.IsAuthenticated(claims)
		if !ok {
			m.util.WriteJSON(w, http.StatusUnauthorized, util.Response{
				Success: false,
				Message: "User is not authenticated",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (m *MiddlewareImpl) IsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := m.service.ExtractClaims(r)
		if err != nil {
			m.util.WriteJSON(w, http.StatusInternalServerError, util.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		ok := m.service.IsAdmin(claims)
		if !ok {
			m.util.WriteJSON(w, http.StatusUnauthorized, util.Response{
				Success: false,
				Message: "User is not an admin",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}
