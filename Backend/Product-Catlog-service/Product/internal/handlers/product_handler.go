package handlers

import (
	"encoding/json"
	"net/http"
	"product-service/internal/models"
	"product-service/internal/services"
	"response"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type ProductHandler struct {
	service services.ProductService
	logger  *zap.Logger
}

func NewProductHandler(service services.ProductService, logger *zap.Logger) *ProductHandler {
	return &ProductHandler{
		service: service,
		logger:  logger,
	}
}

// GetAllProducts godoc
// @Summary Get all products
// @Description Get all products with optional filtering
// @Tags products
// @Produce json
// @Success 200 {array} models.Product
// @Router /products [get]
func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		h.logger.Error("Failed to get products", zap.Error(err))
		response.Error(w, "Failed to get products", http.StatusInternalServerError)
		return
	}

	response.JSON(w, products, http.StatusOK)
}

// GetProduct godoc
// @Summary Get a product by ID
// @Description Get a specific product by its ID with all relations
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.ProductResponse
// @Router /products/{id} [get]
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		response.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := h.service.GetProduct(uint(id))
	if err != nil {
		h.logger.Error("Failed to get product", zap.Error(err))
		response.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	response.JSON(w, product, http.StatusOK)
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with the provided details
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.CreateProductRequest true "Product data"
// @Success 201 {object} models.ProductResponse
// @Router /products [post]
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var request models.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if request.Name == "" || request.Price == 0 || request.SKU == "" || request.Stock == 0 {
		response.Error(w, "Name, price, SKU, and stock are required", http.StatusBadRequest)
		return
	}

	product, err := h.service.CreateProduct(request)
	if err != nil {
		h.logger.Error("Failed to create product", zap.Error(err))
		response.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	response.JSON(w, product, http.StatusCreated)
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Update an existing product
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.UpdateProductRequest true "Product data"
// @Success 200 {object} models.ProductResponse
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		response.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var request models.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	product, err := h.service.UpdateProduct(uint(id), request)
	if err != nil {
		h.logger.Error("Failed to update product", zap.Error(err))
		response.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}

	response.JSON(w, product, http.StatusOK)
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product by ID
// @Tags products
// @Param id path int true "Product ID"
// @Success 204
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		response.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteProduct(uint(id)); err != nil {
		h.logger.Error("Failed to delete product", zap.Error(err))
		response.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	response.JSON(w, nil, http.StatusNoContent)
}

// GetProductsByCategory godoc
// @Summary Get products by category
// @Description Get all products in a specific category
// @Tags products
// @Produce json
// @Param category_id path int true "Category ID"
// @Success 200 {array} models.Product
// @Router /products/category/{category_id} [get]
func (h *ProductHandler) GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID, err := strconv.ParseUint(vars["category_id"], 10, 32)
	if err != nil {
		response.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	products, err := h.service.GetProductsByCategory(uint(categoryID))
	if err != nil {
		h.logger.Error("Failed to get products by category", zap.Error(err))
		response.Error(w, "Failed to get products", http.StatusInternalServerError)
		return
	}

	response.JSON(w, products, http.StatusOK)
}

// SearchProducts godoc
// @Summary Search products
// @Description Search products by name or description
// @Tags products
// @Produce json
// @Param q query string true "Search query"
// @Success 200 {array} models.Product
// @Router /products/search [get]
func (h *ProductHandler) SearchProducts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		response.Error(w, "Search query is required", http.StatusBadRequest)
		return
	}

	products, err := h.service.SearchProducts(query)
	if err != nil {
		h.logger.Error("Failed to search products", zap.Error(err))
		response.Error(w, "Failed to search products", http.StatusInternalServerError)
		return
	}

	response.JSON(w, products, http.StatusOK)
}
