package main

import (
	// import built-in packages
	"fmt"
	"net/http"
	"time"

	// import third-party packages
	"github.com/gin-gonic/gin"

	// import local packages
	"brimstone/internal/utils"
	"brimstone/pkg/config"
	"brimstone/pkg/middleware"
	"brimstone/routes"
)

func main() {
	currentTime := utils.GetTime()
	config.Init()

	host := config.Conf.Host
	port := config.Conf.Port

	// 创建一个gin router
	var router = gin.New()
	router.Use(middleware.Cors())

	// 没有路由即 404返回
	router.NoRoute(func(g *gin.Context) {
		g.JSON(404, gin.H{"code": 404, "msg": fmt.Sprintf("not found '%s:%s'", g.Request.Method, g.Request.URL.Path)})
	})
	// 注册路由处理函数
	router.GET("/ping", routes.PingRouter)
	router.GET("/version", routes.VersionRouter)

	utils.ShowTip(currentTime, port)
	// 运行服务器，监听指定端口
	s := &http.Server{
		Addr:           host + ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
