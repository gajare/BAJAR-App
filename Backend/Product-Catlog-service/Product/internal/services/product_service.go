package services

import (
	"product-service/internal/models"
	"product-service/internal/repository"
)

type ProductService interface {
	GetAllProducts() ([]models.Product, error)
	GetProduct(id uint) (*models.ProductResponse, error)
	CreateProduct(request models.CreateProductRequest) (*models.ProductResponse, error)
	UpdateProduct(id uint, request models.UpdateProductRequest) (*models.ProductResponse, error)
	DeleteProduct(id uint) error
	GetProductsByCategory(categoryID uint) ([]models.Product, error)
	SearchProducts(query string) ([]models.Product, error)
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

func (s *productService) GetProduct(id uint) (*models.ProductResponse, error) {
	return s.repo.GetProductWithRelations(id)
}

func (s *productService) CreateProduct(request models.CreateProductRequest) (*models.ProductResponse, error) {
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
	}

	if err := s.repo.Create(product); err != nil {
		return nil, err
	}

	// Create images
	for _, imgReq := range request.Images {
		image := &models.ProductImage{
			ProductID: product.ID,
			ImageURL:  imgReq.ImageURL,
			AltText:   imgReq.AltText,
			SortOrder: imgReq.SortOrder,
		}
		s.repo.CreateImage(image)
	}

	// Create variants
	for _, varReq := range request.Variants {
		variant := &models.ProductVariant{
			ProductID: product.ID,
			Name:      varReq.Name,
			Value:     varReq.Value,
			Price:     varReq.Price,
			Stock:     varReq.Stock,
			SKU:       varReq.SKU,
		}
		s.repo.CreateVariant(variant)
	}

	// Create attributes
	for _, attrReq := range request.Attributes {
		attribute := &models.ProductAttribute{
			ProductID: product.ID,
			Key:       attrReq.Key,
			Value:     attrReq.Value,
		}
		s.repo.CreateAttribute(attribute)
	}

	// Add categories
	for _, catID := range request.CategoryIDs {
		s.repo.AddCategory(product.ID, catID)
	}

	return s.repo.GetProductWithRelations(product.ID)
}

func (s *productService) UpdateProduct(id uint, request models.UpdateProductRequest) (*models.ProductResponse, error) {
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
	if request.Price != 0 {
		product.Price = request.Price
	}
	product.DiscountPrice = request.DiscountPrice
	if request.SKU != "" {
		product.SKU = request.SKU
	}
	if request.Stock != 0 {
		product.Stock = request.Stock
	}
	product.Weight = request.Weight
	product.Dimensions = request.Dimensions
	if request.Status != "" {
		product.Status = request.Status
	}
	if request.Slug != "" {
		product.Slug = request.Slug
	}

	if err := s.repo.Update(product); err != nil {
		return nil, err
	}

	// TODO: Handle updates for images, variants, attributes, and categories
	// This would require more complex logic to handle additions, updates, and deletions

	return s.repo.GetProductWithRelations(id)
}

func (s *productService) DeleteProduct(id uint) error {
	return s.repo.Delete(id)
}

func (s *productService) GetProductsByCategory(categoryID uint) ([]models.Product, error) {
	return s.repo.FindByCategory(categoryID)
}

func (s *productService) SearchProducts(query string) ([]models.Product, error) {
	return s.repo.Search(query)
}
