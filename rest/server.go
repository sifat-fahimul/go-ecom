package rest

import (
	"ecom/config"
	"ecom/rest/handlers/products"
	"ecom/rest/handlers/users"
	"ecom/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	cnf            *config.Config
	productHandler *products.Handler
	userHandler    *users.Handler
}

func NewServer(
	cnf *config.Config,
	productHandler *products.Handler,
	userHandler *users.Handler,
) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (server *Server) Start() {
	manager := middlewares.NewManger()
	manager.Use(
		middlewares.Preflight,
		middlewares.Cors,
		middlewares.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(server.cnf.HttpPort)

	fmt.Println("Server running on ", addr)
	err := http.ListenAndServe(addr, wrappedMux)

	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}
