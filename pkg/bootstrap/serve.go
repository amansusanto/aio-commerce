package bootstrap

import (
	"aio-commerce/pkg/config"
	"aio-commerce/pkg/database"
	"aio-commerce/pkg/html"
	"aio-commerce/pkg/routing"
	"aio-commerce/pkg/sessions"
	"aio-commerce/pkg/static"
)

func Serve() {
	config.Set()

	database.Connect()

	routing.Init()

	sessions.Start(routing.GetRouter())

	static.LoadStatic(routing.GetRouter())

	html.LoadHTML(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()
}
