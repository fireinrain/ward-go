package utils

import (
	"github.com/go-ini/ini"
	"log"
	"net/http"
	"os"
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

}
