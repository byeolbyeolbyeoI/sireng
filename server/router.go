package server

import (
	"fmt"
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
	createStudySessionHandler := http.HandlerFunc(trackerHandler.CreateStudySessionHandler)
	tes := middleware.Authenticate(createStudySessionHandler)
	mux.HandleFunc("POST /signup", userHandler.SignUp)
	mux.HandleFunc("POST /login", userHandler.Login)
	mux.HandleFunc("POST /createStudySession", trackerHandler.CreateStudySessionHandler)
	mux.HandleFunc("POST /endStudySession", trackerHandler.EndStudySessionHandler)
	// mux.HandleFunc("GET /logout", userHandler.Logout)
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("tes")
	})

	mux.HandleFunc("GET /swagger/*", swaggerHandler)

	mux.Handle("POST /middlewareTest", tes)

	return mux
}
