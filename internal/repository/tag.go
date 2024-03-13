package repository

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pplmx/blog-go/internal/biz"
	"gorm.io/gorm"
)

func (r *tagRepository) SaveTag(tag *biz.Tag) error {
	//TODO implement me
	panic("implement me")
}

func (r *tagRepository) GetTagByID(u uint) (*biz.Tag, error) {
	//TODO implement me
	panic("implement me")
}

func (r *tagRepository) GetTagByName(s string) (*biz.Tag, error) {
	//TODO implement me
	panic("implement me")
}

func (r *tagRepository) GetAllTags() ([]*biz.Tag, error) {
	//TODO implement me
	panic("implement me")
}

func (r *tagRepository) DeleteTag(u uint) error {
	//TODO implement me
	panic("implement me")
}

func (r *tagRepository) DeleteTagByName(s string) error {
	//TODO implement me
	panic("implement me")
}

func NewTagRepository(db *gorm.DB, log *log.Logger) biz.TagRepo {
	return &tagRepository{
		db:  db,
		log: log,
	}
}

type tagRepository struct {
	db  *gorm.DB
	log *log.Logger
}
