package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
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
	config.WriteConfig2File(set.ServerName, set.Theme, set.Port, config.ConfigFile)
	config.RefreshServerConfig()

	c.JSON(http.StatusOK, gin.H{
		"message": "setting saved correctly",
		"data":    "",
	})
	//启动新ginServer
	go utils.StartGinServer(config.Wg, InitRouter)
	//重新启动web
	go GraceStopGin(config.AppServer)

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

// GraceStopGin
//
//	@Description: 关闭Gin
//	@param srv
func GraceStopGin(srv *http.Server) {
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
loop:
	for {
		select {
		case <-ctx.Done():
			log.Println("timeout of 5 seconds.")
			break loop
		default:
			time.Sleep(1 * time.Second)
		}
	}

	log.Println("Server exiting")
}
