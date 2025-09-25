package main

import (
	"log"
	"net/http"

	"user-service/config"
	"user-service/controller"
	"user-service/db"
	"user-service/repository"
	"user-service/service"

	"user-service/router"

	"go.uber.org/zap"
)

func main() {

	// Initialize logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("Failed to initialize logger:", err)
	}
	defer logger.Sync()

	cfg := config.Load()
	db.InitDB(cfg.BuildDSN())

	userRepo := repository.NewUserRepository(db.DB)
	userService := service.NewUserService(userRepo)
	userController := &controller.UserController{
		Service: *userService,
		Logger:  logger,
	}

	// Use grouped router from user_router.go
	r := router.NewUserRouter(userController)

	log.Printf("Server running on :%s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))

}
