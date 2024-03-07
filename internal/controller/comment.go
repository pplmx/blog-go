package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pplmx/blog-go/internal/model"
	"github.com/pplmx/blog-go/internal/service"
)

type CommentController struct {
	service *service.CommentService
}

func NewCommentController(service *service.CommentService) *CommentController {
	return &CommentController{service: service}
}

func (c *CommentController) CreateComment(ctx *fiber.Ctx) error {
	// parse and validate the request body
	var comment model.Comment
	err := ctx.BodyParser(&comment)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// call the service method
	err = c.service.CreateComment(&comment)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// send the response
	return ctx.Status(fiber.StatusCreated).SendString("Comment created")
}

func (c *CommentController) GetCommentsByPostID(ctx *fiber.Ctx) error {
	// get the query parameter
	postID := ctx.Query("postID")
	pid, err := strconv.Atoi(postID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// validate the query parameter
	// ...

	// call the service method
	comments, err := c.service.GetCommentsByPostID(uint(pid))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// send the response
	return ctx.JSON(comments)
}

func (c *CommentController) DeleteComment(ctx *fiber.Ctx) error {
	// get the path parameter
	id := ctx.Params("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// call the service method
	err = c.service.DeleteComment(uint(i))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// send the response
	return ctx.Status(fiber.StatusNoContent).SendString("Comment deleted")
}
