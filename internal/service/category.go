package service

import (
	"github.com/pplmx/blog-go/internal/model"
	"github.com/pplmx/blog-go/internal/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(category *model.Category) error {
	// validate category data
	// add business logic
	// call repo.Create(category)
	return s.repo.Save(category)
}

func (s *CategoryService) GetCategoryByID(id uint) (*model.Category, error) {
	// validate id
	// call repo.GetByID(id)
	// add business logic
	return s.repo.GetByID(id)
}

func (s *CategoryService) GetCategoryByName(name string) (*model.Category, error) {
	// validate name
	// call repo.GetByName(name)
	// add business logic
	return s.repo.GetByName(name)
}

func (s *CategoryService) DeleteCategory(id uint) error {
	// validate id
	// call repo.Delete(id)
	// add business logic
	return s.repo.Delete(id)
}

func (s *CategoryService) GetAllCategories() ([]*model.Category, error) {
	// call repo.GetAll()
	// add business logic
	return s.repo.GetAll()
}
