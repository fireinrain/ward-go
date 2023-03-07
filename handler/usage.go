package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ward-go/service"
)

// UsageHandler
//
//	@Description: usage rest api
//	@param c
func UsageHandler(c *gin.Context) {
	usageInfo := service.GetUsageInfoService()
	c.JSON(http.StatusOK, usageInfo)
}
