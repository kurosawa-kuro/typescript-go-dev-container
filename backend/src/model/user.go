package model

import (
	"time"
)

type User struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Email      string    `gorm:"not null;unique" json:"email"`
	Password   string    `gorm:"not null" json:"password"`
	Token      string    `json:"token"`
	AvatarPath string    `json:"avatar_path"`
	Role       string    `gorm:"not null" json:"role"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
