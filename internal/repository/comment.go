package repository

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pplmx/blog-go/internal/biz"
	"gorm.io/gorm"
)

type commentRepository struct {
	db  *gorm.DB
	log *log.Helper
}

func NewCommentRepository(db *gorm.DB, logger log.Logger) biz.CommentRepo {
	return &commentRepository{
		db:  db,
		log: log.NewHelper(logger),
	}
}

func (r *commentRepository) SaveComment(comment *biz.Comment) error {
	//TODO implement me
	panic("implement me")
}

func (r *commentRepository) GetCommentsByID(u uint) ([]*biz.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (r *commentRepository) GetPostCommentsByPostID(u uint) ([]*biz.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (r *commentRepository) DeleteComment(u uint) error {
	//TODO implement me
	panic("implement me")
}
