package main

import (
	"fmt"
	"gin-template/conf"
	"gin-template/db"
	"gin-template/docs"
	"gin-template/global"
	"gin-template/logger"
	"gin-template/routing"
	"gin-template/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	// 初始化日志
	logger.Init()

	// 加载 .env 文件
	if err := godotenv.Load(); err != nil {
		logger.Log.Fatal("No .env file found, continue with system environment")
	}

	db.Initialize()
	db.ConnectDB()

	global.Log = utils.InitZap() // 初始化zap日志库

	router := gin.Default()
	router.Use(gin.Logger())
	docs.SwaggerInfo.BasePath = "/api"
	routing.Setup(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	err := router.Run(fmt.Sprintf(":%s", conf.Config.GinPort))
	if err != nil {
		return
	}
	logger.Log.Infoln("Start successfully")
}
