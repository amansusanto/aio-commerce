package routes

import (
	homeRoutes "aio-commerce/internal/modules/home/routes"
	productRoutes "aio-commerce/internal/modules/product/routes"
	userRoutes "aio-commerce/internal/modules/user/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
	userRoutes.Routes(router)
	productRoutes.Routes(router)
}
