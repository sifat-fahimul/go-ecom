package products

import (
	"ecom/database"
	"ecom/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	pId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	product := database.Get(pId)

	if product == nil {
		utils.SendError(w, "Product not found", 404)
		return
	}

	utils.SendData(w, product, 200)
}
