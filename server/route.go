package server

import (
	"fmt"
	"github.com/chaaaeeee/sireng/internal/tracker/handler"
	"net/http"
)

func initializeRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("tes")
	})

	return mux
}
