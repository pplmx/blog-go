package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pplmx/blog-go/internal/biz"
	blogv1 "github.com/pplmx/pb/blog/v1"
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
	resp := &blogv1.CreatePostResponse{}

	// TODO: implement the business logic of CreatePost

	return resp, nil
}

func (s *PostService) UpdatePost(ctx context.Context, req *blogv1.UpdatePostRequest) (*blogv1.UpdatePostResponse, error) {
	resp := &blogv1.UpdatePostResponse{}

	// TODO: implement the business logic of UpdatePost

	return resp, nil
}

func (s *PostService) DeletePost(ctx context.Context, req *blogv1.DeletePostRequest) (*blogv1.DeletePostResponse, error) {
	resp := &blogv1.DeletePostResponse{}

	// TODO: implement the business logic of DeletePost

	return resp, nil
}

func (s *PostService) GetPostByID(ctx context.Context, req *blogv1.GetPostByIDRequest) (*blogv1.GetPostByIDResponse, error) {
	resp := &blogv1.GetPostByIDResponse{}

	// TODO: implement the business logic of GetPostByID

	return resp, nil
}

func (s *PostService) ListPosts(ctx context.Context, req *blogv1.ListPostsRequest) (*blogv1.ListPostsResponse, error) {
	resp := &blogv1.ListPostsResponse{}

	// TODO: implement the business logic of ListPosts

	return resp, nil
}
