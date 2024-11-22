package main

import (
	"backend/src/config"
	"backend/src/database"
	"log"
)

func main() {
	// データベース接続
	db := config.SetupDB()

	// マイグレーション実行
	if err := database.Migrate(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("Database migrations completed successfully")

	// シードデータ投入
	if err := database.Seed(db); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}
	log.Println("Database seeding completed successfully")
}
