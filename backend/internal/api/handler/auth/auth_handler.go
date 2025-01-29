package auth

import (
	"Communication/internal/repository"
	"Communication/internal/repository/model"
	"Communication/internal/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 登录
func LoginHandler(c *gin.Context) {
	// 参数绑定
	var user struct {
		DisplayID string `json:"display_id" binding:"required"`
		Password  string `json:"password" binding:"required"`
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
	// 检查是否有其他session是同一个user_id  todo
	// if session, ok := utils.GetUserIDSession(dbEntity.ID); ok {
	// 	// 删除旧session todo
	// 	log.Println("删除旧session：", session)
	// 	utils.DeleteSession(c.Request, c.Writer, session, dbEntity.ID)
	// }
	// 生成用户session，写入响应中
	if err := utils.SetUserSession(c.Request, c.Writer, dbEntity.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
		log.Println("写入session失败", err)
		return
	}

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

	// 登录成功
	c.JSON(http.StatusOK, gin.H{"data": dbUserSetting})
}

// 注册
func RegisterHandler(c *gin.Context) {
	// 参数绑定
	var input struct {
		DisplayID string `json:"display_id"`
		Name      string `json:"name" binding:"required"`
		Password  string `json:"password" binding:"required"`
		Email     string `json:"email" binding:"required"`
		Avatar    string `json:"avatar"`
		Sex       uint8  `json:"sex" binding:"required"`
		Birthday  string `json:"birthday"` // 格式："2006-01-02"
		Signature string `json:"signature"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误"})
		return
	}
	// 检查birthday格式
	var date time.Time
	if input.Birthday != "" {
		d, err := time.Parse("2006-01-02", input.Birthday)
		if err != nil {
			log.Println("日期参数错误：", err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "日期参数错误"})
			return
		}
		date = d
	}

	// 假如有display_id，查询字段，检验账号是否已存在
	var dbEntity model.Entity
	if input.DisplayID != "" && (input.DisplayID == "undefined" || len(input.DisplayID) < 6 || len(input.DisplayID) > 14) {
		return
	}
	if input.DisplayID != "" {
		// 查询entity表
		result := repository.DB.Select("id").Where("display_id = ?", input.DisplayID).First(&dbEntity)
		if result.Error == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "账号已存在"})
			return
		}
	}

	// 注册
	// 获取当前时间
	currentTime := time.Now()
	// 开始事务
	tx := repository.DB.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
		return
	}
	// entity表
	// 自动生成display_id
	if input.DisplayID == "" {
		for {
			displayID := generateRandomString(8)
			if isDisplayIDUnique(displayID) {
				input.DisplayID = displayID
				break
			}
		}
	}
	dbEntity.DisplayID = input.DisplayID
	dbEntity.IsGroup = false
	dbEntity.Name = input.Name
	if input.Avatar == "" {
		// dbEntity.Avatar = "default_avatar.jpg"  // todo
	} else {
		dbEntity.Avatar = input.Avatar
	}
	dbEntity.Introduction = input.Signature
	// 插入
	if err := tx.Create(&dbEntity).Error; err != nil {
		tx.Rollback()
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
		return
	}
	// user_setting表
	var dbUserSetting model.UserSetting
	dbUserSetting.UserID = dbEntity.ID
	dbUserSetting.Password = input.Password
	dbUserSetting.Email = input.Email
	if input.Birthday != "" {
		dbUserSetting.Birthday = date
	}
	dbUserSetting.Sex = input.Sex
	dbUserSetting.CreateTime = currentTime
	// 插入
	if err := tx.Create(&dbUserSetting).Error; err != nil {
		tx.Rollback()
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
		return
	}
	// divide_friend
	var dbDivideFriend model.DivideFriend
	dbDivideFriend.UserID = dbEntity.ID
	dbDivideFriend.Name = "未分组"
	// 插入
	if err := tx.Create(&dbDivideFriend).Error; err != nil {
		tx.Rollback()
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
		return
	}
	// divide_group
	var dbDivideGroup model.DivideGroup
	dbDivideGroup.UserID = dbEntity.ID
	dbDivideGroup.Name = "未分组"
	// 插入
	if err := tx.Create(&dbDivideGroup).Error; err != nil {
		tx.Rollback()
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
		return
	}

	// 提交事务
	tx.Commit()

	// 注册成功
	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})

}
