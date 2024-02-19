package service

import (
	"github.com/pplmx/blog-go/internal/repository"
)

type PostService struct {
	repo *repository.PostRepository
	// if needed, add much more other repos or service directly
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost() error {
	return s.repo.Create()
}
