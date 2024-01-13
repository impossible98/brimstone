package routes

import (
	// import third-party packages
	"github.com/gin-gonic/gin"

	// import local packages
	"brimstone/global"
)

const version = global.VERSION

// 处理GET /version请求
func VersionRouter(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code":    100,
		"version": version,
	})
}
