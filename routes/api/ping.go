package api

import (
	// import third-party packages
	"github.com/gin-gonic/gin"
)

// 处理GET /ping请求
func PingRouter(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 100,
		"msg":  "pong",
	})
}
