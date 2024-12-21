package main

import (
	"log"

	"github.com/abdisetiakawan/go-restfulapi/internal/config"
	"github.com/abdisetiakawan/go-restfulapi/internal/model"
)

func main() {
    // Initialize database
    db, err := config.NewDatabaseConnection()
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Run migration
    err = db.AutoMigrate(&model.Category{})
    if err != nil {
        log.Fatal("Failed to migrate database:", err)
    }
}