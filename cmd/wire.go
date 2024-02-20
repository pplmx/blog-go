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
	db.NewDBConn,
	repository.NewPostRepository,
)

var serviceProviderSet = wire.NewSet(
	service.NewPostService,
)

var controllerProviderSet = wire.NewSet(
	controller.NewPostController,
)

var providerSet = wire.NewSet(
	internal.NewFiberApp,
	repositoryProviderSet,
	serviceProviderSet,
	controllerProviderSet,
)

func NewApp() *fiber.App {
	wire.Build(providerSet)
	return nil
}
