package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pplmx/blog-go/internal/service"
)

type PostController struct {
	service *service.PostService
}

func NewPostController(service *service.PostService) *PostController {
	return &PostController{service: service}
}

func (c *PostController) CreatePosts(ctx *fiber.Ctx) error {
	err := c.service.CreatePosts()
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	return ctx.Status(201).SendString("Post created")
}

func (c *PostController) GetPosts(ctx *fiber.Ctx) error {
	posts, err := c.service.GetPosts()
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	return ctx.JSON(posts)
}
