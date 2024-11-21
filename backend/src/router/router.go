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

	// ヘルスチェックルート
	r.GET("/health", healthHandler.CheckHealth)
	r.GET("/health/db", healthHandler.CheckDBConnection)
	r.GET("/health/db/dev", healthHandler.CheckDevDatabase)
	r.GET("/health/db/test", healthHandler.CheckTestDatabase)

	// マイクロポストルート
	microposts := r.Group("/microposts")
	{
		microposts.POST("", micropostHandler.Create)
		microposts.GET("", micropostHandler.FindAll)
	}
}
