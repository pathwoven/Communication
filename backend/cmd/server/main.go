package main

import (
	"Communication/config"
	"Communication/internal/api/router"
	"Communication/internal/repository"
	"Communication/internal/utils"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// utils.Test() // todo debug
	// 初始化配置
	log.Println("初始化配置...")
	config.InitConfig()
	// 初始化数据库
	log.Println("初始化数据库...")
	repository.InitMysql()
	// 初始化路由
	log.Println("初始化路由...")
	r := gin.Default()
	r.Use(config.Cors())
	router.SetupRouter(r)
	// 初始化session store
	utils.InitSessionStore()

	// 启动服务
	log.Println("服务器开始监听" + config.AppConfig.ServerPort)
	if err := r.Run(":" + config.AppConfig.ServerPort); err != nil {
		log.Fatal("服务器启动失败: ", err)
	}
}
