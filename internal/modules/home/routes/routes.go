package routes

import (
	homeCtrl "aio-commerce/internal/modules/home/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func formatNumber(value string) string {
	s, err := strconv.Atoi(value)
	formatted := ""
	if err == nil {
		formatted = fmt.Sprintf("%.0f", s)
	}
	return formatted
}

func Routes(router *gin.Engine) {
	homeController := homeCtrl.New()
	router.GET("/", homeController.Index)
}
