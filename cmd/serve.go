package cmd

import (
	"ecom/config"
	"ecom/infra/db"
	"ecom/repo"
	"ecom/rest"
	"ecom/rest/handlers/products"
	"ecom/rest/handlers/users"
	"ecom/rest/middlewares"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	db, err := db.NewConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	middlewares := middlewares.NewMiddleWare(cnf)

	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo(db)

	productHandler := products.NewHandler(middlewares, productRepo)
	userHandler := users.NewHandler(cnf, userRepo)

	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start()
}
