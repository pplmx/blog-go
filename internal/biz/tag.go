package biz

import "github.com/go-kratos/kratos/v2/log"

type Tag struct {
	ID          uint
	Name        string
	Slug        string
	Count       uint
	Description string
}

// TagRepo defines the storage interface for tag.
type TagRepo interface {
	SaveTag(*Tag) error
	GetTagByID(uint) (*Tag, error)
	GetTagByName(string) (*Tag, error)
	GetAllTags() ([]*Tag, error)
	DeleteTag(uint) error
	DeleteTagByName(string) error
}

// TagUseCase defines the use case for tag.
type TagUseCase struct {
	logger *log.Logger
	repo   TagRepo
}

func NewTagUseCase(repo TagRepo, logger *log.Logger) *TagUseCase {
	return &TagUseCase{repo: repo, logger: logger}
}

func (uc *TagUseCase) CreateTag(c *Tag) error {
	return uc.repo.SaveTag(c)
}

func (uc *TagUseCase) GetTagByID(id uint) (*Tag, error) {
	return uc.repo.GetTagByID(id)
}

func (uc *TagUseCase) GetTagByName(name string) (*Tag, error) {
	return uc.repo.GetTagByName(name)
}

func (uc *TagUseCase) GetAllTags() ([]*Tag, error) {
	return uc.repo.GetAllTags()
}

func (uc *TagUseCase) DeleteTag(id uint) error {
	return uc.repo.DeleteTag(id)
}

func (uc *TagUseCase) DeleteTagByName(name string) error {
	return uc.repo.DeleteTagByName(name)
}
