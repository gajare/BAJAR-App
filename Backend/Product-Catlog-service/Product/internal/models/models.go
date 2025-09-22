package models

import (
	"encoding/json"
	"time"
)

type Product struct {
	ID            uint        `json:"id"`
	Name          string      `json:"name" binding:"required"`
	Description   string      `json:"description"`
	Price         float64     `json:"price" binding:"required"`
	DiscountPrice float64     `json:"discount_price"`
	SKU           string      `json:"sku" binding:"required"`
	Stock         int         `json:"stock" binding:"required"`
	Weight        float64     `json:"weight"`
	Dimensions    Dimensions  `json:"dimensions"`
	Status        string      `json:"status"` // active, inactive, draft
	Slug          string      `json:"slug"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	Images        []Image     `json:"images,omitempty"`
	Variants      []Variant   `json:"variants,omitempty"`
	Attributes    []Attribute `json:"attributes,omitempty"`
}

type Dimensions struct {
	Length float64 `json:"length"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

type Image struct {
	ID        uint   `json:"id"`
	ProductID uint   `json:"product_id"`
	ImageURL  string `json:"image_url" binding:"required"`
	AltText   string `json:"alt_text"`
	SortOrder int    `json:"sort_order"`
}

type Variant struct {
	ID        uint    `json:"id"`
	ProductID uint    `json:"product_id"`
	Name      string  `json:"name" binding:"required"`  // e.g., "Size", "Color"
	Value     string  `json:"value" binding:"required"` // e.g., "XL", "Red"
	Price     float64 `json:"price"`
	Stock     int     `json:"stock"`
	SKU       string  `json:"sku"`
}

type Attribute struct {
	ID        uint   `json:"id"`
	ProductID uint   `json:"product_id"`
	Key       string `json:"key" binding:"required"`   // e.g., "Material"
	Value     string `json:"value" binding:"required"` // e.g., "Cotton"
}

type CreateProductRequest struct {
	Name          string      `json:"name" binding:"required"`
	Description   string      `json:"description"`
	Price         float64     `json:"price" binding:"required"`
	DiscountPrice float64     `json:"discount_price"`
	SKU           string      `json:"sku" binding:"required"`
	Stock         int         `json:"stock" binding:"required"`
	Weight        float64     `json:"weight"`
	Dimensions    Dimensions  `json:"dimensions"`
	Status        string      `json:"status"`
	Slug          string      `json:"slug"`
	Images        []Image     `json:"images"`
	Variants      []Variant   `json:"variants"`
	Attributes    []Attribute `json:"attributes"`
}

type UpdateProductRequest struct {
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	Price         float64     `json:"price"`
	DiscountPrice float64     `json:"discount_price"`
	SKU           string      `json:"sku"`
	Stock         int         `json:"stock"`
	Weight        float64     `json:"weight"`
	Dimensions    Dimensions  `json:"dimensions"`
	Status        string      `json:"status"`
	Slug          string      `json:"slug"`
}

type ProductResponse struct {
	ID            uint        `json:"id"`
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	Price         float64     `json:"price"`
	DiscountPrice float64     `json:"discount_price"`
	SKU           string      `json:"sku"`
	Stock         int         `json:"stock"`
	Weight        float64     `json:"weight"`
	Dimensions    Dimensions  `json:"dimensions"`
	Status        string      `json:"status"`
	Slug          string      `json:"slug"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	Images        []Image     `json:"images"`
	Variants      []Variant   `json:"variants"`
	Attributes    []Attribute `json:"attributes"`
}

type SearchRequest struct {
	Query    string   `json:"query"`
	Category string   `json:"category"`
	MinPrice float64  `json:"min_price"`
	MaxPrice float64  `json:"max_price"`
	Status   string   `json:"status"`
	Tags     []string `json:"tags"`
	Page     int      `json:"page"`
	Limit    int      `json:"limit"`
}

// Implement Scanner and Valuer interfaces for Dimensions
func (d *Dimensions) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, d)
}

func (d Dimensions) Value() ([]byte, error) {
	return json.Marshal(d)
}