package view

import (
	"aio-commerce/internal/modules/user/helpers"
	"aio-commerce/pkg/converters"
	"aio-commerce/pkg/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func WithGlobalData(c *gin.Context, data gin.H) gin.H {
	data["APP_NAME"] = viper.Get("App.Name")
	data["APP_Slogan"] = viper.Get("App.Slogan")
	data["ERRORS"] = converters.StringToMap(sessions.Flash(c, "errors"))
	data["OLD"] = converters.StringToUrlValues(sessions.Flash(c, "old"))
	data["AUTH"] = helpers.Auth(c)
	return data
}
