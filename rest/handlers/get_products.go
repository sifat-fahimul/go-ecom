package handlers

import (
	"ecom/database"
	"ecom/utils"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, database.List(), 200)
}
