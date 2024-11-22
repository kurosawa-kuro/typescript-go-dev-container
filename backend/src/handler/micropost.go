package handler

import (
	"fmt"
	"net/http"

	"backend/src/middleware"
	"backend/src/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MicropostHandler struct {
	db *gorm.DB
}

func NewMicropostHandler(db *gorm.DB) *MicropostHandler {
	return &MicropostHandler{db: db}
}

func (h *MicropostHandler) Create(c *gin.Context) {
	// GetAuthUser 認証済みユーザー情報を取得
	userID, _, _ := middleware.GetAuthUser(c)
	fmt.Println("userID:", userID)

	var micropost model.Micropost
	if err := c.ShouldBindJSON(&micropost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ユーザーIDを設定
	micropost.UserID = userID

	if err := h.db.Create(&micropost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, micropost)
}

func (h *MicropostHandler) FindAll(c *gin.Context) {
	var microposts []model.Micropost
	if err := h.db.Find(&microposts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, microposts)
}
