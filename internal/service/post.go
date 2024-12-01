package service

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pplmx/blog-go/internal/biz"
	blogv1 "github.com/pplmx/pb/blog/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PostService struct {
	blogv1.UnimplementedPostServiceServer

	log         *log.Helper
	postUseCase *biz.PostUseCase
}

func NewPostService(logger log.Logger, postUseCase *biz.PostUseCase) *PostService {
	return &PostService{log: log.NewHelper(logger), postUseCase: postUseCase}
}

func (s *PostService) CreatePost(ctx context.Context, req *blogv1.CreatePostRequest) (*blogv1.CreatePostResponse, error) {
	s.log.WithContext(ctx).Infof("CreatePost: title=%s", req.Title)

	// 调用业务逻辑层创建文章
	post, err := s.postUseCase.Create(ctx, &biz.Post{
		Title:    req.Title,
		Content:  req.Content,
		Author:   req.GetAuthor(),
		Category: req.GetCategory(),
	})
	if err != nil {
		s.log.WithContext(ctx).Errorf("Failed to create post: %v", err)
		return nil, err
	}

	return &blogv1.CreatePostResponse{
		Id:        post.ID,
		CreatedAt: timestamppb.New(time.Now()),
	}, nil
}

func (s *PostService) UpdatePost(ctx context.Context, req *blogv1.UpdatePostRequest) (*blogv1.UpdatePostResponse, error) {
	s.log.WithContext(ctx).Infof("UpdatePost: id=%d", req.Id)

	// 准备更新数据
	updateData := &biz.Post{
		ID: req.Id,
	}

	// 只更新非空的字段
	if req.Title != nil {
		updateData.Title = *req.Title
	}
	if req.Content != nil {
		updateData.Content = *req.Content
	}
	if req.Category != nil {
		updateData.Category = *req.Category
	}

	// 调用业务逻辑层更新文章
	post, err := s.postUseCase.Update(ctx, updateData)
	if err != nil {
		s.log.WithContext(ctx).Errorf("Failed to update post: %v", err)
		return nil, err
	}

	return &blogv1.UpdatePostResponse{
		Id:        post.ID,
		UpdatedAt: timestamppb.New(time.Now()),
	}, nil
}

func (s *PostService) DeletePost(ctx context.Context, req *blogv1.DeletePostRequest) (*blogv1.DeletePostResponse, error) {
	s.log.WithContext(ctx).Infof("DeletePost: id=%d", req.Id)

	// 调用业务逻辑层删除文章
	err := s.postUseCase.Delete(ctx, req.Id)
	if err != nil {
		s.log.WithContext(ctx).Errorf("Failed to delete post: %v", err)
		return nil, err
	}

	return &blogv1.DeletePostResponse{
		Id:      req.Id,
		Success: true,
	}, nil
}

func (s *PostService) GetPostByID(ctx context.Context, req *blogv1.GetPostByIDRequest) (*blogv1.GetPostByIDResponse, error) {
	s.log.WithContext(ctx).Infof("GetPostByID: id=%d", req.Id)

	// 调用业务逻辑层获取文章
	post, err := s.postUseCase.GetByID(ctx, req.Id)
	if err != nil {
		s.log.WithContext(ctx).Errorf("Failed to get post: %v", err)
		return nil, err
	}

	return &blogv1.GetPostByIDResponse{
		Id:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		Author:    &post.Author,
		Category:  &post.Category,
		CreatedAt: timestamppb.New(post.CreatedAt),
		UpdatedAt: timestamppb.New(post.UpdatedAt),
	}, nil
}

func (s *PostService) ListPosts(ctx context.Context, req *blogv1.ListPostsRequest) (*blogv1.ListPostsResponse, error) {
	s.log.WithContext(ctx).Infof("ListPosts: page=%d, pageSize=%d", req.Page, req.PageSize)

	// 准备查询条件
	listOptions := &biz.ListOptions{
		Page:     req.Page,
		PageSize: req.PageSize,
		Author:   req.GetAuthor(),
		Category: req.GetCategory(),
	}

	// 调用业务逻辑层获取文章列表
	posts, totalCount, err := s.postUseCase.List(ctx, listOptions)
	if err != nil {
		s.log.WithContext(ctx).Errorf("Failed to list posts: %v", err)
		return nil, err
	}

	// 转换为 protobuf 响应
	respPosts := make([]*blogv1.ListPostsResponse_Post, len(posts))
	for i, post := range posts {
		respPosts[i] = &blogv1.ListPostsResponse_Post{
			Id:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			Author:    &post.Author,
			Category:  &post.Category,
			CreatedAt: timestamppb.New(post.CreatedAt),
		}
	}

	return &blogv1.ListPostsResponse{
		Posts:      respPosts,
		TotalCount: totalCount,
	}, nil
}
