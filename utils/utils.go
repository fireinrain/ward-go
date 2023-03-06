package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
	"ward-go/config"
)

// StartGinServer
//
//	@Description: 启动Gin服务器
func StartGinServer(wg *sync.WaitGroup, routerFunc func(engine *gin.Engine)) {
	wg.Add(1)
	// 初始化一个http服务对象
	app := gin.Default()

	// 首先加载templates目录下面的所有模版文件，模版文件扩展名随意
	app.LoadHTMLGlob("website/templates/*")
	app.Static("/static", "website/static")
	routerFunc(app)
	//router.InitRouter(app)

	// 监听并在 0.0.0.0:8888 上启动服务
	addr := fmt.Sprintf(":%d", config.GlobalConfig.Setup.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: app,
	}
	go func() {
		defer wg.Done()
		// service connections
		log.Println("Server start at： ", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	config.AppServer = srv
}
