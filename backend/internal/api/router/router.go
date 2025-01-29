package router

import (
	"Communication/internal/api/handler/auth"
	"Communication/internal/api/handler/chat"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {

	// 设置路由
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			// 用户认证
			authGroup := v1.Group("/auth")
			{
				authGroup.POST("/login", auth.LoginHandler)
				authGroup.POST("/register", auth.RegisterHandler)
			}
			// 聊天
			chatGroup := v1.Group("/chat")
			{
				listGroup := chatGroup.Group("/list")
				{
					listGroup.GET("", chat.GetChatListHandler)
				}
			}
		}

	}

}
