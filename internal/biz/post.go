package biz

import "github.com/go-kratos/kratos/v2/log"

type Post struct {
	ID      uint
	Title   string
	Content *string
	Slug    string
	Status  string
}

// PostRepo defines the storage interface for post.
type PostRepo interface {
	SavePost(*Post) error
	GetPostByID(uint) (*Post, error)
	GetPostBySlug(string) (*Post, error)
	GetAllPosts() ([]*Post, error)
	DeletePost(uint) error
	DeletePostBySlug(string) error
}

// PostUseCase defines the use case for post.
type PostUseCase struct {
	log  *log.Helper
	repo PostRepo
}

func NewPostUseCase(repo PostRepo, logger log.Logger) *PostUseCase {
	return &PostUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *PostUseCase) CreatePost(p *Post) error {
	return uc.repo.SavePost(p)
}

func (uc *PostUseCase) GetPostByID(id uint) (*Post, error) {
	return uc.repo.GetPostByID(id)
}

func (uc *PostUseCase) GetPostBySlug(slug string) (*Post, error) {
	return uc.repo.GetPostBySlug(slug)
}

func (uc *PostUseCase) GetAllPosts() ([]*Post, error) {
	return uc.repo.GetAllPosts()
}

func (uc *PostUseCase) DeletePost(id uint) error {
	return uc.repo.DeletePost(id)
}

func (uc *PostUseCase) DeletePostBySlug(slug string) error {
	return uc.repo.DeletePostBySlug(slug)
}
