package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pplmx/blog-go/internal/model"
	"github.com/pplmx/blog-go/internal/service"
	"strconv"
)

type CategoryController struct {
	service *service.CategoryService
}

func NewCategoryController(service *service.CategoryService) *CategoryController {
	return &CategoryController{service: service}
}

func (c *CategoryController) CreateCategory(ctx *fiber.Ctx) error {
	// parse and validate the request body
	var category model.Category
	err := ctx.BodyParser(&category)
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	// call the service method
	err = c.service.CreateCategory(&category)
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	// send the response
	return ctx.Status(201).SendString("Category created")
}

func (c *CategoryController) GetCategoryByID(ctx *fiber.Ctx) error {
	// get the path parameter
	id := ctx.Params("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	// call the service method
	category, err := c.service.GetCategoryByID(uint(i))
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	// send the response
	return ctx.JSON(category)
}

func (c *CategoryController) GetCategoryByName(ctx *fiber.Ctx) error {
	// get the query parameter
	name := ctx.Query("name")

	// validate the query parameter
	// ...

	// call the service method
	category, err := c.service.GetCategoryByName(name)
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	// send the response
	return ctx.JSON(category)
}

func (c *CategoryController) DeleteCategory(ctx *fiber.Ctx) error {
	// get the path parameter
	id := ctx.Params("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	// call the service method
	err = c.service.DeleteCategory(uint(i))
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	// send the response
	return ctx.Status(200).SendString("Category deleted")
}

func (c *CategoryController) GetAllCategories(ctx *fiber.Ctx) error {
	// call the service method
	categories, err := c.service.GetAllCategories()
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	// send the response
	return ctx.JSON(categories)
}
