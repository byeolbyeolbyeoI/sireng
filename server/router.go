package server

import (
	"fmt"
	userHandler "github.com/chaaaeeee/sireng/internal/user/handler"
	"net/http"
)

func initializeRoutes(userHandler userHandler.UserHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /login", userHandler.Login)
	mux.HandleFunc("POST /signup", userHandler.SignUp)
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("tes")
	})

	return mux
}
