package products

import (
	"ecom/repo"
	"ecom/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ReqUpdateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	pId, err := strconv.Atoi(id)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Please give me a valid product id")
		return
	}

	var req ReqUpdateProduct
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		utils.SendError(w, http.StatusBadRequest, "Please give me valid json")
		return
	}

	product, err := h.productRepo.Update(repo.Product{
		ID:          pId,
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImgUrl:      req.ImgUrl,
	})
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	utils.SendData(w, http.StatusOK, product)

}
