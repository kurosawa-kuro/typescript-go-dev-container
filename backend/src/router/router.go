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

	// ルートの設定
	r.GET("/ping", handler.PingHandler)

	microposts := r.Group("/microposts")
	{
		microposts.POST("", micropostHandler.Create)
		microposts.GET("", micropostHandler.FindAll)
	}
}
