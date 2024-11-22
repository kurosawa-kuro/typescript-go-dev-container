package middleware

import (
	"backend/src/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// IsAuthenticated 認証済みユーザーかチェックするミドルウェア
func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		// クッキーからトークンを取得
		tokenString, err := util.GetAuthCookie(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "No token provided",
			})
			c.Abort()
			return
		}

		// トークンを検証
		claims, err := util.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			c.Abort()
			return
		}

		// 検証済みのユーザー情報をコンテキストに保存
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		c.Next()
	}
}

// GetAuthUser コンテキストから認証済みユーザー情報を取得
func GetAuthUser(c *gin.Context) (uint, string, string) {
	userID, _ := c.Get("user_id")
	email, _ := c.Get("email")
	role, _ := c.Get("role")
	return userID.(uint), email.(string), role.(string)
}
