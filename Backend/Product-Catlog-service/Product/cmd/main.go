package main

import (
	"log"
	"net/http"

	"github.com/gajare/BAJAR-App/Backend/Product-Catalog-Service/Product/internal/config"
	"github.com/gajare/BAJAR-App/Backend/Product-Catalog-Service/Product/internal/handlers"
	"github.com/gajare/BAJAR-App/Backend/Product-Catalog-Service/Product/internal/models"
	"github.com/gajare/BAJAR-App/Backend/Product-Catalog-Service/Product/internal/repository"
	"github.com/gajare/BAJAR-App/Backend/Product-Catalog-Service/Product/internal/services"

	"github.com/gajare/BAJAR-App/Backend/Product-Catalog-Service/database"
	"github.com/gajare/BAJAR-App/Backend/Product-Catalog-Service/middleware"

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
		&models.Product{},
		&models.Image{},
		&models.Variant{},
		&models.Attribute{},
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
	api.HandleFunc("/products/search", productHandler.SearchProducts).Methods("GET")
	api.HandleFunc("/products/{id}/images", productHandler.GetProductImages).Methods("GET")
	api.HandleFunc("/products/{id}/variants", productHandler.GetProductVariants).Methods("GET")

	logger.Info("Product Service starting on :8082")
	if err := http.ListenAndServe(":8082", router); err != nil {
		logger.Fatal("Server failed to start", zap.Error(err))
	}
}
