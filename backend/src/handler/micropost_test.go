package handler

import (
	"backend/src/model"
	"backend/src/test"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMicropostHandler_Create(t *testing.T) {
	// Setup
	db := test.SetupTestDB(t)
	defer test.CleanupTest(t, db)

	handler := NewMicropostHandler(db)
	router := gin.New()
	router.POST("/microposts", handler.Create)

	// Test data
	micropost := model.Micropost{
		Title: "Test Post",
	}
	jsonData, _ := json.Marshal(micropost)

	// Create request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/microposts", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)

	var response model.Micropost
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, micropost.Title, response.Title)
	assert.NotZero(t, response.ID)
}

func TestMicropostHandler_FindAll(t *testing.T) {
	// Setup
	db := test.SetupTestDB(t)
	defer test.CleanupTest(t, db)

	handler := NewMicropostHandler(db)
	router := gin.New()
	router.GET("/microposts", handler.FindAll)

	// Create test data
	testPosts := []model.Micropost{
		{Title: "Test Post 1"},
		{Title: "Test Post 2"},
	}
	for _, post := range testPosts {
		db.Create(&post)
	}

	// Create request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/microposts", nil)
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var response []model.Micropost
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 2)
	assert.Equal(t, testPosts[0].Title, response[0].Title)
	assert.Equal(t, testPosts[1].Title, response[1].Title)
}
