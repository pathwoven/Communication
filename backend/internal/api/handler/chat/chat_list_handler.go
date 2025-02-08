package chat

import (
	"Communication/internal/repository"
	"Communication/internal/repository/model"
	"Communication/internal/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误，无法查询到用户"})
		return
	}
	// 处理查询结果为空的情况
	if len(dbChats) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "没有找到相关聊天记录", "data": []interface{}{}})
		return
	}
	// 查询entity表和contact表
	// 合并 chat 和 entity 数据
	type ChatWithEntity struct {
		ID           uint32 `json:"id"`
		TargetID     uint32 `json:"target_id"`
		IsGroup      bool   `json:"is_group"`
		Avatar       string `json:"avatar"`
		DisplayID    string `json:"display_id"`
		Name         string `json:"name"`
		Remark       string `json:"remark"`
		UnreadCount  uint8  `json:"unread_count"`
		LastMessage  string `json:"last_message"`
		LastTime     string `json:"last_time"`
		LastPerson   string `json:"last_person"`
		OnlineStatus uint8  `json:"online_status"`
		Status       string `json:"status"`
		IsPinned     bool   `json:"is_pinned"`
		IsMuted      bool   `json:"is_muted"`
		IsBlocked    bool   `json:"is_blocked"`
		Tag1         string `json:"tag1"`
		Tag2         string `json:"tag2"`
		Tag3         string `json:"tag3"`
	}
	var chatsWithEntity []ChatWithEntity
	for _, chat := range dbChats {
		var targetEntity model.Entity
		var targetContact model.Contact
		entityResult := repository.DB.Select("is_group", "avatar", "display_id", "name", "online_status", "status").Where("id = ?", chat.TargetID).First(&targetEntity)
		if entityResult.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误，无法查询到用户"})
			return
		}
		contactResult := repository.DB.Select("remark").Where("user_id = ? AND target_id = ?", userID, chat.TargetID).First(&targetContact)
		if contactResult.Error != nil && contactResult.Error != gorm.ErrRecordNotFound {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误，无法查询到用户"})
			return
		}
		var lastTime string
		if chat.LastTime == nil {
			lastTime = ""
		} else {
			lastTime = chat.LastTime.Format("2006-01-02 15:04:05")
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
			LastTime:     lastTime,
			LastPerson:   chat.LastPerson,
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

func CreateChatHandler(c *gin.Context) {
	// 获取参数
	// 获取user_id
	userID, ok := utils.GetUserID(c.Request)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "登录信息失效"})
		return
	}
	// 获取目标用户id
	var input struct {
		TargetID uint32 `json:"target_id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误"})
		return
	}
	// 创建聊天
	err, chat := createChat(userID, input.TargetID)
	if err != nil {
		log.Println("无法创建聊天", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误，无法创建聊天"})
		return
	}
	// 查询entity表和contact表
	var targetEntity model.Entity
	var targetContact model.Contact
	entityResult := repository.DB.Select("is_group", "avatar", "display_id", "name", "online_status", "status").Where("id = ?", input.TargetID).First(&targetEntity)
	if entityResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误，无法查询到用户"})
		return
	}
	contactResult := repository.DB.Select("remark").Where("user_id = ? AND target_id = ?", userID, input.TargetID).First(&targetContact)
	if contactResult.Error != nil && contactResult.Error != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
		return
	}
	// 假如不是联系人，检查是否允许陌生人聊天
	if contactResult.Error == gorm.ErrRecordNotFound {
		// 查询user_setting表
		var userSetting model.UserSetting
		result := repository.DB.Select("allow_stranger_message").Where("user_id = ?", userID).First(&userSetting)
		if result.Error != nil {
			log.Println("无法查询到用户设置", result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
			return
		}
		if !userSetting.AllowStrangerMessage {
			c.JSON(http.StatusBadRequest, gin.H{"message": "不允许陌生人聊天"})
			return
		}
	}
	// 返回结果
	c.JSON(http.StatusOK, gin.H{"data": gin.H{
		"ID":           chat.ID,
		"TargetID":     chat.TargetID,
		"IsGroup":      targetEntity.IsGroup,
		"Avatar":       targetEntity.Avatar,
		"DisplayID":    targetEntity.DisplayID,
		"Name":         targetEntity.Name,
		"Remark":       targetContact.Remark,
		"UnreadCount":  chat.UnreadCount,
		"LastMessage":  chat.LastMessage,
		"LastTime":     chat.LastTime.Format("2006-01-02 15:04:05"),
		"OnlineStatus": targetEntity.OnlineStatus,
		"Status":       targetEntity.Status,
		"IsPinned":     chat.IsPinned,
		"IsMuted":      chat.IsMuted,
		"IsBlocked":    chat.IsBlocked,
		"Tag1":         chat.Tag1,
		"Tag2":         chat.Tag2,
		"Tag3":         chat.Tag3,
	}})
}

func DeleteChatHandler(c *gin.Context) {
	// 获取参数
	// 获取user_id
	userID, ok := utils.GetUserID(c.Request)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "登录信息失效"})
		return
	}
	// 获取chat_id
	var input struct {
		ID uint `json:"chat_id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误"})
		return
	}
	// 查询chat表
	var dbChat model.Chat
	result := repository.DB.Where("id = ?", input.ID).First(&dbChat)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误，无法查询到聊天"})
		return
	}
	// 检查是否是用户的聊天
	if dbChat.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "无权删除该聊天"})
		return
	}
	// 删除聊天
	dbChat.IsDeleted = true
	result = repository.DB.Save(&dbChat)
	if result.Error != nil {
		log.Println("无法删除聊天", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误，无法删除聊天"})
		return
	}
	// 返回结果
	c.JSON(http.StatusOK, gin.H{"message": "删除聊天成功"})
}

