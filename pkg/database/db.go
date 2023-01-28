package database

import (
	"fmt"
	"log"
	"os"
	"user-access/pkg/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Database() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db_port := os.Getenv("POSTGRES_PORT")
	db_password := os.Getenv("POSTGRES_PASSWORD")
	db_user := os.Getenv("POSTGRES_USER")
	db_name := os.Getenv("POSTGRES_DB")

	connectionString := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", db_user, db_password, db_name, db_port)

	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	Migrations(DB)
}

func Migrations(DB *gorm.DB) {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.System{})
}
