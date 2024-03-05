package config

import (
	"catalog/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var host = os.Getenv("DB_HOST")
var port = os.Getenv("DB_PORT")
var user = os.Getenv("DB_USER")
var name = os.Getenv("DB_NAME")
var password = os.Getenv("DB_PASSWORD")

func DatabaseConnection() *gorm.DB {
	c := os.Getenv("DB_USER")
	fmt.Print(c)
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, name, password)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}

	log.Println("Database connection successful...")
	db.AutoMigrate(&models.Restaurant{}, &models.MenuItem{}, &models.Location{})
	return db
}
