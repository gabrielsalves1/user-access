package config

import (
	"fmt"
	"log"
	"os"
	"user-access/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db_port := os.Getenv("POSTGRES_PORT")
	db_password := os.Getenv("POSTGRES_PASSWORD")
	db_user := os.Getenv("POSTGRES_USER")
	db_name := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", db_user, db_password, db_name, db_port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect database")
	}

	log.Println("Connected successfully to the database")

	Migrations(db)
	return db
}

func Migrations(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
