package services

import (
	"github.com/gajare/BAJAR-App/Backend/Product-Catalog-Service/Product/internal/models"
	"github.com/gajare/BAJAR-App/Backend/Product-Catalog-Service/Product/internal/repository"
)

type ProductService interface {
	GetAllProducts() ([]models.Product, error)
	GetProduct(id uint) (*models.Product, error)
	CreateProduct(request models.CreateProductRequest) (*models.Product, error)
	UpdateProduct(id uint, request models.UpdateProductRequest) (*models.Product, error)
	DeleteProduct(id uint) error
	SearchProducts(query string, minPrice, maxPrice float64, status string, page, limit int) ([]models.Product, error)
	GetProductImages(productID uint) ([]models.Image, error)
	GetProductVariants(productID uint) ([]models.Variant, error)
	GetProductBySlug(slug string) (*models.Product, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetAllProducts() ([]models.Product, error) {
	return s.repo.FindAll()
}

func (s *productService) GetProduct(id uint) (*models.Product, error) {
	return s.repo.FindByID(id)
}

func (s *productService) CreateProduct(request models.CreateProductRequest) (*models.Product, error) {
	product := &models.Product{
		Name:          request.Name,
		Description:   request.Description,
		Price:         request.Price,
		DiscountPrice: request.DiscountPrice,
		SKU:           request.SKU,
		Stock:         request.Stock,
		Weight:        request.Weight,
		Dimensions:    request.Dimensions,
		Status:        request.Status,
		Slug:          request.Slug,
		Images:        request.Images,
		Variants:      request.Variants,
		Attributes:    request.Attributes,
	}

	if product.Status == "" {
		product.Status = "active"
	}

	if err := s.repo.Create(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productService) UpdateProduct(id uint, request models.UpdateProductRequest) (*models.Product, error) {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if request.Name != "" {
		product.Name = request.Name
	}
	if request.Description != "" {
		product.Description = request.Description
	}
	if request.Price > 0 {
		product.Price = request.Price
	}
	if request.DiscountPrice >= 0 {
		product.DiscountPrice = request.DiscountPrice
	}
	if request.SKU != "" {
		product.SKU = request.SKU
	}
	if request.Stock >= 0 {
		product.Stock = request.Stock
	}
	if request.Weight >= 0 {
		product.Weight = request.Weight
	}
	if request.Status != "" {
		product.Status = request.Status
	}
	if request.Slug != "" {
		product.Slug = request.Slug
	}

	if err := s.repo.Update(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productService) DeleteProduct(id uint) error {
	return s.repo.Delete(id)
}

func (s *productService) SearchProducts(query string, minPrice, maxPrice float64, status string, page, limit int) ([]models.Product, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	return s.repo.Search(query, minPrice, maxPrice, status, page, limit)
}

func (s *productService) GetProductImages(productID uint) ([]models.Image, error) {
	return s.repo.GetProductImages(productID)
}

func (s *productService) GetProductVariants(productID uint) ([]models.Variant, error) {
	return s.repo.GetProductVariants(productID)
}

func (s *productService) GetProductBySlug(slug string) (*models.Product, error) {
	return s.repo.FindBySlug(slug)
}
