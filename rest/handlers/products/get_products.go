package products

import (
	"ecom/utils"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	list, err := h.productRepo.List()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	utils.SendData(w, http.StatusOK, list)
}
