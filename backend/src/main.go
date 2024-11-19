package main

import (
	"backend/src/config"
	"backend/src/router"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"log"
)

// @title           Your API Title
// @version         1.0
// @description     Your API Description
// @host           localhost:8000
// @BasePath       /
func main() {
	// データベースの初期化
	db := config.SetupDB()
	// defer db.Close()

	// Ginのインスタンスを作成
	r := gin.Default()

	// CORSミドルウェアの設定
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://frontend:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	// Swaggerの初期化
	config.SetupSwagger()

	// ルーターの設定
	router.Setup(db, r)

	// サーバーの起動 (with error handling)
	if err := r.Run(":8000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
