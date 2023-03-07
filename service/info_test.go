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

	fmt.Println("service:", service)

}

func TestConvertUptime2Seperate(t *testing.T) {
	seperate := ConvertUptime2Seperate(1029499)
	for _, value := range seperate {
		fmt.Println(value)
	}
}
