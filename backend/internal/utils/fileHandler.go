package utils

import (
	"github.com/gin-gonic/gin"
)

var baseSavePath = "./"

// 处理文件上传
func HandleFileUpload(c *gin.Context, formName string, prefix string) (string, error) {
	// 获取文件
	file, err := c.FormFile(formName)
	if err != nil {
		return "", err
	}

	// 保存文件
	err = c.SaveUploadedFile(file, prefix+file.Filename)

	return file.Filename, err
}
