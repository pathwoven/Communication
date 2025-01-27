package router

import (
	"Communication/internal/api/handler/auth"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {

	// 设置路由
	api := r.Group("/auth")
	{
		api.POST("/login", auth.LoginHandler)
	}

}
