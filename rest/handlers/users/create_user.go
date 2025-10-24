package users

import (
	"ecom/repo"
	"ecom/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreateUser struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"isShopOwner"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req ReqCreateUser

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		utils.SendError(w, http.StatusBadRequest, "Please give me valid json")
		return
	}

	createUser, err := h.userRepo.Create(repo.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    req.Password,
		IsShopOwner: req.IsShopOwner,
	})
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	utils.SendData(w, http.StatusCreated, createUser)

}
