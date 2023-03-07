package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Error404PageHandler
//
//	@Description: 404页面
//	@param c
func Error404PageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "404.html", gin.H{
		"title": "404 error page",
		"theme": "light",
	})
}

// Error500PageHandler
//
//	@Description: 500 页面
//	@param c
func Error500PageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "500.html", gin.H{
		"title": "500 error page",
		"theme": "light",
	})
}
