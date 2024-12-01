package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type Post struct {
	ID        uint32
	Title     string
	Content   string
	Slug      string
	Status    string
	Author    string
	Category  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// PostRepo defines the storage interface for post.
type PostRepo interface {
	// 创建文章
	CreatePost(ctx context.Context, post *Post) (*Post, error)

	// 更新文章
	UpdatePost(ctx context.Context, post *Post) (*Post, error)

	// 根据 ID 获取文章
	GetPostByID(ctx context.Context, id uint32) (*Post, error)

	// 根据 Slug 获取文章
	GetPostBySlug(ctx context.Context, slug string) (*Post, error)

	// 列出文章
	ListPosts(ctx context.Context, opts *ListOptions) ([]*Post, uint32, error)

	// 删除文章
	DeletePost(ctx context.Context, id uint32) error

	// 根据 Slug 删除文章
	DeletePostBySlug(ctx context.Context, slug string) error
}

// ListOptions 定义列表查询选项
type ListOptions struct {
	Page     uint32
	PageSize uint32
	Author   string
	Category string
	Status   string
	Slug     string
}

// PostUseCase defines the use case for post.
type PostUseCase struct {
	log  *log.Helper
	repo PostRepo
}

func NewPostUseCase(repo PostRepo, logger log.Logger) *PostUseCase {
	return &PostUseCase{repo: repo, log: log.NewHelper(logger)}
}

// Create 创建文章
func (uc *PostUseCase) Create(ctx context.Context, p *Post) (*Post, error) {
	// 可以在这里添加额外的业务逻辑，如生成 Slug、设置默认状态等
	if p.Slug == "" {
		p.Slug = generateSlug(p.Title)
	}

	if p.Status == "" {
		p.Status = "draft"
	}

	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	return uc.repo.CreatePost(ctx, p)
}

// Update 更新文章
func (uc *PostUseCase) Update(ctx context.Context, p *Post) (*Post, error) {
	// 设置更新时间
	p.UpdatedAt = time.Now()

	return uc.repo.UpdatePost(ctx, p)
}

// GetByID 根据 ID 获取文章
func (uc *PostUseCase) GetByID(ctx context.Context, id uint32) (*Post, error) {
	return uc.repo.GetPostByID(ctx, id)
}

// GetBySlug 根据 Slug 获取文章
func (uc *PostUseCase) GetBySlug(ctx context.Context, slug string) (*Post, error) {
	return uc.repo.GetPostBySlug(ctx, slug)
}

// List 列出文章
func (uc *PostUseCase) List(ctx context.Context, opts *ListOptions) ([]*Post, uint32, error) {
	// 可以在这里添加额外的业务逻辑，如权限检查等
	return uc.repo.ListPosts(ctx, opts)
}

// Delete 删除文章
func (uc *PostUseCase) Delete(ctx context.Context, id uint32) error {
	// 可以在这里添加额外的业务逻辑，如删除前的权限检查
	return uc.repo.DeletePost(ctx, id)
}

// DeleteBySlug 根据 Slug 删除文章
func (uc *PostUseCase) DeleteBySlug(ctx context.Context, slug string) error {
	// 可以在这里添加额外的业务逻辑，如删除前的权限检查
	return uc.repo.DeletePostBySlug(ctx, slug)
}

// generateSlug 生成 Slug 的辅助函数（需要实现）
func generateSlug(title string) string {
	// 实现 Slug 生成逻辑
	// 可以使用正则表达式去除特殊字符，转换为小写，空格转换为连字符等
	return ""
}
