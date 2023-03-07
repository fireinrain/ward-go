package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ServerUsageInfo
// @Description: 服务器资源使用信息
type ServerUsageInfo struct {
	Processor int `json:"processor"`
	Ram       int `json:"ram"`
	Storage   int `json:"storage"`
}

func UsageHandler(c *gin.Context) {
	usageInfo := ServerUsageInfo{
		Processor: 100,
		Ram:       100,
		Storage:   100,
	}
	c.JSON(http.StatusOK, usageInfo)
}
