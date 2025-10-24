package users

import (
	"ecom/config"
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

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
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
		return
	}

	cnf := config.GetConfig()
	accessToken, err := utils.CreateJwt(cnf.JwtSecret, utils.Payload{
		Sub:         1,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		IsShopOwner: user.IsShopOwner,
	})

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	utils.SendData(w, accessToken, http.StatusCreated)

}
