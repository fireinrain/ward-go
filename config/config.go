package config

import (
	"github.com/go-ini/ini"
	"log"
	"net/http"
	"os"
	"sync"
)

const DefaultPort = 8888
const ConfigFile = "setup.ini"

var FirstStartUp = true
var GlobalConfig ServerConfig
var AppServer *http.Server
var Wg *sync.WaitGroup = &sync.WaitGroup{}

type Setup struct {
	ServerName string `ini:"serverName"`
	Theme      string `ini:"theme"`
	Port       int    `ini:"port"`
}

type ServerConfig struct {
	Setup Setup `ini:"setup"`
}

// init
//
//	@Description: 初始化配置
func init() {
	config := ServerConfig{}
	serverConfig := config.InitServerConfig()
	GlobalConfig = serverConfig
}

// InitServerConfig
//
//	@Description: 初始化配置
//	@receiver receiver
//	@return ServerConfig
func (receiver ServerConfig) InitServerConfig() ServerConfig {
	if !IsFileExist(ConfigFile) {
		return ServerConfig{
			Setup: Setup{Port: DefaultPort},
		}
	}
	// Load the INI file
	cfg, err := ini.Load("setup.ini")
	if err != nil {
		log.Fatalf("failed to load configuration file: %v", err)
	}
	// Map the INI file to a struct
	var config ServerConfig
	if err := cfg.MapTo(&config); err != nil {
		log.Fatalf("failed to map configuration file to struct: %v", err)
	}
	FirstStartUp = false
	return config
}

// RefreshServerConfig
//
//	@Description: 刷新配置
func RefreshServerConfig() {
	serverConfig := GlobalConfig.InitServerConfig()
	GlobalConfig = serverConfig
}

// IsFileExist
//
//	@Description: 文件是否存在
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
