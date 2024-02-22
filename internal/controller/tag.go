package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pplmx/blog-go/internal/model"
	"github.com/pplmx/blog-go/internal/service"
	"strconv"
)

type TagController struct {
	service *service.TagService
}

func NewTagController(service *service.TagService) *TagController {
	return &TagController{service: service}
}

func (c *TagController) CreateTag(ctx *fiber.Ctx) error {
	// parse and validate the request body
	var tag model.Tag
	err := ctx.BodyParser(&tag)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// call the service method
	err = c.service.CreateTag(&tag)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// send the response
	return ctx.Status(fiber.StatusCreated).SendString("Tag created")
}

func (c *TagController) GetTagByID(ctx *fiber.Ctx) error {
	// get the path parameter
	id := ctx.Params("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// call the service method
	tag, err := c.service.GetTagByID(uint(i))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// send the response
	return ctx.JSON(tag)
}

func (c *TagController) GetTagByName(ctx *fiber.Ctx) error {
	// get the query parameter
	name := ctx.Query("name")

	// validate the query parameter
	// ...

	// call the service method
	tag, err := c.service.GetTagByName(name)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// send the response
	return ctx.JSON(tag)
}

func (c *TagController) DeleteTag(ctx *fiber.Ctx) error {
	// get the path parameter
	id := ctx.Params("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// call the service method
	err = c.service.DeleteTag(uint(i))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// send the response
	return ctx.Status(fiber.StatusNoContent).SendString("Tag deleted")
}

func (c *TagController) GetAllTags(ctx *fiber.Ctx) error {
	// call the service method
	tags, err := c.service.GetAllTags()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// send the response
	return ctx.JSON(tags)
}
