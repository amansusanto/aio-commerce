package database

import (
	"aio-commerce/pkg/config"
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
)

func Connect() {
	cfg := config.Get()

	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	//	cfg.DB.Username,
	//	cfg.DB.Password,
	//	cfg.DB.Host,
	//	cfg.DB.Port,
	//	cfg.DB.Name,
	//)
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)
	log.Println(dsn)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to database")
		return
	}

	DB = db
}
