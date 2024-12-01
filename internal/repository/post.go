package repository

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pplmx/blog-go/internal/biz"
	"gorm.io/gorm"
)

type postRepository struct {
	db  *gorm.DB
	log *log.Helper
}

func NewPostRepository(db *gorm.DB, logger log.Logger) biz.PostRepo {
	return &postRepository{
		db:  db,
		log: log.NewHelper(logger),
	}
}

// CreatePost 创建新文章
func (r *postRepository) CreatePost(ctx context.Context, post *biz.Post) (*biz.Post, error) {
	result := r.db.WithContext(ctx).Create(post)
	if result.Error != nil {
		r.log.Errorf("Failed to create post: %v", result.Error)
		return nil, result.Error
	}
	return post, nil
}

// UpdatePost 更新文章
func (r *postRepository) UpdatePost(ctx context.Context, post *biz.Post) (*biz.Post, error) {
	if post.ID == 0 {
		return nil, errors.New("post ID is required for update")
	}

	result := r.db.WithContext(ctx).Save(post)
	if result.Error != nil {
		r.log.Errorf("Failed to update post: %v", result.Error)
		return nil, result.Error
	}
	return post, nil
}

// GetPostByID 根据 ID 获取文章
func (r *postRepository) GetPostByID(ctx context.Context, id uint32) (*biz.Post, error) {
	var post biz.Post
	result := r.db.WithContext(ctx).First(&post, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		r.log.Errorf("Failed to get post by ID: %v", result.Error)
		return nil, result.Error
	}
	return &post, nil
}

// GetPostBySlug 根据 Slug 获取文章
func (r *postRepository) GetPostBySlug(ctx context.Context, slug string) (*biz.Post, error) {
	var post biz.Post
	result := r.db.WithContext(ctx).Where("slug = ?", slug).First(&post)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		r.log.Errorf("Failed to get post by slug: %v", result.Error)
		return nil, result.Error
	}
	return &post, nil
}

// ListPosts 列出文章，支持分页和过滤
func (r *postRepository) ListPosts(ctx context.Context, opts *biz.ListOptions) ([]*biz.Post, uint32, error) {
	var posts []*biz.Post
	var total int64

	// 构建查询
	query := r.db.WithContext(ctx)

	// 添加过滤条件
	if opts.Author != "" {
		query = query.Where("author = ?", opts.Author)
	}
	if opts.Category != "" {
		query = query.Where("category = ?", opts.Category)
	}
	if opts.Status != "" {
		query = query.Where("status = ?", opts.Status)
	}
	if opts.Slug != "" {
		query = query.Where("slug = ?", opts.Slug)
	}

	// 计算总数
	if err := query.Model(&biz.Post{}).Count(&total).Error; err != nil {
		r.log.Errorf("Failed to count posts: %v", err)
		return nil, 0, err
	}

	// 分页查询
	page := opts.Page
	if page == 0 {
		page = 1
	}
	pageSize := opts.PageSize
	if pageSize == 0 {
		pageSize = 10 // 默认每页10条
	}

	err := query.
		Order("created_at DESC").
		Offset(int((page - 1) * pageSize)).
		Limit(int(pageSize)).
		Find(&posts).Error

	if err != nil {
		r.log.Errorf("Failed to list posts: %v", err)
		return nil, 0, err
	}

	return posts, uint32(total), nil
}

// DeletePost 根据 ID 删除文章
func (r *postRepository) DeletePost(ctx context.Context, id uint32) error {
	result := r.db.WithContext(ctx).Delete(&biz.Post{}, id)
	if result.Error != nil {
		r.log.Errorf("Failed to delete post: %v", result.Error)
		return result.Error
	}
	return nil
}

// DeletePostBySlug 根据 Slug 删除文章
func (r *postRepository) DeletePostBySlug(ctx context.Context, slug string) error {
	result := r.db.WithContext(ctx).
		Where("slug = ?", slug).
		Delete(&biz.Post{})

	if result.Error != nil {
		r.log.Errorf("Failed to delete post by slug: %v", result.Error)
		return result.Error
	}
	return nil
}
