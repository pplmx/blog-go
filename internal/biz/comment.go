package biz

import "github.com/go-kratos/kratos/v2/log"

type Comment struct {
	ID       uint
	PostID   uint
	Content  string
	ParentID uint
	Rating   int
	NickName string
	Email    string
	Website  string
}

// CommentRepo defines the storage interface for comment.
type CommentRepo interface {
	SaveComment(*Comment) error
	GetCommentsByID(uint) ([]*Comment, error) // if the comment has children, its children will be returned as well
	GetPostCommentsByPostID(uint) ([]*Comment, error)
	DeleteComment(uint) error // if the comment has children, its children will be deleted as well
}

// CommentUseCase defines the use case for comment.
type CommentUseCase struct {
	log  *log.Helper
	repo CommentRepo
}

func NewCommentUseCase(repo CommentRepo, logger log.Logger) *CommentUseCase {
	return &CommentUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CommentUseCase) CreateComment(c *Comment) error {
	return uc.repo.SaveComment(c)
}

func (uc *CommentUseCase) GetCommentsByID(id uint) ([]*Comment, error) {
	return uc.repo.GetCommentsByID(id)
}

func (uc *CommentUseCase) GetPostCommentsByPostID(id uint) ([]*Comment, error) {
	return uc.repo.GetPostCommentsByPostID(id)
}

func (uc *CommentUseCase) DeleteComment(id uint) error {
	return uc.repo.DeleteComment(id)
}
