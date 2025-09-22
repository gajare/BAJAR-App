package main

import (
	"database"
	"log"
	"middleware"
	"net/http"
	"product-service/internal/config"
	"product-service/internal/handlers"
	"product-service/internal/repository"
	"product-service/internal/services"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go.uber.org/zap"
)

// @title Product Service API
// @version 1.0
// @description Product management microservice
// @host localhost:8082
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

	// Auto migrate product tables
	db.AutoMigrate(
		&repository.Product{},
		&repository.ProductImage{},
		&repository.ProductVariant{},
		&repository.ProductAttribute{},
		&repository.ProductCategory{},
	)

	// Initialize repository, service and handler
	productRepo := repository.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService, logger)

	// Initialize router
	router := mux.NewRouter()

	// Middleware
	router.Use(middleware.LoggingMiddleware(logger))
	router.Use(middleware.CORSMiddleware)

	// API routes
	api := router.PathPrefix("/api/v1").Subrouter()

	// Product routes
	api.HandleFunc("/products", productHandler.GetAllProducts).Methods("GET")
	api.HandleFunc("/products/{id}", productHandler.GetProduct).Methods("GET")
	api.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	api.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	api.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods("DELETE")
	api.HandleFunc("/products/category/{category_id}", productHandler.GetProductsByCategory).Methods("GET")
	api.HandleFunc("/products/search", productHandler.SearchProducts).Methods("GET")

	logger.Info("Product Service starting on :8082")
	if err := http.ListenAndServe(":8082", router); err != nil {
		logger.Fatal("Server failed to start", zap.Error(err))
	}
}
