package config

import (
	"fmt"
	"os"

	"backend/src/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	return SetupDBWithConfig(false)
}

func SetupTestDB() *gorm.DB {
	return SetupDBWithConfig(true)
}

func SetupDBWithConfig(isTest bool) *gorm.DB {
	var (
		host     string
		user     string
		password string
		dbname   string
		port     string
	)

	if isTest {
		host = os.Getenv("TEST_DATABASE_HOST")
		user = os.Getenv("TEST_POSTGRES_USER")
		password = os.Getenv("TEST_POSTGRES_PASSWORD")
		dbname = os.Getenv("TEST_POSTGRES_DB")
		port = os.Getenv("TEST_POSTGRES_PORT")
	} else {
		host = os.Getenv("DOCKER_DATABASE_HOST")
		user = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		dbname = os.Getenv("POSTGRES_DB")
		port = os.Getenv("POSTGRES_PORT")
	}

	// Verify that all required variables are present
	if host == "" || user == "" || password == "" || dbname == "" {
		panic("Missing required database environment variables")
	}

	fmt.Printf("Connecting to PostgreSQL with:\nHost: %s\nUser: %s\nDB: %s\nPort: %s\n",
		host, user, dbname, port)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		host, user, password, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Connection error: %v\n", err)
		panic("Failed to connect database")
	}

	fmt.Println("Successfully connected to database!")

	// マイグレーション
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Micropost{})

	return db
}
