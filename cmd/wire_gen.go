// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/pplmx/blog-go/internal"
	"github.com/pplmx/blog-go/internal/biz"
	"github.com/pplmx/blog-go/internal/db"
	"github.com/pplmx/blog-go/internal/repository"
	"github.com/pplmx/blog-go/internal/server"
	"github.com/pplmx/blog-go/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(logger log.Logger) (*kratos.App, func(), error) {
	gormDB := db.NewDBConn()
	categoryRepo := repository.NewCategoryRepository(gormDB, logger)
	categoryUseCase := biz.NewCategoryUseCase(categoryRepo, logger)
	categoryService := service.NewCategoryService(logger, categoryUseCase)
	commentRepo := repository.NewCommentRepository(gormDB, logger)
	commentUseCase := biz.NewCommentUseCase(commentRepo, logger)
	commentService := service.NewCommentService(logger, commentUseCase)
	postRepo := repository.NewPostRepository(gormDB, logger)
	postUseCase := biz.NewPostUseCase(postRepo, logger)
	postService := service.NewPostService(logger, postUseCase)
	tagRepo := repository.NewTagRepository(gormDB, logger)
	tagUseCase := biz.NewTagUseCase(tagRepo, logger)
	tagService := service.NewTagService(logger, tagUseCase)
	httpServer := server.NewHTTPServer(logger, categoryService, commentService, postService, tagService)
	grpcServer := server.NewGRPCServer(logger, categoryService, commentService, postService, tagService)
	app := internal.NewKratosApp(logger, httpServer, grpcServer)
	return app, func() {
	}, nil
}

// wire.go:

var repositoryProviderSet = wire.NewSet(repository.NewCategoryRepository, repository.NewCommentRepository, repository.NewPostRepository, repository.NewTagRepository)

var serviceProviderSet = wire.NewSet(service.NewCategoryService, service.NewCommentService, service.NewPostService, service.NewTagService)

var bizProviderSet = wire.NewSet(biz.NewCategoryUseCase, biz.NewCommentUseCase, biz.NewPostUseCase, biz.NewTagUseCase)

var serverProviderSet = wire.NewSet(server.NewHTTPServer, server.NewGRPCServer)

var providerSet = wire.NewSet(db.NewDBConn, internal.NewKratosApp, serverProviderSet,
	repositoryProviderSet,
	serviceProviderSet,
	bizProviderSet,
)
