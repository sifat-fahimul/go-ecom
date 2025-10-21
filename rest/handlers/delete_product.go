package handlers

import (
	"ecom/database"
	"net/http"
	"strconv"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	pId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	database.Delete(pId)

}
