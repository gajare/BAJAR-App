package models

import (
	"time"
)

type Category struct {
	ID          uint      `json:"id"`
	ParentID    *uint     `json:"parent_id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Slug        string    `json:"slug"`
	CreatedAt   time.Time `json:"created_at"`
}

type CategoryResponse struct {
	ID          uint               `json:"id"`
	ParentID    *uint              `json:"parent_id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Slug        string             `json:"slug"`
	CreatedAt   time.Time          `json:"created_at"`
	Children    []CategoryResponse `json:"children,omitempty"`
}

type CreateCategoryRequest struct {
	ParentID    *uint  `json:"parent_id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
}

type UpdateCategoryRequest struct {
	ParentID    *uint  `json:"parent_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
}
