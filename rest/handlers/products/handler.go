package products

import "ecom/rest/middlewares"

type Handler struct {
	middlewares *middlewares.MiddleWares
}

func NewHandler(middlewares *middlewares.MiddleWares) *Handler {
	return &Handler{
		middlewares: middlewares,
	}
}
