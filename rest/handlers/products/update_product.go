package products

import (
	"ecom/database"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	pId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid json", 400)
		return
	}
	newProduct.ID = pId
	database.Update(newProduct)

}
