package repository

import (
	"github.com/pplmx/blog-go/internal/model"
	"gorm.io/gorm"
)

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{db: db}
}

func (r *TagRepository) Save(tag *model.Tag) error {
	return r.db.Save(tag).Error
}

func (r *TagRepository) GetAll() ([]*model.Tag, error) {
	var tags []*model.Tag
	err := r.db.Find(&tags).Error
	return tags, err
}

func (r *TagRepository) GetByID(id uint) (*model.Tag, error) {
	var tag model.Tag
	err := r.db.First(&tag, id).Error
	return &tag, err
}

func (r *TagRepository) GetByName(name string) (*model.Tag, error) {
	var tag model.Tag
	err := r.db.Where("name = ?", name).First(&tag).Error
	return &tag, err
}

func (r *TagRepository) Delete(id uint) error {
	return r.db.Delete(&model.Tag{}, id).Error
}
