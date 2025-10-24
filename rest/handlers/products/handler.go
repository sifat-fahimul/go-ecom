package products

import (
	"ecom/repo"
	"ecom/rest/middlewares"
)

type Handler struct {
	middlewares *middlewares.MiddleWares
	productRepo repo.ProductRepo
}

func NewHandler(
	middlewares *middlewares.MiddleWares,
	productRepo repo.ProductRepo,
) *Handler {
	return &Handler{
		middlewares: middlewares,
		productRepo: productRepo,
	}
}
