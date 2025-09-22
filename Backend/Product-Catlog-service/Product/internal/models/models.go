package models

import (
	"time"
)

type Product struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name" binding:"required"`
	Description   string    `json:"description"`
	Price         float64   `json:"price" binding:"required"`
	DiscountPrice float64   `json:"discount_price"`
	SKU           string    `json:"sku" binding:"required"`
	Stock         int       `json:"stock" binding:"required"`
	Weight        float64   `json:"weight"`
	Dimensions    string    `json:"dimensions"`
	Status        string    `json:"status"`
	Slug          string    `json:"slug"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ProductImage struct {
	ID        uint   `json:"id"`
	ProductID uint   `json:"product_id"`
	ImageURL  string `json:"image_url" binding:"required"`
	AltText   string `json:"alt_text"`
	SortOrder int    `json:"sort_order"`
}

type ProductVariant struct {
	ID        uint    `json:"id"`
	ProductID uint    `json:"product_id"`
	Name      string  `json:"name" binding:"required"`
	Value     string  `json:"value" binding:"required"`
	Price     float64 `json:"price"`
	Stock     int     `json:"stock"`
	SKU       string  `json:"sku"`
}

type ProductAttribute struct {
	ID        uint   `json:"id"`
	ProductID uint   `json:"product_id"`
	Key       string `json:"key" binding:"required"`
	Value     string `json:"value" binding:"required"`
}

type ProductCategory struct {
	ProductID  uint `json:"product_id"`
	CategoryID uint `json:"category_id"`
}

type CreateProductRequest struct {
	Name          string                    `json:"name" binding:"required"`
	Description   string                    `json:"description"`
	Price         float64                   `json:"price" binding:"required"`
	DiscountPrice float64                   `json:"discount_price"`
	SKU           string                    `json:"sku" binding:"required"`
	Stock         int                       `json:"stock" binding:"required"`
	Weight        float64                   `json:"weight"`
	Dimensions    string                    `json:"dimensions"`
	Status        string                    `json:"status"`
	Slug          string                    `json:"slug"`
	Images        []ProductImageRequest     `json:"images"`
	Variants      []ProductVariantRequest   `json:"variants"`
	Attributes    []ProductAttributeRequest `json:"attributes"`
	CategoryIDs   []uint                    `json:"category_ids"`
}

type UpdateProductRequest struct {
	Name          string                    `json:"name"`
	Description   string                    `json:"description"`
	Price         float64                   `json:"price"`
	DiscountPrice float64                   `json:"discount_price"`
	SKU           string                    `json:"sku"`
	Stock         int                       `json:"stock"`
	Weight        float64                   `json:"weight"`
	Dimensions    string                    `json:"dimensions"`
	Status        string                    `json:"status"`
	Slug          string                    `json:"slug"`
	Images        []ProductImageRequest     `json:"images"`
	Variants      []ProductVariantRequest   `json:"variants"`
	Attributes    []ProductAttributeRequest `json:"attributes"`
	CategoryIDs   []uint                    `json:"category_ids"`
}

type ProductImageRequest struct {
	ImageURL  string `json:"image_url" binding:"required"`
	AltText   string `json:"alt_text"`
	SortOrder int    `json:"sort_order"`
}

type ProductVariantRequest struct {
	Name  string  `json:"name" binding:"required"`
	Value string  `json:"value" binding:"required"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
	SKU   string  `json:"sku"`
}

type ProductAttributeRequest struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type ProductResponse struct {
	Product
	Images     []ProductImage     `json:"images"`
	Variants   []ProductVariant   `json:"variants"`
	Attributes []ProductAttribute `json:"attributes"`
	Categories []uint             `json:"categories"`
}
