package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pplmx/blog-go/internal/model"
	"github.com/pplmx/blog-go/internal/service"
	"strconv"
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
	posts, err := c.service.GetAllPosts()
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	return ctx.JSON(posts)
}

func (c *PostController) GetPostByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	post, err := c.service.GetPostByID(uint(i))
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	return ctx.JSON(post)
}

func (c *PostController) UpdatePost(ctx *fiber.Ctx) error {
	// get the path parameter
	id := ctx.Params("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	// parse and validate the request body
	var post model.Post
	err = ctx.BodyParser(&post)
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	post.ID = uint(i)

	// call the service method
	err = c.service.UpdatePost(&post)
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	// send the response
	return ctx.Status(200).SendString("Post updated")
}

func (c *PostController) DeletePost(ctx *fiber.Ctx) error {
	// get the path parameter
	id := ctx.Params("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	// call the service method
	err = c.service.DeletePost(uint(i))
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	// send the response
	return ctx.Status(200).SendString("Post deleted")
}
