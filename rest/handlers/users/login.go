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
	fmt.Println("decoder", reqLogin)

	user, err := h.userRepo.Find(reqLogin.Email, reqLogin.Password)
	if err != nil {
		utils.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	if user == nil {
		utils.SendError(w, http.StatusNotFound, "User not found")
		return
	}
	fmt.Println("user", user)
	accessToken, err := utils.CreateJwt(h.cnf.JwtSecret, utils.Payload{
		Sub:         102,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		IsShopOwner: user.IsShopOwner,
	})

	fmt.Println("accessToken", accessToken)

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	utils.SendData(w, http.StatusCreated, accessToken)

}
