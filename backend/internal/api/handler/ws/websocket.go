package ws

import (
	"Communication/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WebSocketHandler(c *gin.Context) {
	// 获取user_id
	userID, ok := utils.GetUserID(c.Request)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "登录信息失效"})
		return
	}
	// 打开websocket连接
	err, conn := utils.UpgradeWebsocket(userID, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
		return
	}
	defer utils.CloseWebsocket(userID)
	// 保持连接
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}
}
