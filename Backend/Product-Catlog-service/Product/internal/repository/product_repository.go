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
	FindByCategory(categoryID uint) ([]models.Product, error)
	Search(query string) ([]models.Product, error)
	GetProductWithRelations(id uint) (*models.ProductResponse, error)

	// Image methods
	CreateImage(image *models.ProductImage) error
	DeleteImage(id uint) error
	GetProductImages(productID uint) ([]models.ProductImage, error)

	// Variant methods
	CreateVariant(variant *models.ProductVariant) error
	UpdateVariant(variant *models.ProductVariant) error
	DeleteVariant(id uint) error
	GetProductVariants(productID uint) ([]models.ProductVariant, error)

	// Attribute methods
	CreateAttribute(attribute *models.ProductAttribute) error
	UpdateAttribute(attribute *models.ProductAttribute) error
	DeleteAttribute(id uint) error
	GetProductAttributes(productID uint) ([]models.ProductAttribute, error)

	// Category methods
	AddCategory(productID, categoryID uint) error
	RemoveCategory(productID, categoryID uint) error
	GetProductCategories(productID uint) ([]uint, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *productRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error
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
	err := r.db.Where("slug = ?", slug).First(&product).Error
	return &product, err
}

func (r *productRepository) FindBySKU(sku string) (*models.Product, error) {
	var product models.Product
	err := r.db.Where("sku = ?", sku).First(&product).Error
	return &product, err
}

func (r *productRepository) FindByCategory(categoryID uint) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Joins("JOIN product_categories ON products.id = product_categories.product_id").
		Where("product_categories.category_id = ?", categoryID).
		Find(&products).Error
	return products, err
}

func (r *productRepository) Search(query string) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Where("name ILIKE ? OR description ILIKE ?",
		"%"+query+"%", "%"+query+"%").Find(&products).Error
	return products, err
}

func (r *productRepository) GetProductWithRelations(id uint) (*models.ProductResponse, error) {
	var product models.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}

	images, _ := r.GetProductImages(id)
	variants, _ := r.GetProductVariants(id)
	attributes, _ := r.GetProductAttributes(id)
	categories, _ := r.GetProductCategories(id)

	return &models.ProductResponse{
		Product:    product,
		Images:     images,
		Variants:   variants,
		Attributes: attributes,
		Categories: categories,
	}, nil
}

func (r *productRepository) CreateImage(image *models.ProductImage) error {
	return r.db.Create(image).Error
}

func (r *productRepository) DeleteImage(id uint) error {
	return r.db.Delete(&models.ProductImage{}, id).Error
}

func (r *productRepository) GetProductImages(productID uint) ([]models.ProductImage, error) {
	var images []models.ProductImage
	err := r.db.Where("product_id = ?", productID).Order("sort_order").Find(&images).Error
	return images, err
}

func (r *productRepository) CreateVariant(variant *models.ProductVariant) error {
	return r.db.Create(variant).Error
}

func (r *productRepository) UpdateVariant(variant *models.ProductVariant) error {
	return r.db.Save(variant).Error
}

func (r *productRepository) DeleteVariant(id uint) error {
	return r.db.Delete(&models.ProductVariant{}, id).Error
}

func (r *productRepository) GetProductVariants(productID uint) ([]models.ProductVariant, error) {
	var variants []models.ProductVariant
	err := r.db.Where("product_id = ?", productID).Find(&variants).Error
	return variants, err
}

func (r *productRepository) CreateAttribute(attribute *models.ProductAttribute) error {
	return r.db.Create(attribute).Error
}

func (r *productRepository) UpdateAttribute(attribute *models.ProductAttribute) error {
	return r.db.Save(attribute).Error
}

func (r *productRepository) DeleteAttribute(id uint) error {
	return r.db.Delete(&models.ProductAttribute{}, id).Error
}

func (r *productRepository) GetProductAttributes(productID uint) ([]models.ProductAttribute, error) {
	var attributes []models.ProductAttribute
	err := r.db.Where("product_id = ?", productID).Find(&attributes).Error
	return attributes, err
}

func (r *productRepository) AddCategory(productID, categoryID uint) error {
	return r.db.Exec("INSERT INTO product_categories (product_id, category_id) VALUES (?, ?) ON CONFLICT DO NOTHING",
		productID, categoryID).Error
}

func (r *productRepository) RemoveCategory(productID, categoryID uint) error {
	return r.db.Exec("DELETE FROM product_categories WHERE product_id = ? AND category_id = ?",
		productID, categoryID).Error
}

func (r *productRepository) GetProductCategories(productID uint) ([]uint, error) {
	var categoryIDs []uint
	err := r.db.Table("product_categories").
		Where("product_id = ?", productID).
		Pluck("category_id", &categoryIDs).Error
	return categoryIDs, err
}
