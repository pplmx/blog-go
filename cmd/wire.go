//go:build wireinject
// +build wireinject

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

var repositoryProviderSet = wire.NewSet(
	repository.NewCategoryRepository,
	repository.NewCommentRepository,
	repository.NewPostRepository,
	repository.NewTagRepository,
)

var serviceProviderSet = wire.NewSet(
	service.NewCategoryService,
	service.NewCommentService,
	service.NewPostService,
	service.NewTagService,
)

var bizProviderSet = wire.NewSet(
	biz.NewCategoryUseCase,
	biz.NewCommentUseCase,
	biz.NewPostUseCase,
	biz.NewTagUseCase,
)

var serverProviderSet = wire.NewSet(
	server.NewHTTPServer,
	server.NewGRPCServer,
)

var providerSet = wire.NewSet(
	db.NewDBConn,
	internal.NewKratosApp,
	serverProviderSet,
	repositoryProviderSet,
	serviceProviderSet,
	bizProviderSet,
)

// initApp init kratos application.
func initApp(log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(providerSet))
}
