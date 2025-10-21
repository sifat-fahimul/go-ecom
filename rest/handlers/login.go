package handlers

import (
	"ecom/database"
	"ecom/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin ReqLogin

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid json", http.StatusBadRequest)
		return
	}

	user := database.Find(reqLogin.Email, reqLogin.Password)
	if user == nil {
		http.Error(w, "Invalid credential", http.StatusBadRequest)
	}
	utils.SendData(w, user, http.StatusCreated)

}
