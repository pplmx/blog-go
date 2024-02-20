package service

import (
	"github.com/pplmx/blog-go/internal/model"
	"github.com/pplmx/blog-go/internal/repository"
)

type PostService struct {
	repo *repository.PostRepository
	// if needed, add much more other repos or service directly
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePosts() error {
	return s.repo.Create()
}

func (s *PostService) GetPosts() ([]*model.Post, error) {
	return s.repo.GetAll()
}
