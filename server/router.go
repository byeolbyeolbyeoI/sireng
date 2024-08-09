package server

import (
	_ "github.com/chaaaeeee/sireng/docs"
	trackerHandler "github.com/chaaaeeee/sireng/internal/tracker/handler"
	userHandler "github.com/chaaaeeee/sireng/internal/user/handler"
	"github.com/chaaaeeee/sireng/middleware"
	"github.com/swaggo/http-swagger" // http-swagger middleware
	"net/http"
)

func swaggerHandler(w http.ResponseWriter, r *http.Request) {
	httpSwagger.WrapHandler(w, r)
}

func initializeRoutes(mux *http.ServeMux, userHandler userHandler.UserHandler, trackerHandler trackerHandler.TrackerHandler, middleware middleware.Middleware) *http.ServeMux {
	// declare handlers
	createStudySessionHandlerFunc := http.HandlerFunc(trackerHandler.CreateStudySession)
	endStudySessionHandlerFunc := http.HandlerFunc(trackerHandler.EndStudySession)
	getStudySessionsHandlerFunc := http.HandlerFunc(trackerHandler.GetStudySessions)
	getStudySessionsByUserIdHandlerFunc := http.HandlerFunc(trackerHandler.GetStudySessionsByUserId)
	// set middleware
	createStudySessionHandler := middleware.Authenticate(createStudySessionHandlerFunc)
	createStudySessionHandler = middleware.IsUser(createStudySessionHandler)

	endStudySessionHandler := middleware.Authenticate(endStudySessionHandlerFunc)
	endStudySessionHandler = middleware.IsUser(endStudySessionHandler)

	getStudySessionsHandler := middleware.Authenticate(getStudySessionsHandlerFunc)
	getStudySessionsHandler = middleware.IsUser(getStudySessionsHandler)

	getStudySessionsByUserIdHandler := middleware.Authenticate(getStudySessionsByUserIdHandlerFunc)
	getStudySessionsByUserIdHandler = middleware.IsUser(getStudySessionsByUserIdHandler)

	// initialize handlers
	mux.HandleFunc("POST /signup", userHandler.SignUp)
	mux.HandleFunc("POST /login", userHandler.Login)
	// mux.HandleFunc("GET /logout", userHandler.Logout)

	mux.Handle("POST /createStudySession", createStudySessionHandler)
	mux.Handle("GET /endStudySession/{userId}", endStudySessionHandler)
	mux.Handle("GET /getStudySessions", getStudySessionsHandler)
	mux.Handle("GET /getStudySessions/{userId}", getStudySessionsByUserIdHandler)

	// nantian
	// mux.HandleFunc("GET /swagger/*", swaggerHandler)

	return mux
}
