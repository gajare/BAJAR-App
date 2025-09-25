package controller

import (
	"encoding/json"
	"net/http"
	"user-service/service"

	"go.uber.org/zap"
)

type UserController struct {
	Service service.UserService
	Logger  *zap.Logger
}

func (uc *UserController) GetProfile(w http.ResponseWriter, r *http.Request) {
	uid, ok := r.Context().Value("userID").(string)
	if !ok {
		uc.Logger.Error("Invalid user id in context", zap.Any("userID", r.Context().Value("userID")))
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}
	user, err := uc.Service.GetUserByID(r.Context(), uid)
	if err != nil {
		uc.Logger.Error("User not found", zap.String("userID", uid), zap.Error(err))
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	user.Password = ""
	uc.Logger.Info("User profile fetched", zap.String("userID", uid))
	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uc.Service.ListUsers(r.Context())
	if err != nil {
		uc.Logger.Error("Could not fetch users", zap.Error(err))
		http.Error(w, "could not fetch users", http.StatusInternalServerError)
		return
	}
	for i := range users {
		users[i].Password = ""
	}
	uc.Logger.Info("User list fetched", zap.Int("count", len(users)))
	json.NewEncoder(w).Encode(users)
}