func PinChatHandler(c *gin.Context) {
	// 获取参数
	// 获取user_id
	userID, ok := utils.GetUserID(c.Request)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "登录信息失效"})
		return
	}
	// 获取chat_id
	var input struct {
		ID       uint `json:"chat_id"`
		IsPinned bool `json:"is_pinned"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误"})
		return
	}
	// 查询chat表
	var dbChat model.Chat
	result := repository.DB.Where("id = ?", input.ID).First(&dbChat)
	if result.Error != nil {
		log.Println("无法查询到聊天", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误，无法查询到聊天"})
		return
	}
	// 检查是否是用户的聊天
	if dbChat.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "无权修改该聊天"})
		return
	}
	dbChat.IsPinned = input.IsPinned
	result = repository.DB.Save(&dbChat)
	if result.Error != nil {
		log.Println("无法修改聊天", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误，无法修改聊天"})
		return
	}
	// 返回结果
	c.JSON(http.StatusOK, gin.H{"message": "修改聊天成功"})
}

func MuteChatHandler(c *gin.Context) {
	// 获取参数
	// 获取user_id
	userID, ok := utils.GetUserID(c.Request)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "登录信息失效"})
		return
	}
	// 获取chat_id
	var input struct {
		ID      uint `json:"chat_id"`
		IsMuted bool `json:"is_muted"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误"})
		return
	}
	// 查询chat表
	var dbChat model.Chat
	result := repository.DB.Where("id = ?", input.ID).First(&dbChat)
	if result.Error != nil {
		log.Println("无法查询到聊天", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误，无法查询到聊天"})
		return
	}
	// 检查是否是用户的聊天
	if dbChat.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "无权修改该聊天"})
		return
	}
	dbChat.IsMuted = input.IsMuted
	result = repository.DB.Save(&dbChat)
	if result.Error != nil {
		log.Println("无法修改聊天", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误，无法修改聊天"})
		return
	}
	// 返回结果
	c.JSON(http.StatusOK, gin.H{"message": "修改聊天成功"})
}

