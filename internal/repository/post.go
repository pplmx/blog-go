package repository

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pplmx/blog-go/internal/biz"
	"gorm.io/gorm"
)

func (r *postRepository) SavePost(post *biz.Post) error {
	//TODO implement me
	panic("implement me")
}

func (r *postRepository) GetPostByID(u uint) (*biz.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (r *postRepository) GetPostBySlug(s string) (*biz.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (r *postRepository) GetAllPosts() ([]*biz.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (r *postRepository) DeletePost(u uint) error {
	//TODO implement me
	panic("implement me")
}

func (r *postRepository) DeletePostBySlug(s string) error {
	//TODO implement me
	panic("implement me")
}

func NewPostRepository(db *gorm.DB, log *log.Logger) biz.PostRepo {
	return &postRepository{
		db:  db,
		log: log,
	}
}

type postRepository struct {
	db  *gorm.DB
	log *log.Logger
}
