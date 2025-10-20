package rest

import (
	"ecom/config"
	"ecom/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Start(cnf config.Config) {
	manager := middlewares.NewManger()
	manager.Use(middlewares.Preflight, middlewares.Cors, middlewares.Logger)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	initRoute(mux, manager)

	addr := ":" + strconv.Itoa(cnf.HttpPort)

	fmt.Println("Server running on ", addr)
	err := http.ListenAndServe(addr, wrappedMux)

	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}
