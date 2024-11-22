package util

import "golang.org/x/crypto/bcrypt"

const HashCost = 14

// HashPassword パスワードをハッシュ化
func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), HashCost)
	return string(hashedBytes), err
}

// ComparePasswords ハッシュ化されたパスワードと平文パスワードを比較
func ComparePasswords(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}
