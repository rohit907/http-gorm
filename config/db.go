package config

import (
	"log"
	"os"

	"github.com/rohit907/gin-gorm/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	// Construct the connection string with sslmode=disable
	dsn := "user=postgres password=postgres dbname=postgres host=db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	log.Println("connected to database")
	db.AutoMigrate(&models.User{})
	DB = db
}
