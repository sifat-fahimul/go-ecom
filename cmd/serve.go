package cmd

import (
	"ecom/config"
	"ecom/rest"
	"ecom/rest/handlers/products"
	"ecom/rest/handlers/users"
)

func Serve() {
	cnf := config.GetConfig()

	productHandler := products.NewHandler()
	userHandler := users.NewHandler()

	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start()
}
