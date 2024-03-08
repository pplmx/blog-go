package biz

import "github.com/go-kratos/kratos/v2/log"

type Category struct {
	ID          uint
	Name        string
	Slug        string
	Description string
	Count       uint
}

// CategoryRepo defines the storage interface for category.
type CategoryRepo interface {
	SaveCategory(*Category) error
	GetCategoryByID(uint) (*Category, error)
	GetCategoryByName(string) (*Category, error)
	GetAllCategories() ([]*Category, error)
	DeleteCategory(uint) error
	DeleteCategoryByName(string) error
}

// CategoryUseCase defines the use case for category.
type CategoryUseCase struct {
	logger *log.Logger
	repo   CategoryRepo
}

// NewCategoryUseCase creates a new CategoryUseCase.
func NewCategoryUseCase(repo CategoryRepo, logger *log.Logger) *CategoryUseCase {
	return &CategoryUseCase{repo: repo, logger: logger}
}

// CreateCategory creates a new category.
func (uc *CategoryUseCase) CreateCategory(c *Category) error {
	return uc.repo.SaveCategory(c)
}

// GetCategoryByID gets a category by its ID.
func (uc *CategoryUseCase) GetCategoryByID(id uint) (*Category, error) {
	return uc.repo.GetCategoryByID(id)
}

// GetCategoryByName gets a category by its name.
func (uc *CategoryUseCase) GetCategoryByName(name string) (*Category, error) {
	return uc.repo.GetCategoryByName(name)
}

// GetAllCategories gets all categories.
func (uc *CategoryUseCase) GetAllCategories() ([]*Category, error) {
	return uc.repo.GetAllCategories()
}

// DeleteCategory deletes a category by its ID.
func (uc *CategoryUseCase) DeleteCategory(id uint) error {
	return uc.repo.DeleteCategory(id)
}

// DeleteCategoryByName deletes a category by its name.
func (uc *CategoryUseCase) DeleteCategoryByName(name string) error {
	return uc.repo.DeleteCategoryByName(name)
}
