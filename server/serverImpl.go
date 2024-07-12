package server

import (
	"database/sql"
	"fmt"
	"github.com/chaaaeeee/sireng/config"
	"net/http"
)

type HTTPServer struct {
	mux    *http.ServeMux
	config *config.Config // using 'conf' cuz tabrakan with config package, or not really... since we have to call the struct name first.., decided to use config
	db     *sql.DB
}

func NewServer(conf *config.Config, db *sql.DB) Server {
	mux := http.NewServeMux()

	return &HTTPServer{
		mux:    mux,
		config: conf,
		db:     db,
	}
}

func (h *HTTPServer) Start() {
	// initialize routes?
	router := initializeRoutes()
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", h.config.Server.Port),
		Handler: router} // routes

	server.ListenAndServe()
}
