package users

import (
	"ecom/database"
	"ecom/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser database.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid json", http.StatusBadRequest)
		return
	}

	createUser := newUser.Store()
	utils.SendData(w, createUser, http.StatusCreated)

}
