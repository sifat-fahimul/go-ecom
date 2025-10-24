package products

import (
	"ecom/database"
	"ecom/utils"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, database.List(), 200)
}
