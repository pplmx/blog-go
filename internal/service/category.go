package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pplmx/blog-go/internal/biz"
	blogv1 "github.com/pplmx/pb/blog/v1"
)

type CategoryService struct {
	blogv1.UnimplementedCategoryServiceServer

	logger          *log.Logger
	categoryUseCase *biz.CategoryUseCase
}

func NewCategoryService(logger *log.Logger, categoryUseCase *biz.CategoryUseCase) *CategoryService {
	return &CategoryService{logger: logger, categoryUseCase: categoryUseCase}
}

func (s *CategoryService) CreateCategory(ctx context.Context, req *blogv1.CreateCategoryRequest) (*blogv1.CreateCategoryResponse, error) {
	resp := &blogv1.CreateCategoryResponse{}

	// TODO: implement the business logic of CreateCategory

	return resp, nil
}

func (s *CategoryService) UpdateCategory(ctx context.Context, req *blogv1.UpdateCategoryRequest) (*blogv1.UpdateCategoryResponse, error) {
	resp := &blogv1.UpdateCategoryResponse{}

	// TODO: implement the business logic of UpdateCategory

	return resp, nil
}

func (s *CategoryService) DeleteCategory(ctx context.Context, req *blogv1.DeleteCategoryRequest) (*blogv1.DeleteCategoryResponse, error) {
	resp := &blogv1.DeleteCategoryResponse{}

	// TODO: implement the business logic of DeleteCategory

	return resp, nil
}

func (s *CategoryService) GetCategory(ctx context.Context, req *blogv1.GetCategoryRequest) (*blogv1.GetCategoryResponse, error) {
	resp := &blogv1.GetCategoryResponse{}

	// TODO: implement the business logic of GetCategory

	return resp, nil
}

func (s *CategoryService) ListCategory(ctx context.Context, req *blogv1.ListCategoryRequest) (*blogv1.ListCategoryResponse, error) {
	resp := &blogv1.ListCategoryResponse{}

	// TODO: implement the business logic of ListCategory

	return resp, nil
}
