package service

import (
	"github.com/pplmx/blog-go/internal/model"
	"github.com/pplmx/blog-go/internal/repository"
)

type TagService struct {
	repo *repository.TagRepository
}

func NewTagService(repo *repository.TagRepository) *TagService {
	return &TagService{repo: repo}
}

func (s *TagService) CreateTag(tag *model.Tag) error {
	// validate tag data
	// add business logic
	// call repo.Create(tag)
	return s.repo.Save(tag)
}

func (s *TagService) GetTagByID(id uint) (*model.Tag, error) {
	// validate id
	// call repo.GetByID(id)
	// add business logic
	return s.repo.GetByID(id)
}

func (s *TagService) GetTagByName(name string) (*model.Tag, error) {
	// validate name
	// call repo.GetByName(name)
	// add business logic
	return s.repo.GetByName(name)
}

func (s *TagService) DeleteTag(id uint) error {
	// validate id
	// call repo.Delete(id)
	// add business logic
	return s.repo.Delete(id)
}

func (s *TagService) GetAllTags() ([]*model.Tag, error) {
	// call repo.GetAll()
	// add business logic
	return s.repo.GetAll()
}
