package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/pplmx/blog-go/internal/controller"
)

func NewFiberApp(
	postController *controller.PostController,
	commentController *controller.CommentController,
	tagController *controller.TagController,
	categoryController *controller.CategoryController,
) *fiber.App {
	app := fiber.New()

	// load middleware
	app.Use(
		healthcheck.New(), // /livez and /readyz
		cors.New(),
		cache.New(),
		logger.New(),
		pprof.New(), // /debug/pprof
	)

	// load router
	router(app, postController, commentController, tagController, categoryController)

	return app
}

func router(app *fiber.App, postController *controller.PostController, commentController *controller.CommentController, tagController *controller.TagController, categoryController *controller.CategoryController) {
	routerV1(app, postController, commentController, tagController, categoryController)
	routerV2(app, postController, commentController, tagController, categoryController)
}

func routerV1(app *fiber.App, postController *controller.PostController, commentController *controller.CommentController, tagController *controller.TagController, categoryController *controller.CategoryController) {
	v1 := app.Group("/api/v1")
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World, v1!")
	})

	posts := v1.Group("/posts")
	posts.Get("/", postController.GetPosts)
	posts.Post("/", postController.CreatePosts)
	posts.Get("/:id", postController.GetPostByID)
	posts.Put("/:id", postController.UpdatePost)
	posts.Delete("/:id", postController.DeletePost)

	comments := v1.Group("/comments")
	comments.Get("/", commentController.GetCommentsByPostID)
	comments.Post("/", commentController.CreateComment)
	comments.Delete("/:id", commentController.DeleteComment)

	tags := v1.Group("/tags")
	tags.Get("/", tagController.GetAllTags)
	tags.Post("/", tagController.CreateTag)
	tags.Get("/:id", tagController.GetTagByID)
	tags.Get("/name", tagController.GetTagByName)
	tags.Delete("/:id", tagController.DeleteTag)

	categories := v1.Group("/categories")
	categories.Get("/", categoryController.GetAllCategories)
	categories.Post("/", categoryController.CreateCategory)
	categories.Get("/:id", categoryController.GetCategoryByID)
	categories.Get("/name", categoryController.GetCategoryByName)
	categories.Delete("/:id", categoryController.DeleteCategory)
}

func routerV2(app *fiber.App, postController *controller.PostController, commentController *controller.CommentController, tagController *controller.TagController, categoryController *controller.CategoryController) {
	v2 := app.Group("/api/v2")
	v2.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World, v2!")
	})

	posts := v2.Group("/posts")
	posts.Get("/", postController.GetPosts)
	posts.Post("/", postController.CreatePosts)
	posts.Get("/:id", postController.GetPostByID)
	posts.Put("/:id", postController.UpdatePost)
	posts.Delete("/:id", postController.DeletePost)

	comments := v2.Group("/comments")
	comments.Get("/", commentController.GetCommentsByPostID)
	comments.Post("/", commentController.CreateComment)
	comments.Delete("/:id", commentController.DeleteComment)

	tags := v2.Group("/tags")
	tags.Get("/", tagController.GetAllTags)
	tags.Post("/", tagController.CreateTag)
	tags.Get("/:id", tagController.GetTagByID)
	tags.Get("/name", tagController.GetTagByName)
	tags.Delete("/:id", tagController.DeleteTag)

	categories := v2.Group("/categories")
	categories.Get("/", categoryController.GetAllCategories)
	categories.Post("/", categoryController.CreateCategory)
	categories.Get("/:id", categoryController.GetCategoryByID)
	categories.Get("/name", categoryController.GetCategoryByName)
	categories.Delete("/:id", categoryController.DeleteCategory)
}
