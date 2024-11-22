package util

import (
	"github.com/gin-gonic/gin"
)

// SetAuthCookie JWTトークンをHTTPOnlyクッキーに設定
func SetAuthCookie(c *gin.Context, token string) {
	c.SetCookie(
		CookieName,
		token,
		int(TokenExpiry.Seconds()),
		"/",
		"",
		false, // HTTPS only
		true,  // HTTPOnly
	)
}

// ClearAuthCookie 認証クッキーを削除
func ClearAuthCookie(c *gin.Context) {
	c.SetCookie(
		CookieName,
		"",
		-1,
		"/",
		"",
		false,
		true,
	)
}

// GetAuthCookie クッキーからトークンを取得
func GetAuthCookie(c *gin.Context) (string, error) {
	return c.Cookie(CookieName)
}
