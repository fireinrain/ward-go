package config

import (
	"github.com/go-ini/ini"
	"log"
	"net/http"
	"os"
)

const DefaultPort = 8888
const ConfigFile = "setup.ini"

var FirstStartUp = true
var GlobalConfig ServerConfig
var AppServer *http.Server

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
