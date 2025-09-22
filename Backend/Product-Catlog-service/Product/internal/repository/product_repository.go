package repository

import (
	"github.com/gajare/BAJAR-App/Backend/Product-Catalog-Service/Product/internal/models"

	"github.com/jinzhu/gorm"
)

type ProductRepository interface {
	FindAll() ([]models.Product, error)
	FindByID(id uint) (*models.Product, error)
	Create(product *models.Product) error
	Update(product *models.Product) error
	Delete(id uint) error
	FindBySlug(slug string) (*models.Product, error)
	FindBySKU(sku string) (*models.Product, error)
	Search(query string, minPrice, maxPrice float64, status string, page, limit int) ([]models.Product, error)
	GetProductImages(productID uint) ([]models.Image, error)
	GetProductVariants(productID uint) ([]models.Variant, error)
	GetProductAttributes(productID uint) ([]models.Attribute, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("Images").Preload("Variants").Preload("Attributes").Find(&products).Error
	return products, err
}

func (r *productRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("Images").Preload("Variants").Preload("Attributes").
		First(&product, id).Error
	return &product, err
}

func (r *productRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}

func (r *productRepository) FindBySlug(slug string) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("Images").Preload("Variants").Preload("Attributes").
		Where("slug = ?", slug).First(&product).Error
	return &product, err
}

func (r *productRepository) FindBySKU(sku string) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("Images").Preload("Variants").Preload("Attributes").
		Where("sku = ?", sku).First(&product).Error
	return &product, err
}

func (r *productRepository) Search(query string, minPrice, maxPrice float64, status string, page, limit int) ([]models.Product, error) {
	var products []models.Product
	db := r.db.Preload("Images").Preload("Variants").Preload("Attributes")

	if query != "" {
		db = db.Where("name ILIKE ? OR description ILIKE ?", "%"+query+"%", "%"+query+"%")
	}
	if minPrice > 0 {
		db = db.Where("price >= ?", minPrice)
	}
	if maxPrice > 0 {
		db = db.Where("price <= ?", maxPrice)
	}
	if status != "" {
		db = db.Where("status = ?", status)
	}

	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		db = db.Offset(offset).Limit(limit)
	}

	err := db.Find(&products).Error
	return products, err
}

func (r *productRepository) GetProductImages(productID uint) ([]models.Image, error) {
	var images []models.Image
	err := r.db.Where("product_id = ?", productID).Order("sort_order").Find(&images).Error
	return images, err
}

func (r *productRepository) GetProductVariants(productID uint) ([]models.Variant, error) {
	var variants []models.Variant
	err := r.db.Where("product_id = ?", productID).Find(&variants).Error
	return variants, err
}

func (r *productRepository) GetProductAttributes(productID uint) ([]models.Attribute, error) {
	var attributes []models.Attribute
	err := r.db.Where("product_id = ?", productID).Find(&attributes).Error
	return attributes, err
}
