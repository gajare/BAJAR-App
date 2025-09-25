package router

import (
	"net/http"
	"user-service/controller"
	"user-service/middleware"

	"github.com/gorilla/mux"
)

type Route struct {
	Path      string
	Method    string
	Handler   func(http.ResponseWriter, *http.Request)
	Protected bool
}

func NewUserRouter(userController *controller.UserController) *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.RateLimitMiddleware)

	routes := []Route{
		{"/register", "POST", userController.Register, false},
		{"/login", "POST", userController.Login, false},
		{"/api/me", "GET", userController.GetProfile, true},
		{"/api/users", "GET", userController.ListUsers, true},
	}

	for _, route := range routes {
		if route.Protected {
			s := r.PathPrefix("/api").Subrouter()
			s.Use(middleware.JwtAuthMiddleware)
			s.HandleFunc(route.Path[4:], route.Handler).Methods(route.Method)
		} else {
			r.HandleFunc(route.Path, route.Handler).Methods(route.Method)
		}
	}

	return r
}
