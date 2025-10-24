package cmd

import (
	"ecom/config"
	"ecom/repo"
	"ecom/rest"
	"ecom/rest/handlers/products"
	"ecom/rest/handlers/users"
	"ecom/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	middlewares := middlewares.NewMiddleWare(cnf)

	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo()

	productHandler := products.NewHandler(middlewares, productRepo)
	userHandler := users.NewHandler(cnf, userRepo)

	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start()
}
