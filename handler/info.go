package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ward-go/service"
)

func InfoHandler(c *gin.Context) {
	//serverInfo := service.GetServerInfoService()
	serverInfo := service.GetServerInfo()
	c.JSON(http.StatusOK, serverInfo)
}
