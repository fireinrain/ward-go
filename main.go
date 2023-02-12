package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	app := gin.Default()
	gin.SetMode(gin.DebugMode)

	//模板处理
	app.LoadHTMLGlob("website/templates/*/*")
	app.Static("/static", "website/static")
	app.GET("/", func(c *gin.Context) {
		// 子目录的模版文件，需要加上目录名，例如：posts/index.tmpl
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Posts",
		})
	})
	app.GET("/404", func(c *gin.Context) {
		// 子目录的模版文件，需要加上目录名，例如：users/index.tmpl
		c.HTML(http.StatusOK, "404.html", gin.H{
			"title": "Users",
		})
	})

	_ = app.Run(":8080")
}
