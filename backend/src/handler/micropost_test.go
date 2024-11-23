package handler

import (
	"backend/src/model"
	"backend/src/test"
	"backend/src/util"
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func setupTestUserWithAuth(t *testing.T, db *gorm.DB) (model.User, string) {
	// 既存のユーザーを削除（クリーンアップ）
	db.Exec("DELETE FROM users WHERE email = ?", "test@example.com")

	// テストユーザーを作成
	password := "password123"
	hashedPassword, err := util.HashPassword(password)
	assert.NoError(t, err)

	testUser := model.User{
		Email:    "test@example.com",
		Password: hashedPassword,
		Role:     "user",
	}
	err = db.Create(&testUser).Error
	assert.NoError(t, err)

	// JWTトークンを生成
	token, err := util.GenerateToken(testUser.ID, testUser.Email, testUser.Role)
	assert.NoError(t, err)

	// トークンをDBに保存
	err = db.Model(&testUser).Update("token", token).Error
	assert.NoError(t, err)

	return testUser, token
}

func setupTestRouter(t *testing.T, db *gorm.DB, method, path string, handler gin.HandlerFunc, testUser *model.User, token string) *gin.Engine {
	router := gin.New()

	if testUser != nil && token != "" {
		// 認証が必要な場合、モックミドルウェアを追加
		router.Handle(method, path, func(c *gin.Context) {
			c.Set("user_id", testUser.ID)
			c.Set("email", testUser.Email)
			c.Set("role", testUser.Role)
			c.Request.AddCookie(&http.Cookie{
				Name:  "token",
				Value: token,
			})
			c.Next()
		}, handler)
	} else {
		// 認証が不要な場合、直接ハンドラーを設定
		router.Handle(method, path, handler)
	}

	return router
}

func TestMicropostHandler_Create(t *testing.T) {
	// Setup
	db := test.SetupTestDB(t)
	defer test.CleanupTest(t, db)

	// 認証済みユーザーをセットアップ
	testUser, token := setupTestUserWithAuth(t, db)
	handler := NewMicropostHandler(db)
	router := setupTestRouter(t, db, "POST", "/microposts", handler.Create, &testUser, token)

	// マルチパートフォームデータの作成
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_ = writer.WriteField("title", "Test Post")
	writer.Close()

	// Create request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/microposts", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)

	var response model.Micropost
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Test Post", response.Title)
	assert.NotZero(t, response.ID)
	assert.Equal(t, testUser.ID, response.UserID)
}

func TestMicropostHandler_FindAll(t *testing.T) {
	// Setup
	db := test.SetupTestDB(t)
	defer test.CleanupTest(t, db)

	// 認証済みユーザーをセットアップ
	testUser, _ := setupTestUserWithAuth(t, db)

	// Create test data with valid UserID
	testPosts := []model.Micropost{
		{Title: "Test Post 1", UserID: testUser.ID},
		{Title: "Test Post 2", UserID: testUser.ID},
	}

	// マイクロポストを作成し、エラー処理を追加
	for _, post := range testPosts {
		err := db.Create(&post).Error
		assert.NoError(t, err)
	}

	handler := NewMicropostHandler(db)
	router := setupTestRouter(t, db, "GET", "/microposts", handler.FindAll, nil, "")

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

	// レスポンスの長さを確認してからアクセス
	if assert.Len(t, response, 2) {
		assert.Equal(t, testPosts[1].Title, response[0].Title)
		assert.Equal(t, testPosts[0].Title, response[1].Title)
	}
}
