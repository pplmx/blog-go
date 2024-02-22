//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/pplmx/blog-go/internal"
	"github.com/pplmx/blog-go/internal/controller"
	"github.com/pplmx/blog-go/internal/db"
	"github.com/pplmx/blog-go/internal/repository"
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

var controllerProviderSet = wire.NewSet(
	controller.NewCategoryController,
	controller.NewCommentController,
	controller.NewPostController,
	controller.NewTagController,
)

var providerSet = wire.NewSet(
	db.NewDBConn,
	internal.NewFiberApp,
	repositoryProviderSet,
	serviceProviderSet,
	controllerProviderSet,
)

func NewApp() *fiber.App {
	wire.Build(providerSet)
	return nil
}
