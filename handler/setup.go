package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ward-go/config"
	"ward-go/utils"
)

// SelfServerSet
//
//	SelfServerSet
//	@Description: 配置
type SelfServerSet struct {
	ServerName string `json:"serverName"`
	Theme      string `json:"theme"`
	Port       string `json:"port"`
}

// SetUpHandler
//
//	@Description: setup api
//	@param c
func SetUpHandler(c *gin.Context) {
	set := SelfServerSet{}
	if c.ShouldBindJSON(&set) != nil {
		serverName := set.ServerName
		if serverName == "" {
			c.JSON(http.StatusOK, gin.H{
				"message": "Server name is required",
				"data":    "",
			})
			return
		}
		theme := set.Theme
		if theme == "" {
			c.JSON(http.StatusOK, gin.H{
				"message": "theme option is required",
				"data":    "",
			})
			return
		}
		port := set.Port
		if port == "" {
			c.JSON(http.StatusOK, gin.H{
				"message": "port option is required",
				"data":    "",
			})
			return
		}
	}
	//写入配置文件
	utils.WriteConfig2File(set.ServerName, set.Theme, set.Port, config.ConfigFile)

	c.JSON(http.StatusOK, gin.H{
		"message": "setting saved correctly",
		"data":    "",
	})
	//重新启动web
	utils.GraceStopGin(config.AppServer)

}

// SetUpPageHandler
//
//	@Description: 模板也
//	@param c
func SetUpPageHandler(c *gin.Context) {
	if config.FirstStartUp {
		c.HTML(http.StatusOK, "setup.html", gin.H{
			"title": "Setup ward-go",
		})
	} else {
		//404
		c.HTML(http.StatusOK, "404.html", gin.H{
			"title": "404 page",
		})
	}
}