package products

import (
	"ecom/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	pId, err := strconv.Atoi(id)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Please give me a valid product id")
		return
	}

	product, err := h.productRepo.Get(pId)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if product == nil {
		utils.SendError(w, http.StatusNotFound, "Product not found")
		return
	}

	utils.SendData(w, http.StatusOK, product)
}
