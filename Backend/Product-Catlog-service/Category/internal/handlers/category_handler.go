package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gajare/BAJAR-App/Backend/Product-Catalog-Service/Category/internal/models"
	"github.com/gajare/BAJAR-App/Backend/Product-Catalog-Service/Category/internal/services"
	"github.com/gajare/BAJAR-App/Backend/Product-Catalog-Service/pkg/response"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type CategoryHandler struct {
	service services.CategoryService
	logger  *zap.Logger
}

func NewCategoryHandler(service services.CategoryService, logger *zap.Logger) *CategoryHandler {
	return &CategoryHandler{
		service: service,
		logger:  logger,
	}
}

// GetAllCategories godoc
// @Summary Get all categories
// @Description Get all categories with optional tree structure
// @Tags categories
// @Produce json
// @Param tree query bool false "Return as tree structure"
// @Success 200 {array} models.Category
// @Router /categories [get]
func (h *CategoryHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	tree := r.URL.Query().Get("tree") == "true"

	if tree {
		categories, err := h.service.GetCategoryTree()
		if err != nil {
			h.logger.Error("Failed to get category tree", zap.Error(err))
			response.Error(w, "Failed to get categories", http.StatusInternalServerError)
			return
		}
		response.JSON(w, categories, http.StatusOK)
		return
	}

	categories, err := h.service.GetAllCategories()
	if err != nil {
		h.logger.Error("Failed to get categories", zap.Error(err))
		response.Error(w, "Failed to get categories", http.StatusInternalServerError)
		return
	}

	response.JSON(w, categories, http.StatusOK)
}

// GetCategory godoc
// @Summary Get a category by ID
// @Description Get a specific category by its ID
// @Tags categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} models.Category
// @Router /categories/{id} [get]
func (h *CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		response.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	category, err := h.service.GetCategory(uint(id))
	if err != nil {
		h.logger.Error("Failed to get category", zap.Error(err))
		response.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	response.JSON(w, category, http.StatusOK)
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category with the provided details
// @Tags categories
// @Accept json
// @Produce json
// @Param category body models.CreateCategoryRequest true "Category data"
// @Success 201 {object} models.Category
// @Router /categories [post]
func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var request models.CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if request.Name == "" {
		response.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	category, err := h.service.CreateCategory(request)
	if err != nil {
		h.logger.Error("Failed to create category", zap.Error(err))
		response.Error(w, "Failed to create category", http.StatusInternalServerError)
		return
	}

	response.JSON(w, category, http.StatusCreated)
}

// UpdateCategory godoc
// @Summary Update a category
// @Description Update an existing category
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body models.UpdateCategoryRequest true "Category data"
// @Success 200 {object} models.Category
// @Router /categories/{id} [put]
func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		response.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	var request models.UpdateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	category, err := h.service.UpdateCategory(uint(id), request)
	if err != nil {
		h.logger.Error("Failed to update category", zap.Error(err))
		response.Error(w, "Failed to update category", http.StatusInternalServerError)
		return
	}

	response.JSON(w, category, http.StatusOK)
}

// DeleteCategory godoc
// @Summary Delete a category
// @Description Delete a category by ID
// @Tags categories
// @Param id path int true "Category ID"
// @Success 204
// @Router /categories/{id} [delete]
func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		response.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteCategory(uint(id)); err != nil {
		h.logger.Error("Failed to delete category", zap.Error(err))
		response.Error(w, "Failed to delete category", http.StatusInternalServerError)
		return
	}

	response.JSON(w, nil, http.StatusNoContent)
}
