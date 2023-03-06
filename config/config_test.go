package config

import (
	"fmt"
	"testing"
)

func TestInitConfig(t *testing.T) {
	serverConfig := ServerConfig{}
	config := serverConfig.InitServerConfig()
	fmt.Println(config.Setup)
}
