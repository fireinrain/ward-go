package service

import (
	"fmt"
	"github.com/shirou/gopsutil/host"
	"testing"
)

func TestGetServerInfoService(t *testing.T) {
	service := GetServerInfoService()
	info, err := host.Info()
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%v", info)

	fmt.Printf("service: %v\n", service)

}

func TestConvertUptime2Seperate(t *testing.T) {
	seperate := ConvertUptime2Seperate(1029499)
	for _, value := range seperate {
		fmt.Println(value)
	}
}

func TestGetMainHardDriveDiskName(t *testing.T) {
	name := GetMainHardDriveInfo()
	fmt.Println(name)
}

func TestGetMacosDiskCountAndTotalSize(t *testing.T) {
	size, i := GetMacosDiskCountAndTotalSize()
	fmt.Println(size, i)
}

func TestGetServerInfo(t *testing.T) {
	_ = &ServerInfo{}
	serverInfo := GetServerInfo()
	fmt.Println(serverInfo.Storage.MainStorage)

}

func TestMapKeyExist(t *testing.T) {
	dict := map[string]int{"key1": 1, "key2": 2}
	if value, ok := dict["key1"]; ok {
		fmt.Printf(string(value))
	} else {
		fmt.Println("key1 不存在")
	}
}
