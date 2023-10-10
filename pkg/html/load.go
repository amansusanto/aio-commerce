package html

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func formatNumber(value string) string {
	s, err := strconv.Atoi(value)
	formatted := ""
	if err == nil {
		formatted = fmt.Sprintf("%.2f", s)
	}
	return formatted
}

func LoadHTML(router *gin.Engine) {
	router.LoadHTMLGlob("internal/**/**/**/*tmpl")
}
