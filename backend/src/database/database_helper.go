package database

import (
	"backend/src/model"
	"backend/src/util"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Micropost{})
	return nil
}

// Seed データベースにシードデータを投入
func Seed(db *gorm.DB) error {
	// 既存のデータを削除（テーブル構造は保持）
	if err := cleanDatabase(db); err != nil {
		return err
	}

	// 管理者ユーザーの作成
	admin, err := createAdminUser(db)
	if err != nil {
		return err
	}

	// テストユーザーの作成
	users, err := createTestUsers(db)
	if err != nil {
		return err
	}

	log.Printf("Created admin user: %s", admin.Email)
	log.Printf("Created %d test users", len(users))
	log.Println("Seed data has been successfully loaded")
	return nil
}

// cleanDatabase 既存のデータを削除
func cleanDatabase(db *gorm.DB) error {
	// PostgreSQL用の外部キー制約の無効化
	if err := db.Exec("SET CONSTRAINTS ALL DEFERRED").Error; err != nil {
		return err
	}

	// テーブルの削除
	if err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Micropost{}).Error; err != nil {
		return err
	}
	if err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.User{}).Error; err != nil {
		return err
	}

	// シーケンスのリセット
	if err := db.Exec("ALTER SEQUENCE microposts_id_seq RESTART WITH 1").Error; err != nil {
		return err
	}
	if err := db.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1").Error; err != nil {
		return err
	}

	// PostgreSQL用の外部キー制約の有効化
	if err := db.Exec("SET CONSTRAINTS ALL IMMEDIATE").Error; err != nil {
		return err
	}

	return nil
}

// createAdminUser 管理者ユーザーを作成
func createAdminUser(db *gorm.DB) (*model.User, error) {
	hashedPassword, err := util.HashPassword("admin123")
	if err != nil {
		return nil, err
	}

	admin := &model.User{
		Email:      "admin@example.com",
		Password:   hashedPassword,
		Role:       "admin",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		AvatarPath: "/avatars/admin.png",
	}

	if err := db.Create(admin).Error; err != nil {
		return nil, err
	}

	return admin, nil
}

// createTestUsers テストユーザーを作成
func createTestUsers(db *gorm.DB) ([]*model.User, error) {
	users := make([]*model.User, 0)

	for i := 1; i <= 5; i++ {
		hashedPassword, err := util.HashPassword("password")
		if err != nil {
			return nil, err
		}

		user := &model.User{
			Email:      fmt.Sprintf("user%d@example.com", i),
			Password:   hashedPassword,
			Role:       "user",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			AvatarPath: fmt.Sprintf("/avatars/user%d.png", i),
		}

		if err := db.Create(user).Error; err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// createMicroposts マイクロポストを作成
// func createMicroposts(db *gorm.DB, users []*model.User) error {
// 	titles := []string{
// 		"First Post",
// 		"Hello World",
// 		"Testing Micropost",
// 		"Another Post",
// 		"Final Test",
// 	}

// 	for _, user := range users {
// 		for _, title := range titles {
// 			micropost := &model.Micropost{
// 				Title:     title,
// 				CreatedAt: time.Now(),
// 				UpdatedAt: time.Now(),
// 			}

// 			if err := db.Create(micropost).Error; err != nil {
// 				return err
// 			}
// 		}
// 	}

// 	return nil
// }
