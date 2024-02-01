package db

import (
	"github.com/pplmx/blog-go/internal/model"
	"gorm.io/gorm"
)

type PostRepo struct {
	db *gorm.DB
}

func (repo *PostRepo) Create() error {
	return repo.db.Create(&model.Post{}).Error
}
