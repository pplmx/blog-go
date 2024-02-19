package controller

import "github.com/pplmx/blog-go/internal/service"

type PostController struct {
	service *service.PostService
}

func NewPostController(service *service.PostService) *PostController {
	return &PostController{service: service}
}

func (c *PostController) CreatePost() error {
	return c.service.CreatePost()
}
