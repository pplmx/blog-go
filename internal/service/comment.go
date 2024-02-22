package service

import (
	"github.com/pplmx/blog-go/internal/model"
	"github.com/pplmx/blog-go/internal/repository"
)

type CommentService struct {
	repo *repository.CommentRepository
}

func NewCommentService(repo *repository.CommentRepository) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) CreateComment(comment *model.Comment) error {
	// validate comment data
	// add business logic
	// call repo.Create(comment)
	return s.repo.Save(comment)
}

func (s *CommentService) GetCommentsByPostID(postID uint) ([]*model.Comment, error) {
	// validate postID
	// call repo.GetAllByPostID(postID)
	// add business logic
	return s.repo.GetAllByPostID(postID)
}

func (s *CommentService) DeleteComment(id uint) error {
	// validate id
	// call repo.Delete(id)
	// add business logic
	return s.repo.Delete(id)
}
