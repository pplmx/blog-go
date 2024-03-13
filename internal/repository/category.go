package repository

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pplmx/blog-go/internal/biz"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db  *gorm.DB
	log *log.Logger
}

func (r *CategoryRepository) SaveCategory(category *biz.Category) error {
	//TODO implement me
	panic("implement me")
}

func (r *CategoryRepository) GetCategoryByID(u uint) (*biz.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (r *CategoryRepository) GetCategoryByName(s string) (*biz.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (r *CategoryRepository) GetAllCategories() ([]*biz.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (r *CategoryRepository) DeleteCategory(u uint) error {
	//TODO implement me
	panic("implement me")
}

func (r *CategoryRepository) DeleteCategoryByName(s string) error {
	//TODO implement me
	panic("implement me")
}

func NewCategoryRepository(db *gorm.DB, log *log.Logger) biz.CategoryRepo {
	return &CategoryRepository{
		db:  db,
		log: log,
	}
}
