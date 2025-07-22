package models

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Get database credentials from environment variables
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "root" // fallback for development
	}
	
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		log.Println("Warning: DB_PASSWORD not set, using empty password")
	}
	
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "3306"
	}
	
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "API_superapps"
	}
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
		dbUser, dbPassword, dbHost, dbPort, dbName)
	
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = database.AutoMigrate(&Product{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	DB = database
	log.Println("Database connected successfully")
}
