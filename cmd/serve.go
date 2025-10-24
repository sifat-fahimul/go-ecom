package cmd

import (
	"ecom/config"
	"ecom/rest"
	"ecom/rest/handlers/products"
	"ecom/rest/handlers/users"
	"ecom/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	middlewares := middlewares.NewMiddleWare(cnf)
	productHandler := products.NewHandler(middlewares)
	userHandler := users.NewHandler()

	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start()
}
