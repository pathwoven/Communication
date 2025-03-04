package router

import (
	"Communication/internal/api/handler/auth"
	"Communication/internal/api/handler/chat"
	"Communication/internal/api/handler/ws"
	"Communication/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {

	// 设置路由
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			// websocket
			v1.GET("/ws", ws.WebSocketHandler)
			// 用户认证
			authGroup := v1.Group("/auth")
			{
				authGroup.POST("/login", auth.LoginHandler)
				authGroup.POST("/register", auth.RegisterHandler)
			}
			// 聊天
			chatGroup := v1.Group("/chat")
			{
				chatGroup.Use(middleware.AuthMiddleWare())
				listGroup := chatGroup.Group("/list")
				{
					listGroup.GET("", chat.GetChatListHandler)
					listGroup.POST("/create", chat.CreateChatHandler)
					listGroup.POST("/delete", chat.DeleteChatHandler)
					listGroup.POST("/pin", chat.PinChatHandler)
					listGroup.POST("/mute", chat.MuteChatHandler)
					listGroup.POST("/block", chat.BlockChatHandler)
					listGroup.POST("/read", chat.ReadChatHandler)
					listGroup.POST("/add-tag", chat.AddTagToChatHandler)
					listGroup.POST("/remove-tag", chat.RemoveTagFromChatHandler)
				}
				tagGroup := chatGroup.Group("/tag")
				{
					tagGroup.GET("/list", chat.GetTagListHandler)
					tagGroup.POST("/create", chat.CreateTagHandler)
					tagGroup.POST("/delete", chat.DeleteTagHandler)
					tagGroup.POST("/rename", chat.RenameTagHandler)
				}
				messageGroup := chatGroup.Group("/message")
				{
					messageGroup.POST("/list", chat.GetMessagesHandler)
					messageGroup.POST("/single/send", chat.SendSingleMessageHandler)
					messageGroup.POST("/group/send", chat.SendGroupMessageHandler)
					messageGroup.POST("/download/file", chat.DownloadMessageFileHandler)
				}
			}
		}

	}

}
