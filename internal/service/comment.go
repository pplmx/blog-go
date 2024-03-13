package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pplmx/blog-go/internal/biz"
	blogv1 "github.com/pplmx/pb/blog/v1"
)

type CommentService struct {
	blogv1.UnimplementedCommentServiceServer

	log            *log.Helper
	commentUseCase *biz.CommentUseCase
}

func NewCommentService(logger log.Logger, commentUseCase *biz.CommentUseCase) *CommentService {
	return &CommentService{log: log.NewHelper(logger), commentUseCase: commentUseCase}
}

func (s *CommentService) CreateComment(ctx context.Context, req *blogv1.CreateCommentRequest) (*blogv1.CreateCommentResponse, error) {
	resp := &blogv1.CreateCommentResponse{}

	// TODO: implement the business logic of CreateComment

	return resp, nil
}
