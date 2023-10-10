package html

import (
	"aio-commerce/internal/providers/view"
	"github.com/gin-gonic/gin"
)

func Render(c *gin.Context, ID int, name string, data gin.H) {
	data = view.WithGlobalData(c, data)

	c.HTML(ID, name, data)
}
