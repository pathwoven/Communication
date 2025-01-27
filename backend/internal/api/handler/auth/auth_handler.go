package auth

import (
	"Communication/internal/repository"
	"Communication/internal/repository/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 登录
func LoginHandler(c *gin.Context) {
	// 参数绑定
	var user struct {
		DisplayID string `json:"display_id"`
		Password  string `json:"password"`
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误"})
		return
	}

	// 查询字段，检验账号密码，及是否已登录
	// 查询entity表
	var dbEntity model.Entity
	result := repository.DB.Select("id", "online_status").Where("display_id = ?", user.DisplayID).First(&dbEntity)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "账号或密码错误"})
		return
	}
	// 检查是否已登录
	// todo
	// 查询user_setting表
	var dbUserSetting model.UserSetting
	result = repository.DB.Where("user_id = ?", dbEntity.ID).First(&dbUserSetting)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "账号或密码错误"})
		return
	}
	// 验证密码
	if dbUserSetting.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "账号或密码错误"})
		return
	}

	// 账号密码正确，修改登录信息
	// 获取当前时间
	currentTime := time.Now()
	dbUserSetting.LastLoginTime = currentTime
	// 开始事务
	tx := repository.DB.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
		return
	}
	// 修改entity表
	if dbUserSetting.IsInvisible {
		result = tx.Model(&dbEntity).Update("online_status", 2)
	} else {
		result = tx.Model(&dbEntity).Update("online_status", 1)
	}
	if result.Error != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
		return
	}
	// 修改user_setting表
	result = tx.Model(&dbUserSetting).Update("last_login_time", currentTime)
	if result.Error != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
		return
	}
	// 提交事务
	tx.Commit()

	// 生成session

	// 登录成功
	c.JSON(http.StatusOK, gin.H{"data": dbUserSetting})
}
