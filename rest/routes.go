package rest

import (
	"ecom/rest/handlers"
	"ecom/rest/middlewares"
	"net/http"
)

func initRoute(mux *http.ServeMux, manager *middlewares.Manager) {

	mux.Handle("GET /products",
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
		),
	)

	mux.Handle("GET /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.GetProduct),
		),
	)

	mux.Handle("POST /products",
		manager.With(
			http.HandlerFunc(handlers.CreateProduct),
		),
	)
	mux.Handle("PUT /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.UpdateProduct),
		),
	)

	mux.Handle("POST /users",
		manager.With(
			http.HandlerFunc(handlers.CreateUser),
		),
	)

	mux.Handle("POST /users/login",
		manager.With(
			http.HandlerFunc(handlers.Login),
		),
	)
}
