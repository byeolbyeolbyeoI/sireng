package server

import (
	_ "github.com/chaaaeeee/sireng/docs"
	trackerHandler "github.com/chaaaeeee/sireng/internal/tracker/handler"
	userHandler "github.com/chaaaeeee/sireng/internal/user/handler"
	ws "github.com/chaaaeeee/sireng/internal/ws"
	"github.com/chaaaeeee/sireng/middleware"
	"github.com/swaggo/http-swagger" // http-swagger middleware
	"net/http"
)

func swaggerHandler(w http.ResponseWriter, r *http.Request) {
	httpSwagger.WrapHandler(w, r)
}

// i js be wrapping shit
func AuthUser(h func(w http.ResponseWriter, r *http.Request), m middleware.Middleware) http.Handler {
	return m.IsUser(m.Authenticate(http.HandlerFunc(h)))
}

func initializeRoutes(mux *http.ServeMux, userHandler userHandler.UserHandler, trackerHandler trackerHandler.TrackerHandler, wsHandler ws.WsHandler, middleware middleware.Middleware) *http.ServeMux {
	// initialize handlers
	mux.HandleFunc("POST /signup", userHandler.SignUp)
	mux.HandleFunc("POST /login", userHandler.Login)
	// mux.HandleFunc("GET /logout", userHandler.Logout)

	mux.Handle("POST /createStudySession", AuthUser(trackerHandler.CreateStudySession, middleware))
	mux.Handle("GET /endStudySession/{userId}", AuthUser(trackerHandler.EndStudySession, middleware))
	mux.Handle("GET /getStudySessions", AuthUser(trackerHandler.GetStudySessions, middleware))
	mux.Handle("GET /getStudySessions/{userId}", AuthUser(trackerHandler.GetStudySessionsByUserId, middleware))

	mux.HandleFunc("POST /ws/createRoom", wsHandler.CreateRoom)
	mux.HandleFunc("GET /ws/getRooms", wsHandler.GetRooms)

	// nantian
	// mux.HandleFunc("GET /swagger/*", swaggerHandler)

	return mux
}
