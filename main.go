package main

import (
	// import built-in packages
	"log"

	// import third-party packages
	"github.com/dotenv-org/godotenvvault"
	"github.com/gin-gonic/gin"

	// import local packages
	"brimstone/internal/utils"
)

func main() {
	currentTime := utils.GetTime()
	// 加载环境变量
	err := godotenvvault.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// 设置端口
	port := utils.GetPort()

	// 创建一个gin router
	r := gin.Default()

	// 注册路由处理函数
	r.GET("/ping", handlePing)

	utils.ShowTip(currentTime)
	// 运行服务器，监听指定端口
	r.Run(":" + port)
}

// 处理GET /ping请求
func handlePing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
