package server

import (
	"database/sql"
	"fmt"
	"github.com/chaaaeeee/sireng/config"
	trackerRepo "github.com/chaaaeeee/sireng/internal/tracker/domain/repository"
	trackerService "github.com/chaaaeeee/sireng/internal/tracker/domain/service"
	trackerHandler "github.com/chaaaeeee/sireng/internal/tracker/handler"
	userRepo "github.com/chaaaeeee/sireng/internal/user/domain/repository"
	userService "github.com/chaaaeeee/sireng/internal/user/domain/service"
	userHandler "github.com/chaaaeeee/sireng/internal/user/handler"
	"github.com/chaaaeeee/sireng/util"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type HTTPServer struct {
	mux       *http.ServeMux
	config    *config.Config // using 'conf' cuz tabrakan with config package, or not really... since we have to call the struct name first.., decided to use config
	db        *sql.DB
	util      util.Util
	validator *validator.Validate
}

func NewServer(conf *config.Config, db *sql.DB, util util.Util) Server {
	mux := http.NewServeMux()
	validator := validator.New()

	return &HTTPServer{
		mux:       mux,
		config:    conf,
		db:        db,
		util:      util,
		validator: validator,
	}
}

func (h *HTTPServer) Start() {
	userRepo := userRepo.NewUserRepository(h.db, h.util)
	userService := userService.NewUserService(userRepo, h.util, h.validator)
	userHandler := userHandler.NewUserHandler(userService, h.util)

	trackerRepo := trackerRepo.NewTrackerRepository(h.db, h.util)
	trackerService := trackerService.NewTrackerService(trackerRepo, h.util, h.validator)
	trackerHandler := trackerHandler.NewTrackerHandler(trackerService, h.util)

	// initialize routes?
	router := initializeRoutes(userHandler, trackerHandler)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", h.config.Server.Port),
		Handler: router} // routes

	fmt.Println("server running in :", h.config.Server.Port)
	server.ListenAndServe()
}
