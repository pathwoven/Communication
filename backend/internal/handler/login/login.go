package login

import (
	"Communication/internal/repository"
	"Communication/internal/repository/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 登录
// 操作：
//
//	成功则返回用户设置信息，且生成session
//	同时根据user_setting的is_invisible修改entity的online_status
//	修改user_setting的last_login_time
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

	// 查询字段，检验账号密码
	// 查询entity表
	var dbEntity model.Entity
	result := repository.DB.Select("id").Where("display_id = ?", user.DisplayID).First(&dbEntity)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "账号或密码错误"})
		return
	}
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

	// 登录成功
	c.JSON(http.StatusOK, gin.H{"data": dbUserSetting})
}
