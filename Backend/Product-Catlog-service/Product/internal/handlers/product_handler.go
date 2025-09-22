package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gajare/BAJAR-App/Backend/Product-Catalog-Service/Product/internal/models"
	"github.com/gajare/BAJAR-App/Backend/Product-Catalog-Service/Product/internal/services"
	"github.com/gajare/BAJAR-App/Backend/Product-Catalog-Service/pkg/response"

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
// @Description Get all products with pagination and filtering
// @Tags products
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Param status query string false "Filter by status"
// @Success 200 {array} models.Product
// @Router /products [get]
func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	//	status := r.URL.Query().Get("status")
	_ = r.URL.Query().Get("status")

	products, err := h.service.GetAllProducts()
	if err != nil {
		h.logger.Error("Failed to get products", zap.Error(err))
		response.Error(w, "Failed to get products", http.StatusInternalServerError)
		return
	}

	// Simple pagination (in real implementation, do this at database level)
	if page > 0 && limit > 0 {
		start := (page - 1) * limit
		end := start + limit
		if start >= len(products) {
			products = []models.Product{}
		} else if end > len(products) {
			products = products[start:]
		} else {
			products = products[start:end]
		}
	}

	response.JSON(w, products, http.StatusOK)
}

// GetProduct godoc
// @Summary Get a product by ID
// @Description Get a specific product by its ID
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product
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
// @Success 201 {object} models.Product
// @Router /products [post]
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var request models.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if request.Name == "" || request.SKU == "" || request.Price <= 0 {
		response.Error(w, "Name, SKU, and positive Price are required", http.StatusBadRequest)
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
// @Success 200 {object} models.Product
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

// SearchProducts godoc
// @Summary Search products
// @Description Search products with various filters
// @Tags products
// @Produce json
// @Param query query string false "Search query"
// @Param min_price query number false "Minimum price"
// @Param max_price query number false "Maximum price"
// @Param status query string false "Product status"
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Success 200 {array} models.Product
// @Router /products/search [get]
func (h *ProductHandler) SearchProducts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	minPrice, _ := strconv.ParseFloat(r.URL.Query().Get("min_price"), 64)
	maxPrice, _ := strconv.ParseFloat(r.URL.Query().Get("max_price"), 64)
	status := r.URL.Query().Get("status")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	products, err := h.service.SearchProducts(query, minPrice, maxPrice, status, page, limit)
	if err != nil {
		h.logger.Error("Failed to search products", zap.Error(err))
		response.Error(w, "Failed to search products", http.StatusInternalServerError)
		return
	}

	response.JSON(w, products, http.StatusOK)
}

// GetProductImages godoc
// @Summary Get product images
// @Description Get all images for a specific product
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {array} models.Image
// @Router /products/{id}/images [get]
func (h *ProductHandler) GetProductImages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		response.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	images, err := h.service.GetProductImages(uint(id))
	if err != nil {
		h.logger.Error("Failed to get product images", zap.Error(err))
		response.Error(w, "Failed to get product images", http.StatusInternalServerError)
		return
	}

	response.JSON(w, images, http.StatusOK)
}

// GetProductVariants godoc
// @Summary Get product variants
// @Description Get all variants for a specific product
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {array} models.Variant
// @Router /products/{id}/variants [get]
func (h *ProductHandler) GetProductVariants(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		response.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	variants, err := h.service.GetProductVariants(uint(id))
	if err != nil {
		h.logger.Error("Failed to get product variants", zap.Error(err))
		response.Error(w, "Failed to get product variants", http.StatusInternalServerError)
		return
	}

	response.JSON(w, variants, http.StatusOK)
}
