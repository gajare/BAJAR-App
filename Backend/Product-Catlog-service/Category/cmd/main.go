package main

import (
	"log"
	"net/http"
	"product-catalog-service/Category/internal/config"
	"product-catalog-service/Category/internal/handlers"
	"product-catalog-service/Category/internal/repository"
	"product-catalog-service/Category/internal/services"
	"product-catalog-service/database"
	"product-catalog-service/middleware"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go.uber.org/zap"
)

// @title Category Service API
// @version 1.0
// @description Category management microservice
// @host localhost:8081
// @BasePath /api/v1
func main() {
	// Initialize logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("Failed to initialize logger:", err)
	}
	defer logger.Sync()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Failed to load config", zap.Error(err))
	}

	// Initialize database
	db, err := database.InitDB(cfg.DatabaseURL)
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}
	defer db.Close()

	// Auto migrate category tables
	db.AutoMigrate(&repository.Category{})

	// Initialize repository, service and handler
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService, logger)

	// Initialize router
	router := mux.NewRouter()

	// Middleware
	router.Use(middleware.LoggingMiddleware(logger))
	router.Use(middleware.CORSMiddleware)

	// API routes
	api := router.PathPrefix("/api/v1").Subrouter()

	// Category routes
	api.HandleFunc("/categories", categoryHandler.GetAllCategories).Methods("GET")
	api.HandleFunc("/categories/{id}", categoryHandler.GetCategory).Methods("GET")
	api.HandleFunc("/categories", categoryHandler.CreateCategory).Methods("POST")
	api.HandleFunc("/categories/{id}", categoryHandler.UpdateCategory).Methods("PUT")
	api.HandleFunc("/categories/{id}", categoryHandler.DeleteCategory).Methods("DELETE")

	logger.Info("Category Service starting on :8081")
	if err := http.ListenAndServe(":8081", router); err != nil {
		logger.Fatal("Server failed to start", zap.Error(err))
	}
}
