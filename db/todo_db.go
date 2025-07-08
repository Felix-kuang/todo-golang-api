package db

import (
	"fmt"
	"log"
	"os"

	"github.com/Felix-kuang/todo-golang-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DbClient *gorm.DB
)

func Connect() {
	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Todo{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DbClient = db
	fmt.Print("Connected to database")
}
