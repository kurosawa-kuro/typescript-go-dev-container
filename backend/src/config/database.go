package config

import (
	"fmt"
	"os"

	"backend/src/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	// 環境変数名を.envファイルに合わせる
	host := os.Getenv("POSTGRES_HOSTNAME")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := "5432" // PostgreSQLのデフォルトポート

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
	db.AutoMigrate(&model.Micropost{})

	return db
}
