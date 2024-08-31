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
	ws "github.com/chaaaeeee/sireng/internal/ws"
	middleware "github.com/chaaaeeee/sireng/middleware"
	"github.com/chaaaeeee/sireng/util"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type HTTPServer struct {
	mux      *http.ServeMux
	config   *config.Config // using 'conf' cuz tabrakan with config package, or not really...
	db       *sql.DB
	util     util.Util
	validate *validator.Validate
}

func NewServer(conf *config.Config, db *sql.DB, util util.Util) Server {
	mux := http.NewServeMux()
	validate := validator.New(validator.WithRequiredStructEnabled())

	return &HTTPServer{
		mux:      mux,
		config:   conf,
		db:       db,
		util:     util,
		validate: validate,
	}
}

func (h *HTTPServer) Start() {
	middlewareService := middleware.NewMiddlewareService(h.config)
	middlewareInstance := middleware.NewMiddleware(middlewareService, h.util)
	userRepoInstance := userRepo.NewUserRepository(h.db, h.util)
	userServiceInstance := userService.NewUserService(userRepoInstance, h.util, h.validate)
	userHandlerInstance := userHandler.NewUserHandler(userServiceInstance, h.util)

	trackerRepoInstance := trackerRepo.NewTrackerRepository(h.db, h.util)
	trackerServiceInstance := trackerService.NewTrackerService(trackerRepoInstance, h.util, h.validate)
	trackerHandlerInstance := trackerHandler.NewTrackerHandler(trackerServiceInstance, h.util)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub, h.util)
	go hub.Run()

	// initialize routes?
	// pass mw here
	router := initializeRoutes(h.mux, userHandlerInstance, trackerHandlerInstance, wsHandler, middlewareInstance)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", h.config.Server.Port),
		Handler: router} // routes

	fmt.Println("server running in :", h.config.Server.Port)
	server.ListenAndServe()
}
