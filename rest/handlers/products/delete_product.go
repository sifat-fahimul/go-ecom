package products

import (
	"ecom/utils"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	pId, err := strconv.Atoi(id)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Please give me a valid product id")
		return
	}

	err = h.productRepo.Delete(pId)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	utils.SendData(w, http.StatusOK, "Product Delete Successfully")
}
