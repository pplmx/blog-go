package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewFiberApp() *fiber.App {
	app := fiber.New()

	// load middleware
	app.Use(
		healthcheck.New(),
		cors.New(),
		cache.New(),
		logger.New(),
	)

	// /api
	api := app.Group("/api")

	// /api/v1
	v1 := api.Group("/v1")
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World, v1!")
	})

	// /api/v2
	v2 := api.Group("/v2")
	v2.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World, v2!")
	})

	return app
}
