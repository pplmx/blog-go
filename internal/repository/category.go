package repository

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pplmx/blog-go/internal/biz"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db  *gorm.DB
	log *log.Helper
}

func NewCategoryRepository(db *gorm.DB, logger log.Logger) biz.CategoryRepo {
	return &categoryRepository{
		db:  db,
		log: log.NewHelper(logger),
	}
}

func (r *categoryRepository) SaveCategory(category *biz.Category) error {
	//TODO implement me
	panic("implement me")
}

func (r *categoryRepository) GetCategoryByID(u uint) (*biz.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (r *categoryRepository) GetCategoryByName(s string) (*biz.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (r *categoryRepository) GetAllCategories() ([]*biz.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (r *categoryRepository) DeleteCategory(u uint) error {
	//TODO implement me
	panic("implement me")
}

func (r *categoryRepository) DeleteCategoryByName(s string) error {
	//TODO implement me
	panic("implement me")
}
