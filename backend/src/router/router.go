package router

import (
	"backend/src/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Setup initializes the router and its routes
func Setup(db *gorm.DB, r *gin.Engine) {
	// ハンドラーの初期化
	micropostHandler := handler.NewMicropostHandler(db)
	healthHandler := handler.NewHealthHandler(db)
	authHandler := handler.NewAuthHandler(db)

	// ヘルスチェックルート
	health := r.Group("/api/health")
	{
		health.GET("", healthHandler.CheckHealth)
		health.GET("/db", healthHandler.CheckDBConnection)
		health.GET("/db/dev", healthHandler.CheckDevDatabase)
		health.GET("/db/test", healthHandler.CheckTestDatabase)
	}

	// マイクロポストルート
	microposts := r.Group("/api/microposts")
	{
		microposts.POST("", micropostHandler.Create)
		microposts.GET("", micropostHandler.FindAll)
	}

	// Auth
	auth := r.Group("/api/auth")
	{
		// 新規ユーザー登録
		auth.POST("/register", authHandler.Register)

		// ログイン

		auth.POST("/login", authHandler.Login)

		// ログアウト
		auth.POST("/logout", authHandler.Logout)

		// 認証確認

		auth.GET("/user", authHandler.User)

	}

}
