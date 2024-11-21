package handler

import (
	"net/http"

	"backend/src/model"
	"backend/src/util"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthHandler struct {
	db *gorm.DB
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{db: db}
}

func (h *AuthHandler) Register(c *gin.Context) {
	// リクエストのバインドとバリデーション
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// メールアドレス重複チェック
	var existingUser model.User
	if err := h.db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	// パスワードハッシュ化
	const hashCost = 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), hashCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process password"})
		return
	}

	// ユーザー作成
	user := model.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     "user",
	}

	if err := h.db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// パスワードを除外してレスポンスを返す
	user.Password = ""
	c.JSON(http.StatusCreated, user)
}

func (h *AuthHandler) Login(c *gin.Context) {
	// リクエストのバインドとバリデーション
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// メールアドレスでユーザー検索
	var user model.User
	if err := h.db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// パスワード照合
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// JWTトークン生成
	tokenString, err := util.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// トークンをDBに保存
	user.Token = tokenString
	if err := h.db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save token"})
		return
	}

	// HTTPOnlyクッキーに保存
	c.SetCookie(
		util.CookieName,
		tokenString,
		int(util.TokenExpiry.Seconds()),
		"/",
		"",    // ドメイン
		false, // HTTPS only
		true,  // HTTPOnly
	)

	// レスポンス返却
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged in",
		"user": gin.H{
			"id":    user.ID,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
}

func (h *AuthHandler) User(c *gin.Context) {
	// クッキーからJWT取得
	tokenString, err := c.Cookie(util.CookieName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
		return
	}

	// JWTトークン検証
	claims, err := util.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// ユーザー情報取得
	var user model.User
	if err := h.db.Where("id = ?", claims.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// パスワードとトークンを除外してレスポンスを返す
	user.Password = ""
	user.Token = ""
	c.JSON(http.StatusOK, user)
}
