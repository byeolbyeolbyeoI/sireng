package server

import (
	"fmt"
	userHandler "github.com/chaaaeeee/sireng/internal/user/handler"
	"net/http"
)

func initializeRoutes(userHandler userHandler.UserHandler) *http.ServeMux {
	mux := http.NewServeMux()

	http.HandleFunc("/login", userHandler.Login)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("tes")
	})

	return mux
}
