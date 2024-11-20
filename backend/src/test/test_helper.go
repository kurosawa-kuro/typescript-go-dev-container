package test

import (
	"backend/src/config"
	"backend/src/model"
	"testing"
	"time"

	"gorm.io/gorm"
)

func SetupTestDB(t *testing.T) *gorm.DB {
	// テスト用の環境変数を設定
	t.Setenv("DOCKER_DATABASE_HOST", "test-db")
	t.Setenv("POSTGRES_USER", "postgres")
	t.Setenv("POSTGRES_PASSWORD", "postgres")
	t.Setenv("POSTGRES_DB", "test_db")
	t.Setenv("POSTGRES_PORT", "5432")

	// データベースへの接続を試行（リトライ付き）
	var db *gorm.DB
	maxRetries := 5

	for i := 0; i < maxRetries; i++ {
		db = config.SetupDB()
		if db != nil {
			// 接続テスト
			sqlDB, err := db.DB()
			if err == nil {
				if err := sqlDB.Ping(); err == nil {
					break
				}
			}
		}
		t.Logf("Waiting for database to be ready... (attempt %d/%d)", i+1, maxRetries)
		time.Sleep(2 * time.Second)
	}

	if db == nil {
		t.Fatal("Failed to connect to database after multiple attempts")
	}

	// テスト前にデータベースをクリーン
	if err := db.Migrator().DropTable(&model.Micropost{}); err != nil {
		t.Fatalf("Failed to drop table: %v", err)
	}
	if err := db.AutoMigrate(&model.Micropost{}); err != nil {
		t.Fatalf("Failed to migrate table: %v", err)
	}

	return db
}

func CleanupTest(t *testing.T, db *gorm.DB) {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			t.Logf("Failed to get underlying sql.DB: %v", err)
			return
		}
		sqlDB.Close()
	}
}
