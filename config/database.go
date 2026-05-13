package config

import (
	"fmt"
	"log"

	"learn-echo/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsn := "host=localhost user=postgres password=password123 dbname=learn-echo port=5432 sslmode=disable"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	fmt.Println("database connected")

	database.AutoMigrate(
		&model.User{},
		&model.Task{},
	)

	DB = database
}