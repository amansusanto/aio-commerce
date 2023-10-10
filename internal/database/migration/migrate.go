package migration

import (
	productModels "aio-commerce/internal/modules/product/models"
	userModels "aio-commerce/internal/modules/user/models"
	"aio-commerce/pkg/database"
	"fmt"
	"log"
)

func Migrate() {
	db := database.Connection()

	err := db.AutoMigrate(&userModels.User{}, &productModels.Product{})

	if err != nil {
		log.Fatal("Cant migrate")
		return
	}

	fmt.Println("Migration done ..")
}
