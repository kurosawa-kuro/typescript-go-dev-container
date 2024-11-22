package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

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
	userID, _, _ := middleware.GetAuthUser(c)

	// フォームデータを取得
	title := c.PostForm("title")
	file, err := c.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var imagePath string
	if file != nil {
		// ファイル名をユニークにする
		filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
		// 保存先のパスを設定
		imagePath = filepath.Join("uploads", filename)

		// uploadsディレクトリが存在しない場合は作成
		if err := os.MkdirAll("uploads", 0755); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// ファイルを保存
		if err := c.SaveUploadedFile(file, imagePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// Micropostを作成
	micropost := model.Micropost{
		UserID:    userID,
		Title:     title,
		ImagePath: imagePath,
	}

	if err := h.db.Create(&micropost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, micropost)
}

func (h *MicropostHandler) FindAll(c *gin.Context) {
	var microposts []model.Micropost

	// created_at でDESC（降順）ソート
	if err := h.db.Order("created_at DESC").Find(&microposts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, microposts)
}
