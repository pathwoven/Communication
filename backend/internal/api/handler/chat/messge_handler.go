package chat

import (
	"Communication/internal/repository"
	"Communication/internal/repository/model"
	"Communication/internal/utils"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 获取指定条数消息（todo 重写群聊部分及删除部分）
func GetMessagesHandler(c *gin.Context) {
	// 获取参数
	// 获取user_id
	userID, ok := utils.GetUserID(c.Request)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "登录信息失效"})
		return
	}
	// 绑定参数
	var input struct { // note: 注意，go会把零值当作没有传值，所以验证的时候要小心
		TargetID uint32 `json:"target_id" binding:"required"`
		Start    uint32 `json:"start"`
		Count    uint32 `json:"count" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("参数错误", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误"})
		return
	}
	// 查询数据库
	// 查询message表
	var dbMessages []model.Message
	result := repository.DB.
		Where("((sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)) AND (operation is null OR (operation != 0 AND operation != 1 AND operation != 2 AND operation != 3))",
			userID, input.TargetID, input.TargetID, userID).
		Order("create_time asc").Offset(int(input.Start)).Limit(int(input.Count)).Find(&dbMessages)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		log.Println("无法获取message", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
		return
	}
	// 返回数据
	c.JSON(http.StatusOK, gin.H{"data": dbMessages})
}

// 发送单人消息
func SendSingleMessageHandler(c *gin.Context) {
	// 获取参数
	// 获取user_id
	userID, ok := utils.GetUserID(c.Request)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "登录信息失效"})
		return
	}
	// 绑定参数
	var input struct {
		ReceiverID  uint32  `json:"receiver_id" binding:"required"`
		Content     string  `json:"content" binding:"required"`
		ContentType *uint8  `json:"content_type" binding:"required"`
		Operation   *uint8  `json:"operation"`
		OpMessageID *uint32 `json:"op_message_id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("参数错误: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误"})
		return
	}
	// 修改数据库
	// 开始事务
	tx := repository.DB.Begin()
	// 查找contact表，确定是否是好友，是否被拉黑，获取备注
	var dbContact model.Contact
	result := tx.Select("is_blacklisted", "remark").Where("user_id = ? AND target_id = ?", input.ReceiverID, userID).First(&dbContact)
	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			log.Println("无法获取contact", result.Error)
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
			return
		} else {
			// 不是好友
			// 查找user_setting表，确定是否允许陌生人私信
			var dbUserSetting model.UserSetting
			result = tx.Select("allow_stranger_message").Where("user_id = ?", input.ReceiverID).First(&dbUserSetting)
			if result.Error != nil {
				log.Println("无法获取user_setting", result.Error)
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
				return
			}
			if !dbUserSetting.AllowStrangerMessage {
				tx.Rollback()
				c.JSON(http.StatusBadRequest, gin.H{"message": "对方不允许陌生人私信"})
				return
			}
		}
	}
	if dbContact.IsBlacklisted {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"message": "对方已将您拉黑"})
		return
	}
	// 保存至message表
	var message model.Message
	message.SenderID = userID
	message.ReceiverID = input.ReceiverID
	message.Content = input.Content
	message.ContentType = *input.ContentType
	message.CreateTime = time.Now()
	if input.Operation != nil {
		message.Operation = *input.Operation
	}
	if input.OpMessageID != nil {
		message.OpMessageID = *input.OpMessageID
	}
	result = tx.Create(&message)
	if result.Error != nil {
		log.Println("无法保存message", result.Error)
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
		return
	}
	// 更新chat表
	var lastMessage string
	if message.ContentType == 1 {
		lastMessage = "[图片]"
	} else if message.ContentType == 2 {
		lastMessage = "[文件]"
	} else if message.ContentType == 3 {
		lastMessage = "[群投票]"
	} else if message.ContentType == 4 {
		lastMessage = "[聊天记录]"
	} else {
		lastMessage = message.Content
	}
	// 查询发送者的chat表
	var dbChat model.Chat
	result = tx.Where("user_id = ? AND target_id = ?", userID, input.ReceiverID).First(&dbChat)
	if result.Error != nil {
		log.Println("无法获取chat", result.Error)
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
		return
	}
	if dbChat.IsDeleted {
		log.Println("chat已被删除")
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"message": "聊天已被删除"})
		return
	}
	dbChat.LastTime = &message.CreateTime
	dbChat.LastMessage = lastMessage
	dbChat.LastPerson = ""
	result = tx.Save(&dbChat)
	if result.Error != nil {
		log.Println("无法更新chat", result.Error)
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
		return
	}
	// 更新对方的chat表
	isNewChat := false
	dbChat = model.Chat{}
	// 查询chat表
	result = tx.Where("user_id = ? AND target_id = ?", input.ReceiverID, userID).First(&dbChat)
	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			log.Println("无法获取chat", result.Error)
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
			return
		} else {
			// 创建chat
			isNewChat = true
			dbChat.UserID = input.ReceiverID
			dbChat.TargetID = userID
			result = tx.Create(&dbChat)
			if result.Error != nil {
				log.Println("无法创建chat", result.Error)
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
				return
			}
		}
	}
	// 如果chat被删除了
	if dbChat.IsDeleted {
		isNewChat = true
		dbChat.IsDeleted = false
	}
	dbChat.LastTime = &message.CreateTime
	dbChat.LastMessage = lastMessage
	// 填写last_person（对方对我的备注、昵称）
	var last_person string
	if dbContact.Remark != "" {
		last_person = dbContact.Remark
	} else {
		// 查询昵称 查询entity表
		var dbEntity model.Entity
		result = tx.Select("name, online_status").Where("id = ?", userID).First(&dbEntity)
		if result.Error != nil {
			log.Println("无法获取entity", result.Error)
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
			return
		}
		last_person = dbEntity.Name
	}
	dbChat.LastPerson = last_person
	// 未读消息条数
	if dbChat.UnreadCount < 100 {
		dbChat.UnreadCount++
	}
	result = tx.Save(&dbChat)
	if result.Error != nil {
		log.Println("无法更新chat", result.Error)
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
		return
	}
	// 提交事务
	tx.Commit()
	// 返回数据
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"id": message.ID, "create_time": message.CreateTime}})
	// websocket推送(message或chatAndMessage)
	if isNewChat {
		// todo
	} else {
		var newData = map[string]interface{}{
			"type": "message",
			"data": map[string]interface{}{
				"message": message,
				"chat": map[string]interface{}{
					"last_message": lastMessage,
					"last_person":  last_person,
					"target_id":    userID,
				},
			},
		}
		data, err := json.Marshal(newData)
		if err != nil {
			log.Println("无法序列化message", err)
			return
		}
		if err = utils.WsSendMessage(input.ReceiverID, data); err != nil {
			log.Println("无法发送消息", err)
		}
	}

}

