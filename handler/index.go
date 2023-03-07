package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ward-go/config"
)

// IndexPageHandler
//
//	@Description: 默认页面
//	@param c
func IndexPageHandler(c *gin.Context) {
	if config.FirstStartUp {
		c.HTML(http.StatusOK, "setup.html", gin.H{
			"title": "Welcome to ward-go setup page",
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"theme": config.GlobalConfig.Setup.Theme,
			"title": config.GlobalConfig.Setup.ServerName,
		})
	}

}

// PingHandler
//
//	@Description: ping
//	@param c
func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
