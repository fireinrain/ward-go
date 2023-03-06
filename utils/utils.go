package utils

import (
	"context"
	"github.com/go-ini/ini"
	"log"
	"net/http"
	"os"
	"time"
)

// IsFileExist
//
//	@Description: 判断文件是否存在
//	@param path
//	@return bool
func IsFileExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
	return false
}

// WriteConfig2File
//
//	@Description: 将配置写入ini文件
//	@param serverName
//	@param theme
//	@param port
//	@param configPath
func WriteConfig2File(serverName string, theme string, port string, configPath string) {
	empty := ini.Empty()
	section := empty.Section("setup")
	section.Key("serverName").SetValue(serverName)
	section.Key("theme").SetValue(theme)
	section.Key("port").SetValue(port)

	if err := empty.SaveTo(configPath); err != nil {
		log.Fatalf("failed to save config file: %v", err)
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
