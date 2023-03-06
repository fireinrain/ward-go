package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"ward-go/config"
	"ward-go/router"
)

// StartGinServer
//
//	@Description: 启动Gin服务器
func StartGinServer() {
	// 初始化一个http服务对象
	app := gin.Default()

	// 首先加载templates目录下面的所有模版文件，模版文件扩展名随意
	app.LoadHTMLGlob("website/templates/*")
	app.Static("/static", "website/static")
	router.InitRouter(app)

	// 监听并在 0.0.0.0:8888 上启动服务
	addr := fmt.Sprintf(":%d", config.GlobalConfig.Setup.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: app,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	StartGinServer()
	waitGroup.Wait()

}
