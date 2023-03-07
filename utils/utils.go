package utils

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
	"time"
	"ward-go/config"
)

// StartGinServer
//
//	@Description: 启动Gin服务器
func StartGinServer(wg *sync.WaitGroup, routerFunc func(engine *gin.Engine)) {
	wg.Add(1)
	// 初始化一个http服务对象
	app := gin.Default()
	gin.SetMode(gin.ReleaseMode)

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
