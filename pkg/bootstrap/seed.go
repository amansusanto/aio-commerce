package bootstrap

import (
	"aio-commerce/internal/database/seeder"
	"aio-commerce/pkg/config"
	"aio-commerce/pkg/database"
)

func Seed() {
	config.Set()

	database.Connect()

	seeder.Seed()
}
