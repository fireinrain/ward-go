package handler

import "github.com/gin-gonic/gin"

func InitRouter(app *gin.Engine) {
	// 设置一个get请求的路由，url为/ping, 处理函数（或者叫控制器函数）是一个闭包函数。
	app.GET("/ping", PingHandler)
	app.GET("/setup", SetUpPageHandler)
	app.GET("/", IndexPageHandler)
	app.POST("/api/setup", SetUpHandler)
	app.GET("/api/info", InfoHandler)
	app.GET("/api/usage", UsageHandler)
}
