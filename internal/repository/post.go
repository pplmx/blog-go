package repository

import (
	"github.com/pplmx/blog-go/internal/model"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create() error {
	return r.db.Create(&model.Post{}).Error
}

func (r *PostRepository) GetAll() ([]*model.Post, error) {
	var posts []*model.Post
	err := r.db.Find(&posts).Error
	return posts, err
}
