package users

import (
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
		utils.SendError(w, http.StatusBadRequest, "Please give me valid json")
		return
	}

	user, err := h.userRepo.Find(reqLogin.Email, reqLogin.Password)
	if err != nil {
		utils.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	accessToken, err := utils.CreateJwt(h.cnf.JwtSecret, utils.Payload{
		Sub:         1,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		IsShopOwner: user.IsShopOwner,
	})

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	utils.SendData(w, http.StatusCreated, accessToken)

}