func BlockChatHandler(c *gin.Context) {
	// 获取参数
	// 获取user_id
	userID, ok := utils.GetUserID(c.Request)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "登录信息失效"})
		return
	}
	// 获取chat_id
	var input struct {
		ID        uint `json:"chat_id"`
		IsBlocked bool `json:"is_blocked"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误"})
		return
	}
	// 查询chat表
	var dbChat model.Chat
	result := repository.DB.Where("id = ?", input.ID).First(&dbChat)
	if result.Error != nil {
		log.Println("无法查询到聊天", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误，无法查询到聊天"})
		return
	}
	// 检查是否是用户的聊天
	if dbChat.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "无权修改该聊天"})
		return
	}
	dbChat.IsBlocked = input.IsBlocked
	result = repository.DB.Save(&dbChat)
	if result.Error != nil {
		log.Println("无法修改聊天", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误，无法修改聊天"})
		return
	}
	// 返回结果
	c.JSON(http.StatusOK, gin.H{"message": "修改聊天成功"})
}

func ReadChatHandler(c *gin.Context) {
	// 获取参数
	// 获取user_id
	userID, ok := utils.GetUserID(c.Request)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "登录信息失效"})
		return
	}
	// 获取chat_id
	var input struct {
		ID     uint `json:"chat_id"`
		IsRead bool `json:"is_read"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误"})
		return
	}
	// 查询chat表
	var dbChat model.Chat
	result := repository.DB.Where("id = ?", input.ID).First(&dbChat)
	if result.Error != nil {
		log.Println("无法查询到聊天", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误，无法查询到聊天"})
		return
	}
	// 检查是否是用户的聊天
	if dbChat.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "无权修改该聊天"})
		return
	}
	if input.IsRead {
		dbChat.UnreadCount = 0
	} else {
		dbChat.UnreadCount = 1
	}
	result = repository.DB.Save(&dbChat)
	if result.Error != nil {
		log.Println("无法修改聊天", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误，无法修改聊天"})
		return
	}
	// 返回结果
	c.JSON(http.StatusOK, gin.H{"message": "修改聊天成功"})
}

func AddTagToChatHandler(c *gin.Context) {
	// 获取参数
	// 获取user_id
	userID, ok := utils.GetUserID(c.Request)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "登录信息失效"})
		return
	}
	// 获取chat_id
	var input struct {
		ID      uint   `json:"chat_id"`
		TagName string `json:"tag_name"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误"})
		return
	}
	// 查询chat表
	var dbChat model.Chat
	result := repository.DB.Where("id = ?", input.ID).First(&dbChat)
	if result.Error != nil {
		log.Println("无法查询到聊天", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误，无法查询到聊天"})
		return
	}
	// 判断用户是否有权限
	if dbChat.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "无权修改该聊天"})
		return
	}
	// 判断标签是否重复
	if dbChat.Tag1 == input.TagName || dbChat.Tag2 == input.TagName || dbChat.Tag3 == input.TagName {
		c.JSON(http.StatusBadRequest, gin.H{"message": "标签已存在"})
		return
	}
	if dbChat.Tag1 == "" {
		dbChat.Tag1 = input.TagName
	} else if dbChat.Tag2 == "" {
		dbChat.Tag2 = input.TagName
	} else if dbChat.Tag3 == "" {
		dbChat.Tag3 = input.TagName
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "标签已满"})
		return
	}
	result = repository.DB.Save(&dbChat)
	if result.Error != nil {
		log.Println("无法修改聊天", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误，无法修改聊天"})
		return
	}
	// 返回结果
	c.JSON(http.StatusOK, gin.H{"message": "修改聊天成功"})

}

func RemoveTagFromChatHandler(c *gin.Context) {
	// 获取参数
	// 获取user_id
	userID, ok := utils.GetUserID(c.Request)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "登录信息失效"})
		return
	}
	// 获取chat_id
	var input struct {
		ID      uint   `json:"chat_id"`
		TagName string `json:"tag_name"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误"})
		return
	}
	// 查询chat表
	var dbChat model.Chat
	result := repository.DB.Where("id = ?", input.ID).First(&dbChat)
	if result.Error != nil {
		log.Println("无法查询到聊天", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误，无法查询到聊天"})
		return
	}
	// 判断用户是否有权限
	if dbChat.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "无权修改该聊天"})
		return
	}
	if dbChat.Tag1 == input.TagName {
		dbChat.Tag1 = ""
	} else if dbChat.Tag2 == input.TagName {
		dbChat.Tag2 = ""
	} else if dbChat.Tag3 == input.TagName {
		dbChat.Tag3 = ""
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "标签不存在"})
		return
	}
	result = repository.DB.Save(&dbChat)
	if result.Error != nil {
		log.Println("无法修改聊天", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误，无法修改聊天"})
		return
	}
	// 返回结果
	c.JSON(http.StatusOK, gin.H{"message": "修改聊天成功"})
}
