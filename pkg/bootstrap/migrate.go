package bootstrap

import (
	"aio-commerce/internal/database/migration"
	"aio-commerce/pkg/config"
	"aio-commerce/pkg/database"
)

func Migrate() {
	config.Set()

	database.Connect()

	migration.Migrate()
}
