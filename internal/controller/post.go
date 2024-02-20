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
	return ctx.SendString("CreatePosts")
}

func (c *PostController) GetPosts(ctx *fiber.Ctx) error {
	return ctx.SendString("GetPosts")
}
