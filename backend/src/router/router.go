package router

import (
	"backend/src/handler"
	"backend/src/middleware"

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

	// Auth routes - パブリックルート
	auth := r.Group("/api/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	// Auth routes - 認証が必要なルート
	authProtected := r.Group("/api/auth")
	authProtected.Use(middleware.IsAuthenticated())
	{
		authProtected.POST("/logout", authHandler.Logout)
		authProtected.GET("/user", authHandler.User)
	}

	// マイクロポストルート
	microposts := r.Group("/api/microposts")
	{
		// パブリックエンドポイント
		microposts.GET("", micropostHandler.FindAll)

		// 認証が必要なエンドポイント
		authenticated := microposts.Group("")
		authenticated.Use(middleware.IsAuthenticated())
		{
			authenticated.POST("", micropostHandler.Create)
		}
	}
}
