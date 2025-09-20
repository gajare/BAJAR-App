package services

import (
	"product-catalog-service/Category/internal/models"
	"product-catalog-service/Category/internal/repository"
)

type CategoryService interface {
	GetAllCategories() ([]models.Category, error)
	GetCategory(id uint) (*models.Category, error)
	CreateCategory(request models.CreateCategoryRequest) (*models.Category, error)
	UpdateCategory(id uint, request models.UpdateCategoryRequest) (*models.Category, error)
	DeleteCategory(id uint) error
	GetCategoryTree() ([]models.CategoryResponse, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) GetAllCategories() ([]models.Category, error) {
	return s.repo.FindAll()
}

func (s *categoryService) GetCategory(id uint) (*models.Category, error) {
	return s.repo.FindByID(id)
}

func (s *categoryService) CreateCategory(request models.CreateCategoryRequest) (*models.Category, error) {
	category := &models.Category{
		ParentID:    request.ParentID,
		Name:        request.Name,
		Description: request.Description,
		Slug:        request.Slug,
	}

	if err := s.repo.Create(category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) UpdateCategory(id uint, request models.UpdateCategoryRequest) (*models.Category, error) {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if request.Name != "" {
		category.Name = request.Name
	}
	if request.Description != "" {
		category.Description = request.Description
	}
	if request.Slug != "" {
		category.Slug = request.Slug
	}
	if request.ParentID != nil {
		category.ParentID = request.ParentID
	}

	if err := s.repo.Update(category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) DeleteCategory(id uint) error {
	return s.repo.Delete(id)
}

func (s *categoryService) GetCategoryTree() ([]models.CategoryResponse, error) {
	categories, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return s.buildTree(categories, nil), nil
}

func (s *categoryService) buildTree(categories []models.Category, parentID *uint) []models.CategoryResponse {
	var tree []models.CategoryResponse

	for _, cat := range categories {
		if (parentID == nil && cat.ParentID == nil) || (parentID != nil && cat.ParentID != nil && *cat.ParentID == *parentID) {
			node := models.CategoryResponse{
				ID:          cat.ID,
				ParentID:    cat.ParentID,
				Name:        cat.Name,
				Description: cat.Description,
				Slug:        cat.Slug,
				CreatedAt:   cat.CreatedAt,
			}

			children := s.buildTree(categories, &cat.ID)
			if len(children) > 0 {
				node.Children = children
			}

			tree = append(tree, node)
		}
	}

	return tree
}
