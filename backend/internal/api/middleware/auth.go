package middleware

import (
	"Communication/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取user_id
		userID, ok := utils.GetUserID(c.Request)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "登录信息失效"})
			c.Abort()
			return
		}
		c.Set("user_id", userID)
		c.Next()
	}
}
