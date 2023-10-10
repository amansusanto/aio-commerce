package routes

import (
	ProductController "aio-commerce/internal/modules/product/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	productController := ProductController.New()
	router.GET("products/:code", productController.Show)

}
