package controller

import (
	"encoding/json"
	"net/http"
	"user-service/models"
	"user-service/service"

	"go.uber.org/zap"
)

type UserController struct {
	UserService service.UserService
	Logger      *zap.Logger
}



func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var userReq models.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	resp, err := uc.UserService.Register(r.Context(), &userReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(resp)
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var userReq models.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	resp, err := uc.UserService.Login(r.Context(), &userReq)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(resp)
}
