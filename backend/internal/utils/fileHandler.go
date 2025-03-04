package utils

import (
	"Communication/internal/repository"
	"Communication/internal/repository/model"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

var baseSavePath = "./file"

// 处理文件上传，返回在file表中的id和错误
func HandleFileUpload(c *gin.Context, formName string, prefix string) (uint32, error) {
	// 获取文件
	file, err := c.FormFile(formName)
	if err != nil {
		return 0, err
	}
	var dbFile model.File
	dbFile.Name = file.Filename
	dbFile.Size = file.Size

	// 事务开始
	tx := repository.DB.Begin()
	if tx.Error != nil {
		return 0, tx.Error
	}
	// 插入文件信息
	if result := tx.Create(&dbFile); result.Error != nil {
		tx.Rollback()
		return 0, result.Error
	}
	// 保存路径
	savePath := filepath.Join(baseSavePath, prefix, strconv.Itoa(int(dbFile.ID)))
	// 保存文件
	if err = c.SaveUploadedFile(file, savePath); err != nil {
		tx.Rollback()
		return 0, err
	}
	dbFile.Path = savePath
	if result := tx.Save(&dbFile); result.Error != nil {
		tx.Rollback()
		return 0, result.Error
	}
	// 事务提交
	tx.Commit()

	return dbFile.ID, err
}

func HandleFileDownload(c *gin.Context, id uint32) error {
	var dbFile model.File
	result := repository.DB.Where("id = ?", id).First(&dbFile)
	if result.Error != nil {
		return result.Error
	}
	c.Header("Content-Disposition", "inline; filename="+dbFile.Name)
	c.Header("Content-Type", "application/octet-stream")
	c.File(dbFile.Path)

	return nil
}
