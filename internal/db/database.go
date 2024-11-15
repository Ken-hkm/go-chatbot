// database.go
package db

import (
	"fmt"
	"go-chatbot/internal/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"os"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PWD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		TablePrefix: "chat.",
	}})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Database connected!")
	DB = db

	// Auto-migrate models
	err_migration := db.AutoMigrate(&models.User{}, &models.ChatMessage{})
	if err_migration != nil {
		log.Fatalf("Could not migrate database: %v", err_migration)
	}
	log.Println("Database migration completed!")
}
