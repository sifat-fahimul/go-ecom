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

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	middlewares := middlewares.NewMiddleWare(cnf)

	if err := db.MigrateDB(dbCon, "./migrations"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	productHandler := products.NewHandler(middlewares, productRepo)
	userHandler := users.NewHandler(cnf, userRepo)

	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start()
}
