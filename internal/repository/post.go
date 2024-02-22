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

// Save saves a post with new data, or creates a new post
func (r *PostRepository) Save(post *model.Post) error {
	return r.db.Save(post).Error
}

// GetAll returns all posts
func (r *PostRepository) GetAll() ([]*model.Post, error) {
	var posts []*model.Post
	err := r.db.Find(&posts).Error
	return posts, err
}

// GetByID returns a post by its ID
func (r *PostRepository) GetByID(id uint) (*model.Post, error) {
	var post model.Post
	err := r.db.First(&post, id).Error
	return &post, err
}

// Delete deletes a post by its ID
func (r *PostRepository) Delete(id uint) error {
	return r.db.Delete(&model.Post{}, id).Error
}
