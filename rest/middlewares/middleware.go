package middlewares

import "ecom/config"

type MiddleWares struct {
	cnf *config.Config
}

func NewMiddleWare(cnf *config.Config) *MiddleWares {
	return &MiddleWares{
		cnf: cnf,
	}
}
