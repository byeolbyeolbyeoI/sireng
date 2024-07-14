package server

import (
	"database/sql"
	"fmt"
	"github.com/chaaaeeee/sireng/config"
	userRepo "github.com/chaaaeeee/sireng/internal/user/domain/repository"
	userService "github.com/chaaaeeee/sireng/internal/user/domain/service"
	userHandler "github.com/chaaaeeee/sireng/internal/user/handler"
	"github.com/chaaaeeee/sireng/util"
	"net/http"
)

type HTTPServer struct {
	mux    *http.ServeMux
	config *config.Config // using 'conf' cuz tabrakan with config package, or not really... since we have to call the struct name first.., decided to use config
	db     *sql.DB
	util   util.Util
}

func NewServer(conf *config.Config, db *sql.DB, util util.Util) Server {
	mux := http.NewServeMux()

	return &HTTPServer{
		mux:    mux,
		config: conf,
		db:     db,
		util:   util,
	}
}

func (h *HTTPServer) Start() {
	userRepo := userRepo.NewUserRepository(h.db)
	userService := userService.NewUserService(userRepo)
	userHandler := userHandler.NewUserHandler(userService)

	// initialize routes?
	router := initializeRoutes(userHandler)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", h.config.Server.Port),
		Handler: router} // routes

	fmt.Println("server running in :", h.config.Server.Port)
	server.ListenAndServe()
}
