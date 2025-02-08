package chat

import (
	"Communication/internal/repository"
	"Communication/internal/repository/model"

	"gorm.io/gorm"
)

// 创建chat，有可能之前chat已经存在，但是被删除了
func createChat(userID uint32, targetID uint32) (error, *model.Chat) {
	var dbChat model.Chat
	// 查询是否已经存在
	result := repository.DB.Where("user_id = ? AND target_id = ?", userID, targetID).First(&dbChat)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return result.Error, nil
	}
	if result.RowsAffected != 0 {
		// 如果已经存在，但是被删除了
		if dbChat.IsDeleted {
			dbChat.IsDeleted = false
			result = repository.DB.Save(&dbChat)
			if result.Error != nil {
				return result.Error, nil
			}
			return nil, &dbChat
		} else {
			return nil, &dbChat
		}
	}
	// 不存在，创建
	dbChat.UserID = userID
	dbChat.TargetID = targetID
	result = repository.DB.Create(&dbChat)
	if result.Error != nil {
		return result.Error, nil
	}

	return nil, &dbChat
}
