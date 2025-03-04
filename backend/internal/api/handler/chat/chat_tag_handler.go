package chat

import (
	"Communication/internal/repository"
	"Communication/internal/repository/model"
	"Communication/internal/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTagListHandler(c *gin.Context) {
	// 获取参数
	// 获取user_id
	userID, ok := utils.GetUserID(c.Request)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "登录信息失效"})
		return
	}
	// 查询tag表
	var dbTags []model.ChatTag
	result := repository.DB.Where("user_id = ?", userID).Find(&dbTags)
	if result.Error != nil {
		log.Println("无法获取tag", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
		return
	}
	// 提取tag_name组成的数组
	var tagNames []string
	for _, tag := range dbTags {
		tagNames = append(tagNames, tag.Name)
	}
	// 返回数据
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"tags": tagNames}})

}

func CreateTagHandler(c *gin.Context) {
	// 获取参数
	// 获取user_id
	userID, ok := utils.GetUserID(c.Request)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "登录信息失效"})
		return
	}
	// 获取tag_name
	var input struct {
		TagName string `json:"tag_name"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误"})
		return
	}
	// 查询tag表
	var dbTag model.ChatTag
	result := repository.DB.Where("user_id = ? AND name = ?", userID, input.TagName).First(&dbTag)
	if result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "标签已存在"})
		return
	}
	// 创建tag
	dbTag = model.ChatTag{
		UserID: userID,
		Name:   input.TagName,
	}
	result = repository.DB.Create(&dbTag)
	if result.Error != nil {
		log.Println("无法创建tag", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
		return
	}
	// 返回数据
	c.JSON(http.StatusOK, gin.H{"message": "创建成功"})
}

func DeleteTagHandler(c *gin.Context) {
	// 获取参数
	// 获取user_id
	userID, ok := utils.GetUserID(c.Request)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "登录信息失效"})
		return
	}
	// 获取tag_name
	tagName := c.Param("tag_name")
	// 查询tag表
	var dbTag model.ChatTag
	result := repository.DB.Where("user_id = ? AND name = ?", userID, tagName).First(&dbTag)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "标签不存在"})
		return
	}
	// 删除tag
	result = repository.DB.Delete(&dbTag)
	if result.Error != nil {
		log.Println("无法删除tag", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
		return
	}
	// 返回数据
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

func RenameTagHandler(c *gin.Context) {
	// 获取参数
	// 获取user_id
	userID, ok := utils.GetUserID(c.Request)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "登录信息失效"})
		return
	}
	var input struct {
		OldTagName string `json:"old_tag_name"`
		NewTagName string `json:"new_tag_name"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误"})
		return
	}
	// 查询tag表
	var dbTag model.ChatTag
	result := repository.DB.Where("user_id = ? AND name = ?", userID, input.OldTagName).First(&dbTag)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "标签不存在"})
		return
	}
	// 更新tag
	result = repository.DB.Model(&dbTag).Update("name", input.NewTagName)
	if result.Error != nil {
		log.Println("无法更新tag", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
		return
	}
	// 返回数据
	c.JSON(http.StatusOK, gin.H{"message": "修改成功"})
}
