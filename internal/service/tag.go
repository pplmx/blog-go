package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pplmx/blog-go/internal/biz"
	blogv1 "github.com/pplmx/pb/blog/v1"
)

type TagService struct {
	blogv1.UnimplementedTagServiceServer

	logger     *log.Logger
	tagUseCase *biz.TagUseCase
}

func NewTagService(logger *log.Logger, tagUseCase *biz.TagUseCase) *TagService {
	return &TagService{logger: logger, tagUseCase: tagUseCase}
}

func (s *TagService) CreateTag(ctx context.Context, req *blogv1.CreateTagRequest) (*blogv1.CreateTagResponse, error) {
	resp := &blogv1.CreateTagResponse{}

	// TODO: implement the business logic of CreateTag

	return resp, nil
}

func (s *TagService) UpdateTag(ctx context.Context, req *blogv1.UpdateTagRequest) (*blogv1.UpdateTagResponse, error) {
	resp := &blogv1.UpdateTagResponse{}

	// TODO: implement the business logic of UpdateTag

	return resp, nil
}

func (s *TagService) DeleteTag(ctx context.Context, req *blogv1.DeleteTagRequest) (*blogv1.DeleteTagResponse, error) {
	resp := &blogv1.DeleteTagResponse{}

	// TODO: implement the business logic of DeleteTag

	return resp, nil
}

func (s *TagService) ListTags(ctx context.Context, req *blogv1.ListTagsRequest) (*blogv1.ListTagsResponse, error) {
	resp := &blogv1.ListTagsResponse{}

	// TODO: implement the business logic of ListTags

	return resp, nil
}
