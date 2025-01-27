package main

import (
	"log"

	"Communication/config"
	"Communication/internal/api/router"
	"Communication/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	log.Println("初始化配置...")
	config.InitConfig()
	// 初始化数据库
	log.Println("初始化数据库...")
	repository.InitMysql()
	// 初始化路由
	log.Println("初始化路由...")
	r := gin.Default()
	router.SetupRouter(r)

	// 启动服务
	log.Println("服务器开始监听" + config.AppConfig.ServerPort)
	if err := r.Run(config.AppConfig.ServerPort); err != nil {
		log.Fatal("服务器启动失败: ", err)
	}
}
