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

func (controller *Controller) Show(c *gin.Context) {
	code := c.Param("code")
	product, err := controller.productService.Find(code)
	if err != nil {
		html.Render(c, http.StatusNotFound, "templates/errors/html/404", gin.H{"title": "Data " + code + " tidak ditemukan", "message": err.Error()})
		return
	}
	html.Render(c, http.StatusOK, "modules/product/html/show", gin.H{"title": "Data List", "product": product})

}
