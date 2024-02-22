package repository

import (
	"github.com/pplmx/blog-go/internal/model"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) Save(comment *model.Comment) error {
	return r.db.Save(comment).Error
}

func (r *CommentRepository) GetAllByPostID(postID uint) ([]*model.Comment, error) {
	var comments []*model.Comment
	err := r.db.Where("post_id = ?", postID).Find(&comments).Error
	return comments, err
}

func (r *CommentRepository) Delete(id uint) error {
	return r.db.Delete(&model.Comment{}, id).Error
}
