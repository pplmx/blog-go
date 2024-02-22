package service

import (
	"github.com/pplmx/blog-go/internal/model"
	"github.com/pplmx/blog-go/internal/repository"
)

type PostService struct {
	repo *repository.PostRepository
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePosts() error {
	return s.repo.Save(&model.Post{})
}

func (s *PostService) CreatePost(post *model.Post) error {
	// validate post data
	// add business logic
	// call repo.Create(post)
	return s.repo.Save(post)
}

func (s *PostService) GetPostByID(id uint) (*model.Post, error) {
	// validate id
	// call repo.GetByID(id)
	// add business logic
	return s.repo.GetByID(id)
}

func (s *PostService) UpdatePost(post *model.Post) error {
	// validate post data
	// call repo.Update(post)
	// add business logic
	return s.repo.Save(post)
}

func (s *PostService) DeletePost(id uint) error {
	// validate id
	// call repo.Delete(id)
	// add business logic
	return s.repo.Delete(id)
}

func (s *PostService) GetAllPosts() ([]*model.Post, error) {
	// call repo.GetAll()
	// add business logic
	return s.repo.GetAll()
}
