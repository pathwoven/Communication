package chat

import (
	"Communication/internal/repository"
	"Communication/internal/repository/model"
	"Communication/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetChatListHandler(c *gin.Context) {
	// 获取参数
	// 获取user_id
	userID, ok := utils.GetUserID(c.Request)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "登录信息失效"})
		return
	}
	// 查询数据库
	// 查询chat表，不显示已删除的聊天
	var dbChats []model.Chat
	result := repository.DB.Where("user_id = ? AND is_deleted = false", userID).Find(&dbChats)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误，无法查询到用户"})
		return
	}
	// 处理查询结果为空的情况
	if len(dbChats) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "没有找到相关聊天记录", "data": []interface{}{}})
		return
	}
	// 查询entity表
	// 合并 chat 和 entity 数据
	type ChatWithEntity struct {
		ID           uint
		TargetID     uint
		IsGroup      bool
		Avatar       string
		DisplayID    string
		Name         string
		Remark       string
		UnreadCount  uint8
		LastMessage  string
		LastTime     string
		OnlineStatus uint8
		Status       string
		IsPinned     bool
		IsMuted      bool
		IsBlocked    bool
		Tag1         string
		Tag2         string
		Tag3         string
	}
	var chatsWithEntity []ChatWithEntity
	for _, chat := range dbChats {
		var targetEntity model.Entity
		var targetContact model.Contact
		entityResult := repository.DB.Select("is_group", "avatar", "display_id", "name", "remark", "online_status", "status").Where("id = ?", chat.TargetID).First(&targetEntity)
		if entityResult.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误，无法查询到用户"})
			return
		}
		chatsWithEntity = append(chatsWithEntity, ChatWithEntity{
			ID:           chat.ID,
			TargetID:     chat.TargetID,
			IsGroup:      targetEntity.IsGroup,
			Avatar:       targetEntity.Avatar,
			DisplayID:    targetEntity.DisplayID,
			Name:         targetEntity.Name,
			Remark:       targetContact.Remark,
			UnreadCount:  chat.UnreadCount,
			LastMessage:  chat.LastMessage,
			LastTime:     chat.LastTime.Format("2006-01-02 15:04:05"),
			OnlineStatus: targetEntity.OnlineStatus,
			Status:       targetEntity.Status,
			IsPinned:     chat.IsPinned,
			IsMuted:      chat.IsMuted,
			IsBlocked:    chat.IsBlocked,
			Tag1:         chat.Tag1,
			Tag2:         chat.Tag2,
			Tag3:         chat.Tag3,
		})
	}
	// 返回结果
	c.JSON(http.StatusOK, gin.H{"data": chatsWithEntity})
}