// 发送群消息
func SendGroupMessageHandler(c *gin.Context) {
	// 获取参数
	// 获取user_id
	userID, ok := utils.GetUserID(c.Request)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "登录信息失效"})
		return
	}
	// 绑定参数
	var input struct {
		GroupID     uint32  `json:"group_id" binding:"required"`
		Content     string  `json:"content" binding:"required"`
		ContentType uint8   `json:"content_type" binding:"required"`
		Operation   *uint8  `json:"operation"`
		OpMessageID *uint32 `json:"op_message_id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误"})
		return
	}
	// 修改数据库
	// 开始事务
	tx := repository.DB.Begin()
	// 查找group_member_info表，获取所有成员的id,群昵称，是否被禁言
	var dbGroupMemberInfo []model.GroupMemberInfo
	result := tx.Select("user_id, is_banned, group_nickname").Where("group_id = ?", input.GroupID).Find(&dbGroupMemberInfo)
	if result.Error != nil {
		log.Println("无法获取group_member", result.Error)
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
		return
	}
	var senderGroupNickname string
	isInGroup := false
	for _, member := range dbGroupMemberInfo {
		if member.UserID == userID {
			senderGroupNickname = member.GroupNickname
			if member.IsBanned {
				tx.Rollback()
				c.JSON(http.StatusBadRequest, gin.H{"message": "您已被禁言"})
				return
			}
			isInGroup = true
			break
		}
	}
	if !isInGroup {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"message": "您不在该群中"})
		return
	}
	// 保存至message表
	var message model.Message
	message.SenderID = userID
	message.ReceiverID = input.GroupID
	message.Content = input.Content
	message.ContentType = input.ContentType
	message.CreateTime = time.Now()
	if input.Operation != nil {
		message.Operation = *input.Operation
	}
	if input.OpMessageID != nil {
		message.OpMessageID = *input.OpMessageID
	}
	result = tx.Create(&message)
	if result.Error != nil {
		log.Println("无法保存message", result.Error)
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
		return
	}
	// 更新chat表
	var lastMessage string
	if message.ContentType == 1 {
		lastMessage = "[图片]"
	} else if message.ContentType == 2 {
		lastMessage = "[文件]"
	} else if message.ContentType == 3 {
		lastMessage = "[群投票]"
	} else if message.ContentType == 4 {
		lastMessage = "[聊天记录]"
	} else {
		lastMessage = message.Content
	}
	// 查询发送者的chat表
	var dbChat model.Chat
	result = tx.Where("user_id = ? AND target_id = ?", userID, input.GroupID).First(&dbChat)
	if result.Error != nil {
		log.Println("无法获取chat", result.Error)
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
		return
	}
	if dbChat.IsDeleted {
		log.Println("chat已被删除")
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"message": "聊天已被删除"})
		return
	}
	dbChat.LastTime = &message.CreateTime
	dbChat.LastMessage = lastMessage
	dbChat.LastPerson = ""
	result = tx.Save(&dbChat)
	if result.Error != nil {
		log.Println("无法更新chat", result.Error)
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
		return
	}
	// 返回消息
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"id": message.ID, "create_time": message.CreateTime}})
	var dbEntity model.Entity
	var senderName string
	// 查询entity表
	result = tx.Select("name").Where("id = ?", userID).First(&dbEntity)
	if result.Error != nil {
		log.Println("无法获取entity", result.Error)
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
		return
	}
	// 更新所有成员的chat表
	for _, member := range dbGroupMemberInfo {
		dbChat = model.Chat{}
		if member.UserID == userID {
			continue
		}
		// 查找contact表
		var dbContact model.Contact
		result = tx.Select("remark").Where("user_id = ? AND target_id = ?", member.UserID, userID).First(&dbContact)
		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			log.Println("无法获取contact", result.Error)
			continue

		}
		// 查找entity表，看是否在线（todo）
		result = tx.Select("online_status").Where("id = ?", member.UserID).First(&dbEntity)
		if result.Error != nil {
			log.Println("无法获取entity", result.Error)
			continue
		}
		isNewChat := false
		// 查询chat表
		result = tx.Where("user_id = ? AND target_id = ?", member.UserID, input.GroupID).First(&dbChat)
		if result.Error != nil {
			if result.Error != gorm.ErrRecordNotFound {
				log.Println("无法获取chat", result.Error)
				continue
			} else {
				isNewChat = true
				// 创建chat
				dbChat.UserID = member.UserID
				dbChat.TargetID = input.GroupID
				result = tx.Create(&dbChat)
				if result.Error != nil {
					log.Println("无法创建chat", result.Error)
					continue
				}
			}
		}
		// 如果chat被删除了
		if dbChat.IsDeleted {
			isNewChat = true
			dbChat.IsDeleted = false
		}
		dbChat.LastTime = &message.CreateTime
		dbChat.LastMessage = lastMessage
		// 填写last_person（对方对我的备注、昵称）
		var last_person string
		if dbContact.Remark != "" {
			last_person = dbContact.Remark
		} else if senderGroupNickname != "" {
			last_person = senderGroupNickname
		} else {
			last_person = senderName
		}
		dbChat.LastPerson = last_person
		// 未读消息条数
		if dbChat.UnreadCount < 100 {
			dbChat.UnreadCount++
		}
		result = tx.Save(&dbChat)
		if result.Error != nil {
			log.Println("无法更新chat", result.Error)
			continue
		}
		// websocket推送(message或chatAndMessage)
		if isNewChat {
			// todo
		} else {

		}
	}
}
