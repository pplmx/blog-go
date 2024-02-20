package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/pplmx/blog-go/internal/controller"
)

type controllers struct {
	post *controller.PostController
}

func NewFiberApp(postController *controller.PostController) *fiber.App {
	app := fiber.New()

	// load middleware
	app.Use(
		healthcheck.New(),
		cors.New(),
		cache.New(),
		logger.New(),
	)

	// load router
	router(app, postController)

	return app
}

func router(app *fiber.App, postController *controller.PostController) {
	routerV1(app, postController)
	routerV2(app, postController)
}

func routerV1(app *fiber.App, postController *controller.PostController) {
	v1 := app.Group("/api/v1")
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World, v1!")
	})

	posts := v1.Group("/posts")
	posts.Get("/", postController.GetPosts)
	posts.Post("/", postController.CreatePosts)
}

func routerV2(app *fiber.App, postController *controller.PostController) {
	v2 := app.Group("/api/v2")
	v2.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World, v2!")
	})

	posts := v2.Group("/posts")
	posts.Get("/", postController.GetPosts)
	posts.Post("/", postController.CreatePosts)
}
