package server

import (
	"fmt"
	"github.com/chaaaeeee/sireng/config"
	"net/http"
)

type HTTPServer struct {
	mux    *http.ServeMux
	config *config.Config // using 'conf' cuz tabrakan with config package, or not really... since we had to call the struct name first.., decided to use config
}

func NewServer(conf *config.Config) Server {
	mux := http.NewServeMux()

	return &HTTPServer{
		mux:    mux,
		config: conf}
}

func (h *HTTPServer) Start() {
	// initialize routes?
	server := http.Server{
		Addr: fmt.Sprintf(":%d", h.config.Server.Port)}

	server.ListenAndServe()
}
