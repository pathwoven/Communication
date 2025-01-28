package auth

import (
	"Communication/internal/repository"
	"Communication/internal/repository/model"
	"math/rand"
	"time"
)

// 生成随机字符串
func generateRandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[randSource.Intn(len(letters))]
	}
	return string(b)
}

// 检查 display_id 是否唯一
func isDisplayIDUnique(displayID string) bool {
	var dbEntity model.Entity
	result := repository.DB.Select("id").Where("display_id = ?", displayID).First(&dbEntity)
	return result.Error != nil
}
