package handler

import (
	"net/http"

	"backend/src/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandler struct {
	db *gorm.DB
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{db: db}
}

const (
	defaultRole = "user"
	hashCost    = 10
)

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), hashCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process password"})
		return
	}

	// ユーザー作成
	user := model.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     defaultRole,
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
	// メール認証
	// メールアドレスでユーザー検索
	// パスワード照合
	// セッション管理
	// JWTトークン生成
	// HTTPOnlyクッキーに24時間保存
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
}

func (h *AuthHandler) User(c *gin.Context) {
	// クッキーからJWT取得
	// トークン検証
	// データ取得
	// ユーザーID抽出
	// プロフィール情報取得
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
}
