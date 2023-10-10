package controllers

import (
	ProductService "aio-commerce/internal/modules/product/services"
	"aio-commerce/pkg/html"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	productService ProductService.ProductServiceInterface
}

func New() *Controller {
	return &Controller{
		productService: ProductService.New(),
	}
}

func (controller *Controller) Index(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
		"title":    "Home page",
		"featured": controller.productService.GetFeaturedProducts(),
		"product":  controller.productService.GetStoriesProducts(),
	})
}
