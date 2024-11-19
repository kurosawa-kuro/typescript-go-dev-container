package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingHandler(t *testing.T) {
	// テスト用のGinコンテキストを設定
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	// テストケースの定義
	tests := []struct {
		name           string
		expectedCode   int
		expectedBody   map[string]interface{}
		setupMock      func()
		validateResult func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:         "正常系: 200とpongメッセージを返す",
			expectedCode: http.StatusOK,
			expectedBody: map[string]interface{}{
				"message": "pong",
			},
			setupMock: func() {},
			validateResult: func(t *testing.T, w *httptest.ResponseRecorder) {
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)

				assert.NoError(t, err)
				assert.Equal(t, "pong", response["message"])
			},
		},
	}

	// テストの実行
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ルートの設定
			r.GET("/ping", PingHandler)

			// モックのセットアップ
			tt.setupMock()

			// リクエストの実行
			req := httptest.NewRequest(http.MethodGet, "/ping", nil)
			r.ServeHTTP(w, req)

			// アサーション
			assert.Equal(t, tt.expectedCode, w.Code)
			tt.validateResult(t, w)
		})
	}
}
