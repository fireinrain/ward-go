package router

import (
	"github.com/gin-gonic/gin"
	"ward-go/handler"
)

func InitRouter(app *gin.Engine) {
	// 设置一个get请求的路由，url为/ping, 处理函数（或者叫控制器函数）是一个闭包函数。
	app.GET("/ping", handler.PingHandler)
	app.GET("/setup", handler.SetUpPageHandler)
	app.GET("/", handler.IndexPageHandler)
	app.POST("/api/setup", handler.SetUpHandler)
	app.GET("/api/info", handler.InfoHandler)
	app.GET("/api/usage", handler.UsageHandler)
}
