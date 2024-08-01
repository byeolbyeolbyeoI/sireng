package server

import (
	"fmt"
	trackerHandler "github.com/chaaaeeee/sireng/internal/tracker/handler"
	userHandler "github.com/chaaaeeee/sireng/internal/user/handler"
	"net/http"
)

func initializeRoutes(userHandler userHandler.UserHandler, trackerHandler trackerHandler.TrackerHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /signup", userHandler.SignUp)
	mux.HandleFunc("POST /login", userHandler.Login)
	mux.HandleFunc("POST /createStudySession", trackerHandler.CreateStudySessionHandler)
	mux.HandleFunc("POST /endStudySession", trackerHandler.EndStudySessionHandler)
	// mux.HandleFunc("GET /logout", userHandler.Logout)
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("tes")
	})

	return mux
}
