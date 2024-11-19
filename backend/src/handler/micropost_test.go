package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"backend/src/config"
	"backend/src/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	// Set test environment
	os.Setenv("GO_ENV", "test")

	// Get database connection
	db := config.SetupDB()

	// Migrate the schema
	db.AutoMigrate(&model.Micropost{})

	return db
}

func TestCreateMicropost(t *testing.T) {
	db := setupTestDB()
	handler := NewMicropostHandler(db)

	router := gin.Default()
	router.POST("/microposts", handler.Create)

	micropost := model.Micropost{Title: "Test content"}
	body, _ := json.Marshal(micropost)

	req, _ := http.NewRequest("POST", "/microposts", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var createdMicropost model.Micropost
	json.Unmarshal(w.Body.Bytes(), &createdMicropost)
	assert.Equal(t, micropost.Title, createdMicropost.Title)
}

func TestFindAllMicroposts(t *testing.T) {
	db := setupTestDB()
	handler := NewMicropostHandler(db)

	// Insert a test micropost
	db.Create(&model.Micropost{Title: "Test content"})

	router := gin.Default()
	router.GET("/microposts", handler.FindAll)

	req, _ := http.NewRequest("GET", "/microposts", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var microposts []model.Micropost
	json.Unmarshal(w.Body.Bytes(), &microposts)
	assert.NotEmpty(t, microposts)
	assert.Equal(t, "Test content", microposts[0].Title)
}
